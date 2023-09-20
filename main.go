package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/fwilhelm92/gha-demo/lib"
)

type MissingInput struct {
	input string
}

func (m *MissingInput) Error() string {
	return fmt.Sprintf("fehlender input '%v'", m.input)
}

func main() {
	if err := doMain(); err != nil {
		log.Fatalf("gha-demo failed: %v", err)
	}
}

func doMain() error {

	input1 := flag.String("zahl1", "", "1. Zahl")
	input2 := flag.String("zahl2", "", "2. Zahl")
	operator := flag.String("operator", "", "Operator (+, -, x, /)")

	flag.Parse()

	// Validation
	if *input1 == "" {
		return &MissingInput{"zahl1"}
	}
	zahl1, err := strconv.Atoi(*input1)
	if err != nil {
		return err
	}

	if *input2 == "" {
		return &MissingInput{"zahl2"}
	}
	zahl2, err := strconv.Atoi(*input2)
	if err != nil {
		return err
	}

	switch *operator {
	case "+":
		fmt.Println(lib.Sum(zahl1, zahl2))
	case "-":
		fmt.Println(lib.Sub(zahl1, zahl2))
	case "x":
		fmt.Println(lib.Mult(zahl1, zahl2))
	case "/":
		fmt.Println(lib.Div(zahl1, zahl2))
	default:
		return fmt.Errorf("operator ist nicht implementiert: %v", *operator)
	}

	return nil
}
