package main

import (
	"fmt"
	"time"

	handler "./handlers"
)

//Command to compile in real time and build
//  ./CompileDaemon -directory=/Users/vinayaga/Documents/GolangElasticAPI -command=/Users/vinayaga/Documents/GolangElasticAPI/./GolangElasticAPI

func main() {

	fmt.Println("Starting GolangElasticAPI Application...")

	handler.HandleRequests()

	time.Sleep(10 * time.Minute)

}
