package main

import (
	"html/template"
	"log"
	"net/http"
)

// 1. Kita siapkan struktur datanya (Mirip struct di C++)
type SocialLink struct {
	Platform string
	URL      string
	Icon     string // Kita pakai emoji dulu biar simpel
}

type UserProfile struct {
	Name     string
	Bio      string
	ImageURL string
	Links    []SocialLink // Array of links (One-to-Many)
}

// 2. Handler untuk halaman utama
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Isi data profil kamu di sini
	data := UserProfile{
		Name:     "M. Nabil Fabian",
		Bio:      "IT Student | Gym Enthusiast",
		ImageURL: "https://avatars.githubusercontent.com/u/240920051?s=400&u=2d9a60ce5f27ccf9c90365d476e17fdcc7a16ef0&v=4", // Ganti link foto kamu nanti
		Links: []SocialLink{
			{Platform: "GitHub", URL: "https://github.com/blowmamain-spec", Icon: "💻"},
			{Platform: "LinkedIn", URL: "https://www.linkedin.com/in/m-nabil-fabian", Icon: "g"},
			{Platform: "Instagram", URL: "https://www.instagram.com/nblfbiann", Icon: "📸"},
			{Platform: "Email", URL: "mail to: blowmamain@gmail.com", Icon: "📧"},
			{Platform: "WhatsApp", URL: "https://wa.me/6281262362940", Icon: "📞"},
		},
	}

	// Parsing file HTML
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Gabungkan data Go ke dalam HTML
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)

	log.Println("Server jalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
