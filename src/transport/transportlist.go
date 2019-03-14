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
