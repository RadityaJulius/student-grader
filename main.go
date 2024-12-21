package main

import (
	"golang-mahasiswa/config"
	"golang-mahasiswa/controllers/mahasiswacontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", mahasiswacontroller.Index)
	http.HandleFunc("/add", mahasiswacontroller.Add)

	log.Println("Server running on port: 8080")
	http.ListenAndServe(":8080", nil)
}
