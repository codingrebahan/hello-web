package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type User struct {
	ID    int
	Name  string
	Email string
}

var users = make(map[int]User)
var nextID = 1

func usersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		id := r.PathValue("id")
		if id == "" {
			userList := make([]User, 0)
			for _, user := range users {
				userList = append(userList, user)
			}
			w.Header().Set("Content-Type", "application/json")
			encoder := json.NewEncoder(w)
			encoder.Encode(userList)
			return
		}
		userID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Input Id yang sesuai", http.StatusBadRequest)
			return
		}
		user, ok := users[userID]
		if ok == false {
			http.Error(w, "Id tidak ditemukan", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(user)

	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var user User
		err := decoder.Decode(&user)
		if err != nil {
			http.Error(w, "JSON tidak valid", http.StatusBadRequest)
			return
		}
		if user.Name == "" || user.Email == "" {
			http.Error(w, "Data nama atau email tidak boleh kosong", http.StatusBadRequest)
			return
		}
		user.ID = nextID
		users[user.ID] = user
		nextID++
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		encoder := json.NewEncoder(w)
		err = encoder.Encode(&user)
	default:
		http.Error(w, "Hanya menerima GET dan POST", http.StatusMethodNotAllowed)
	}
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

	http.HandleFunc("/", rootHandler)            //register welcomeHandler ke root path "/"
	http.HandleFunc("/hello", helloHandler)      //register helloHandler ke path "/hello" handlerfunc
	http.HandleFunc("/about", aboutHandler)      //register abouthandler ke path "/about" handlerfunc
	http.HandleFunc("/users", usersHandler)      //register usershandler ke path "/users" handlerfunc
	http.HandleFunc("/users/{id}", usersHandler) //register userhandler ke path "/users{id}"

	fmt.Println("Server berjalan di http://localhost:8080") //cetak informasi web server berjalan di localhost

	err := http.ListenAndServe(":8080", nil) // Menjalankan server localhost di port 8080
	if err != nil {                          //lakukan erro handling jika server tidak jalan berikan pesan
		fmt.Println(err)
	}
}
