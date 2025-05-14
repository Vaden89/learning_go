package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	//We need to dd a json name to the struct to make it a valid json object
	Name string `json:"name"`
}

var userCache = make(map[int]User)

var cacheMutex sync.RWMutex

func main(){
	//mux is what handles our request, allowing requests to be mapped to the particular handler they need to go to
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	mux.HandleFunc("POST /users", createUser)

	//{id} allows for us to have wildcards in our request url
	mux.HandleFunc("GET /users/{id}", getUser)

	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	fmt.Println("Server listening to :8080")
	//This starts up the http server, taking in the port as the first param and the request multiplexier as the second param 
	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world")
}

func deleteUser (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := userCache[id]; !ok {
		http.Error(w, "User not found!", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	delete(userCache, id)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)

}

func getUser (w http.ResponseWriter, r *http.Request){
	//Allows us to get the wildcard value of the id, since it's part of the url path, this also converts the id to a number
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cacheMutex.RLock()
	user, ok := userCache[id]
	cacheMutex.RUnlock()

	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	//Tells the client what type of data is being sent
	w.Header().Set("Content-Type", "application/json")
	//Converts data to valid json
	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.WriteHeader(http.StatusOK)
	// Return the bytes of data to the client
	w.Write(j)
}

func createUser(w http.ResponseWriter, r *http.Request){
	var user User 

	//Decodes the request body which contains the information about the user, and stores it inside of the user we created
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	userCache[len(userCache) + 1] = user
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}