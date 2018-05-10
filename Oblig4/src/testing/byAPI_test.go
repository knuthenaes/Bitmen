package main

import (
	"net/http"
	"encoding/json"
	"log"
	"testing"
)

// Returns an API response
func OsloResponse(url string) Index {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?id=6453366&APPID=b50b6d19b643bd8438490d345969a894")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var data Index
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data

}

func TestGoodResponseses(t *testing.T) {
	r := OsloResponse("http://api.openweathermap.org/data/2.5/weather?id=6453366&APPID=b50b6d19b643bd8438490d345969a894")
	if r.ID != 6453366 {
		// tests if city ID is correct, for correct weather output.
		t.Errorf("OpenWeatherMap Oslo returnerte id kode %d, expected 6453366", r.ID)
	}
}

func TestBadResponseses(t *testing.T) {
	r := OsloResponse("http://api.openweathermap.org/data/2.5/weather?id=6453366&APPID=b50b6d19b643bd8438490d345969a894")
	if r.ID == 6453366 {
		// ensures that the correct id is used for the actual city.
		t.Errorf("OpenWeatherMap Oslo returnerte 6453366, det er riktig ID.")
	}
}

// Returns an API response
func TrondheimResponse(url string) Index2 {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?id=3133881&APPID=b50b6d19b643bd8438490d345969a894")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var data Index2
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data

}

func TestGoodResponses(t *testing.T) {
	r := TrondheimResponse("http://api.openweathermap.org/data/2.5/weather?id=3133881&APPID=b50b6d19b643bd8438490d345969a894")
	if r.ID != 3133881 {
		// tests if city ID is correct, for correct weather output.
		t.Errorf("OpenWeatherMap Trondheim returnerte id kode %d, expected 3133881", r.ID)
	}
}

func TestBadResponses(t *testing.T) {
	r := TrondheimResponse("http://api.openweathermap.org/data/2.5/weather?id=3133881&APPID=b50b6d19b643bd8438490d345969a894")
	if r.ID == 3133881 {
		// ensures that the correct id is used for the actual city.
		t.Errorf("OpenWeatherMap Trondheim returnerte 3133881, det er riktig ID.")
	}
}

// Returns an API response
func StavangerResponse(url string) Index3 {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?id=3137115&APPID=b50b6d19b643bd8438490d345969a894")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var data Index3
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data

}

func TestGoodResponse(t *testing.T) {
	r := StavangerResponse("http://api.openweathermap.org/data/2.5/weather?id=3137115&APPID=b50b6d19b643bd8438490d345969a894")
	if r.ID != 3137115 {
		// tests if city ID is correct, for correct weather output.
		t.Errorf("OpenWeatherMap Stavanger returnerte id kode %d, expected 3137115", r.ID)
	}
}

func TestBadResponse(t *testing.T) {
	r := StavangerResponse("http://api.openweathermap.org/data/2.5/weather?id=3137115&APPID=b50b6d19b643bd8438490d345969a894")
	if r.ID == 3137115 {
		// ensures that the correct id is used for the actual city.
		t.Errorf("OpenWeatherMap Stavanger returnerte 3137115, det er riktig ID.")
	}
}
