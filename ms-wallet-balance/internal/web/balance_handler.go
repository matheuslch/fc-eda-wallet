package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	find_balance "github.com/matheuslch/fc-ms-wallet-balance/internal/usecases/find_balance"
)

type WebBalanceHandler struct {
	FindBalanceUseCase find_balance.FindBalanceUseCase
}

func NewWebBalanceHandler(findBalanceUseCase find_balance.FindBalanceUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		FindBalanceUseCase: findBalanceUseCase,
	}
}

func (h *WebBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	AccountId := chi.URLParam(r, "account_id")

	if AccountId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.FindBalanceUseCase.Execute(AccountId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
