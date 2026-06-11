package main

import "fmt"

type arrInt [100]int

func digitPuluhan(n int) int {
	return (n / 10) % 10
}

func Insertion(a []int, n int) {
	var pass, i, temp int
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = a[pass]
		for i > 0 && ((digitPuluhan(temp) < digitPuluhan(a[i-1])) ||
			(digitPuluhan(temp) == digitPuluhan(a[i-1]) && temp < a[i-1])) {
			a[i] = a[i-1]
			i = i - 1
		}
		a[i] = temp
		pass = pass + 1
	}
}

func main() {
	var n, i int
	var a arrInt

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println("Data sebelum sorting:")
	for i = 0; i < n; i++ {
		fmt.Println(a[i])
	}

	Insertion(a[:n], n)

	fmt.Println("\nData setelah sorting:")
	for i = 0; i < n; i++ {
		fmt.Println(a[i])
	}
}