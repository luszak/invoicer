package main

import (
	"fmt"

	"github.com/luszak/invoicer/invoice"
)

func main() {
	conf := invoice.Config{}
	conf.ReadFromFile("test_config.yml")
	fmt.Printf("%+v", conf)
}
