package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ID      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{ID: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{ID: 2, Title: "Hello2", Desc: "Article Description2", Content: "Article Content2"},
	}
	fmt.Println("EndPoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint Hit:Single article")
	fmt.Fprintf(w, "Keys : "+key)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the HomePage")
	fmt.Println("Endpoint Hit: Homepage")
}
