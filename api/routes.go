package main

import (
	"fmt"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Edge REST API : V0.0.1")
}

var routes = Routes{
	Route{"homePage", "GET", "/", homepage},

	//Articles Information
	Route{"All Articles", "GET", "/allArticles", returnAllArticles},
	Route{"Single Article", "GET", "/article/{id}", returnSingleArticle},
}
