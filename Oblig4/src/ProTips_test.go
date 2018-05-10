package main

import (
	"testing"
	"reflect"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
)

//Collecting JSON-file and executes an unmarshal-command
func OsloTest (url string)  {
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
}

func TrondheimTest (url string){
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
}

func StavangerTest (url string){
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
}

//Testing if the function is returning a message
func TestProtipOslo(t *testing.T) {
	OsloTest("http://api.openweathermap.org/data/2.5/weather?id=6453366&APPID=b50b6d19b643bd8438490d345969a894")
	ProTipsOslo()
	y := index.Main.TempA
	if reflect.DeepEqual(y, reflect.Zero(reflect.TypeOf(y)).Interface()){
		t.Errorf("TempA returnerer ingen string-verdi, forventet beskjed")
	} else {
		t.Errorf("TempA returnerer en beskjed")
	}
}

func TestProtipTrondheim(t *testing.T) {
	TrondheimTest("http://api.openweathermap.org/data/2.5/weather?id=3133881&APPID=b50b6d19b643bd8438490d345969a894")
	ProTipsTrondheim()
	x := index2.Main.TempB
	if reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface()){
		t.Errorf("TempB returnerer ingen string-verdi, forventet beskjed")
	} else {
		t.Errorf("TempB returnerer en beskjed")
	}
}

func TestProtipStavanger(t *testing.T) {
	StavangerTest("http://api.openweathermap.org/data/2.5/weather?id=3137115&APPID=b50b6d19b643bd8438490d345969a894")
	ProTipsStavanger()
	z := index3.Main.TempC
	if reflect.DeepEqual(z, reflect.Zero(reflect.TypeOf(z)).Interface()){
		t.Errorf("TempC returnerer ingen string-verdi, forventet beskjed")
	} else {
		t.Errorf("TempC returnerer en beskjed")}
}
