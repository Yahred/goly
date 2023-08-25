package main

import (
	"fmt"

	"goly.com/app/models"
	"goly.com/app/server"
)

func main() {
	models.Setup()

	defer server.SetupAndListen()

	fmt.Println("Servicio iniciado en el puerto 3000")
}
