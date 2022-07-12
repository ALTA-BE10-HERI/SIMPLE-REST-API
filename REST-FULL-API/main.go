package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	Nama     string
	Email    string
	Password string
}

var arrData []User

func main() {
	var isRunning = true
	var inputMenu int
	scanner := bufio.NewScanner(os.Stdin)

	for isRunning {
		fmt.Println("1. register")
		fmt.Println("2. lihat semua data")
		fmt.Println("99. keluar")

		fmt.Println("masukkan pilihan")
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case 1:
			var newUser User
			fmt.Println("Registrasi")
			fmt.Println("Input nama ")
			scanner.Scan()
			newUser.Nama = scanner.Text()
			fmt.Println("Input email ")
			scanner.Scan()
			newUser.Email = scanner.Text()
			fmt.Println("Input password ")
			scanner.Scan()
			newUser.Password = scanner.Text()

			arrData = append(arrData, newUser)
		case 2:
			fmt.Println(arrData)
		case 99:
			isRunning = false
		default:
			continue
		}
	}
}
