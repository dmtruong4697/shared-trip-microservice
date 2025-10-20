package main

import "ride-sharing/shared/types"

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

// Read implements io.Reader.
func (previewTripRequest) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}
