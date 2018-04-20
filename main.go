package main

//uncomment this --
import (
	// "encoding/json"
	"encoding/json"
	"log" //error logging
	"math/rand"
	"net/http" //random number gen
	"os"
	"strconv" //string conversion

	"github.com/gorilla/mux" //router
	"github.com/rs/cors"
)

//Roll is model for sushi
type Roll struct {
	ID          string `json:"id"`
	ImageNumber string `json:"imageNumber"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

//Init rolls var as a slice
var rolls []Roll

//Get All Books
func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls)
}

//Get Single Roll
func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Roll{})
}

//Create a New Roll
func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("localhost:3000"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	var roll Roll
	_ = json.NewDecoder(r.Body).Decode(&roll)
	roll.ID = strconv.Itoa(len(rolls) + 1) //fake an Id for new roll
	roll.ImageNumber = strconv.Itoa(rand.Intn(10) + 1)
	rolls = append(rolls, roll)
	json.NewEncoder(w).Encode(roll)
}

//Update
func updateRoll(w http.ResponseWriter, r *http.Request) {

}

//Delete Roll
func deleteRoll(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//initialize router
	router := mux.NewRouter()

	//stump data TODO: Implement database
	rolls = append(rolls, Roll{ID: "1", ImageNumber: "8", Name: "Spicy Tuna Roll", Ingredients: "Tuna, Chili sauce, Nori, Rice"})

	rolls = append(rolls, Roll{ID: "2", ImageNumber: "4", Name: "California Roll", Ingredients: "Crab, Cucumber, Avocado, Nori, Rice"})

	rolls = append(rolls, Roll{ID: "3", ImageNumber: "10", Name: "Caterpillar Roll", Ingredients: "Tempura Shrimp, Cucumber, Avocado, Nori, Rice, Eel Sauce"})

	rolls = append(rolls, Roll{ID: "4", ImageNumber: "7", Name: "Spider Roll", Ingredients: "Tempura Crab, Nori, Rice, Avocado, Spicy mayonnaise"})

	rolls = append(rolls, Roll{ID: "5", ImageNumber: "2", Name: "Unagi Roll", Ingredients: "Broiled Eel, Avocado, Nori, Rice, Eel sauce"})

	rolls = append(rolls, Roll{ID: "6", ImageNumber: "1", Name: "Philidelphia Roll", Ingredients: "Smoked Salmon, Cream Cheese, Cucumber, Nori, Rice"})

	rolls = append(rolls, Roll{ID: "7", ImageNumber: "3", Name: "Salmon Roll", Ingredients: "Salmon, Nori, Rice, Tobiko"})

	rolls = append(rolls, Roll{ID: "8", ImageNumber: "6", Name: "Avocado Roll", Ingredients: "Avocado, Nori, Rice"})

	rolls = append(rolls, Roll{ID: "9", ImageNumber: "9", Name: "Rainbow Roll", Ingredients: "Yellow Tail, Salmon, Cucumber, Nori, Rice"})

	rolls = append(rolls, Roll{ID: "10", ImageNumber: "5", Name: "Tobiko Roll", Ingredients: "Tobiko, Nori, Sushi Rice"})

	rolls = append(rolls, Roll{ID: "11", ImageNumber: "11", Name: "Yellow Tail Sushi", Ingredients: "Yellow Tail, Wasabi, Sushi Rice"})

	//endpoints
	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	// router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}
