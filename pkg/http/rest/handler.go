package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	response := ""
	value, _ := io.ReadAll(req.Body)
	var property models.Property
	err := json.Unmarshal(value, &property)
	if err != nil {
		json.NewEncoder(w).Encode("Error in json Format")
		return
	}
	property.SetCreatedAt(time.Now())
	property.SetUpdatedAt(time.Now())
	property.SetStatus("")
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
	if property.GetStatus() != "INVALID" {
		status := validation.CountryValidation(property, validate)
		property.SetStatus(status)
		if property.GetStatus() == "INACTIVE" {
			response = "Property is not active, check the country"

		} else if property.GetStatus() == "INVALID" {
			response = "Property is not valid, check the price"
		}
	}
	_, err = UserService.Create(property)
	if err != nil {
		json.NewEncoder(w).Encode("Error creating property")
		return
	}
	json.NewEncoder(w).Encode(property.PropertyType + " created, " + response)
}
func updateProperty(w http.ResponseWriter, req *http.Request) {
	value, _ := io.ReadAll(req.Body)
	var property models.Property
	err := json.Unmarshal(value, &property)
	if err != nil {
		json.NewEncoder(w).Encode("Error in json Format")
		return
	}
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Error in json Format")
		return
	}
	propertyForUpdate, err := UserService.GetById(id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	property.SetCreatedAt(propertyForUpdate.GetCreatedAt())
	property.SetUpdatedAt(time.Now())
	property.SetStatus("")
	err = validate.Struct(property)
	if err != nil {
		json.NewEncoder(w).Encode("Error in json Format")
		return
	}
	errs := validation.ChoicePropertyType(property, validate)
	if errs != nil {
		json.NewEncoder(w).Encode(errs)
		return
	}
	if property.GetStatus() != "INVALID" {
		status := validation.CountryValidation(property, validate)
		property.SetStatus(status)

	}
	err = UserService.Update(property)
	if err != nil {
		json.NewEncoder(w).Encode("Error updating property")
		return
	}
	json.NewEncoder(w).Encode(strconv.FormatUint(uint64(property.ID), 10) + " updated")
}

func getProperty(w http.ResponseWriter, req *http.Request) {
	status := req.URL.Query().Get("status")
	bbox := req.URL.Query().Get("bbox")
	var params map[string]interface{}
	if bbox != "" && status != "" {
		splitBbox := strings.Split(bbox, ",")
		x0, err := strconv.ParseFloat(splitBbox[0], 64)
		if err != nil {
			json.NewEncoder(w).Encode("Error in bbox minLongitude")
			return
		}
		y0, err := strconv.ParseFloat(splitBbox[1], 64)
		if err != nil {
			json.NewEncoder(w).Encode("Error in bbox minLatitude")
			return
		}
		x1, err := strconv.ParseFloat(splitBbox[2], 64)
		if err != nil {
			json.NewEncoder(w).Encode("Error in bbox maxLongitude")
			return
		}
		y1, err := strconv.ParseFloat(splitBbox[3], 64)
		if err != nil {
			json.NewEncoder(w).Encode("Error in bbox maxLatitude")
			return
		}
		params = map[string]interface{}{
			"status": status,
			"x0":     x0,
			"y0":     y0,
			"x1":     x1,
			"y1":     y1,
		}

	} else if bbox == "" && status != "" {
		params = map[string]interface{}{
			"status": status,
		} //Voy aqui, revisar status y condiciones

	} else if bbox != "" && status == "" {
		splitBbox := strings.Split(bbox, ",")
		x0, err := strconv.ParseFloat(splitBbox[0], 64)
		if err != nil {
			json.NewEncoder(w).Encode("Error in bbox minLongitude")
			return
		}
		y0, err := strconv.ParseFloat(splitBbox[1], 64)
		if err != nil {
			json.NewEncoder(w).Encode("Error in bbox minLatitude")
			return
		}
		x1, err := strconv.ParseFloat(splitBbox[2], 64)
		if err != nil {
			json.NewEncoder(w).Encode("Error in bbox maxLongitude")
			return
		}
		y1, err := strconv.ParseFloat(splitBbox[3], 64)
		if err != nil {
			json.NewEncoder(w).Encode("Error in bbox maxLatitude")
			return
		}
		params = map[string]interface{}{
			"x0": x0,
			"y0": y0,
			"x1": x1,
			"y1": y1,
		}
	} else {
		params = map[string]interface{}{
			"status": "ALL",
		}
	}

	/*page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(req.URL.Query().Get("pageSize"))*/
	paginated, _ := UserService.GetByParams(params)
	json.NewEncoder(w).Encode(paginated)

}
func Controller() {
	validate = validator.New()
	router := mux.NewRouter()
	router.HandleFunc("/v1/properties", createProperty).Methods("POST")
	router.HandleFunc("/v1/properties/{id}", updateProperty).Methods("PUT")
	router.HandleFunc("/v1/properties", getProperty).Methods("GET")
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
