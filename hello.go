package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	config "config"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/bookme/rest_api", HeadersAuthorization(BookmeRest, config.BookmeAPIKey, config.BookmeAuthorization))
	fmt.Println(config.INFO, "Server started, listening at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// WriteJSONResponse - helper function for writing JSON response to Response Writer
func WriteJSONResponse(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
	return
}

// ValidRequestMethod - This method checks if request method is allowed or not
func ValidRequestMethod(w http.ResponseWriter, r *http.Request, method string) bool {

	if r.Method != method {
		w.WriteHeader(405)
		w.Write([]byte("405 - Method not allowed"))
		fmt.Printf("%s - Forbidden, method not allowed!\n", r.RequestURI)
		return false
	}
	return true
}

// HeadersAuthorization - This method checks if correct Headers are provided in the requests
func HeadersAuthorization(handler http.HandlerFunc, bookmeAPIKey, bookmeAuthorization string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("=================================================================")

		var apikey = r.Header.Get("API_KEY")
		var authorization = r.Header.Get("Authorization")

		if apikey != bookmeAPIKey || authorization != bookmeAuthorization {
			w.WriteHeader(401)
			WriteJSONResponse(w, "Unauthorised.\nInvalid Headers provided (API_KEY, Authorization)")
			return
		}

		handler(w, r)
	}
}

// IsValidAPIKey - Checks if valid API key was passed in form-data
func IsValidAPIKey(w http.ResponseWriter, r *http.Request, APIKey string) bool {
	if APIKey != config.BookmeAPIKey {
		WriteJSONResponse(w, "{\"status\":false,\"error\":\"Invalid API Key.\"}")
		fmt.Println(config.ERROR, r.RequestURI, "Invalid API key!")
		return false
	}
	fmt.Println(config.INFO, r.RequestURI, "'api_key' validated!")
	return true
}

// AreValidCinemaDetails - This method checks if movie details passed in the request are valid
func AreValidCinemaDetails(w http.ResponseWriter, r *http.Request, movieid string, showid string, cinemaid string, date string) bool {

	result, found := GetMovieDetails(movieid)

	if !found {
		WriteJSONResponse(w, fmt.Sprintf(`{"show_id":%s,"hall_id":null,"hall_name":null,"rows":null,"cols":null,"seat_plan":null,"booked_seats":""}`, showid))
		fmt.Println(config.DEBUG, r.RequestURI, "Movieid not found")
		return false
	}

	//Iterate in movie shows list to find show_id
	_result := result["shows"].([]map[string]interface{})
	for k := range _result {
		_item := _result[k]
		// fmt.Println(config.DEBUG, "Show id : ", showid, " Mock show id : ", _item["show_id"])
		if showid == _item["show_id"] {
			return true
		}
	}

	return false
}

// BookmeRest - View function for all bookme API requests
func BookmeRest(w http.ResponseWriter, r *http.Request) {

	if !ValidRequestMethod(w, r, "POST") {
		return
	}

	if err := r.ParseMultipartForm(0); err != nil {
		fmt.Fprintf(w, "\nThere was an error in parsing the form data\n")
		fmt.Fprintf(w, "Error: %v\n", err)
		return
	}

	// Check if valid api_key is passed in the form-data
	apikey := r.FormValue("api_key")
	if !IsValidAPIKey(w, r, apikey) {
		return
	}

	var params = r.URL.Query()

	if params["play_movies"] != nil {
		PlayMovies(w, r)
	} else if params["play_movie_shows"] != nil {
		PlayMovieShows(w, r)
	} else if params["cinema_seatplan"] != nil {
		CinemaSeatPlan(w, r)
	} else if params["cinema_reserve_seats"] != nil {
		CinemaReserveSeats(w, r)
	} else if params["save_cinema"] != nil {
		SaveCinema(w, r)
	} else {
		fmt.Fprintf(w, "Invalid query parameter (endpoint)")
	}

	// fmt.Printf("Got Data! r.PostFrom = %v\n", r.PostForm)
}

// stringInSlice - helper function which provides 'in' functionality as in python
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}
	return false
}

// getRandomSeq - helper function for generating random sequence
func getRandomSeq(n int) string {

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcde0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
