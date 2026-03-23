package main

import (
	"fmt"
	"os"
)

func callbackExit(){

	os.Exit(0)
	fmt.Print("check something")
}