package transport

import (
	"config"
	"encoding/json"
	"fmt"
	"net/http"
	"utils"
)

// GetTransportServices - function for get_transport_services
func GetTransportServices(w http.ResponseWriter, r *http.Request) {

	js, err := json.Marshal(transportList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return json response
	utils.WriteJSONResponse(w, string(js))
	return
}

func getServiceDepartureCities(serviceID string) []byte {
	for _, serviceMap := range departureCities {
		if serviceMap["service_id"] == serviceID {
			js, _ := json.Marshal(serviceMap["departure_cities"])
			return js
		}
	}
	return []byte("[]")
}

// GetDepartureCities - function for get_deparature_cities
func GetDepartureCities(w http.ResponseWriter, r *http.Request) {

	serviceID := r.FormValue("service_id")

	response := getServiceDepartureCities(serviceID)

	utils.WriteJSONResponse(w, string(response))

}

func getAllDestinationCitiesOfAService(serviceID string) ([]byte, bool) {
	allDestinationCitiesOfAService := []map[string]string{}
	cityIDs := make(map[string]bool)

	for _, serviceMap := range destinationCities {
		if serviceID == serviceMap["service_id"] {
			for _, destinationCity := range serviceMap["destination_cities"].([]map[string]string) {
				// Check if we already appended this destination city in array of maps
				if !cityIDs[destinationCity["destination_city_id"]] {
					allDestinationCitiesOfAService = append(allDestinationCitiesOfAService, destinationCity)
					cityIDs[destinationCity["destination_city_id"]] = true
				}
			}
		}
	}

	if len(allDestinationCitiesOfAService) > 0 {
		js, _ := json.Marshal(allDestinationCitiesOfAService)

		// if we pass invalid origin_city_id in real bookme API, they return the map of all destination cities of that service_id
		return js, true
	}

	// return an empty list incase of an invalid service_id
	return []byte("[]"), false
}

func getServiceDestinationCities(serviceID, originCityID string) []byte {
	for _, serviceMap := range destinationCities {
		if serviceID == serviceMap["service_id"] && originCityID == serviceMap["origin_city_id"] {
			js, _ := json.Marshal(serviceMap["destination_cities"])
			return js
		}
	}

	fmt.Println(config.WARNING, "origin city id:", originCityID, "not found for service id: ", serviceID)

	response, found := getAllDestinationCitiesOfAService(serviceID)
	if found {
		fmt.Println(config.WARNING, "returning all destination cities for service id:", serviceID)
	} else {
		fmt.Println(config.WARNING, "invalid service id:", serviceID, "returning empty list in response")
	}
	return response
}

// GetDestinationCities - function for /get_destination_cities
func GetDestinationCities(w http.ResponseWriter, r *http.Request) {

	serviceID := r.FormValue("service_id")
	originCityID := r.FormValue("origin_city_id")

	response := getServiceDestinationCities(serviceID, originCityID)

	utils.WriteJSONResponse(w, string(response))
	return
}
