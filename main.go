package main

import (
	"fmt"
	"restaurant-app/database"
	"restaurant-app/utils"
	"restaurant-app/view"

	_ "github.com/lib/pq"
)

func main() {
	db, err := database.InitDb()
	if err != nil {
		fmt.Println("Gagal menginisialisasi database:", err)
		return
	}
	defer db.Close()

	for {

		var endpoint string
		fmt.Print("Masukkan endpoint: ")
		fmt.Scan(&endpoint)
		utils.ClearScreen()
		switch endpoint {
		case "login":
			handler.Login(db)
		case "add-order":
			handler.AddOrder(db)
		case "update-status":
			handler.UpdateStatus(db)
		// case "update-stock":
		// 	handler.UpdateItemStock(db)
		// case "add-transaction":
		// 	handler.AddTransaction(db)
		// case "get-transaction":
		// 	handler.GetTransactions(db)
		// case "add-category":
		// 	handler.AddCategory(db)
		// case "add-location":
		// 	handler.AddLocation(db)
		case "logout":
			handler.Logout()
			return
		default:
			fmt.Println("Endpoint tidak dikenal")
		}
	}
}