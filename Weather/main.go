package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type MainTemp struct {
	Temp float64 `json:"temp"`
}

type Response struct {
	MainTemp `json:"main"`
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("home.html")
	if err != nil {
		log.Fatal(err)
	}
	err = html.Execute(writer, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getWeatherHandler(writer http.ResponseWriter, request *http.Request) {
	city := request.FormValue("city")
	var ans Response
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=49a76d2465245d1581a522d89f08d632&units=metric&lang=ru")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusNotFound {
		_, err = writer.Write([]byte("Нет такого города"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = json.Unmarshal(body, &ans)
		if err != nil {
			log.Fatal(err)
		}
		answer := fmt.Sprintf("%.1f", ans.MainTemp.Temp)
		_, err = writer.Write([]byte(answer))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/home/weather", getWeatherHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
