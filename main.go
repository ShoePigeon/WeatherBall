package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type GeoResult struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

type ForecastEntry struct {
	Dt   int64 `json:"dt"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	DtTxt string `json:"dt_txt"`
}

type ForecastResponse struct {
	List []ForecastEntry `json:"list"`
}

type CoolHoursPerDay map[string][]string // date string -> list of cool hour strings

type RequestPayload struct {
	Location string `json:"location"`
}

type ResponsePayload struct {
	Location  string          `json:"location"`
	CoolTimes CoolHoursPerDay `json:"cool_times"`
}

func main() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	fmt.Println("API KEY:", apiKey)
	if apiKey == "" {
		log.Fatal("WEATHER_API_KEY env var missing")
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/api/cooltimes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}

		var req RequestPayload
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || strings.TrimSpace(req.Location) == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Step 1: Geocode location
		geoURL := "https://nominatim.openstreetmap.org/search?format=json&q=" + strings.ReplaceAll(req.Location, " ", "+")
		respGeo, err := http.Get(geoURL)
		if err != nil || respGeo.StatusCode != 200 {
			fmt.Println("Geocoding failed:", err, "Status:", respGeo.StatusCode)
			http.Error(w, "Geocoding failed", 500)
			return
		}
		defer respGeo.Body.Close()

		var geoResults []GeoResult
		bodyGeo, _ := io.ReadAll(respGeo.Body)
		err = json.Unmarshal(bodyGeo, &geoResults)
		if err != nil || len(geoResults) == 0 {
			http.Error(w, "Location not found", 404)
			return
		}
		lat, lon := geoResults[0].Lat, geoResults[0].Lon
		locationName := geoResults[0].DisplayName
		fmt.Println("Geocoded location:", req.Location, "Lat:", lat, "Lon:", lon)

		// Step 2: Fetch weather data (One Call API)
		weatherURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%s&lon=%s&appid=%s", lat, lon, apiKey)
		fmt.Printf("Fetching weather data from: %s\n", weatherURL)
		respWeather, err := http.Get(weatherURL)
		if err != nil || respWeather.StatusCode != 200 {
			fmt.Println("Weather API error:", err, "Status:", respWeather.StatusCode)
			http.Error(w, "Weather API error", 500)
			return
		}
		defer respWeather.Body.Close()

		var forecast ForecastResponse
		err = json.NewDecoder(respWeather.Body).Decode(&forecast)
		if err != nil {
			fmt.Println("Failed to parse forecast data:", err)
			http.Error(w, "Failed to parse forecast data", 500)
			return
		}

		const coolThreshold = 294.26 // 70°F in Kelvin
		coolTimes := make(CoolHoursPerDay)
		for _, entry := range forecast.List {
			if entry.Main.Temp <= coolThreshold {
				t, _ := time.Parse("2006-01-02 15:04:05", entry.DtTxt)
				dateStr := t.Format("2006-01-02")
				hourStr := t.Format("15:04")
				coolTimes[dateStr] = append(coolTimes[dateStr], hourStr)
			}
		}

		// Step 4: Respond with JSON
		response := ResponsePayload{
			Location:  locationName,
			CoolTimes: coolTimes,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
