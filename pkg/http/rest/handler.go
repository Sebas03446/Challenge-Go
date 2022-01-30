package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Sebas03446/Challenge-Go/pkg/domain/models"
	"github.com/Sebas03446/Challenge-Go/pkg/domain/validation"
	"github.com/Sebas03446/Challenge-Go/pkg/service"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var UserService *service.Service
var validate *validator.Validate

func createProperty(w http.ResponseWriter, req *http.Request) {
	value, _ := io.ReadAll(req.Body)
	var property models.Property
	err := json.Unmarshal(value, &property)
	if err != nil {
		json.NewEncoder(w).Encode("Error in json Format")
		return
	}
	err = validate.Struct(property)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Error in json Format")
		return
	}
	errs := validation.ChoicePropertyType(property, validate)
	if errs != nil {
		json.NewEncoder(w).Encode(errs)
		return
	}
	_, err = UserService.Create(property)
	if err != nil {
		json.NewEncoder(w).Encode("Error creating property")
		return
	}
	json.NewEncoder(w).Encode(property.PropertyType + " created")
}
func updateProperty(w http.ResponseWriter, req *http.Request) {
	value, _ := io.ReadAll(req.Body)
	var property models.Property
	err := json.Unmarshal(value, &property)
	if err != nil {
		json.NewEncoder(w).Encode("Error in json Format")
		return
	}
	err = validate.Struct(property)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Error in json Format")
		return
	}
	errs := validation.ChoicePropertyType(property, validate)
	if errs != nil {
		json.NewEncoder(w).Encode(errs)
		return
	}
	err = UserService.Update(property)
	if err != nil {
		json.NewEncoder(w).Encode("Error updating property")
		return
	}
	json.NewEncoder(w).Encode(strconv.FormatUint(uint64(property.ID), 10) + " updated")
}
func Controller() {
	validate = validator.New()
	router := mux.NewRouter()
	router.HandleFunc("/v1/properties", createProperty).Methods("POST")
	router.HandleFunc("/v1/properties/{id}", updateProperty).Methods("PUT")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"},
		AllowCredentials: true,
	})
	log.Print("Listening on port 3030")
	handler := c.Handler(router)
	err := http.ListenAndServe(":3030", handler)
	if err != nil {
		fmt.Println("pasos")
		log.Fatal(err)
	}

}
