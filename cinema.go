package main

import (
	config "config"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetMovieDetails - function for getting movie details of a particular movie_id
func GetMovieDetails(movieid string) (map[string]interface{}, bool) {

	for k := range playMovieShowsList {
		if movieid == playMovieShowsList[k]["movie_id"] {
			return playMovieShowsList[k], true
		}
	}
	return nil, false
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

// PlayMovies - function for bookme /play_movies API
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

// PlayMovieShows - function for bookme /play_movie_shows API
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

// CheckIfSeatsAvailable - helper function to check if all seats are available before we can mark them booked
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

// ChangeSeatStatus - helper function which changes the status of the seat of a particular show
func ChangeSeatStatus(showID, seatID string, status int) {
	for i, showDict := range CinemaSeatPlanMock {
		if showID == showDict["show_id"] {
			// check if we are removing the seat from being reserved
			if status == 0 {
				for j, seat := range CinemaSeatPlanMock[i]["booked_seats"].([]string) {
					if seatID == seat {
						// Remove the seat at index j from booked_seats
						CinemaSeatPlanMock[i]["booked_seats"].([]string)[j] = CinemaSeatPlanMock[i]["booked_seats"].([]string)[len(CinemaSeatPlanMock[i]["booked_seats"].([]string))-1] // Copy last element to index j.
						CinemaSeatPlanMock[i]["booked_seats"].([]string)[len(CinemaSeatPlanMock[i]["booked_seats"].([]string))-1] = ""                                                  // Erase last element (write zero value).
						CinemaSeatPlanMock[i]["booked_seats"] = CinemaSeatPlanMock[i]["booked_seats"].([]string)[:len(CinemaSeatPlanMock[i]["booked_seats"].([]string))-1]              // Truncate slice.

						fmt.Println(config.INFO, "Removed seat_id:", seatID, "from the booked seats")
						fmt.Println(config.DEBUG, "Booked seats now are: ", CinemaSeatPlanMock[i]["booked_seats"])
					}
				}
			}

			for _, seatPlan := range showDict["seat_plan"].([]map[string]interface{}) {
				for _, seat := range seatPlan["seats"].([]map[string]interface{}) {
					if seat["seat_id"] == seatID {
						// CinemaSeatPlanMock[i]["seat_plan"].([]map[string]interface{})[k]["seats"].([]map[string]interface{})[l]["status"] = status
						seat["status"] = status
						fmt.Println(config.INFO, "changed seat status of 'seat_id':", seatID, "to", status)
						return
					}
				}
			}
		} else {
			fmt.Println(config.ERROR, "'show_id':", showID, "not found for changing seat status")
		}
	}
}

// ResetReservedSeat - helper function which is called after some time to reset a reserved seat
func ResetReservedSeat(seatID string) {

	for i, reservedSeat := range permanentlyReservedSeats {
		if seatID == reservedSeat.SeatID {

			// Change the status of the seat so that it is not reserved anymore
			ChangeSeatStatus(reservedSeat.ShowID, reservedSeat.SeatID, 0)

			// Remove the seat at index i from permanentlyReservedSeats
			permanentlyReservedSeats[i] = permanentlyReservedSeats[len(permanentlyReservedSeats)-1]                                           // Copy last element to index i.
			permanentlyReservedSeats[len(permanentlyReservedSeats)-1] = reservedSeats{ShowID: "-1", SeatID: "-1", PermanentlyReserved: false} // Erase last element (write zero value).
			permanentlyReservedSeats = permanentlyReservedSeats[:len(permanentlyReservedSeats)-1]                                             // Truncate slice.
			fmt.Println(config.INFO, "Removed the reserved seat '", seatID, "' from the list")
		}
	}
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

					// Make sure to remove the seat from the permanently reserved seats list (if user doesnt save it within a minute)
					_seatID := seatDict["seat_id"].(string)
					ResetTimer := time.AfterFunc(config.SeatResetTimer*time.Second, func() { ResetReservedSeat(_seatID) })

					// Temporarily add the seat in permanently reserved seats list (only for a minute)
					permanentlyReservedSeats = append(permanentlyReservedSeats, reservedSeats{ShowID: seatPlanMap["show_id"].(string), SeatID: seatDict["seat_id"].(string), PermanentlyReserved: false, SeatResetTimer: ResetTimer})

				} else {
					fmt.Println(config.WARNING, "seat number:", seatDict["seat_id"], "already reserved [WHY]")
				}
			}
		}
	}
	return true
}

// MarkSeatsBooked - function for marking seats status as booked
func MarkSeatsBooked(seatPlanIndex int, seatNumbers string) {

	splittedSeatNumbers := strings.Split(strings.Replace(seatNumbers, " ", "", -1), ",")

	// Change the status of the seats provided here to booked
	for _, seat := range splittedSeatNumbers {
		CinemaSeatPlanMock[seatPlanIndex]["booked_seats"] = append(CinemaSeatPlanMock[seatPlanIndex]["booked_seats"].([]string), seat)
	}

	fmt.Println(config.INFO, seatNumbers, "have been marked as reserved")
	fmt.Println(config.DEBUG, "Reserved seats are: ", CinemaSeatPlanMock[seatPlanIndex]["booked_seats"])
	return
}

