package handler

import (
	"app-booking/internal/model"
	"encoding/json"
	"log/slog"
	"net/http"
)

type OrderHandler struct {
	repo OrderRepository
}

func NewOrderHandler(repo OrderRepository) *OrderHandler {
	return &OrderHandler{
		repo: repo,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		slog.Error("Incorrect request body", err)
		http.Error(w, "Incorrect request body", http.StatusBadRequest)
		return
	}

	if err := h.repo.CreateOrder(&order); err != nil {
		slog.Error("Can`t create order", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.repo.GetAllOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		return
	}
}
