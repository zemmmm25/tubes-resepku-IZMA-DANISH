package main

import "fmt"
const NMAX = 100 

type Peserta struct {
	id     string
	nama   string
	nilai  int
	durasi int
}
type TabPeserta [NMAX]Peserta

func Insertion(a []Peserta, n int) {
	var pass, i int
	var temp Peserta
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = a[pass]
		for i > 0 && ((temp.nilai > a[i-1].nilai) ||
			(temp.nilai == a[i-1].nilai && temp.durasi < a[i-1].durasi)) {
			a[i] = a[i-1]
			i = i - 1
		}
		a[i] = temp
		pass = pass + 1
	}
}
func main() {
	var n, i, total, count int
	var a TabPeserta
	var rata float64

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&a[i].id, &a[i].nama, &a[i].nilai, &a[i].durasi)
	}

	Insertion(a[:n], n)

	fmt.Println("Data setelah diurutkan:")
	for i = 0; i < n; i++ {
		fmt.Println(a[i].id, a[i].nama, a[i].nilai, a[i].durasi)
	}

	fmt.Println("\nPeserta terbaik:")
	fmt.Println(a[0].id, a[0].nama, a[0].nilai, a[0].durasi)

	var rata float64
	for i = 0; i < n; i++ {
		total = total + a[i].nilai
	}
	rata = float64(total) / float64(n)

	for i = 0; i < n; i++ {
		if float64(a[i].nilai) > rata {
			count = count + 1
		}
	}
	fmt.Println("\nJumlah peserta di atas rata-rata:", count)
}