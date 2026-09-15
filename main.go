package main

import (
	"fmt"
	"net/http"
)

// buat handlernya untuk mencetak hellow world
func helloHandler(w http.ResponseWriter, r *http.Request) {
	//deklarasi dan inisialisasi variable name dari query parameter
	name := r.URL.Query().Get("name")

	// Jika name kosong, gunakan "Guest" sebagai nilai default
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintln(w, "hello", name)
}

// buat fungsi main
func main() {

	//regist helloHandler ke path "/" handlerfunc
	http.HandleFunc("/hello", helloHandler)

	//cetak informasi web server berjalan di localhost
	fmt.Println("Server berjalan di http://localhost:8080")

	err := http.ListenAndServe(":8080", nil) // Menjalankan server localhost di port 8080
	if err != nil {                          //lakukan erro handling jika server tidak jalan berikan pesan
		fmt.Println(err)
	}
}
