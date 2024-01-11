package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/sendToBankHostSimulator", handleTransaction)
	http.ListenAndServe(":8084", nil)
}

type RequestBody struct {
	CardNumber     string  `json:"cardNumber"`
	ExpirationDate string  `json:"expirationDate"`
	Balance        float64 `json:"balance"`
	CVC            int     `json:"cvc"`
}

func handleTransaction(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//APIURL for docker
	apiURL := "http://host.docker.internal:8083/api/v1/transactions/processTransaction"

	//APIURL for localhost
	//apiURL := "http://localhost:8083/api/v1/transactions/processTransaction"

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)
}
