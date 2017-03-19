package models

// ZomatoRestaurantResponse Zomato response
type ZomatoRestaurantResponse struct {
	Restaurants []ItemRestaurant `json:"restaurants"`
}

type ItemRestaurant struct {
	Restaurant Restaurant `json:"restaurant"`
}

// Restaurant restaurant entity
type Restaurant struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

// Location location entity
type Location struct {
	Address   string `json:"address"`
	Locality  string `json:"locality"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
