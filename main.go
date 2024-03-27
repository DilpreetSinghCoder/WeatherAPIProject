func main() {
	http.HandleFunc("/city", handleCityRequest)
	fmt.Printf("Server is starting on port: %v\n", Dport)
	http.ListenAndServe(Dport, nil)
}

func postWeather(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	cityName := requestData.Name
	if cityName == "" {
		http.Error(w, "City name parameter is missing", http.StatusBadRequest)
		return
	}

	weather, err := weatherAPI(cityName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(weather)
}

