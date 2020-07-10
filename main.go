package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		type requestJSON struct {
			Int1 int `json:"int1"`
			Int2 int `json:"int2"`
		}
		// accept JSON
		request := new(requestJSON)
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			w.WriteHeader(400)
			return
		}

		// return JSON back
		result := request.Int1 + request.Int2
		returnMap := map[string]int{
			"result": result,
		}
		if err := json.NewEncoder(w).Encode(&returnMap); err != nil {
			w.WriteHeader(500)
			return
		}
	})

	log.Printf("startin server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
