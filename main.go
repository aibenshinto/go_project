package main

import (
	
	"main.go/connection"
)

type Mob_Struct struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Spec   string   `json:"spec"`
	Images []string `json:"images"`
	Price  float64  `json:"price"`
}

func main() {

	db := connection.InitDB()
	defer db.Close()
}
