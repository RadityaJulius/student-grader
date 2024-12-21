package mahasiswacontroller

import (
	"golang-mahasiswa/entities"
	"golang-mahasiswa/models/mahasiswamodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	mahasiswa := mahasiswamodel.GetAll()
	data := map[string]any{
		"mahasiswa": mahasiswa,
	}

	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var mahasiswa entities.Mahasiswa

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			panic(err)
		}

		uts, err := strconv.Atoi(r.FormValue("uts"))
		if err != nil {
			panic(err)
		}

		uas, err := strconv.Atoi(r.FormValue("uas"))
		if err != nil {
			panic(err)
		}

		mahasiswa.Id = int64(id)
		mahasiswa.Nama = r.FormValue("nama")
		mahasiswa.Uts = float32(uts)
		mahasiswa.Uas = float32(uas)
		mahasiswa.CreatedAt = time.Now()
		mahasiswa.UpdatedAt = time.Now()
		mahasiswa.Grade = (mahasiswa.Uts + mahasiswa.Uas) / 2

		ok := mahasiswamodel.Create(mahasiswa)
		if !ok {
			temp, _ := template.ParseFiles("views/index.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
