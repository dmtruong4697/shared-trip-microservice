package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/api-gateway/grpc_clients"
	"ride-sharing/shared/contracts"
)

func handleTripStart(w http.ResponseWriter, r *http.Request) {
	var reqBody startTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed parse Json data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	tripService, err := grpc_clients.NewTripServiceClient()
	if err != nil {
		log.Fatal(err)
	}
	defer tripService.Close()

	trip, err := tripService.Client.CreateTrip(r.Context(), reqBody.ToProto())
	if err != nil {
		log.Printf("failed to start trip %v", err)
		http.Error(w, "failed to start trip", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: trip}

	writeJSON(w, http.StatusCreated, response)

}

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {

	// }

	// time.Sleep(time.Second * 9)
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

	tripService, err := grpc_clients.NewTripServiceClient()
	if err != nil {
		log.Fatal(err)
	}
	defer tripService.Close()

	// tripService.Client.PreviewTrip()

	//TODO: Call trip service
	tripPreview, err := tripService.Client.PreviewTrip(r.Context(), reqBody.ToProto())
	if err != nil {
		log.Printf("failed to preview trip %v", err)
		http.Error(w, "failed to preview trip", http.StatusInternalServerError)
		return
	}
	response := contracts.APIResponse{Data: tripPreview}

	writeJSON(w, http.StatusCreated, response)
}
