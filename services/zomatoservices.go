package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/fajarpnugroho/omelettebackend/models"
)

const (
	// ZomatoAPIKey Zomato API KEY
	ZomatoAPIKey = "8f579904c450d535003969488316607c"
)

// SearchRestaurantByKeyword Search restaurant using keyword in zomato
func SearchRestaurantByKeyword(zomatoRestaurantResponse *models.ZomatoRestaurantResponse, keyword string) {
	log.Printf("%s", keyword)

	urlString := fmt.Sprintf("https://developers.zomato.com/api/v2.1/search?q=%s", url.QueryEscape(keyword))

	log.Printf("%s", urlString)

	req, errRequest := http.NewRequest("GET", urlString, nil)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-key", ZomatoAPIKey)

	log.Printf("%s", req.Header.Get("user-key"))
	log.Printf("%s", req.Header.Get("Accept"))

	if errRequest != nil {
		log.Fatalln("NewRequest: ", errRequest)
		return
	}

	client := &http.Client{}

	resp, errResponse := client.Do(req)
	if errResponse != nil {
		log.Fatalln("Do: ", errResponse)
		return
	}

	if resp.StatusCode > 200 {
		log.Println("Response unsuccessful")
		return
	}

	defer resp.Body.Close()

	if errDecoder := json.NewDecoder(resp.Body).Decode(&zomatoRestaurantResponse); errDecoder != nil {
		log.Fatalln("Decode Error: ", errDecoder)
		return
	}
}

// RestaurantDetail Get detail restaurant by id
func RestaurantDetail(merchantID int64) models.Restaurant {

	urlString := fmt.Sprintf("https://developers.zomato.com/api/v2.1/restaurant?res_id=%d", merchantID)

	log.Printf("%s", urlString)

	req, errRequest := http.NewRequest("GET", urlString, nil)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-key", ZomatoAPIKey)

	log.Printf("%s", req.Header.Get("user-key"))
	log.Printf("%s", req.Header.Get("Accept"))

	if errRequest != nil {
		log.Fatalln("NewRequest: ", errRequest)
		return models.Restaurant{}
	}

	client := &http.Client{}

	resp, errResponse := client.Do(req)
	if errResponse != nil {
		log.Fatalln("Do: ", errResponse)
		return models.Restaurant{}
	}

	if resp.StatusCode > 200 {
		log.Println("Response unsuccessful")
		return models.Restaurant{}
	}

	defer resp.Body.Close()

	var merchantProfile models.Restaurant

	if errDecoder := json.NewDecoder(resp.Body).Decode(&merchantProfile); errDecoder != nil {
		log.Fatalln("Decode Error: ", errDecoder)
		return models.Restaurant{}
	}

	return merchantProfile
}
