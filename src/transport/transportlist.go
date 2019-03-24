package transport

var transportList = []map[string]string{
	{
		"service_id":     "10",
		"service_name":   "Bilal Travels",
		"address":        "Band Road near Yateem Khana Lahore.",
		"phone":          "042-111-287-444",
		"c_phone":        "03422874444",
		"status":         "1",
		"cod":            "0",
		"flexifare":      "1",
		"thumbnail":      "https://bookme.pk/custom/upload/transport/Bilal-Travels.jpg",
		"background":     "",
		"background_img": "",
		"facilities":     "",
		"careem":         "no"},
	{
		"service_id":     "19",
		"service_name":   "QConnect",
		"address":        "Fortress Lahore",
		"phone":          "000",
		"c_phone":        "",
		"status":         "1",
		"cod":            "0",
		"flexifare":      "0",
		"thumbnail":      "https://bookme.pk/custom/upload/transport/qconnectlogo.png",
		"background":     "",
		"background_img": "",
		"facilities":     "",
		"careem":         "no"}}

var departureCities = []map[string]interface{}{
	{
		"service_id": "10",
		"departure_cities": []map[string]string{
			{
				"origin_city_id":   "1",
				"origin_city_name": "Lahore"},
			{
				"origin_city_id":   "6",
				"origin_city_name": "Karachi"}}},
	{
		"service_id": "19",
		"departure_cities": []map[string]string{
			{
				"origin_city_id":   "2",
				"origin_city_name": "Islamabad"},
			{
				"origin_city_id":   "1",
				"origin_city_name": "Lahore"}}}}

var destinationCities = []map[string]interface{}{
	{
		"service_id":     "10",
		"origin_city_id": "1",
		"destination_cities": []map[string]string{
			{
				"destination_city_id":   "38",
				"destination_city_name": "Rawalpindi"},
			{
				"destination_city_id":   "6",
				"destination_city_name": "Karachi"}}},
	{
		"service_id":     "10",
		"origin_city_id": "6",
		"destination_cities": []map[string]string{
			{
				"destination_city_id":   "1",
				"destination_city_name": "Lahore"},
			{
				"destination_city_id":   "38",
				"destination_city_name": "Rawalpindi"}}},
	{
		"serive_id":      "19",
		"origin_city_id": "2",
		"destination_cities": []map[string]string{
			{
				"destination_city_id":   "1",
				"destination_city_name": "Lahore"}}},
	{
		"service_id":     "19",
		"origin_city_id": "1",
		"destination_cities": []map[string]string{
			{
				"destination_city_id":   "2",
				"destination_city_name": "Islamabad"}}}}

var timeIDs = []map[string]interface{}{
	{
		"date":                "2019-03-23",
		"departure_city_id":   1,
		"departure_city_name": "Lahore",
		"arrival_city_id":     6,
		"arrival_city_name":   "Karachi",
		"service_id":          10,
		"service_name":        "Bilal Travels",
		"time_id":             48160,
		"schedule_id":         17478,
		"route_id":            28,
		"time":                "23:00",
		"arrtime":             "03:15:00",
		"original_fare":       1200,
		"fare":                840,
		"handling_charges":    60,
		"easypaisa_charges":   0,
		"thumb":               "https://bookme.pk/custom/upload/transport/Bilal-Travels.jpg",
		"seats":               45,
		"busname":             "Luxury",
		"bustype":             "MAN-45",
		"btype_id":            2,
		"seat_info":           "",
		"facilities": []map[string]string{
			{
				"id":   "2",
				"name": "Deals",
				"img":  "https://bookme.pk/assets/img/transport_facility/deals.png",
			},
			{
				"id":   "5",
				"name": "Free WIFI",
				"img":  "https://bookme.pk/assets/img/transport_facility/free_wifi.png",
			},
			{
				"id":   "6",
				"name": "Headphones",
				"img":  "https://bookme.pk/assets/img/transport_facility/headphones.png",
			},
			{
				"id":   "7",
				"name": "Individual entertainment system",
				"img":  "https://bookme.pk/assets/img/transport_facility/individual_entertainment_system.png",
			},
			{
				"id":   "9",
				"name": "Mobile Charging",
				"img":  "https://bookme.pk/assets/img/transport_facility/mobile_charging.png",
			},
			{
				"id":   "10",
				"name": "More Leg room",
				"img":  "https://bookme.pk/assets/img/transport_facility/more_leg_room.png",
			},
			{
				"id":   "12",
				"name": "Regular seat",
				"img":  "https://bookme.pk/assets/img/transport_facility/regular_seat.png",
			},
		},
	},
	{
		"date":                "2019-03-23",
		"departure_city_id":   1,
		"departure_city_name": "Lahore",
		"arrival_city_id":     38,
		"arrival_city_name":   "Rawalpindi",
		"service_id":          10,
		"service_name":        "Bilal Travels",
		"time_id":             48161,
		"schedule_id":         17479,
		"route_id":            29,
		"time":                "16:00",
		"arrtime":             "23:15:00",
		"original_fare":       1300,
		"fare":                940,
		"handling_charges":    60,
		"easypaisa_charges":   0,
		"thumb":               "https://bookme.pk/custom/upload/transport/Bilal-Travels.jpg",
		"seats":               45,
		"busname":             "Luxury",
		"bustype":             "MAN-45",
		"btype_id":            2,
		"seat_info":           "",
		"facilities": []map[string]string{
			{
				"id":   "2",
				"name": "Deals",
				"img":  "https://bookme.pk/assets/img/transport_facility/deals.png",
			},
			{
				"id":   "5",
				"name": "Free WIFI",
				"img":  "https://bookme.pk/assets/img/transport_facility/free_wifi.png",
			},
			{
				"id":   "6",
				"name": "Headphones",
				"img":  "https://bookme.pk/assets/img/transport_facility/headphones.png",
			},
			{
				"id":   "7",
				"name": "Individual entertainment system",
				"img":  "https://bookme.pk/assets/img/transport_facility/individual_entertainment_system.png",
			},
			{
				"id":   "9",
				"name": "Mobile Charging",
				"img":  "https://bookme.pk/assets/img/transport_facility/mobile_charging.png",
			},
			{
				"id":   "10",
				"name": "More Leg room",
				"img":  "https://bookme.pk/assets/img/transport_facility/more_leg_room.png",
			},
			{
				"id":   "12",
				"name": "Regular seat",
				"img":  "https://bookme.pk/assets/img/transport_facility/regular_seat.png",
			},
		},
	},
}

