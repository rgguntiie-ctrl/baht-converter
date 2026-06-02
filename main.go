package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/shopspring/decimal"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}
	for _, input := range inputs {
		fmt.Println(input)
		fmt.Println(DecimalToThaiText(input))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/convert", convertHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":"ok"}`)
	})

	log.Printf("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
