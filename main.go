package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	. "github.com/giovapanasiti/logbucket/config"
	. "github.com/giovapanasiti/logbucket/dao"
	"github.com/giovapanasiti/logbucket/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var dao = LogPortaleDAO{}
var config = Config{}

func CreateLogPortale(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var logportale models.LogPortale

	auth := r.Header.Get("auth")

	if auth != "questo_sara_il_fantastico_token" {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&logportale); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	logportale.ID = bson.NewObjectId()
	logportale.CreatedAt = time.Now()

	if err := dao.Insert(logportale); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, logportale)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {

	dao.Server = os.Getenv("MONGO_URL")
	dao.Database = os.Getenv("MONGO_DB")
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/portale", CreateLogPortale).Methods("POST")
	fmt.Println("Starting.....")
	if err := http.ListenAndServe(":1314", r); err != nil {
		log.Fatal(err)
	}
}
