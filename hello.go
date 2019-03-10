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

var playMovieShowsList []playMovieShows

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

// AreValidMovieDetails - This method checks if movie details passed in the request are valid
func AreValidMovieDetails(w http.ResponseWriter, r *http.Request, movieid string, showid string, cinemaid string, date string) bool {

	result := GetMovieDetails(movieid)
	if string(result) == "None" {
		w.Write([]byte(fmt.Sprintf(`{"show_id":%s,"hall_id":null,"hall_name":null,"rows":null,"cols":null,"seat_plan":null,"booked_seats":""}`, showid)))
		fmt.Printf("%s - Movieid not found\n", r.RequestURI)
		return false
	}

	return true
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
	} else {
		fmt.Fprintf(w, "FAILED")
	}

	// fmt.Printf("Got Data! r.PostFrom = %v\n", r.PostForm)
}

// GetMovieDetails - function for getting movie details of a particular movie_id
func GetMovieDetails(movieid string) []byte {
	var response []byte
	if movieid == "934" {
		response = []byte(`[{"movie_id":"934","imdb_id":"tt4154664","title":"Captain Marvel","genre":"Action, Adventure","language":"English","director":"Anna Boden, Ryan Fleck","producer":"Victoria Alonso","release_date":"2019-03-08","music_director":"Pinar Toprak","country":"USA","cast":"Brie Larson","synopsis":"","details":"","ranking":"6.0","length":"124","trailer_link":"https:\/\/bookme.pk\/custom\/videoupload\/Captain_Marvel_T.mp4","thumbnail":"https:\/\/bookme.pk\/custom\/upload\/marvil_IMDB.jpg","date":"2019-03-01 18:39:21","booking_type":"0","points":"0","update_date":"2019-03-01 16:56:31","close_date":null,"status":"1","shows":[{"city_id":"11","city_name":"Rawalpindi","show_id":"477974","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall General","show_date":"2019-03-08","show_start_time":"2019-03-08 15:45:00","show_time":"15:45","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477975","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall Vip Box","show_date":"2019-03-08","show_start_time":"2019-03-08 15:45:00","show_time":"15:45","ticket_price":"3000","handling_charges":150,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477981","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall Vip Box","show_date":"2019-03-08","show_start_time":"2019-03-08 18:30:00","show_time":"18:30","ticket_price":"3000","handling_charges":150,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477980","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall General","show_date":"2019-03-08","show_start_time":"2019-03-08 18:30:00","show_time":"18:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477987","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall Vip Box","show_date":"2019-03-08","show_start_time":"2019-03-08 20:45:00","show_time":"20:45","ticket_price":"3000","handling_charges":150,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477986","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall General","show_date":"2019-03-08","show_start_time":"2019-03-08 20:45:00","show_time":"20:45","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477993","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall Vip Box","show_date":"2019-03-08","show_start_time":"2019-03-08 23:00:00","show_time":"23:00","ticket_price":"3000","handling_charges":150,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477992","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall General","show_date":"2019-03-08","show_start_time":"2019-03-08 23:00:00","show_time":"23:00","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477977","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall Vip Box","show_date":"2019-03-09","show_start_time":"2019-03-09 15:45:00","show_time":"15:45","ticket_price":"3000","handling_charges":150,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477976","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall General","show_date":"2019-03-09","show_start_time":"2019-03-09 15:45:00","show_time":"15:45","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477983","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall Vip Box","show_date":"2019-03-09","show_start_time":"2019-03-09 18:30:00","show_time":"18:30","ticket_price":"3000","handling_charges":150,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477982","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall General","show_date":"2019-03-09","show_start_time":"2019-03-09 18:30:00","show_time":"18:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477989","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall Vip Box","show_date":"2019-03-09","show_start_time":"2019-03-09 20:45:00","show_time":"20:45","ticket_price":"3000","handling_charges":150,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477988","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall General","show_date":"2019-03-09","show_start_time":"2019-03-09 20:45:00","show_time":"20:45","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477995","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall Vip Box","show_date":"2019-03-09","show_start_time":"2019-03-09 23:00:00","show_time":"23:00","ticket_price":"3000","handling_charges":150,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477994","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall General","show_date":"2019-03-09","show_start_time":"2019-03-09 23:00:00","show_time":"23:00","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477979","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall Vip Box","show_date":"2019-03-10","show_start_time":"2019-03-10 15:45:00","show_time":"15:45","ticket_price":"3000","handling_charges":150,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477978","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall General","show_date":"2019-03-10","show_start_time":"2019-03-10 15:45:00","show_time":"15:45","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477984","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall General","show_date":"2019-03-10","show_start_time":"2019-03-10 18:30:00","show_time":"18:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477985","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall Vip Box","show_date":"2019-03-10","show_start_time":"2019-03-10 18:30:00","show_time":"18:30","ticket_price":"3000","handling_charges":150,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477990","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall General","show_date":"2019-03-10","show_start_time":"2019-03-10 20:45:00","show_time":"20:45","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477991","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall Vip Box","show_date":"2019-03-10","show_start_time":"2019-03-10 20:45:00","show_time":"20:45","ticket_price":"3000","handling_charges":150,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477997","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall Vip Box","show_date":"2019-03-10","show_start_time":"2019-03-10 23:00:00","show_time":"23:00","ticket_price":"3000","handling_charges":150,"easypaisa_charges":0,"house_full":"0","eticket":""},{"city_id":"11","city_name":"Rawalpindi","show_id":"477996","show_movie_id":"934","show_cenima_id":"10009","cinema_name":"Odeon Cineplex","hall_id":"203","hall_name":"Red Hall General","show_date":"2019-03-10","show_start_time":"2019-03-10 23:00:00","show_time":"23:00","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":""}]}]`)
	} else if movieid == "901" {
		response = []byte(`[{"movie_id":"901","imdb_id":"","title":"3 BAHADUR","genre":"Adventure,  Animation, Family","language":"Urdu","director":"Sharmeen Obaid Chinoy","producer":"Waadi Animations","release_date":"2018-12-14","music_director":"","country":"pakistan","cast":" Mehwish HayatFahad MustafaSarwat GillaniNimra BuchaBehroze Sabzwari","synopsis":"3 Bahadur: Rise of the Warriors","details":"3 Bahadur: Rise of the Warriors","ranking":"0.9","length":"120","trailer_link":"https:\/\/bookme.pk\/custom\/videoupload\/Rise_of_The_Warriors.mp4","thumbnail":"https:\/\/bookme.pk\/custom\/upload\/3_Bahadur_RiseWarriors_IMDB.jpeg","date":"2019-01-11 11:26:50","booking_type":"0","points":"0","update_date":"2018-12-11 00:00:00","close_date":null,"status":"1","shows":[{"city_id":"1","city_name":"Lahore","show_id":"480497","show_movie_id":"901","show_cenima_id":"30","cinema_name":"Imperial Cinema","hall_id":"113","hall_name":"IMP 2 STAN 500","show_date":"2019-03-08","show_start_time":"2019-03-08 16:30:00","show_time":"16:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"0"},{"city_id":"1","city_name":"Lahore","show_id":"480498","show_movie_id":"901","show_cenima_id":"30","cinema_name":"Imperial Cinema","hall_id":"113","hall_name":"IMP 2 LUXURY TICKET","show_date":"2019-03-08","show_start_time":"2019-03-08 16:30:00","show_time":"16:30","ticket_price":"1500","handling_charges":75,"easypaisa_charges":0,"house_full":"0","eticket":"0"},{"city_id":"1","city_name":"Lahore","show_id":"480504","show_movie_id":"901","show_cenima_id":"30","cinema_name":"Imperial Cinema","hall_id":"113","hall_name":"IMP 2 LUXURY TICKET","show_date":"2019-03-09","show_start_time":"2019-03-09 16:30:00","show_time":"16:30","ticket_price":"1500","handling_charges":75,"easypaisa_charges":0,"house_full":"0","eticket":"0"},{"city_id":"1","city_name":"Lahore","show_id":"480503","show_movie_id":"901","show_cenima_id":"30","cinema_name":"Imperial Cinema","hall_id":"113","hall_name":"IMP 2 STAN 500","show_date":"2019-03-09","show_start_time":"2019-03-09 16:30:00","show_time":"16:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"0"},{"city_id":"1","city_name":"Lahore","show_id":"480510","show_movie_id":"901","show_cenima_id":"30","cinema_name":"Imperial Cinema","hall_id":"113","hall_name":"IMP 2 LUXURY TICKET","show_date":"2019-03-10","show_start_time":"2019-03-10 16:30:00","show_time":"16:30","ticket_price":"1500","handling_charges":75,"easypaisa_charges":0,"house_full":"0","eticket":"0"},{"city_id":"1","city_name":"Lahore","show_id":"480509","show_movie_id":"901","show_cenima_id":"30","cinema_name":"Imperial Cinema","hall_id":"113","hall_name":"IMP 2 STAN 500","show_date":"2019-03-10","show_start_time":"2019-03-10 16:30:00","show_time":"16:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"0"}]}]`)
	} else if movieid == "898" {
		response = []byte(`[{"movie_id":"898","imdb_id":"tt1477834","title":"AQUAMAN","genre":"Action, Adventure, Fantasy","language":"English","director":"James Wan","producer":"Peter Safran","release_date":"2018-12-21","music_director":"Rupert Gregson Williams","country":"Australia               |         USA","cast":"Jason Momoa, Amber Heard, Nicole Kidman","synopsis":"Arthur Curry learns that he is the heir to the underwater kingdom of Atlantis, and must step forward to lead his people and be a hero to the world.","details":"","ranking":"7.5","length":"143","trailer_link":"https:\/\/bookme.pk\/custom\/videoupload\/Aquaman_Trailer.mp4","thumbnail":"https:\/\/bookme.pk\/custom\/upload\/Aquaman_IMDB.jpg","date":"2019-02-07 19:02:35","booking_type":"0","points":"0","update_date":"2018-12-10 11:20:05","close_date":null,"status":"1","shows":[{"city_id":"1","city_name":"Lahore","show_id":"479601","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"148","hall_name":"Screen 3 Premium B","show_date":"2019-03-07","show_start_time":"2019-03-07 22:30:00","show_time":"22:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"479600","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"148","hall_name":"Screen 3 Premium A","show_date":"2019-03-07","show_start_time":"2019-03-07 22:30:00","show_time":"22:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"479599","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"148","hall_name":"Screen 3 Regular","show_date":"2019-03-07","show_start_time":"2019-03-07 22:30:00","show_time":"22:30","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480195","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Premium B","show_date":"2019-03-08","show_start_time":"2019-03-08 14:00:00","show_time":"14:00","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480194","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Premium A","show_date":"2019-03-08","show_start_time":"2019-03-08 14:00:00","show_time":"14:00","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480193","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Regular","show_date":"2019-03-08","show_start_time":"2019-03-08 14:00:00","show_time":"14:00","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480197","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Premium A","show_date":"2019-03-08","show_start_time":"2019-03-08 23:30:00","show_time":"23:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480196","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Regular","show_date":"2019-03-08","show_start_time":"2019-03-08 23:30:00","show_time":"23:30","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480198","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Premium B","show_date":"2019-03-08","show_start_time":"2019-03-08 23:30:00","show_time":"23:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480201","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Premium B","show_date":"2019-03-09","show_start_time":"2019-03-09 14:00:00","show_time":"14:00","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480200","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Premium A","show_date":"2019-03-09","show_start_time":"2019-03-09 14:00:00","show_time":"14:00","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480199","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Regular","show_date":"2019-03-09","show_start_time":"2019-03-09 14:00:00","show_time":"14:00","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480203","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Premium A","show_date":"2019-03-09","show_start_time":"2019-03-09 23:30:00","show_time":"23:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480202","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Regular","show_date":"2019-03-09","show_start_time":"2019-03-09 23:30:00","show_time":"23:30","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480204","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Premium B","show_date":"2019-03-09","show_start_time":"2019-03-09 23:30:00","show_time":"23:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480207","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Premium B","show_date":"2019-03-10","show_start_time":"2019-03-10 14:00:00","show_time":"14:00","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480206","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Premium A","show_date":"2019-03-10","show_start_time":"2019-03-10 14:00:00","show_time":"14:00","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480205","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Regular","show_date":"2019-03-10","show_start_time":"2019-03-10 14:00:00","show_time":"14:00","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480208","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Regular","show_date":"2019-03-10","show_start_time":"2019-03-10 23:30:00","show_time":"23:30","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480210","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Premium B","show_date":"2019-03-10","show_start_time":"2019-03-10 23:30:00","show_time":"23:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480209","show_movie_id":"898","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"147","hall_name":"Screen 2 Premium A","show_date":"2019-03-10","show_start_time":"2019-03-10 23:30:00","show_time":"23:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"3","city_name":"Islamabad","show_id":"477337","show_movie_id":"898","show_cenima_id":"20","cinema_name":"Centaurus Cineplex","hall_id":"74","hall_name":"Cinema 3 3D ","show_date":"2019-03-07","show_start_time":"2019-03-07 22:15:00","show_time":"22:15","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"3","city_name":"Islamabad","show_id":"477353","show_movie_id":"898","show_cenima_id":"20","cinema_name":"Centaurus Cineplex","hall_id":"72","hall_name":"Cinema 1 3D ","show_date":"2019-03-07","show_start_time":"2019-03-07 22:30:00","show_time":"22:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"479397","show_movie_id":"898","show_cenima_id":"25","cinema_name":"Super Cinema Vogue Tower","hall_id":"36","hall_name":"SC SUPER 2 Regular","show_date":"2019-03-07","show_start_time":"2019-03-07 23:30:00","show_time":"23:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"1"}]}]`)
	} else if movieid == "881" {
		response = []byte(`[{"movie_id":"881","imdb_id":"tt1727824","title":"Bohemian Rhapsody","genre":"Drama, Biography  , Music","language":"English","director":"Bryan Singer","producer":"","release_date":"2018-11-02","music_director":"","country":"UK               |         USA","cast":"Rami Malek, Lucy Boynton, Gwilym Lee","synopsis":"Bohemian Rhapsody","details":"Bohemian Rhapsody","ranking":"8.1","length":"134","trailer_link":"https:\/\/bookme.pk\/custom\/videoupload\/Bohemian_Rhapsody_Trailer.mp4","thumbnail":"https:\/\/bookme.pk\/custom\/upload\/Bohemian IMDB.jpg","date":"2019-03-01 11:24:35","booking_type":"0","points":"0","update_date":null,"close_date":null,"status":"1","shows":[{"city_id":"3","city_name":"Islamabad","show_id":"477361","show_movie_id":"881","show_cenima_id":"20","cinema_name":"Centaurus Cineplex","hall_id":"69","hall_name":"Cinema 2 2D ","show_date":"2019-03-07","show_start_time":"2019-03-07 22:15:00","show_time":"22:15","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"}]}]`)
	} else if movieid == "855" {
		response = []byte(`[{"movie_id":"855","imdb_id":"","title":"THE DONKEY KING","genre":" Animation","language":"Urdu","director":"Aziz Jindani","producer":"Talisman Studios","release_date":"2018-10-13","music_director":"Various artists, Shani Arshad","country":"Pakistan","cast":"Jan Rambo, Ismail Tara, Hina Dilpazeer, Ghulam Mohiuddin, Jawed Sheikh","synopsis":"The Donkey King is an upcoming Pakistani computer animated comedy film directed by Aziz Jindani. The film stars the voices of Jan Rambo, Ismail Tara, Hina Dilpazeer, Ghulam Mohiuddin, and Jawed Sheikh.","details":"","ranking":"7.2","length":"120","trailer_link":"https:\/\/bookme.pk\/custom\/videoupload\/TDK_Trailer.mp4","thumbnail":"https:\/\/bookme.pk\/custom\/upload\/TDK_IMDB.jpg","date":"2019-03-01 18:57:37","booking_type":"0","points":"0","update_date":"2018-11-22 10:26:47","close_date":null,"status":"1","shows":[{"city_id":"1","city_name":"Lahore","show_id":"480211","show_movie_id":"855","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"152","hall_name":"Screen 7 Regular","show_date":"2019-03-08","show_start_time":"2019-03-08 14:45:00","show_time":"14:45","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480213","show_movie_id":"855","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"152","hall_name":"Screen 7 Premium B","show_date":"2019-03-08","show_start_time":"2019-03-08 14:45:00","show_time":"14:45","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480212","show_movie_id":"855","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"152","hall_name":"Screen 7 Premium A","show_date":"2019-03-08","show_start_time":"2019-03-08 14:45:00","show_time":"14:45","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480216","show_movie_id":"855","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"152","hall_name":"Screen 7 Premium B","show_date":"2019-03-09","show_start_time":"2019-03-09 14:45:00","show_time":"14:45","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480215","show_movie_id":"855","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"152","hall_name":"Screen 7 Premium A","show_date":"2019-03-09","show_start_time":"2019-03-09 14:45:00","show_time":"14:45","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480214","show_movie_id":"855","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"152","hall_name":"Screen 7 Regular","show_date":"2019-03-09","show_start_time":"2019-03-09 14:45:00","show_time":"14:45","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480219","show_movie_id":"855","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"152","hall_name":"Screen 7 Premium B","show_date":"2019-03-10","show_start_time":"2019-03-10 14:45:00","show_time":"14:45","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480218","show_movie_id":"855","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"152","hall_name":"Screen 7 Premium A","show_date":"2019-03-10","show_start_time":"2019-03-10 14:45:00","show_time":"14:45","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480217","show_movie_id":"855","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"152","hall_name":"Screen 7 Regular","show_date":"2019-03-10","show_start_time":"2019-03-10 14:45:00","show_time":"14:45","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"}]}]`)
	} else if movieid == "835" {
		response = []byte(`[{"movie_id":"835","imdb_id":"tt8032912","title":"PARWAAZ HAI JUNOON","genre":"Action, Romance, War","language":"Urdu","director":"Haseeb Hasan","producer":"Momina Duraid","release_date":"2018-08-22","music_director":"-","country":"Pakistan","cast":"Hamza Ali Abbasi, Hania Aamir, Ahad Raza Mir, Shaz Khan, Kubra Khan and others","synopsis":"Parwaaz Hai Junoon, a new movie by Momina and Duraid Films in collaboration with the Pakistan Air Force to be directed by Haseeb Hasan, written by Farhat Ishtiaq and starring Hamza Ali Abbasi, Ahad Raza Mir, Hania Amir, and Kubra Khan.","details":"","ranking":"8.5","length":"130","trailer_link":"https:\/\/bookme.pk\/custom\/videoupload\/PHJ_Trailer.mp4","thumbnail":"https:\/\/bookme.pk\/custom\/upload\/parwaz_IMDB.jpg","date":"2019-03-01 18:56:22","booking_type":"0","points":"0","update_date":null,"close_date":null,"status":"1","shows":[{"city_id":"1","city_name":"Lahore","show_id":"479604","show_movie_id":"835","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"152","hall_name":"Screen 7 Premium B","show_date":"2019-03-07","show_start_time":"2019-03-07 22:45:00","show_time":"22:45","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"40","city_name":"Quetta","show_id":"480543","show_movie_id":"835","show_cenima_id":"10011","cinema_name":"Weplex","hall_id":"205","hall_name":"Screen 1 ","show_date":"2019-03-09","show_start_time":"2019-03-09 19:30:00","show_time":"19:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"40","city_name":"Quetta","show_id":"480544","show_movie_id":"835","show_cenima_id":"10011","cinema_name":"Weplex","hall_id":"205","hall_name":"Screen 1 ","show_date":"2019-03-10","show_start_time":"2019-03-10 19:30:00","show_time":"19:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"40","city_name":"Quetta","show_id":"480545","show_movie_id":"835","show_cenima_id":"10011","cinema_name":"Weplex","hall_id":"205","hall_name":"Screen 1 ","show_date":"2019-03-11","show_start_time":"2019-03-11 19:30:00","show_time":"19:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"40","city_name":"Quetta","show_id":"480546","show_movie_id":"835","show_cenima_id":"10011","cinema_name":"Weplex","hall_id":"205","hall_name":"Screen 1 ","show_date":"2019-03-12","show_start_time":"2019-03-12 19:30:00","show_time":"19:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"40","city_name":"Quetta","show_id":"480547","show_movie_id":"835","show_cenima_id":"10011","cinema_name":"Weplex","hall_id":"205","hall_name":"Screen 1 ","show_date":"2019-03-13","show_start_time":"2019-03-13 19:30:00","show_time":"19:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"40","city_name":"Quetta","show_id":"480548","show_movie_id":"835","show_cenima_id":"10011","cinema_name":"Weplex","hall_id":"205","hall_name":"Screen 1 ","show_date":"2019-03-14","show_start_time":"2019-03-14 19:30:00","show_time":"19:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"1"}]}]`)
	} else if movieid == "829" {
		response = []byte(`[{"movie_id":"829","imdb_id":"tt7816386","title":"JAWANI PHIR NAHI ANI 2","genre":"Comedy","language":"Urdu","director":"Nadeem Beyg","producer":"Salman Iqbal,Humayun Saeed,Shahzad NasibJarjees Seja","release_date":"2018-08-21","music_director":"Shani ArshadSahir Ali BaggaAhmad Ali Butt","country":"Pakistan","cast":"Mawra Hocane,Kanwaljit Singh,Humayun Saeed","synopsis":"Remake of JPNA.","details":"","ranking":"8.0","length":"165","trailer_link":"https:\/\/bookme.pk\/custom\/videoupload\/JPNA2_Trailer.mp4","thumbnail":"https:\/\/bookme.pk\/custom\/upload\/JPNA2_IMDB.jpg","date":"2019-03-01 18:45:03","booking_type":"0","points":"0","update_date":"2018-08-17 23:24:21","close_date":null,"status":"1","shows":[{"city_id":"1","city_name":"Lahore","show_id":"479609","show_movie_id":"829","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"151","hall_name":"Screen 6 Regular","show_date":"2019-03-07","show_start_time":"2019-03-07 23:00:00","show_time":"23:00","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480144","show_movie_id":"829","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"146","hall_name":"Screen 1 Ultra Regular StudentNN","show_date":"2019-03-10","show_start_time":"2019-03-10 23:15:00","show_time":"23:15","ticket_price":"299","handling_charges":14.95,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"477572","show_movie_id":"829","show_cenima_id":"10000","cinema_name":"Cinestar ATC","hall_id":"160","hall_name":"Screen 3 Regular Ali Trade Onlline","show_date":"2019-03-07","show_start_time":"2019-03-07 22:00:00","show_time":"22:00","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"3","city_name":"Islamabad","show_id":"477760","show_movie_id":"829","show_cenima_id":"20","cinema_name":"Centaurus Cineplex","hall_id":"71","hall_name":"Cinema 4 2D ","show_date":"2019-03-07","show_start_time":"2019-03-07 22:30:00","show_time":"22:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480516","show_movie_id":"829","show_cenima_id":"30","cinema_name":"Imperial Cinema","hall_id":"111","hall_name":"IMP 1 LUXURY TICKET","show_date":"2019-03-08","show_start_time":"2019-03-08 21:00:00","show_time":"21:00","ticket_price":"1500","handling_charges":75,"easypaisa_charges":0,"house_full":"0","eticket":"0"},{"city_id":"1","city_name":"Lahore","show_id":"480515","show_movie_id":"829","show_cenima_id":"30","cinema_name":"Imperial Cinema","hall_id":"111","hall_name":"IMP 1 PRIME 600","show_date":"2019-03-08","show_start_time":"2019-03-08 21:00:00","show_time":"21:00","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"0"},{"city_id":"1","city_name":"Lahore","show_id":"480517","show_movie_id":"829","show_cenima_id":"30","cinema_name":"Imperial Cinema","hall_id":"111","hall_name":"IMP 1 PRIME 600","show_date":"2019-03-09","show_start_time":"2019-03-09 21:00:00","show_time":"21:00","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"0"},{"city_id":"1","city_name":"Lahore","show_id":"480518","show_movie_id":"829","show_cenima_id":"30","cinema_name":"Imperial Cinema","hall_id":"111","hall_name":"IMP 1 LUXURY TICKET","show_date":"2019-03-09","show_start_time":"2019-03-09 21:00:00","show_time":"21:00","ticket_price":"1500","handling_charges":75,"easypaisa_charges":0,"house_full":"0","eticket":"0"},{"city_id":"1","city_name":"Lahore","show_id":"480520","show_movie_id":"829","show_cenima_id":"30","cinema_name":"Imperial Cinema","hall_id":"111","hall_name":"IMP 1 LUXURY TICKET","show_date":"2019-03-10","show_start_time":"2019-03-10 21:00:00","show_time":"21:00","ticket_price":"1500","handling_charges":75,"easypaisa_charges":0,"house_full":"0","eticket":"0"},{"city_id":"1","city_name":"Lahore","show_id":"480519","show_movie_id":"829","show_cenima_id":"30","cinema_name":"Imperial Cinema","hall_id":"111","hall_name":"IMP 1 PRIME 600","show_date":"2019-03-10","show_start_time":"2019-03-10 21:00:00","show_time":"21:00","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"0"},{"city_id":"1","city_name":"Lahore","show_id":"479433","show_movie_id":"829","show_cenima_id":"25","cinema_name":"Super Cinema Vogue Tower","hall_id":"37","hall_name":"SC SUPER 3 Recliner","show_date":"2019-03-07","show_start_time":"2019-03-07 23:00:00","show_time":"23:00","ticket_price":"1000","handling_charges":50,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"479432","show_movie_id":"829","show_cenima_id":"25","cinema_name":"Super Cinema Vogue Tower","hall_id":"37","hall_name":"SC SUPER 3 Regular","show_date":"2019-03-07","show_start_time":"2019-03-07 23:00:00","show_time":"23:00","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"479814","show_movie_id":"829","show_cenima_id":"21","cinema_name":"Sozo World Cinema","hall_id":"78","hall_name":"Sozo World ","show_date":"2019-03-07","show_start_time":"2019-03-07 22:30:00","show_time":"22:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"479813","show_movie_id":"829","show_cenima_id":"21","cinema_name":"Sozo World Cinema","hall_id":"78","hall_name":"Sozo World ","show_date":"2019-03-07","show_start_time":"2019-03-07 22:30:00","show_time":"22:30","ticket_price":"500","handling_charges":25,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"40","city_name":"Quetta","show_id":"477910","show_movie_id":"829","show_cenima_id":"10011","cinema_name":"Weplex","hall_id":"205","hall_name":"Screen 1 ","show_date":"2019-03-07","show_start_time":"2019-03-07 22:15:00","show_time":"22:15","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"}]}]`)
	} else if movieid == "824" {
		response = []byte(`[{"movie_id":"824","imdb_id":"tt4912910","title":"Mission Impossible Fallout","genre":"Action, Thriller, Adventure","language":"English","director":"Christopher McQuarrie","producer":"Tom Cruise","release_date":"2018-07-27","music_director":"Lorne Balfe","country":"USA","cast":"Tom Cruise,Rebecca Ferguson,Henry Cavill","synopsis":"Ethan Hunt and his IMF team, along with some familiar allies, race against time after a mission gone wrong.","details":"","ranking":"7.8","length":"227","trailer_link":"https:\/\/bookme.pk\/custom\/videoupload\/MIF_Trailer.mp4","thumbnail":"https:\/\/bookme.pk\/custom\/upload\/mif_IMDB.jpg","date":"2019-03-01 18:58:29","booking_type":"0","points":"0","update_date":"2018-07-23 20:28:37","close_date":null,"status":"1","shows":[{"city_id":"1","city_name":"Lahore","show_id":"479598","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"146","hall_name":"Screen 1 Gold","show_date":"2019-03-07","show_start_time":"2019-03-07 22:30:00","show_time":"22:30","ticket_price":"1100","handling_charges":55,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480070","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"149","hall_name":"Screen 4 Premium B","show_date":"2019-03-09","show_start_time":"2019-03-09 11:30:00","show_time":"11:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480073","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"149","hall_name":"Screen 4 Premium B","show_date":"2019-03-09","show_start_time":"2019-03-09 17:30:00","show_time":"17:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480072","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"149","hall_name":"Screen 4 Premium A","show_date":"2019-03-09","show_start_time":"2019-03-09 17:30:00","show_time":"17:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480071","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"149","hall_name":"Screen 4 Regular","show_date":"2019-03-09","show_start_time":"2019-03-09 17:30:00","show_time":"17:30","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480067","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"153","hall_name":"Screen 8 Gold","show_date":"2019-03-09","show_start_time":"2019-03-09 20:30:00","show_time":"20:30","ticket_price":"1100","handling_charges":55,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480075","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"149","hall_name":"Screen 4 Regular","show_date":"2019-03-10","show_start_time":"2019-03-10 11:30:00","show_time":"11:30","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480077","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"149","hall_name":"Screen 4 Premium B","show_date":"2019-03-10","show_start_time":"2019-03-10 11:30:00","show_time":"11:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480076","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"149","hall_name":"Screen 4 Premium A","show_date":"2019-03-10","show_start_time":"2019-03-10 11:30:00","show_time":"11:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480080","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"149","hall_name":"Screen 4 Premium B","show_date":"2019-03-10","show_start_time":"2019-03-10 17:30:00","show_time":"17:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480079","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"149","hall_name":"Screen 4 Premium A","show_date":"2019-03-10","show_start_time":"2019-03-10 17:30:00","show_time":"17:30","ticket_price":"700","handling_charges":35,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480078","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"149","hall_name":"Screen 4 Regular","show_date":"2019-03-10","show_start_time":"2019-03-10 17:30:00","show_time":"17:30","ticket_price":"600","handling_charges":30,"easypaisa_charges":0,"house_full":"0","eticket":"1"},{"city_id":"1","city_name":"Lahore","show_id":"480074","show_movie_id":"824","show_cenima_id":"8888","cinema_name":"Universal Cinema lahore Emporium Mall","hall_id":"153","hall_name":"Screen 8 Gold","show_date":"2019-03-10","show_start_time":"2019-03-10 20:30:00","show_time":"20:30","ticket_price":"1100","handling_charges":55,"easypaisa_charges":0,"house_full":"0","eticket":"1"}]}]`)
	} else {
		response = []byte("None")
	}

	return response
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
	response := GetMovieDetails(movieid)

	if string(response) != "None" {

		// Return json response
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
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

	apikey := r.FormValue("api_key")
	if !IsValidAPIKey(w, r, apikey) {
		return
	}

	movieid := r.FormValue("movie_id")
	showid := r.FormValue("movie_id")
	cinemaid := r.FormValue("cinema_id")
	date := r.FormValue("date")

	if !AreValidMovieDetails(w, r, movieid, showid, cinemaid, date) {
		return
	}

	// if string(response) != "None" {

	// 	// Return json response
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(response)
	// 	return
	// }

	// Send an empty response as sent by Bookme API (if movie id isn't found)
	w.WriteHeader(200)
	w.Write([]byte("[[]]"))
	fmt.Println(r.RequestURI, "- Error, Invalid Movie ID!")
	return
}
