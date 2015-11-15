package main

import (
	"encoding/csv"
	"os"
	"strconv"

	"./src/redmart"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func pv(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func Products() []redmart.Product {
	var ps []redmart.Product

	f, err := os.Open("products.csv")
	check(err)
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	check(err)

	var p redmart.Product
	for _, v := range records {
		p.ID = pv(v[0])
		p.Price = pv(v[1])
		p.Length = pv(v[2])
		p.Width = pv(v[3])
		p.Height = pv(v[4])
		p.Weight = pv(v[5])
		ps = append(ps, p)
	}
	return ps
}

func main() {
	redmart.PackTote(redmart.Tote{45, 30, 35}, Products())
}
