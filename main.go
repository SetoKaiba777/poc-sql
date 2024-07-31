package main

import (
	"fmt"
	"poc-mysql/database"
)

func main() {
	db := database.NewSqlDatabase("root","123","localhost:3306","recordings")
	outList, _ :=db.SelectQuery("* FROM album WHERE artist = ?","John Coltrane",)
	for _,out := range outList{
		fmt.Println(out)
	}
}