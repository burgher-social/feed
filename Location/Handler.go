package Location

import (
	Utils "burgher/Utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	var locationReq LocationRequest
	err := json.NewDecoder(r.Body).Decode(&locationReq)
	fmt.Println(locationReq)
	if err != nil {
		// return HTTP 400 bad request
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	location := Location{
		Id:        Utils.GenerateId(),
		UserId:    locationReq.UserId,
		Latitude:  locationReq.Latitude,
		Longitude: locationReq.Longitude,
		City:      locationReq.City,
	}
	loc, err := create(location)
	if err != nil {
		w.WriteHeader(503)
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(LocationResponse{
		Id:        loc.Id,
		UserId:    loc.UserId,
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
		City:      loc.City,
	})
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	type ReadRequest struct {
		UserId string `json:"userId"`
	}
	var readRequest ReadRequest
	err := json.NewDecoder(r.Body).Decode(&readRequest)
	fmt.Println(readRequest)
	if err != nil {
		// return HTTP 400 bad request
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	locations, notfound := Read(readRequest.UserId)

	if notfound != nil {
		w.WriteHeader(503)
		fmt.Println(notfound)
		w.Write([]byte(notfound.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	resp := LocationResponse{
		Id:        locations.Id,
		UserId:    locations.UserId,
		Latitude:  locations.Latitude,
		Longitude: locations.Longitude,
		City:      locations.City,
	}
	json.NewEncoder(w).Encode(&resp)
}
