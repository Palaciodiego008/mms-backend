package actions

import (
	"bufio"
	"fmt"
	"mms-project/internal"
	"net/http"
	"strconv"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var lambda, mu, s float64
	scanner := bufio.NewScanner(r.Body)

	fmt.Fprint(w, "Enter the value of lambda: ")
	if scanner.Scan() {
		lambda, _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	fmt.Fprint(w, "Enter the value of mu: ")
	if scanner.Scan() {
		mu, _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	fmt.Fprint(w, "Enter the value of s: ")
	if scanner.Scan() {
		s, _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	L, Lq, W, Wq := internal.MmsModel(lambda, mu, s)

	fmt.Fprintf(w, "Average Number of Customers in the System (L): %.6f\n", L)
	fmt.Fprintf(w, "Average Number of Customers in the Queue (Lq): %.6f\n", Lq)
	fmt.Fprintf(w, "Average Time a Customer Spends in the System (W): %.6f\n", W)
	fmt.Fprintf(w, "Average Time a Customer Spends in the Queue (Wq): %.6f\n", Wq)
}
