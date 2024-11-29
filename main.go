package main

import (
    "encoding/json"
    "html/template"
    "log"
    "net/http"
    "path/filepath"
)

type ConversionRequest struct {
    Amount float64 `json:"amount"`
}

type ConversionResponse struct {
    USDAmount float64 `json:"usd_amount"`
    NPRAmount float64 `json:"npr_amount"`
    Rate      float64 `json:"rate"`
}

const EXCHANGE_RATE = 132.50

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", handleHome)
    http.HandleFunc("/convert", handleConversion)

    log.Println("Server starting on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))
    tmpl.Execute(w, nil)
}

func handleConversion(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req ConversionRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    response := ConversionResponse{
        USDAmount: req.Amount,
        NPRAmount: req.Amount * EXCHANGE_RATE,
        Rate:      EXCHANGE_RATE,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
