package main

import (
	"fmt"
	"github.com/darthlukan/goconvtemps"
	"os"
	"strconv"
	"strings"
)

func main() {

	var unit string
	var converted string

	in := strings.ToLower(os.Args[1])

	if strings.Index(in, "c") != -1 {
		unit = string(in[strings.Index(in, "c")])
	} else if strings.Index(in, "f") != -1 {
		unit = string(in[strings.Index(in, "f")])
	} else {
		fmt.Println("You must supply an argument of format '32F' or '23C'.")
		os.Exit(1)
	}

	temp, err := strconv.ParseFloat(fmt.Sprintf("%v", string(strings.Split(in, unit)[0])), 64)

	if err != nil {
		panic(err)
	}

	converted = goconvtemps.ConvertTemps(temp, unit)
	fmt.Printf("Converted %v to %v.\n", strings.ToUpper(in), converted)
}
