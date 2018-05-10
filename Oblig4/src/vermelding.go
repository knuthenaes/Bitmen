package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"html/template"
	"runtime"
	"os/exec"
)

type Index struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempA    string
	} `json:"main"`
	Wind       struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	ID int `json:"id"`
	Name string `json:"name"`
}
type Index2 struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempB    string
	} `json:"main"`
	Wind       struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	ID int `json:"id"`
	Name string `json:"name"`
}
type Index3 struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempC    string
	} `json:"main"`
	Wind       struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	ID int `json:"id"`
	Name string `json:"name"`
}

var index Index
var index2 Index2
var index3 Index3

func ProTipsOslo () {

	var utenKelvin = index.Main.Temp - 273.15
	switch {
	case utenKelvin >= 18 && index.Wind.Speed < 4:
		index.Main.TempA = "God norsk sommmer i dag, med lite vind. På tide å finne frem badetøy"
	case utenKelvin >= 18 && index.Wind.Speed >= 4:
		index.Main.TempA = "I dag blir det varmt, men mye vind. Noe vindtett kan være lurt å ha med seg"
	case utenKelvin <= 18 && utenKelvin > 10:
		index.Main.TempA = "Det kan bli kaldt utover kvelden, ta gjerne på en buff for å unngå forkjølelse!"
	case utenKelvin <= 10 && utenKelvin >= 4:
		index.Main.TempA = "Ikke spesielt varmt i dag. En god ytterjakke er lurt å ha på seg"
	case utenKelvin < 4:
		index.Main.TempA = "Hold deg inne, eller bruk ulla!"
	}
}
func ProTipsTrondheim () {

	var utenKelvin = index2.Main.Temp - 273.15
	switch {
	case utenKelvin >= 18 && index2.Wind.Speed < 4:
		index2.Main.TempB = "I dag er det norsk sommmer med lav vindstyrke. På tide å dra frem shortsen"
	case utenKelvin >= 18 && index2.Wind.Speed >= 4:
		index2.Main.TempB = "I dag blir det varmt med høy vindstyrke, pass på å kle dere"
	case utenKelvin <= 18 && utenKelvin > 10:
		index2.Main.TempB = "Det kan bli kaldt utover kvelden, ta gjerne på et skjerf for å unngå forkjølelse!"
	case utenKelvin <= 10 && utenKelvin >= 4:
		index2.Main.TempB = "Husk D-vitaminer da solen ikke er like sterk."
	case utenKelvin < 4:
		index2.Main.TempB = "Hold deg inne, eller bruk ulla!"
	}
}
func ProTipsStavanger () {

	var utenKelvin = index3.Main.Temp - 273.15

	switch {
	case utenKelvin >= 18 && index3.Wind.Speed < 4:
		index3.Main.TempC = "I dag er det norsk sommmer med lav vindstyrke. På tide å dra frem shortsen"
	case utenKelvin >= 18 && index3.Wind.Speed >= 4:
		index3.Main.TempC = "I dag blir det varmt med høy vindstyrke, pass på å kle dere"
	case utenKelvin <= 18 && utenKelvin > 10:
		index3.Main.TempC = "Det kan bli kaldt utover kvelden, ta gjerne på et skjerf for å unngå forkjølelse!"
	case utenKelvin <= 10 && utenKelvin >= 4:
		index3.Main.TempC = "Husk D-vitaminer da solen ikke er like sterk."
	case utenKelvin < 4:
		index3.Main.TempC = "Hold deg inne, eller bruk ulla!"
	}
}

func Oslo (w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?id=6453366&APPID=b50b6d19b643bd8438490d345969a894")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	jsonErr := json.Unmarshal(body, &index)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	ProTipsOslo()

	temp, err := template.ParseFiles("Oslo.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, index)
	if err != nil {
		log.Fatal(err)
	}

}
func Trondheim (w http.ResponseWriter, r *http.Request) {

	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?id=3133881&APPID=b50b6d19b643bd8438490d345969a894")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	jsonErr := json.Unmarshal(body, &index2)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	ProTipsTrondheim()

	temp, err := template.ParseFiles("Trondheim.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, index2)
	if err != nil {
		log.Fatal(err)
	}

}
func Stavanger (w http.ResponseWriter, r *http.Request) {

	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?id=3137115&APPID=b50b6d19b643bd8438490d345969a894")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	jsonErr := json.Unmarshal(body, &index3)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	ProTipsStavanger()

	temp, err := template.ParseFiles("Stavanger.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, index3)
	if err != nil {
		log.Fatal(err)
	}
}

//Åpner nettleseren så fort koden blir kjørt
func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

func main() {

	openBrowser("http://localhost:8080/1")
	http.HandleFunc("/1", Oslo)
	http.HandleFunc("/2", Trondheim)
	http.HandleFunc("/3", Stavanger)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
