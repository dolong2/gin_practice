package main

import "gin_practice/src/practice/gin/domain/controller"

func main() {
	r := controller.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
