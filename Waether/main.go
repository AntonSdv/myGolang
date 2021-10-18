package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	MakeRequest()
}

func MakeRequest() {
	var ans Response
	var city string
	fmt.Println("Введите город")
	fmt.Scan(&city)
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=49a76d2465245d1581a522d89f08d632&units=metric&lang=ru")
	if err != nil {
		log.Fatalln(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	status := resp.Status
	if status == "404 Not Found" {
		fmt.Println("Нет такого города")
	} else {
		err = json.Unmarshal(body, &ans)
		fmt.Printf("%v\n", ans.MainTemp.Temp)
	}

}
