package main

import (
	"fmt"

	handler "./handlers"
)

//Command to compile in real time and build
//  ./CompileDaemon -directory=/Users/vinayaga/Documents/Github/GolangElasticAPI -command=/Users/vinayaga/Documents/Github/GolangElasticAPI/./GolangElasticAPI

func main() {

	fmt.Println("Starting GolangElasticAPI Application...")

	handler.HandleRequests()

}
