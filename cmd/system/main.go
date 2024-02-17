package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/configs"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/event/handler"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/infra/graph"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/infra/grpc/pb"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/infra/grpc/service"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/infra/web/webserver"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dbConn, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("CreatedOrder", &handler.CreateOrderHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	//Usecases
	createOrderUsecase := NewOrderCreateUseCase(dbConn, eventDispatcher)
	listOrderUsecase := NewOrderListUseCase(dbConn)

	//Webserver
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(dbConn, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.HandleOrder)
	fmt.Println("Server is running on port", configs.WebServerPort)
	go webserver.Start()

	//GrpcServer
	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService(*createOrderUsecase, *listOrderUsecase)
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	fmt.Println("gRPC server is running on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	//GraphQL Server
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUsecase: *createOrderUsecase,
		ListOrdersUsecase:  *listOrderUsecase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("GraphQL server is running on port", configs.GraphQLServerPort)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GraphQLServerPort), nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
