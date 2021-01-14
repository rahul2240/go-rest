package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

var articles Articles

func allArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("All articles api hit")
	json.NewEncoder(w).Encode(articles)
}

func singleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range articles {
		id, err := strconv.Atoi(key)
		if err == nil && article.Id == id {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	json.NewEncoder(w).Encode("{err: error found}")
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article

	json.Unmarshal(reqBody, &article)

	articles = append(articles, article)

	json.NewEncoder(w).Encode(article)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
	fmt.Println("Endpoint Hit: HomePage")
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for index, article := range articles {

		k, err := strconv.Atoi(key)
		if err == nil && article.Id == k {
			articles = append(articles[:index], articles[index+1:]...)
		}
	}

}

func handleRequests() {

	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", homePage)
	myrouter.HandleFunc("/articles", allArticles)
	myrouter.HandleFunc("/article", createArticle).Methods("POST")
	myrouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")

	myrouter.HandleFunc("/article/{id}", singleArticle)

	log.Fatal(http.ListenAndServe(":5000", myrouter))

}

func main() {

	articles = Articles{
		Article{
			Id:      1,
			Title:   "Monk who sold his ferrari",
			Desc:    "Monk learned life lesson",
			Content: "No content, ferrari sold",
		},
		Article{
			Id:      2,
			Title:   "Tandav",
			Desc:    "Ad aari thi",
			Content: "Dimag",
		},
	}

	handleRequests()
}