// CinemaReserveSeats - function for bookme /cinema_reserve_seats API
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
				seatsReserved := ReserveCinemaSeats(k, seatNumbers)
				if !seatsReserved {
					WriteJSONResponse(w, `[{"status":"failed", "msg":"Seats are aready allocated."}]`)
				} else {
					WriteJSONResponse(w, fmt.Sprintf(`[{"status":"success", "msg":"Seats are allocated successfully.", "booking_no":"%d"}]`, config.BookingNumber))

					// increment the BookingNumber variable
					config.BookingNumber++

					// change status of reserved seats to booked
					MarkSeatsBooked(k, seatNumbers)
				}
				return
			}
		}
		WriteJSONResponse(w, fmt.Sprintf("No seatplan found for show_id: %s", showid))
		return
	}
}

// SaveCinemaResponse - helper function for formatting the response to be sent in /save_cinema API
func SaveCinemaResponse(cinema, screen, movie, seats, ticketPrice, name, phone, email, city, address, seatNumbers, Time string, handlingCharges int) (string, bool) {
	status := "success"
	msg := ""
	bookingID := fmt.Sprintf("AC%d", config.BookingNumber*100000)
	orderRefNumber := fmt.Sprintf("AC%d", config.BookingNumber*100000)
	order := fmt.Sprintf("mov%d", config.BookingNumber*100)
	bookmeBookingID := getRandomSeq(32)
	bookingReference := ""
	netAmount := 0
	discount := 0
	if _ticketPrice, err := strconv.ParseInt(ticketPrice, 10, 0); err == nil {
		netAmount = int(_ticketPrice)
	} else {
		fmt.Println(config.ERROR, "There was an error converting 'ticket_price' to integer")
		return "There was an error converting 'ticket_price' to integer", true
	}
	if _seats, err := strconv.ParseInt(seats, 10, 0); err == nil {
		netAmount *= int(_seats)
	} else {
		fmt.Println(config.ERROR, "There was an error converting 'seats' to integer")
		return "There was an error converting 'seats' to integer", true
	}

	totalAmount := netAmount + handlingCharges
	seatPreference := ""
	date := time.Now().Format("02th January 2006 03:04:05 PM")
	_time, _ := time.Parse("2006-01-02 15:04:05", Time)

	return fmt.Sprintf(`[
		{
			"status": "%s",
			"msg": "%s",
			"booking_id": "%s",
			"orderRefNumber": "%s",
			"order": "%s",
			"bookme_booking_id": "%s",
			"booking_reference": "%s",
			"cinema": "%s",
			"screen": "%s",
			"movie": "%s",
			"seats": "%s",
			"net_amount": "%d",
			"handling_charges": "%d",
			"Discount": "%d",
			"total_amount": "%d",
			"name": "%s",
			"phone": "%s",
			"email": "%s",
			"city": "%s",
			"address": "%s",
			"seat_numbers": "%s",
			"seat_preference": "%s",
			"date": "%s",
			"time": "%s"
		}
	]`,
		status, msg, bookingID, orderRefNumber, order, bookmeBookingID, bookingReference, cinema, screen, movie, seats, netAmount,
		handlingCharges, discount, totalAmount, name, phone, email, city, address, seatNumbers, seatPreference, date, _time.Format("Jan 02 2006 03:04 PM")), true

}

// MarkSeatPermanentlyBooked - helper function for marking seats permanently booked and it also stops the timer
func MarkSeatPermanentlyBooked(showID, seatID string) {
	for _, seat := range permanentlyReservedSeats {
		if showID == seat.ShowID {

			// mark seat permanently reserved
			seat.PermanentlyReserved = true

			// stop the reset timer
			seat.SeatResetTimer.Stop()
		}
	}
}

// SaveCinema - function for bookme /save_cinema API
func SaveCinema(w http.ResponseWriter, r *http.Request) {

	// bookingNumber := r.FormValue("booking_no")
	showID := r.FormValue("show_id")
	// bookingType := r.FormValue("booking_type")
	seats := r.FormValue("seats")
	seatNumbers := r.FormValue("seat_numbers")
	// amount := r.FormValue("amount")
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	city := r.FormValue("city")
	address := r.FormValue("address")
	// paymentType := r.FormValue("payment_type")

	response := ""
	success := false
	splittedSeatNumbers := strings.Split(strings.Replace(seatNumbers, " ", "", -1), ",")

	for _, movie := range playMovieShowsList {
		for _, show := range movie["shows"].([]map[string]interface{}) {
			if showID == show["show_id"] {
				response, success = SaveCinemaResponse(show["cinema_name"].(string), show["hall_name"].(string), movie["title"].(string), seats, show["ticket_price"].(string), name, phone, email, city, address, seatNumbers, show["show_start_time"].(string), show["handling_charges"].(int))
				for _, seat := range splittedSeatNumbers {
					MarkSeatPermanentlyBooked(showID, seat)
				}
			}
		}
	}

	if success {
		WriteJSONResponse(w, response)
	} else {
		WriteJSONResponse(w, `[{"status":"failed", "msg":"Invalid parameters"}]`)
	}
	return

}
