package main

import "fmt"

func main() {
	products := make(map[string]float64, 10)

	products["magic"] = 100.10
	products["wealthy"] = 200.20

	fmt.Println("MAP Products size:", len(products))
	fmt.Println("Products Price:", products["magic"])
	fmt.Println("Products Price:", products["wealthy"])

	products_literal := map[string]float64{
		"magic2":   20.2,
		"wealthy2": 100.3,
		"keyboard": 0,
	}

	fmt.Println("MAP Literal Products size:", len(products_literal))
	fmt.Println("Literal Products Price:", products_literal["magic2"])
	fmt.Println("Literal Products Price:", products_literal["wealthy2"])

	if value, ok := products_literal["keyboard"]; ok {
		fmt.Println("Stored Value:", value)
	} else {
		fmt.Println("Not Stored Value")
	}

	for key, value := range products_literal {
		fmt.Println("Key:", key, "Value:", value)
	}

	delete(products_literal, "keyboard")

	if value, ok := products_literal["keyboard"]; ok {
		fmt.Println("Stored Value", value)
	} else {
		fmt.Println("Not Stored Value")
	}

}
