package main

import (
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {

	// }

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

	//TODO: Call trip service

	res := contracts.APIResponse{Data: "ok"}

	writeJSON(w, http.StatusCreated, res)
}
