package main

import (
	"fmt"
	"os"
)

func main(){
	//os.Setenv("APP_ENV","development")
	fmt.Println(os.Getenv("APP_ENV"))
}