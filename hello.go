package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	fmt.Println(config.INFO, "Server started, listening at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// WriteJSONResponse - helper function for writing JSON response to Response Writer
func WriteJSONResponse(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
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
		fmt.Println("=================================================================")
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
		fmt.Println(config.DEBUG, "Show id : ", showid, " Mock show id : ", _item["show_id"])
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
	} else {
		fmt.Fprintf(w, "Invalid query parameter (endpoint)")
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
	WriteJSONResponse(w, string(js))
	return
}

// PlayMovieShows - function for bookme /play_movie_shows
func PlayMovieShows(w http.ResponseWriter, r *http.Request) {

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
		WriteJSONResponse(w, string(_response))
		return
	}

	// Send an empty response as sent by Bookme API (if movie id isn't found)
	WriteJSONResponse(w, "[[]]")
	fmt.Println(config.WARNING, r.RequestURI, "Invalid Movie ID", movieid)
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
				WriteJSONResponse(w, string(js))
				return
			}
		}
		w.Write([]byte(fmt.Sprintf("No seatplan found for show_id: %s", showid)))
		return
	}

	// Send an empty response as sent by Bookme API (if movie id isn't found)
	WriteJSONResponse(w, "[[]]")
	fmt.Println(config.ERROR, r.RequestURI, "Invalid Show ID or Movie ID")
	return
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}
	return false
}

// SeatAvailable - helper function to check seat availability
func SeatAvailable(seatPlanIndex int, seatID string) bool {

	for _, seatPlan := range CinemaSeatPlanMock[seatPlanIndex]["seat_plan"].([]map[string]interface{}) {
		SeatsIterable := seatPlan["seats"].([]map[string]interface{})
		for _, seatDict := range SeatsIterable {
			if seatDict["seat_id"] == seatID {
				return seatDict["status"] == 0
			}
		}
	}
	return false
}

// CheckIfSeatsAvailable - helper function to check all seats availability
func CheckIfSeatsAvailable(seatPlanIndex int, seatNumbers []string) bool {

	// Check if seats user want to book are not reserved
	for _, seatNumber := range seatNumbers {
		if !SeatAvailable(seatPlanIndex, seatNumber) {
			fmt.Println(config.WARNING, "seat number", seatNumber, "is already booked")
			return false
		}
	}
	return true
}

// ReserveCinemaSeats - helper function for reserving cinema seats
func ReserveCinemaSeats(seatPlanIndex int, seatNumbers string) bool {

	// Strip whitespaces from the string and then split coma separated seat numbers
	splittedSeatNumbers := strings.Split(strings.Replace(seatNumbers, " ", "", -1), ",")
	fmt.Println(config.DEBUG, "splittedseatnum:", splittedSeatNumbers)

	seatPlanMap := CinemaSeatPlanMock[seatPlanIndex]
	seatPlanMapIterable := seatPlanMap["seat_plan"].([]map[string]interface{})

	// Check if all the seats user requested for booking are available
	if !CheckIfSeatsAvailable(seatPlanIndex, splittedSeatNumbers) {
		fmt.Println(config.WARNING, "seats are already reserved (ref1)")
		return false
	}

	for _, seatPlan := range seatPlanMapIterable {
		SeatsIterable := seatPlan["seats"].([]map[string]interface{})
		for _, seatDict := range SeatsIterable {
			if stringInSlice(seatDict["seat_id"].(string), splittedSeatNumbers) {
				if seatDict["status"] == 0 {
					seatDict["status"] = 1
					fmt.Println(config.INFO, "reserved seat:", seatDict["seat_id"])
				} else {
					fmt.Println(config.WARNING, "seat number:", seatDict["seat_id"], "already reserved [WHY]")
				}
			}
		}
	}
	return true
}

// CinemaReserveSeats - function for bookme /cinema_reserve_seats
func CinemaReserveSeats(w http.ResponseWriter, r *http.Request) {

	showid := r.FormValue("show_id")
	movieid := r.FormValue("movie_id")
	cinemaid := r.FormValue("cinema_id")
	showDate := r.FormValue("show_date")
	// showTime := r.FormValue("show_time")
	seatNumbers := r.FormValue("seat_numbers")
	// seats := r.FormValue("seats")
	// ticketPrice := r.FormValue("ticket_price")

	if AreValidCinemaDetails(w, r, movieid, showid, cinemaid, showDate) {
		for k := range CinemaSeatPlanMock {
			if CinemaSeatPlanMock[k]["show_id"] == showid {
				reserved := ReserveCinemaSeats(k, seatNumbers)
				if !reserved {
					WriteJSONResponse(w, `[{"status":"failed", "msg":"Seats are aready allocated."}]`)
				} else {
					WriteJSONResponse(w, fmt.Sprintf(`[{"status":"success", "msg":"Seats are allocated successfully.", "booking_no":"%d"}]`, config.BookingNumber))
					// increment the BookingNumber variable
					config.BookingNumber++
				}
				return
			}
		}
		WriteJSONResponse(w, fmt.Sprintf("No seatplan found for show_id: %s", showid))
		return
	}

}
