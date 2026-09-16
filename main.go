package main

import (
	"fmt"
	"io"
	"net/http"
)

// buat handlernya untuk mencetak hellow world
func helloHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		//deklarasi dan inisialisasi variable name dari query parameter
		name := r.URL.Query().Get("name")

		// Jika name kosong, gunakan "Guest" sebagai nilai default
		if name == "" {
			name = "Guest"
		}
		fmt.Fprintln(w, "hello", name)
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)

		if err != nil {
			fmt.Fprintln(w, "Error!", err)
			return
		}
		fmt.Fprintln(w, string(body))
	default:
		fmt.Fprintln(w, "Hanya menerima GET dan POST")

	}

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Ini adalah halaman About")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Selamat Datang !")
}

// buat fungsi main
func main() {

	http.HandleFunc("/", rootHandler)       //register welcomeHandler ke root path "/"
	http.HandleFunc("/hello", helloHandler) //register helloHandler ke path "/hello" handlerfunc
	http.HandleFunc("/about", aboutHandler) //register abouthandler ke path "/about" handlerfunc

	//cetak informasi web server berjalan di localhost
	fmt.Println("Server berjalan di http://localhost:8080")

	err := http.ListenAndServe(":8080", nil) // Menjalankan server localhost di port 8080
	if err != nil {                          //lakukan erro handling jika server tidak jalan berikan pesan
		fmt.Println(err)
	}
}
