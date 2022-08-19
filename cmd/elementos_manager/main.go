package main

import (
	"fmt"

	"github.com/ad9311/elementos_manager/internal/environment"
)

func main() {
	fmt.Println("Elementos API")

	conf, err := environment.New("development")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(conf.Development.DatabaseURL)
}
