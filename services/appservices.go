package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fajarpnugroho/omelettebot/models"
)

func SearchMerchantByLocation(merchants *models.Merchants, latitude float64, longitude float64) {
	url := fmt.Sprintf("https://powerful-river-36528.herokuapp.com/api/v1/merchants?lat=%f&long=%f", latitude, longitude)
	req, errRequest := http.NewRequest("GET", url, nil)

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

	defer resp.Body.Close()

	if errDecoder := json.NewDecoder(resp.Body).Decode(&merchants); errDecoder != nil {
		log.Fatalln(errDecoder)
		return
	}
}

func GetDetailMerchant(merchant *models.Merchant, merchantId string) {
	url := fmt.Sprintf("https://powerful-river-36528.herokuapp.com/api/v1/merchant/%s", merchantId)
	req, errRequest := http.NewRequest("GET", url, nil)

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

	defer resp.Body.Close()

	if errDecoder := json.NewDecoder(resp.Body).Decode(&merchant); errDecoder != nil {
		log.Fatalln(errDecoder)
		return
	}
}

func GetMenuCategory(menuCategories *models.MenuCategories, merchantId string) {
	url := fmt.Sprintf("https://powerful-river-36528.herokuapp.com/api/v1/categories/%s", merchantId)
	req, errRequest := http.NewRequest("GET", url, nil)

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

	defer resp.Body.Close()

	if errDecoder := json.NewDecoder(resp.Body).Decode(&menuCategories); errDecoder != nil {
		log.Fatalln(errDecoder)
		return
	}
}

func GetMenuByCategory(menus *models.Menus, menuCategoryId string) {
	url := fmt.Sprintf("https://powerful-river-36528.herokuapp.com/api/v1/menus/%s", menuCategoryId)
	req, errRequest := http.NewRequest("GET", url, nil)

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

	defer resp.Body.Close()

	if errDecoder := json.NewDecoder(resp.Body).Decode(&menus); errDecoder != nil {
		log.Fatalln(errDecoder)
		return
	}
}
