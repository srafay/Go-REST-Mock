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
