package main

import (
	"fmt"
	"os"

	"github.com/nt2311-vn/Blog_Aggregator/internal/config"
)

func main() {
	c := config.Read()
	c.SetUser()

	byteData, err := os.ReadFile("./gatorconfig.json")
	if err != nil {
		fmt.Printf("cannot reafile config %s", err)
	}

	fmt.Println(string(byteData))
}
