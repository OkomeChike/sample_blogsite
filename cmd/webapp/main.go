package main

import "webapp/controller"

func main() {
	router := controller.GetRouter()
	router.Run(":8000")
}
