package main

import (
	"fmt"
	"net/http"
)

func main() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := Routes(db)
	http.ListenAndServe(":8000", r)
}
