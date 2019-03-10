package main

// var playMoviesList []map[string]interface{}

var playMoviesList = []map[string]interface{}{
	{
		"movie_id":     "934",
		"imdb_id":      "tt4154664",
		"title":        "Captain Marvel",
		"genre":        "Action, Adventure",
		"language":     "English",
		"director":     "Anna Boden, Ryan Fleck",
		"producer":     "Victoria Alonso",
		"release_date": "2019-03-08",
		"cast":         "Brie Larson",
		"ranking":      "6.0",
		"length":       "124",
		"thumbnail":    "https://bookme.pk/custom/upload/marvil_IMDB.jpg"},
	{
		"movie_id":     "901",
		"imdb_id":      "",
		"title":        "3 BAHADUR",
		"genre":        "Adventure,  Animation, Family",
		"language":     "Urdu",
		"director":     "Sharmeen Obaid Chinoy",
		"producer":     "Waadi Animations",
		"release_date": "2018-12-14",
		"cast":         " Mehwish HayatFahad MustafaSarwat GillaniNimra BuchaBehroze Sabzwari",
		"ranking":      "0.9",
		"length":       "120",
		"thumbnail":    "https://bookme.pk/custom/upload/3_Bahadur_RiseWarriors_IMDB.jpeg"}}

// playMoviesList = append(playMoviesList, map[string]string{
// 	"movie_id":     "934",
// 	"imdb_id":      "tt4154664",
// 	"title":        "Captain Marvel",
// 	"genre":        "Action, Adventure",
// 	"language":     "English",
// 	"director":     "Anna Boden, Ryan Fleck",
// 	"producer":     "Victoria Alonso",
// 	"release_date": "2019-03-08",
// 	"cast":         "Brie Larson",
// 	"ranking":      "6.0",
// 	"length":       "124",
// 	"thumbnail":    "https://bookme.pk/custom/upload/marvil_IMDB.jpg"})

// playMoviesList = append(playMoviesList, map[string]string{
// 	"movie_id":     "901",
// 	"imdb_id":      "",
// 	"title":        "3 BAHADUR",
// 	"genre":        "Adventure,  Animation, Family",
// 	"language":     "Urdu",
// 	"director":     "Sharmeen Obaid Chinoy",
// 	"producer":     "Waadi Animations",
// 	"release_date": "2018-12-14",
// 	"cast":         " Mehwish HayatFahad MustafaSarwat GillaniNimra BuchaBehroze Sabzwari",
// 	"ranking":      "0.9",
// 	"length":       "120",
// 	"thumbnail":    "https://bookme.pk/custom/upload/3_Bahadur_RiseWarriors_IMDB.jpeg"})

// playMoviesList = append(playMoviesList, map[string]string{
// 	"movie_id":     "898",
// 	"imdb_id":      "tt1477834",
// 	"title":        "AQUAMAN",
// 	"genre":        "Action, Adventure, Fantasy",
// 	"language":     "English",
// 	"director":     "James Wan",
// 	"producer":     "Peter Safran",
// 	"release_date": "2018-12-21",
// 	"cast":         "Jason Momoa, Amber Heard, Nicole Kidman",
// 	"ranking":      "7.5",
// 	"length":       "143",
// 	"thumbnail":    "https://bookme.pk/custom/upload/Aquaman_IMDB.jpg"})

// playMoviesList = append(playMoviesList, map[string]string{
// 	"movie_id":     "881",
// 	"imdb_id":      "tt1727824",
// 	"title":        "Bohemian Rhapsody",
// 	"genre":        "Drama, Biography  , Music",
// 	"language":     "English",
// 	"director":     "Bryan Singer",
// 	"producer":     "",
// 	"release_date": "2018-11-02",
// 	"cast":         "Rami Malek, Lucy Boynton, Gwilym Lee",
// 	"ranking":      "8.1",
// 	"length":       "134",
// 	"thumbnail":    "https://bookme.pk/custom/upload/Bohemian IMDB.jpg"})

// playMoviesList = append(playMoviesList, map[string]string{
// 	"movie_id":     "855",
// 	"imdb_id":      "",
// 	"title":        "THE DONKEY KING",
// 	"genre":        " Animation",
// 	"language":     "Urdu",
// 	"director":     "Aziz Jindani",
// 	"producer":     "Talisman Studios",
// 	"release_date": "2018-10-13",
// 	"cast":         "Jan Rambo, Ismail Tara, Hina Dilpazeer, Ghulam Mohiuddin, Jawed Sheikh",
// 	"ranking":      "7.2",
// 	"length":       "120",
// 	"thumbnail":    "https://bookme.pk/custom/upload/TDK_IMDB.jpg"})

// playMoviesList = append(playMoviesList, map[string]string{
// 	"movie_id":     "835",
// 	"imdb_id":      "tt8032912",
// 	"title":        "PARWAAZ HAI JUNOON",
// 	"genre":        "Action, Romance, War",
// 	"language":     "Urdu",
// 	"director":     "Haseeb Hasan",
// 	"producer":     "Momina Duraid",
// 	"release_date": "2018-08-22",
// 	"cast":         "Hamza Ali Abbasi, Hania Aamir, Ahad Raza Mir, Shaz Khan, Kubra Khan and others",
// 	"ranking":      "8.5",
// 	"length":       "130",
// 	"thumbnail":    "https://bookme.pk/custom/upload/parwaz_IMDB.jpg"})

// playMoviesList = append(playMoviesList, map[string]string{
// 	"movie_id":     "829",
// 	"imdb_id":      "tt7816386",
// 	"title":        "JAWANI PHIR NAHI ANI 2",
// 	"genre":        "Comedy",
// 	"language":     "Urdu",
// 	"director":     "Nadeem Beyg",
// 	"producer":     "Salman Iqbal,Humayun Saeed,Shahzad NasibJarjees Seja",
// 	"release_date": "2018-08-21",
// 	"cast":         "Mawra Hocane,Kanwaljit Singh,Humayun Saeed",
// 	"ranking":      "8.0",
// 	"length":       "165",
// 	"thumbnail":    "https://bookme.pk/custom/upload/JPNA2_IMDB.jpg"})

// playMoviesList = append(playMoviesList, map[string]string{
// 	"movie_id":     "824",
// 	"imdb_id":      "tt4912910",
// 	"title":        "Mission Impossible Fallout",
// 	"genre":        "Action, Thriller, Adventure",
// 	"language":     "English",
// 	"director":     "Christopher McQuarrie",
// 	"producer":     "Tom Cruise",
// 	"release_date": "2018-07-27",
// 	"cast":         "Tom Cruise,Rebecca Ferguson,Henry Cavill",
// 	"ranking":      "7.8",
// 	"length":       "227",
// 	"thumbnail":    "https://bookme.pk/custom/upload/mif_IMDB.jpg"})
