package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Ashish2345/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.Test()
	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
