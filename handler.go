package main

import (
	"encoding/json"
	"net/http"

	"github.com/shopspring/decimal"
)

type ConvertRequest struct {
	Amount string `json:"amount"`
}

type ConvertResponse struct {
	Amount string `json:"amount"`
	Text   string `json:"text"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "method not allowed"})
		return
	}

	var req ConvertRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid JSON body"})
		return
	}

	if req.Amount == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "amount is required"})
		return
	}

	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid amount"})
		return
	}

	if amount.IsNegative() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "amount must not be negative"})
		return
	}

	json.NewEncoder(w).Encode(ConvertResponse{
		Amount: amount.String(),
		Text:   DecimalToThaiText(amount),
	})
}
