package main

import (
	"fmt"
	"io/ioutil"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"

	var total float64 = 0

	// list items
	for key, val := range b.items {
		fs += fmt.Sprintf("%-20v ...$%v \n", key+":", val)
		total += val
	}
	fs += fmt.Sprintf("%-20v ...$%v \n", "tip", b.tip)
	total += b.tip
	// total
	fs += fmt.Sprintf("-----------------------------\n%-20v ...$%0.2f \n", "total:", total)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())
	err := ioutil.WriteFile(""+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file")
}
