package redmart

var testPS = []Product{
	Product{1, 10, 1, 1, 1, 8},
	Product{1, 10, 1, 1, 1, 8},
	Product{2, 5, 1, 1, 1, 8},
	Product{3, 5, 1, 1, 1, 8},
	Product{4, 5, 1, 1, 1, 8},
	Product{5, 15, 1, 1, 1, 8},
	Product{6, 3, 1, 1, 1, 8},
	Product{7, 9, 1, 1, 1, 8},
	Product{999, 15, 1, 1, 1, 1},
	Product{9, 12, 1, 1, 1, 8},
	Product{10, 1, 1, 1, 1, 8},
}

// This will act both as test and example in documentation
func ExamplePackTote() {
	PackTote(Tote{1, 1, 5}, testPS)
	// Output:
	// Optimized value:  62
	// Sum of product IDs:  1015

}
