package main

import(
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


func ClientDo(){
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := client.Do(req)
	req.Header.Set("Client-ID" ,"hq22yaqr8c24w8ymv4q0vhh9wrxzst")
}



func TopStreamHandler(w http.ResponseWriter, re *http.Request){

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func VideoByIDHander(w http.ResponseWriter, re *http.Request){
	vars := mux.Vars(re)
	id :=vars["videoid"]	

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func UserByIdFunc(w http.ResponseWriter, re *http.Request){
	vars := mux.Vars(re)
	id :=vars["userid"]

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func GameByIdFunc(w http.ResponseWriter, re *http.Request){
	vars := mux.Vars(re)
	id :=vars["gameid"]

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}