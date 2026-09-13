package main

import (
	"fmt"
	"net/http"
)

// buat handlernya untuk mencetak hellow world
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

// buat fungsi main
func main() {
	//regist helloHandler ke path "/" handlerfunc
	http.HandleFunc("/", helloHandler)

	//cetak informasi web server berjalan di localhost
	fmt.Println("Server berjalan di http://localhost:8080")

	err := http.ListenAndServe(":8080", nil) // Menjalankan server localhost di port 8080
	if err != nil {                          //lakukan erro handling jika server tidak jalan berikan pesan
		fmt.Println(err)
	}
}
