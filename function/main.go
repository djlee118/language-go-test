package main

import "fmt"

// Swap 되지 않는 함수
func swapValues(oneValue, twoValue int) {
	fmt.Println("Before swap:", oneValue, twoValue)

	temp := oneValue
	oneValue = twoValue
	twoValue = temp

	fmt.Println("After swap:", oneValue, twoValue)
}

// 포인터를 사용해서 Swap 되는 함수
func swapPointerValues(oneValue, twoValue *int) {
	fmt.Println("Before swap:", *oneValue, *twoValue)

	temp := *oneValue
	*oneValue = *twoValue
	*twoValue = temp

	fmt.Println("After swap:", *oneValue, *twoValue)
}

func main() {
	oneValue, twoValue := 100, 500

	fmt.Println("swapValues Before calling Function", oneValue, twoValue)
	swapValues(oneValue, twoValue)
	fmt.Println("swapValues After calling Function", oneValue, twoValue)

	fmt.Println("swapPointerValues Before calling Function", oneValue, twoValue)
	swapPointerValues(&oneValue, &twoValue)
	fmt.Println("swapPointerValues After calling Function", oneValue, twoValue)

}
