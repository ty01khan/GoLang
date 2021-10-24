package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

// func addItems() bill {
// 	reader := bufio.Re
// }

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip) ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		p, _ := getInput("Item price: ", reader)
		price, err := strconv.ParseFloat(p, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, price)
		fmt.Println("Items added- ", name, price)
		fmt.Println()
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you saved the file - ", b.name)
	case "t":
		t, _ := getInput("Enter tip amount ($): ", reader)
		tip, err := strconv.ParseFloat(t, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(tip)
		fmt.Printf("Tip updated - %0.2f\n\n", tip)
		promptOptions(b)
	default:
		fmt.Printf("that was not a valid option...\n\n")
		promptOptions(b)
	}
}

func main() {
	// myBill := newBill("Talha's bill")
	// fmt.Println(myBill.format())

	// myBill.updateTip(1.24)
	// myBill.addItem("mango-icecream", 12.50)
	// fmt.Println("-------- Updated bill -------")
	// fmt.Println(myBill.format())

	b := createBill()
	promptOptions(b)
}
