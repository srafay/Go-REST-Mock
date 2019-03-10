package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	config "config"

	"github.com/gorilla/mux"
)

// show - struct for the show field
type show struct {
	CityID           string `json:city_id`
	CityName         string `json:city_name`
	ShowID           string `json:show_id`
	ShowMovieID      string `json:show_movie_id`
	ShowCinemaID     string `json:show_cenima_id`
	CinemaName       string `json:cinema_name`
	HallID           string `json:hall_id`
	HallName         string `json:hall_name`
	ShowDate         string `json:show_date`
	ShowStartTime    string `json:show_start_time`
	ShowTime         string `json:show_time`
	TicketPrice      string `json:ticket_price`
	HandlingCharges  int    `json:handling_charges`
	EasyPaisaCharges int    `json:easypaisa_charges`
	HouseFull        string `json:house_full`
	ETicket          string `json:eticket`
}

// playMovieShows - Struct for returning in Bookme play_movie_shows request
type playMovieShows struct {
	MovieID       string `json:movie_id`
	IMDBID        string `json:imdb_id`
	Title         string `json:title`
	Genre         string `json:genre`
	Language      string `json:language`
	Director      string `json:director`
	Producer      string `json:producer`
	ReleaseDate   string `json:release_date`
	Cast          string `json:cast`
	Ranking       string `json:ranking`
	Length        string `json:length`
	Thumbnail     string `json:thumbnail`
	MusicDirector string `json:music_director`
	Country       string `json:country`
	Synopsis      string `json:synopsis`
	Details       string `json:details`
	TrailerLink   string `json:trailer_link`
	Date          string `json:date`
	BookingType   string `json:booking_type`
	Points        string `json:points`
	UpdateDate    string `json:update_date`
	CloseDate     string `json:close_date`
	Status        string `json:status`
	Shows         []show `json:shows`
}

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/bookme/rest_api", HeadersAuthorization(BookmeRest, config.BookmeAPIKey, config.BookmeAuthorization))
	fmt.Println("Server started, listening at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
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

		var apikey = r.Header.Get("API_KEY")
		var authorization = r.Header.Get("Authorization")

		if apikey != bookmeAPIKey || authorization != bookmeAuthorization {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\nInvalid Headers provided (API_KEY, Authorization)"))
			return
		}

		handler(w, r)
	}
}

// IsValidAPIKey - Checks if valid API key was passed in form-data
func IsValidAPIKey(w http.ResponseWriter, r *http.Request, APIKey string) bool {
	if APIKey != config.BookmeAPIKey {
		w.WriteHeader(200)
		w.Write([]byte("{\"status\":false,\"error\":\"Invalid API Key.\"}"))
		fmt.Println(r.RequestURI, "- Error, Invalid API key!")
		return false
	}
	fmt.Println(r.RequestURI, "- 'api_key' validated!")
	return true
}

// AreValidCinemaDetails - This method checks if movie details passed in the request are valid
func AreValidCinemaDetails(w http.ResponseWriter, r *http.Request, movieid string, showid string, cinemaid string, date string) bool {

	result, found := GetMovieDetails(movieid)

	if !found {
		w.Write([]byte(fmt.Sprintf(`{"show_id":%s,"hall_id":null,"hall_name":null,"rows":null,"cols":null,"seat_plan":null,"booked_seats":""}`, showid)))
		fmt.Printf("%s - Movieid not found\n", r.RequestURI)
		return false
	}

	//Iterate in movie shows list to find show_id
	_result := result["shows"].([]map[string]interface{})
	for k := range _result {
		_item := _result[k]
		fmt.Println("Show id : ", showid, " Mock show id : ", _item["show_id"])
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
	} else {
		fmt.Fprintf(w, "Invalid query parameter")
	}

	// fmt.Printf("Got Data! r.PostFrom = %v\n", r.PostForm)
}

// GetMovieDetails - function for getting movie details of a particular movie_id
func GetMovieDetails(movieid string) (map[string]interface{}, bool) {

	for k := range playMovieShowsList {
		if movieid == playMovieShowsList[k]["movie_id"] {
			return playMovieShowsList[k], true
		}
	}
	return nil, false
}

// PlayMovies - function for bookme /play_movies
func PlayMovies(w http.ResponseWriter, r *http.Request) {

	js, err := json.Marshal(playMoviesList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return json response
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

// PlayMovieShows - function for bookme /play_movie_shows
func PlayMovieShows(w http.ResponseWriter, r *http.Request) {

	apikey := r.FormValue("api_key")
	if !IsValidAPIKey(w, r, apikey) {
		return
	}

	movieid := r.FormValue("movie_id")

	response, found := GetMovieDetails(movieid)

	if found {

		var resp = []map[string]interface{}{response}
		// Return json response
		_response, err := json.Marshal(resp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(_response)
		return
	}

	// Send an empty response as sent by Bookme API (if movie id isn't found)
	w.WriteHeader(200)
	w.Write([]byte("[[]]"))
	fmt.Println(r.RequestURI, "- [Warning], Invalid Movie ID!")
	return
}

// CinemaSeatPlan - function for bookme /play_movie_shows
func CinemaSeatPlan(w http.ResponseWriter, r *http.Request) {

	movieid := r.FormValue("movie_id")
	showid := r.FormValue("show_id")
	cinemaid := r.FormValue("cinema_id")
	date := r.FormValue("date")

	if AreValidCinemaDetails(w, r, movieid, showid, cinemaid, date) {
		for k := range CinemaSeatPlanMock {
			if CinemaSeatPlanMock[k]["show_id"] == showid {
				js, err := json.Marshal(CinemaSeatPlanMock[k])
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
				return
			}
		}
		w.Write([]byte(fmt.Sprintf("No seatplan found for show_id: %s", showid)))
		return
	}

	// if string(response) != "None" {

	// 	// Return json response
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(response)
	// 	return
	// }

	// Send an empty response as sent by Bookme API (if movie id isn't found)
	w.Write([]byte("[[]]"))
	fmt.Println(r.RequestURI, "- Error, Invalid Show ID or Movie ID")
	return
}