var seatsInfo = []map[string]interface{}{
	{
		"service_id":            "10",
		"origin_city_id":        "1",
		"arrival_city_id":       "6",
		"date":                  "2019-03-23",
		"deptime":               "23:00",
		"time_id":               "48160",
		"schedule_id":           "17478",
		"route_id":              "28",
		"seats":                 "40",
		"total_seats":           "40",
		"total_available":       40,
		"available_seats":       []string{},
		"total_occupied":        0,
		"occupied_seats_male":   []string{},
		"occupied_seats_female": []string{},
		"total_reserved":        0,
		"reserved_seats_male":   []string{},
		"reserved_seats_female": []string{},
		"seat_plan": map[string]interface{}{
			"rows":  "4",
			"cols":  "10",
			"seats": "40",
			"seatplan": [][]interface{}{
				{
					map[string]string{
						"seat_id":   "1",
						"seat_name": "1",
					},
					map[string]string{
						"seat_id":   "5",
						"seat_name": "5",
					},
					map[string]string{
						"seat_id":   "9",
						"seat_name": "9",
					},
					map[string]string{
						"seat_id":   "13",
						"seat_name": "13",
					},
					map[string]string{
						"seat_id":   "17",
						"seat_name": "17",
					},
					map[string]string{
						"seat_id":   "21",
						"seat_name": "21",
					},
					map[string]string{
						"seat_id":   "25",
						"seat_name": "25",
					},
					map[string]string{
						"seat_id":   "29",
						"seat_name": "29",
					},
					map[string]string{
						"seat_id":   "33",
						"seat_name": "33",
					},
					map[string]string{
						"seat_id":   "37",
						"seat_name": "37",
					},
				},
				{
					map[string]string{
						"seat_id":   "2",
						"seat_name": "2",
					},
					map[string]string{
						"seat_id":   "6",
						"seat_name": "6",
					},
					map[string]string{
						"seat_id":   "10",
						"seat_name": "10",
					},
					map[string]string{
						"seat_id":   "14",
						"seat_name": "14",
					},
					map[string]string{
						"seat_id":   "18",
						"seat_name": "18",
					},
					map[string]string{
						"seat_id":   "22",
						"seat_name": "22",
					},
					map[string]string{
						"seat_id":   "26",
						"seat_name": "26",
					},
					map[string]string{
						"seat_id":   "30",
						"seat_name": "30",
					},
					map[string]string{
						"seat_id":   "34",
						"seat_name": "34",
					},
					map[string]string{
						"seat_id":   "38",
						"seat_name": "38",
					},
				},
				{
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
				},
				{
					map[string]string{
						"seat_id":   "0",
						"seat_name": "0",
					},
					map[string]string{
						"seat_id":   "7",
						"seat_name": "7",
					},
					map[string]string{
						"seat_id":   "11",
						"seat_name": "11",
					},
					map[string]string{
						"seat_id":   "15",
						"seat_name": "15",
					},
					map[string]string{
						"seat_id":   "19",
						"seat_name": "19",
					},
					map[string]string{
						"seat_id":   "23",
						"seat_name": "23",
					},
					map[string]string{
						"seat_id":   "27",
						"seat_name": "27",
					},
					map[string]string{
						"seat_id":   "31",
						"seat_name": "31",
					},
					map[string]string{
						"seat_id":   "35",
						"seat_name": "35",
					},
					map[string]string{
						"seat_id":   "39",
						"seat_name": "39",
					},
				},
				{
					map[string]string{
						"seat_id":   "4",
						"seat_name": "4",
					},
					map[string]string{
						"seat_id":   "8",
						"seat_name": "8",
					},
					map[string]string{
						"seat_id":   "12",
						"seat_name": "12",
					},
					map[string]string{
						"seat_id":   "16",
						"seat_name": "16",
					},
					map[string]string{
						"seat_id":   "20",
						"seat_name": "20",
					},
					map[string]string{
						"seat_id":   "24",
						"seat_name": "24",
					},
					map[string]string{
						"seat_id":   "28",
						"seat_name": "28",
					},
					map[string]string{
						"seat_id":   "32",
						"seat_name": "32",
					},
					map[string]string{
						"seat_id":   "36",
						"seat_name": "36",
					},
					map[string]string{
						"seat_id":   "40",
						"seat_name": "40",
					},
				},
			},
		},
	},
	{
		"service_id":            "10",
		"origin_city_id":        "1",
		"arrival_city_id":       "38",
		"date":                  "2019-03-23",
		"deptime":               "16:00",
		"time_id":               "48161",
		"schedule_id":           "17479",
		"route_id":              "29",
		"seats":                 "40",
		"total_seats":           "40",
		"total_available":       40,
		"available_seats":       []string{},
		"total_occupied":        0,
		"occupied_seats_male":   []string{},
		"occupied_seats_female": []string{},
		"total_reserved":        0,
		"reserved_seats_male":   []string{},
		"reserved_seats_female": []string{},
		"seat_plan": map[string]interface{}{
			"rows":  "4",
			"cols":  "10",
			"seats": "40",
			"seatplan": [][]interface{}{
				{
					map[string]string{
						"seat_id":   "1",
						"seat_name": "1",
					},
					map[string]string{
						"seat_id":   "5",
						"seat_name": "5",
					},
					map[string]string{
						"seat_id":   "9",
						"seat_name": "9",
					},
					map[string]string{
						"seat_id":   "13",
						"seat_name": "13",
					},
					map[string]string{
						"seat_id":   "17",
						"seat_name": "17",
					},
					map[string]string{
						"seat_id":   "21",
						"seat_name": "21",
					},
					map[string]string{
						"seat_id":   "25",
						"seat_name": "25",
					},
					map[string]string{
						"seat_id":   "29",
						"seat_name": "29",
					},
					map[string]string{
						"seat_id":   "33",
						"seat_name": "33",
					},
					map[string]string{
						"seat_id":   "37",
						"seat_name": "37",
					},
				},
				{
					map[string]string{
						"seat_id":   "2",
						"seat_name": "2",
					},
					map[string]string{
						"seat_id":   "6",
						"seat_name": "6",
					},
					map[string]string{
						"seat_id":   "10",
						"seat_name": "10",
					},
					map[string]string{
						"seat_id":   "14",
						"seat_name": "14",
					},
					map[string]string{
						"seat_id":   "18",
						"seat_name": "18",
					},
					map[string]string{
						"seat_id":   "22",
						"seat_name": "22",
					},
					map[string]string{
						"seat_id":   "26",
						"seat_name": "26",
					},
					map[string]string{
						"seat_id":   "30",
						"seat_name": "30",
					},
					map[string]string{
						"seat_id":   "34",
						"seat_name": "34",
					},
					map[string]string{
						"seat_id":   "38",
						"seat_name": "38",
					},
				},
				{
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
					map[string]string{
						"seat_id":   "0",
						"seat_name": "",
					},
				},
				{
					map[string]string{
						"seat_id":   "0",
						"seat_name": "0",
					},
					map[string]string{
						"seat_id":   "7",
						"seat_name": "7",
					},
					map[string]string{
						"seat_id":   "11",
						"seat_name": "11",
					},
					map[string]string{
						"seat_id":   "15",
						"seat_name": "15",
					},
					map[string]string{
						"seat_id":   "19",
						"seat_name": "19",
					},
					map[string]string{
						"seat_id":   "23",
						"seat_name": "23",
					},
					map[string]string{
						"seat_id":   "27",
						"seat_name": "27",
					},
					map[string]string{
						"seat_id":   "31",
						"seat_name": "31",
					},
					map[string]string{
						"seat_id":   "35",
						"seat_name": "35",
					},
					map[string]string{
						"seat_id":   "39",
						"seat_name": "39",
					},
				},
				{
					map[string]string{
						"seat_id":   "4",
						"seat_name": "4",
					},
					map[string]string{
						"seat_id":   "8",
						"seat_name": "8",
					},
					map[string]string{
						"seat_id":   "12",
						"seat_name": "12",
					},
					map[string]string{
						"seat_id":   "16",
						"seat_name": "16",
					},
					map[string]string{
						"seat_id":   "20",
						"seat_name": "20",
					},
					map[string]string{
						"seat_id":   "24",
						"seat_name": "24",
					},
					map[string]string{
						"seat_id":   "28",
						"seat_name": "28",
					},
					map[string]string{
						"seat_id":   "32",
						"seat_name": "32",
					},
					map[string]string{
						"seat_id":   "36",
						"seat_name": "36",
					},
					map[string]string{
						"seat_id":   "40",
						"seat_name": "40",
					},
				},
			},
		},
	},
}
