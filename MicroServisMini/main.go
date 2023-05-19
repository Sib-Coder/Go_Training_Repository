package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Robot struct {
	mass     float64
	velocity float64
}
type IRobot interface {
	getKineticEnergy() float64
}

// я приму только структуру Robot
func (r *Robot) getKineticEnergy() float64 {
	return (r.mass * math.Pow(r.velocity, 2)) * 0.5
}
func example(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mass, err := strconv.ParseFloat(vars["mass"], 64)
	if err != nil {
		log.Println("Error 00001")
	}
	velocity, err := strconv.ParseFloat(vars["velocity"], 64)
	if err != nil {
		log.Println("Error 00002")
	}
	bot := Robot{mass, velocity}
	jsonResponse, jsonError := json.Marshal(bot.getKineticEnergy())
	if jsonError != nil {
		log.Println("Error 00003")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/exaple/{mass:[0-9]+}/{velocity:[0-9]+}/", example)
	log.Fatal(http.ListenAndServe(":8090", r))
}

///"exaple/21/32" - проверка работоспособности сервиса
