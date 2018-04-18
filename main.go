package main

//uncomment this --
import (
	// "encoding/json"
	"encoding/json"
	"log"      //error logging
	"net/http" //random number gen
	// "strconv"
	// math/rand
	"github.com/gorilla/mux" //router
)

//Roll is model for sushi
type Roll struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Ingredients  string `json:"ingredients"`
	Instructions string `json:"instructions"`
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

}

//Create a New Roll
func createRoll(w http.ResponseWriter, r *http.Request) {

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
	rolls = append(rolls, Roll{ID: "1", Name: "Spicy Tuna Roll", Ingredients: "Big Eye Tuna, Chili sauce, Nori, Sushi Rice, Shichimi Togarashi, Tobiko (optional), Sesame oil (optional)", Instructions: "roll using mat"})

	rolls = append(rolls, Roll{ID: "2", Name: "California Roll", Ingredients: "Crab, Sliced Cucumber, Avocado, Nori, Sushi Rice, Tobiko (optional)", Instructions: "roll using mat"})

	rolls = append(rolls, Roll{ID: "3", Name: "Caterpillar Roll", Ingredients: "4 pieces of Tempura Shrimp, Sliced Cucumber, Avocado, Nori, Sushi Rice, Eel Sauce, Sesame seeds", Instructions: "roll using mat"})

	rolls = append(rolls, Roll{ID: "4", Name: "Spider Roll", Ingredients: "Tempura Soft-shell Crab, Nori, Sushi Rice, Avocado, Spicy mayonnaise (optional)", Instructions: "roll using mat"})

	rolls = append(rolls, Roll{ID: "5", Name: "Unagi Roll", Ingredients: "Broiled Eel, Avocado, Nori, Sushi Rice, Eel sauce", Instructions: "roll using mat"})

	rolls = append(rolls, Roll{ID: "6", Name: "Philidelphia Roll", Ingredients: "Smoked Salmon, Cream Cheese, Cucumber, Nori, Sushi Rice", Instructions: "roll using mat"})

	rolls = append(rolls, Roll{ID: "7", Name: "Salmon Roll", Ingredients: "Salmon, Nori, Sushi Rice, Tobiko (optional)", Instructions: "roll using mat"})

	rolls = append(rolls, Roll{ID: "8", Name: "Avocado Roll", Ingredients: "Avocado, Nori, Sushi Rice", Instructions: "roll using mat"})

	rolls = append(rolls, Roll{ID: "9", Name: "Sweet Potato Roll", Ingredients: "Tempura Sweet Potato, Nori, Sushi Rice", Instructions: "roll using mat"})

	//endpoints
	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi/", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
