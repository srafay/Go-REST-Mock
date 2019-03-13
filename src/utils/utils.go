package utils

import (
	"math/rand"
	"net/http"
	"time"
)

// WriteJSONResponse - helper function for writing JSON response to Response Writer
func WriteJSONResponse(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
	return
}

// StringInSlice - helper function which provides 'in' functionality as in python
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}
	return false
}

// GetRandomSeq - helper function for generating random sequence
func GetRandomSeq(n int) string {

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcde0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
