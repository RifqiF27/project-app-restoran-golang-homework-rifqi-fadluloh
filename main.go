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
		fmt.Print("Enter endpoint: ")
		fmt.Scan(&endpoint)
		utils.ClearScreen()
		switch endpoint {
		case "login":
			handler.Login(db)
		case "add-order":
			handler.AddOrder(db)
		case "update-status":
			handler.UpdateStatus(db)
		case "get-order":
			handler.GetOrderItems(db)
		case "add-menu":
			handler.AddMenu(db)
		case "delete-order":
			handler.DeleteOrder(db)
		case "add-rating":
			handler.AddRating(db)
		case "get-rating":
			handler.GetRatings(db)
		case "logout":
			handler.Logout()
			return
		default:
			fmt.Println("Endpoint invalid")
		}
	}
}