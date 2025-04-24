package main

import (
	"fmt"

	"github.com/itsmandrew/central-finance/internal/config"
)

func main() {
	cgf := config.Load()
	fmt.Println(cgf)
}
