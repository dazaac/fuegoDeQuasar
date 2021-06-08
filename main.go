/*
 * author: Andrea Daza
 * email: andreacdazar1@gmail.com
 * topic: Operación fuego de Quasar
 * date: 07/06/2021
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func TopSecret(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, _ := ioutil.ReadAll(r.Body)
		var satell []Satellite
		erro := json.Unmarshal(body, &satell)
		if erro != nil {
			errorHandler(w, r, http.StatusNotFound, "el request")
			return
		}
		distances := make([]float64, 0)
		var messages [][]string
		for _, satellit := range satell {
			distances = append(distances, satellit.Distance)
			messages = append(messages, satellit.Message)
		}
		x, y, err := GetLocation(distances...)
		message, errmsg := GetMessage(messages...)
		if err {
			errorHandler(w, r, http.StatusNotFound, "la posición")
			return
		} else if errmsg {
			errorHandler(w, r, http.StatusNotFound, "el mensaje")
			return
		} else {
			var response Response
			response.Position.X = x
			response.Position.Y = y
			response.Message = message
			json.NewEncoder(w).Encode(response)
		}
	default:
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int, err string) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 ACDR: No se puede determinar ", err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Autor: andreacdazar1@gmail.com")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/topsecret", TopSecret)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
