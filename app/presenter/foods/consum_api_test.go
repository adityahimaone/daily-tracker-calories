package foods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

type FoodNutrients struct {
	Title  string  `json:"title"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}
type FoodNutrion struct {
	FoodNutrients FoodNutrients
}
type FoodDesc struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	ImageType   string `json:"imageType"`
	FoodNutrion []FoodNutrion
}
type Food struct {
	Results      []FoodDesc `json:"results"`
	Offset       int        `json:"offset"`
	Number       int        `json:"number"`
	TotalResults int        `json:"totalResults"`
}

func TestGetFood(t *testing.T) {
	apikey := "f85868cf3e9f448c851d46fe687a40ac"
	query := "cup cake"
	splitQuery := strings.Split(query, " ")
	joinQuery := strings.Join(splitQuery, "%20")
	log.Println(joinQuery)
	minCal := 0
	number := 1
	urlString := fmt.Sprintf("https://api.spoonacular.com/recipes/complexSearch?apiKey=%s&query=%s&minCalories=%d&number=%d", apikey, joinQuery, minCal, number)
	log.Println(urlString)
	response, err := http.Get(urlString)
	if err != nil {
		log.Println(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	food := Food{}
	err = json.Unmarshal(responseData, &food)
	if err != nil {
		panic(err)
	}
	log.Println(food.Results)
	log.Println(food.Number)
}

func TestSplit(t *testing.T) {
	query := "cup cake"
	splitQuery := strings.Split(query, " ")
	joinQuery := strings.Join(splitQuery, "%20")
	log.Println(joinQuery)

	foodNts := FoodNutrients{"Calories", "calories", 200.0, "kcal"}
	foodNto := FoodNutrion{foodNts}
	food1 := FoodDesc{12, "Anggur", "", "", []FoodNutrion{foodNto}}
	res := Food{[]FoodDesc{food1}, 1, 2, 3}
	log.Println(res)
}
