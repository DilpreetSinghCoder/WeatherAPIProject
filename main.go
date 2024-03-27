package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const Dport = ":8012"

// WeatherResponse represents the structure of weather data to be returned
type WeatherResponse struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"`
}

func main() {
	http.HandleFunc("/city", handleCityRequest)
	fmt.Printf("Server is starting on port: %v\n", Dport)
	http.ListenAndServe(Dport, nil)
}

func handleCityRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getWeather(w, r)
	case "POST":
		postWeather(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	cityName := r.URL.Query().Get("name")
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

// OpenWeatherAPI function
func weatherAPI(cityName string) (*WeatherResponse, error) {
	//API key
	weatherapiKey := "4d3f146cc59bcfa91195b7d935b57d87"
	if weatherapiKey == "" {
		return nil, fmt.Errorf("OpenWeatherMap API key not provided")
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", cityName, weatherapiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error while fetching weather data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error while fetching weather data: %s", resp.Status)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode weather data: %v", err)
	}

	weather := &WeatherResponse{
		City:        cityName,
		Temperature: fmt.Sprintf("%.1f°C", data["main"].(map[string]interface{})["temp"].(float64)),
		Weather:     data["weather"].([]interface{})[0].(map[string]interface{})["main"].(string),
	}

	return weather, nil
}
