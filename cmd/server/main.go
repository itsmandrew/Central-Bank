package main

import (
	"fmt"

	"github.com/itsmandrew/Central-Bank/internal/config"
)

func main() {
	cgf := config.Load()
	fmt.Println(cgf)
}
