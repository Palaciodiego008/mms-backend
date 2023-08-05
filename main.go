package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var lambda, mu, s float64
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the value of lambda: ")
	if scanner.Scan() {
		lambda, _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	fmt.Print("Enter the value of mu: ")
	if scanner.Scan() {
		mu, _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	fmt.Print("Enter the value of s: ")
	if scanner.Scan() {
		s, _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	L, Lq, W, Wq := mmsModel(lambda, mu, s)

	fmt.Printf("Average Number of Customers in the System (L): %.6f\n", L)
	fmt.Printf("Average Number of Customers in the Queue (Lq): %.6f\n", Lq)
	fmt.Printf("Average Time a Customer Spends in the System (W): %.6f\n", W)
	fmt.Printf("Average Time a Customer Spends in the Queue (Wq): %.6f\n", Wq)

}
