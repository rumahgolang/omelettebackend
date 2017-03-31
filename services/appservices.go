package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	modelsOmelette "github.com/fajarpnugroho/omelettebackend/models"
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

func GetDetailMerchant(merchantId string) (*models.Merchant, error) {
	url := fmt.Sprintf("https://powerful-river-36528.herokuapp.com/api/v1/merchant/%s", merchantId)
	req, errRequest := http.NewRequest("GET", url, nil)

	if errRequest != nil {
		log.Fatalln("NewRequest: ", errRequest)
		return nil, errRequest
	}

	client := &http.Client{}

	resp, errResponse := client.Do(req)
	if errResponse != nil {
		log.Fatalln("Do: ", errResponse)
		return nil, errResponse
	}

	defer resp.Body.Close()

	var merchant models.Merchant
	if errDecoder := json.NewDecoder(resp.Body).Decode(&merchant); errDecoder != nil {
		log.Fatalln(errDecoder)
		return nil, errDecoder
	}

	return &merchant, nil
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

// GetAllMerchant get all merchants
func GetAllMerchant(merchants *models.Merchants) {
	url := "https://powerful-river-36528.herokuapp.com/api/v1/merchants"
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

// AddedNewMerchant add new merchant to omelette
func AddedNewMerchant(merchant *models.Merchant) error {
	url := "https://powerful-river-36528.herokuapp.com/api/v1/merchant/register"

	var jsonStr = []byte(`{"name":"` + merchant.Name + `",
	"range_price":"` + merchant.RangePrice + `",
	"location":"` + merchant.Location + `",
	"pict":"",
	"open_close_info":"` + merchant.OpenCloseInfo + `",
	"latitude":` + fmt.Sprintf("%f", merchant.Latitude) + `,
	"longitude":` + fmt.Sprintf("%f", merchant.Longitude) + `}`)

	log.Printf("%s", jsonStr)

	req, errRequest := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if errRequest != nil {
		log.Fatalln("NewRequest: ", errRequest)
		return errRequest
	}

	client := &http.Client{}

	resp, errResponse := client.Do(req)
	if errResponse != nil {
		log.Fatalln("Do: ", errResponse)
		return errResponse
	}

	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)

	if errRead != nil {
		log.Fatalln("Read: ", errRead)
		return errRead
	}

	fmt.Println("response Body:", string(body))

	return nil
}

// Add Category
func AddedNewCategory(cat *modelsOmelette.MenuCategory) error {
	url := "https://powerful-river-36528.herokuapp.com/api/v1/category"
	catJson, _ := json.Marshal(cat)

	req, errRequest := http.NewRequest("POST", url, bytes.NewBuffer(catJson))
	req.Header.Set("Content-Type", "application/json")

	if errRequest != nil {
		log.Fatalln("NewRequest: ", errRequest)
		return errRequest
	}

	client := &http.Client{}

	resp, errResponse := client.Do(req)
	if errResponse != nil {
		log.Fatalln("Do: ", errResponse)
		return errResponse
	}

	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)

	if errRead != nil {
		log.Fatalln("Read: ", errRead)
		return errRead
	}

	fmt.Println("response Body:", string(body))

	return nil
}

// Get Menus By Merchant Id
func GetAllMenuByMerchantId(merchantId string) (*modelsOmelette.MenuCategories, error) {
	url := fmt.Sprintf("https://powerful-river-36528.herokuapp.com/api/v1/categories/%s", merchantId)

	req, errRequest := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	if errRequest != nil {
		log.Fatalln("NewRequest: ", errRequest)
		return nil, errRequest
	}

	client := &http.Client{}

	resp, errResponse := client.Do(req)
	if errResponse != nil {
		log.Fatalln("Do: ", errResponse)
		return nil, errResponse
	}

	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)

	if errRead != nil {
		log.Fatalln("Read: ", errRead)
		return nil, errRead
	}

	fmt.Println("response Body:", string(body))

	var categories modelsOmelette.MenuCategories

	errJson := json.Unmarshal(body, &categories)

	if errJson != nil {
		return nil, errJson
	}

	return &categories, nil
}

// Update Category By Merchant Id
func UpdateCategoryByMerchantId(cat *modelsOmelette.MenuCategory) (*[]modelsOmelette.MenuCategory, error) {
	url := fmt.Sprintf("https://powerful-river-36528.herokuapp.com/api/v1/category/%s/%s", cat.MerchantID, cat.ID)
	catJson, _ := json.Marshal(cat)
	req, errRequest := http.NewRequest("PUT", url, bytes.NewBuffer(catJson))
	req.Header.Set("Content-Type", "application/json")

	if errRequest != nil {
		log.Fatalln("NewRequest: ", errRequest)
		return nil, errRequest
	}

	client := &http.Client{}

	resp, errResponse := client.Do(req)
	if errResponse != nil {
		log.Fatalln("Do: ", errResponse)
		return nil, errResponse
	}

	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)

	if errRead != nil {
		log.Fatalln("Read: ", errRead)
		return nil, errRead
	}

	fmt.Println("response Body:", string(body))

	var categories []modelsOmelette.MenuCategory

	errJson := json.Unmarshal(body, &categories)

	if errJson != nil {
		return nil, errJson
	}

	return &categories, nil
}

// Delete Category By Merchant Id
func DeleteCategoryByMerchantId(cat *modelsOmelette.MenuCategory) error {
	url := fmt.Sprintf("https://powerful-river-36528.herokuapp.com/api/v1/category/%s/%s", cat.MerchantID, cat.ID)

	req, errRequest := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")

	if errRequest != nil {
		log.Fatalln("NewRequest: ", errRequest)
		return errRequest
	}

	client := &http.Client{}

	resp, errResponse := client.Do(req)
	if errResponse != nil {
		log.Fatalln("Do: ", errResponse)
		return errResponse
	}

	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)

	if errRead != nil {
		log.Fatalln("Read: ", errRead)
		return errRead
	}

	fmt.Println("response Body:", string(body))

	var categories []modelsOmelette.MenuCategory

	errJson := json.Unmarshal(body, &categories)

	if errJson != nil {
		return errJson
	}

	return nil
}
