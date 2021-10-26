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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("home.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func getWeatherHandler(writer http.ResponseWriter, request *http.Request) {
	city := request.FormValue("city")
	var ans Response
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=49a76d2465245d1581a522d89f08d632&units=metric&lang=ru")
	check(err)
	body, _ := ioutil.ReadAll(resp.Body)
	status := resp.Status
	if status == "404 Not Found" {
		_, err = writer.Write([]byte("Нет такого города"))
		check(err)
	} else {
		err = json.Unmarshal(body, &ans)
		answer := fmt.Sprintf("%f", ans.MainTemp.Temp)
		_, err = writer.Write([]byte(answer))
		check(err)
	}
}

func main() {
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/home/weather", getWeatherHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
