package transport

import (
	"config"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"utils"
)

// GetTransportServices - function for /get_transport_services
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

// GetDepartureCities - function for /get_deparature_cities
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

// BusTimes - function for /bus_times
func BusTimes(w http.ResponseWriter, r *http.Request) {

	serviceID := r.FormValue("service_id")
	departureCityID := r.FormValue("departure_city_id")
	arrivalCityID := r.FormValue("arrival_city_id")
	date := r.FormValue("date")

	_serviceID, _ := strconv.Atoi(serviceID)
	_departureCityID, _ := strconv.Atoi(departureCityID)
	_arrivalCityID, _ := strconv.Atoi(arrivalCityID)

	for _, timeDict := range timeIDs {
		if _serviceID == timeDict["service_id"] && _departureCityID == timeDict["departure_city_id"] &&
			_arrivalCityID == timeDict["arrival_city_id"] && date == timeDict["date"] {
			facilities, _ := json.Marshal(timeDict["facilities"])
			utils.WriteJSONResponse(w, fmt.Sprintf(`{
				"times": [{
					"departure_city_id": %d,
					"departure_city_name": "%s",
					"arrival_city_id": %d,
					"arrival_city_name": "%s",
					"service_id": %d,
					"service_name": "%s",
					"time_id": %d,
					"schedule_id": %d,
					"route_id": %d,
					"time": "%s",
					"arrtime": "%s",
					"original_fare": %d,
					"fare": %d,
					"handling_charges": %d,
					"easypaisa_charges": %d,
					"thumb": "%s",
					"seats": %d,
					"busname": "%s",
					"bustype": "%s",
					"btype_id": %d,
					"facilities": %s,
					"seat_info": "%s"
				}]
			}`,
				timeDict["departure_city_id"], timeDict["departure_city_name"], timeDict["arrival_city_id"],
				timeDict["arrival_city_name"], timeDict["service_id"], timeDict["service_name"], timeDict["time_id"],
				timeDict["schedule_id"], timeDict["route_id"], timeDict["time"], timeDict["arrtime"],
				timeDict["original_fare"], timeDict["fare"], timeDict["handling_charges"],
				timeDict["easypaisa_charges"], timeDict["thumb"], timeDict["seats"], timeDict["busname"],
				timeDict["bustype"], timeDict["btype_id"], string(facilities), timeDict["seat_info"]))

			return
		}
	}

	utils.WriteJSONResponse(w, `{"times":[]}`)
	return
}
