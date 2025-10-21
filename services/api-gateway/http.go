package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
	"time"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {

	// }

	time.Sleep(time.Second * 9)
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed parse Json data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	//validation
	if reqBody.UserID == "" {
		http.Error(w, "user id is required", http.StatusBadRequest)
		return
	}

	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(jsonBody)

	//TODO: Call trip service
	res, err := http.Post("http://trip-service:8083/preview", "application/json", reader)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()

	var respBody any
	if err := json.NewDecoder(res.Body).Decode(&respBody); err != nil {
		http.Error(w, "failed parse Json data from trip-service", http.StatusBadRequest)
		return
	}

	response := contracts.APIResponse{Data: respBody}

	writeJSON(w, http.StatusCreated, response)
}
