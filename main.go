package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CovidData struct {
	Centers []struct {
		CenterID     int    `json:"center_id"`
		Name         string `json:"name"`
		StateName    string `json:"state_name"`
		DistrictName string `json:"district_name"`
		BlockName    string `json:"block_name"`
		Pincode      int    `json:"pincode"`
		Lat          int    `json:"lat"`
		Long         int    `json:"long"`
		From         string `json:"from"`
		To           string `json:"to"`
		FeeType      string `json:"fee_type"`
		Sessions     []struct {
			SessionID         string   `json:"session_id"`
			Date              string   `json:"date"`
			AvailableCapacity int      `json:"available_capacity"`
			MinAgeLimit       int      `json:"min_age_limit"`
			Vaccine           string   `json:"vaccine"`
			Slots             []string `json:"slots"`
		} `json:"sessions"`
		VaccineFees []struct {
			Vaccine string `json:"vaccine"`
			Fee     string `json:"fee"`
		} `json:"vaccine_fees,omitempty"`
	} `json:"centers"`
}

func main() {

	url := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByPin?pincode=332001&date=01-05-2021"
	method := "GET"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "cdn-api.co-vin.in")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("origin", "https://www.cowin.gov.in")
	req.Header.Add("referer", "https://www.cowin.gov.in/")
	req.Header.Add("accept-language", "en-IN,en-GB;q=0.9,en-US;q=0.8,en;q=0.7")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var covidData CovidData
	err = json.Unmarshal(body, &covidData)
	if err != nil {
		fmt.Println("-----------")
		fmt.Println(err)
		fmt.Println("-----------")
	}
	totalCenters := len(covidData.Centers)
	for i := 0; i < totalCenters; i++ {
		totalSessions := len(covidData.Centers[i].Sessions)
		for j := 0; j < totalSessions; j++ {
			minAgeLimit := covidData.Centers[i].Sessions[j].MinAgeLimit
			if minAgeLimit == 18 {
				fmt.Println("+++++++++++++++Center Information+++++++++++++++++")
				fmt.Printf("Center Name:\t %s \n", covidData.Centers[i].Name)
				fmt.Printf("Center Name:\t %s \n", covidData.Centers[i].Name)
				fmt.Printf("Date:\t %s \n", covidData.Centers[i].Sessions[j].Date)
				fmt.Printf("Available Capacity:\t %d \n", covidData.Centers[i].Sessions[j].AvailableCapacity)
			}

		}
	}
}