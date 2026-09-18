package main

import (
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		name := r.URL.Query().Get("name") //deklarasi dan inisialisasi variable name dari query parameter
		if name == "" {                   // Jika name kosong, gunakan "Guest" sebagai nilai default
			name = "Guest"
		}
		fmt.Fprintln(w, "hello", name)
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error!", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, string(body))
	default:
		http.Error(w, "Hanya menerima GET dan POST", http.StatusMethodNotAllowed)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Ini adalah halaman About")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Selamat Datang !")
}

func main() {

	http.HandleFunc("/", rootHandler)       //register welcomeHandler ke root path "/"
	http.HandleFunc("/hello", helloHandler) //register helloHandler ke path "/hello" handlerfunc
	http.HandleFunc("/about", aboutHandler) //register abouthandler ke path "/about" handlerfunc

	fmt.Println("Server berjalan di http://localhost:8080") //cetak informasi web server berjalan di localhost

	err := http.ListenAndServe(":8080", nil) // Menjalankan server localhost di port 8080
	if err != nil {                          //lakukan erro handling jika server tidak jalan berikan pesan
		fmt.Println(err)
	}
}
