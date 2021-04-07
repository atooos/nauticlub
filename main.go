package main

import (
	"github.com/atooos/nauticlub/db/moke"
	"github.com/atooos/nauticlub/service"
)

func main() {
	db := moke.New()
	service.Init("8080", db)
}
