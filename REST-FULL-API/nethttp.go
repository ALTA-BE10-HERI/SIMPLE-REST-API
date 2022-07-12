package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Nama     string
	Email    string
	Password string
}

var arrData []User

func HalloWorld(writer http.ResponseWriter, request *http.Request) { //format dari HandleFunc
	switch request.Method {
	case "GET":
		writer.Header().Set("contect-type", "aplication/json")
		msg := "Hallo world!"
		writer.Write([]byte(msg))
	}
}

func UserHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "aplication/json") //untuk menentukan deader contentype
	switch request.Method {
	case "POST":

		//baca input dari postman
		var tmp User
		decode := json.NewDecoder(request.Body)
		err := decode.Decode(&tmp)
		if err != nil {
			log.Println("Cannot parse", err.Error())
		}

		//menyiapkan respon untuk di tampilkan ke postmen
		log.Println(tmp)
		arrData = append(arrData, tmp)
		res := map[string]interface{}{
			"message": "Succes input data",
			"data":    tmp,
		}
		send, err := json.Marshal(res) //mengubah data biasa ke json
		if err != nil {
			log.Println("Cannot send ", err.Error())
		}

		// untuk menampilkan respon
		writer.Write(send)

	case "GET":
		res := map[string]interface{}{
			"message": "Get all data",
			"data":    arrData,
		}
		send, err := json.Marshal(res) //mengubah data biasa ke json
		if err != nil {
			log.Println("Cannot send ", err.Error())
		}

		// untuk menampilkan respon
		writer.Write(send)
	}
}
func UserHandlerByid(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "aplication/json")
	switch request.Method {
	case "GET":
		path := request.URL.Path
		params := strings.Split(path, "/")
		log.Println(params)

		cnv, err := strconv.Atoi(params[len(params)-1])
		log.Println(cnv)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("internal server eror"))
			log.Println("canot convert to int ", err.Error())
			return
		}
		if cnv > len(arrData) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("internal server eror"))
			log.Println("Index out off range")
			return
		}

		res := map[string]interface{}{
			"message": "Get data by id",
			"data":    arrData[cnv-1],
		}
		send, err := json.Marshal(res) //mengubah data biasa ke json
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("internal server eror"))
			log.Println("Cannot send ", err.Error())
		}

		// untuk menampilkan respon
		writer.Write(send)
	}
}
func main() {
	http.HandleFunc("/halloworld", HalloWorld)
	http.HandleFunc("/user", UserHandler)
	http.HandleFunc("/user/", UserHandlerByid)

	fmt.Println("menjalankan program ...")
	err := http.ListenAndServe(":8000", nil) //buka port
	if err != nil {
		log.Fatal(err.Error()) //untuk memperhentikan program saat program ada yang eror
	}
}
