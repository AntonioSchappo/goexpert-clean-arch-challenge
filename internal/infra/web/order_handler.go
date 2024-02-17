package web

import (
	"encoding/json"
	"net/http"

	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/entity"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/usecase"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/pkg/events"
)

type WebOrderHandler struct {
	OrderRepository   entity.IOrderRepository
	CreatedOrderEvent events.IEvent
	EventDispatcher   events.IEventDispatcher
}

func NewWebOrderHandler(
	orderRepository entity.IOrderRepository,
	createdOrderEvent events.IEvent,
	eventDispatcher events.IEventDispatcher) *WebOrderHandler {
	return &WebOrderHandler{
		OrderRepository:   orderRepository,
		CreatedOrderEvent: createdOrderEvent,
		EventDispatcher:   eventDispatcher,
	}
}

func (h *WebOrderHandler) HandleOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.List(w, r)
	} else {
		h.Create(w, r)
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewOrderCreateUseCase(h.OrderRepository, h.CreatedOrderEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *WebOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	listOrder := usecase.NewOrderListUseCase(h.OrderRepository)
	output, err := listOrder.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
