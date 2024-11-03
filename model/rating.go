package model

type Rating struct {
    ID         int    `json:"id"`
    OrderID    int    `json:"order_id"`   
    CustomerID int    `json:"customer_id"`
    Rating     int    `json:"rating"`     
    Comment    string `json:"comment"`    
}
