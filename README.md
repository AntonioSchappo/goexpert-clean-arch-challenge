# goexpert-clean-arch-challenge

This project is an implementation of a simple application structured according to the Clean Architecture paradigm using Go.

In this application there are two basic usecases (`CreateOrder` and `ListOrders`) that can be accessed through a gRPC, REST and GraphQL servers.

The `CreateOrder` usecase also propagates a message to a RabbitMQ queue.

## Commands to start the application

1. It is advised to execute the command below at the root of the project:

```sh
go mod tidy
```

2. The following command is necessary to start the RabbitMQ and MySQL containers:

```sh
docker compose up
```

3. Run the migration up command in order to create the `orders` database table:
```sh
make migrateup
```

4. In order to run the application, please execute the command below to go to the `/cmd/system` folder:
```sh
cd cmd/system/
```

5. Run the application:
```sh
go run main.go wire_gen.go
```

## Instructions for running the GraphQL playground

With the application running access the [GraphQL playground](http://localhost:8080/) on your browser.

Please check below an example of a GraphQL mutation and a query:
```sh
mutation createOrder {
  createOrder(input: {id:"e", Price: 139.5, Tax: 21.5}) {
    id
    Price
    Tax
    FinalPrice
  }
}

query listOrders {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

## Instructions for running the Evans reflection server for gRPC on REPL

1. With the application running type the following command to start the Evans REPL:
```sh
evans -r repl
```

2. Select the package by typing:
```sh
package pb
```
3. Select the service by typing:
```sh
service OrderService
```
4. Choose between the `CreateOrder` or `ListOrders` procedures. The `CreateOrder` procedure will require you to type the parameters in the specified order:
```sh
call CreateOrder
```
```sh
call ListOrders
```

## Instructions for making requests to the REST server

On the `/api` folder there are three http files with examples of requests that can be made to the Rest server

If you are using VSCode as your IDE downloading the [Rest Client Plugin](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) is an excellent way execute these requests.

## Instructions for creating and binding a queue to an exchange in RabbitMQ:

1. Access the [RabbitMQ interface](http://localhost:15672/) through your browser and type in `guest` as both **Username** and **Password**.

2. Select the **Queues and Streams** tab at the top of the screen

3. Press on the dropdown menu titled **Add a new queue**, add a name to the queue and press the **Add queue** button.

4. Select the recently created queue by pressing on its name at the overview table

5. Press on the dropdown menu titled **Bindings** and on the field *From exchange* type `amq.direct` and press the **Bind** button.

6. Using either the gRPC REPL, the REST Server requests in the `/api` folder or the GraphQL playground produce a message by executing the `CreateOrder` usecase.

7. To check if the message was successfully created, press on the dropdown menu titled **Get Messages** and press the button **Get Message(s)**.
