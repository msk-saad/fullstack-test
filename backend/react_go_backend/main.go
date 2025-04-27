// This is a simple Go backend server that servers a greeting message in JSON format.
// It also implements CORS to allow requests from a React frontend running on port 5173.

package main

import (
	"fmt"
	"net/http"
)

func greetings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message": "Greetings from Go Backend!"}`)
}

func main() {
	// http.HandleFunc("/api/greet", greetings) // Without CORS implementation
	http.HandleFunc("/api/greet", withCORS(greetings)) // Implementing CORS to allow requests from the React frontend.
	fmt.Println("Backend is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Allowing requests from the React frontend running on port 5173
		//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allowing GET, POST, and OPTIONS methods (might be needed for complex projects but not necessary for this simple example)
		//w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Allowing Content-Type header in requests (might be needed for complex projects but not necessary for this simple example)

		// Handling preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		h(w, r) // Calling the original handler function
		// Note: The above CORS implementation is basic and might need to be adjusted for more complex scenarios.
	}
}
