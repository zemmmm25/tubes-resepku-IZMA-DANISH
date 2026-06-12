package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const NMAX = 999

type dataresep struct {
	judulresep       string
	bahanMakanan     string
	kategoriMasakan  string
	komposisiBahan   string
	langkahPembuatan string
	estimasiWaktu    int
	jumlahDicari     int
	favorit          bool
	noAsli		 int
}

type tabresep [NMAX]dataresep

func main() {
	var resep tabresep
	var n int

	var bahanYgDicari string
	var caraPencarianData int
	var caraResepDikeluarkan int
	var pilihanKelola int
	var pilihanmenu int
	var angkafav int
	var reader *bufio.Reader
	var arah int
	var CaraCari int
	var temp tabresep

	reader = bufio.NewReader(os.Stdin)

	datadummy(&resep, &n)

	for {
		clearScreen()
		tampilanmenuUtama()
		fmt.Println("🍓 Apa yang mau anda lakukannn?? 🍓 𐔌՞ ܸ.ˬ.ܸ՞𐦯")
		fmt.Print("> masukan angkanya ya!! ₍ᵔ.˛.ᵔ₎ :  ")
		fmt.Scanln(&pilihanmenu)
		clearScreen()
		fmt.Println(" ")
		switch pilihanmenu {
		//═════════════════════════════════════════
		// input data
		//═════════════════════════════════════════
		case 1:
			input(&resep, &n, reader)
			clearScreen()
		//═════════════════════════════════════════
		// tampilan data
		//═════════════════════════════════════════
		case 2:
			if n == 0 {
				fmt.Println("╔══════════════════════════════════════════╗")
				fmt.Println("║      𐙚⋆🍓 Belum Ada Data Resep 🍰⋆𐙚      ║")
				fmt.Println("╠══════════════════════════════════════════╣")
				fmt.Println("║    Silakan tambahkan resep terlebih      ║")
				fmt.Println("║    dahulu yaa ₍ᐢ. .ᐢ₎♡                   ║")
				fmt.Println("╚══════════════════════════════════════════╝")
			} else {
				tampilanresep(&resep, n)
			}
			wait()
			clearScreen()
			//═════════════════════════════════════════ß
			// 	searching
			//═════════════════════════════════════════
		case 3:
			tampilansearching()
			fmt.Print("> masukan angkanya ya!! ₍ᵔ.˛.ᵔ₎ :  ")
			fmt.Scanln(&caraPencarianData)
			clearScreen()
			tampilanCara()
			fmt.Print("> masukan angkanya ya!! ₍ᵔ.˛.ᵔ₎ :  ")
			fmt.Scanln(&CaraCari)
			switch caraPencarianData {
			case 1:
				clearScreen()
				fmt.Println("╔══════════════════════════════════════════╗")
				fmt.Println("║    𐙚⋆🍓  Bahan makan utama 🍰⋆𐙚          ║")
				fmt.Println("╚══════════════════════════════════════════╝")
				fmt.Print("> bahan utama yang ingin dicari 𓂃۶ৎ : ")
				bahanYgDicari = inputString(reader)
				clearScreen()
			case 2:
				clearScreen()
				fmt.Println("╔══════════════════════════════════════════╗")
				fmt.Println("║         𐙚⋆🍓  judul makanan🍰⋆𐙚          ║")
				fmt.Println("╚══════════════════════════════════════════╝")
				fmt.Print("> judul makanan yang ingin dicari 𓂃۶ৎ : ")
				bahanYgDicari = inputString(reader)
			}

			// percabangan sequential dan binary/ judul dan durasi
			if caraPencarianData == 1 {
				if CaraCari == 1 {
					fmt.Println(" ")
					(pencarSequential(n, &resep, bahanYgDicari))
				} else {
					temp = resep
					sortBahanMakanan(&temp, n)
					fmt.Println(" ")
					binaryBahanMakanan(n, &temp, bahanYgDicari)
				}
			} else if caraPencarianData == 2 {
				if CaraCari == 1 {
					fmt.Println(" ")
					judulSequential(n, &resep, bahanYgDicari)
				} else {
					temp = resep
					sortAbjad(&temp, n)
					fmt.Println(" ")
					(pencarianBinary(n, &temp, bahanYgDicari))
				}
			}

			wait()
			clearScreen()
			//═════════════════════════════════════════
			// 	kelola data
			//═════════════════════════════════════════
		case 4:
			pilihanKelola = menukelola()

			if pilihanKelola == 1 {
				tambahData(&resep, &n, reader)
			} else if pilihanKelola == 2 {
				ubahData(&resep, n, reader)
			} else if pilihanKelola == 3 {
				hapusData(&resep, &n)
			} else if pilihanKelola == 4 {
				Statistik(resep, n)
				StatistikKategori(resep, n)
				wait()
			}
			clearScreen()
			//═════════════════════════════════════════
			// pengurutan data
			//═════════════════════════════════════════
		case 5:
			fmt.Println("╔══════════════════════════════════════════╗")
			fmt.Println("║        𐙚⋆🍓 Urutkan Resep 🍰⋆𐙚           ║")
			fmt.Println("╠══════════════════════════════════════════╣")
			fmt.Println("║  Berdasarkan apa resep ingin ditampilkan?║")
			fmt.Println("║  [1] Durasi              [2] judul       ║")
			fmt.Println("╚══════════════════════════════════════════╝")
			fmt.Print("> masukan angkanya ya!! ₍ᵔ.˛.ᵔ₎ : ")
			fmt.Scanln(&caraResepDikeluarkan)
			clearScreen()

			fmt.Println("╔══════════════════════════════════════════╗")
			fmt.Println("║         𐙚⋆🍓 Arah Urutan 🍰⋆𐙚            ║")
			fmt.Println("╠══════════════════════════════════════════╣")
			fmt.Println("║  [1] Ascending                           ║")
			fmt.Println("║  [2] Descending                          ║")
			fmt.Println("╚══════════════════════════════════════════╝")
			fmt.Print("> masukan angkanya ya!! ₍ᵔ.˛.ᵔ₎ : ")
			fmt.Scanln(&arah)
			// kondisi output
			if caraResepDikeluarkan == 1 {
				if arah == 1 {
					sortWaktu(&resep, n)
				} else {
					sortWaktuDesc(&resep, n)
				}
			} else if caraResepDikeluarkan == 2 {
				if arah == 1 {
					sortAbjad(&resep, n)
				} else {
					sortAbjadDesc(&resep, n)
				}
			}
			tampilanSorting(&resep, &n)
			wait()
			clearScreen()
			//═════════════════════════════════════════
			// 	favorit
			//═════════════════════════════════════════
		case 6:
			tampilanfavorit()
			fmt.Println("apa yang mau anda lakukan??: ")
			fmt.Scanln(&angkafav)
			switch angkafav {
			case 1:
				favorite(&resep, n)
			case 2:
				lihatfavori(resep, n)
				wait()
			}
			clearScreen()
			//═════════════════════════════════════════
			// 	keluar
			//═════════════════════════════════════════
		case 7:
			fmt.Println("terimakasihh telahhh menggunakan aplikasi resepku")
			wait()
			fmt.Println("keluar")
			return

		}
	}
}

// ═════════════════════════════════════════
// tampilan menu
// ═════════════════════════════════════════
func tampilanmenuUtama() {
	fmt.Println()
	fmt.Println("           🍓  ≽^• ˕ • ྀི≼ 🍰 ")
	fmt.Printf("        > !! APLIKASI RESEPKU!! <")
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║      ⋆.𐙚 ̊ 🍰 MENU UTAMA ૮₍˃̵֊ ˂̵₎ა 🍓      ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║ 𑣲⋆ [1] Input Data Resep                  ║")
	fmt.Println("║ 𑣲⋆ [2] Lihat Semua Resep                 ║")
	fmt.Println("║ 𑣲⋆ [3] Cari Resep                        ║")
	fmt.Println("║ 𑣲⋆ [4] Kelola Data                       ║")
	fmt.Println("║ 𑣲⋆ [5] Urutan Resep                      ║")
	fmt.Println("║ 𑣲⋆ [6] Kelola Favorit                    ║")
	fmt.Println("║ 𑣲⋆ [7] Keluar                            ║")
	fmt.Println("╚══════════════════════════════════════════╝")
}

// ═════════════════════════════════════════
// function case 1 "inputan"
// ═════════════════════════════════════════
func input(resep *tabresep, n *int, reader *bufio.Reader) {
	var i int
	var konfirmasi int
	var jumlah int
	var mulai int

	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║          𐙚⋆🍓 Input Data 🍰⋆𐙚            ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Print(" > Masukan jumlah data yang ingin diinput ya!! : ")

	fmt.Scanln(&jumlah)
	clearScreen()

	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║      𐙚⋆🍓 Konfirmasi Input 🍰⋆𐙚          ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Printf("║  Jumlah data yang akan diinput : %-3d     ║\n", jumlah)
	fmt.Println("║  [1] Ya          [2] Tidak               ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Print("> masukan angkanya ya!! ₍ᵔ.˛.ᵔ₎ :  ")

	fmt.Scanln(&konfirmasi)
	clearScreen()

	if konfirmasi != 1 {
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║      𐙚⋆🍓 Input Dibatalkan 🍰⋆𐙚          ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║     Tidak ada data yang ditambahkan      ║")
		fmt.Println("║                ૮◞ ‸ ◟ ა                  ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		fmt.Println()
		wait()
		return
	}
	if *n+jumlah > NMAX {
	fmt.Println("Data melebihi kapasitas penyimpanan")
	wait()
	return
}

	mulai = *n
	*n = *n + jumlah
	clearScreen()

	fmt.Println(" ")
	fmt.Println("             🍓  ≽^• ˕ • ྀི≼ 🍰 ")
	fmt.Printf("         > !! MASUKAN DATANYA YA!! <")
	fmt.Println()
	for i = mulai; i < *n; i++ {
		fmt.Println("══════════════════════════════════════════")
		fmt.Printf("        𐙚🍓 Resep ke-%-2d 🍰⋆𐙚            \n", i+1)
		fmt.Println("══════════════════════════════════════════")
		fmt.Print("  📖 Judul Resep       : ")
		resep[i].judulresep = inputString(reader)

		fmt.Print("  🥬 Bahan Makanan     : ")
		resep[i].bahanMakanan = inputString(reader)

		fmt.Print("  🏷️  Kategori Masakan  : ")
		resep[i].kategoriMasakan = inputString(reader)

		fmt.Print("  🧄 Komposisi Bahan   : ")
		resep[i].komposisiBahan = inputString(reader)

		fmt.Print("  📝 Langkah Pembuatan : ")
		resep[i].langkahPembuatan = inputString(reader)

		fmt.Print("  ⏰ Estimasi Waktu    : ")
		fmt.Scanln(&resep[i].estimasiWaktu)
		reader.ReadString('\n')

		resep[i].noAsli = i + 1

		wait()
	}
}

// ═════════════════════════════════════════
// function case 2 "tampilan resep"
// ═════════════════════════════════════════
func tampilanresep(resep *tabresep, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println("══════════════════════════════════════════")
		fmt.Printf("        𐙚⋆🍓 Resep ke-%-2d 🍰⋆𐙚            \n", i+1)
		fmt.Println("══════════════════════════════════════════")
		fmt.Printf("  📖 Judul Resep       : %s\n", (*resep)[i].judulresep)
		fmt.Printf("  🥬 Bahan Makanan     : %s\n", (*resep)[i].bahanMakanan)
		fmt.Printf("  🏷️  Kategori Masakan  : %s\n", (*resep)[i].kategoriMasakan)
		fmt.Printf("  🧄 Komposisi Bahan   : %s\n", (*resep)[i].komposisiBahan)
		fmt.Printf("  📝 Langkah Pembuatan : %s\n", (*resep)[i].langkahPembuatan)
		fmt.Printf("  ⏰ Estimasi Waktu    : %d menit\n", (*resep)[i].estimasiWaktu)
		fmt.Println()
	}
}

// ═════════════════════════════════════════
// function case 3 "searching"
// ═════════════════════════════════════════
func tampilansearching() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║        𐙚⋆🍓 Cari Resep 🍰⋆𐙚              ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Ingin berdasarkan apa data dicari?      ║")
	fmt.Println("║  [1] Bahan makanan utama                 ║")
	fmt.Println("║  [2] Judul                               ║")
	fmt.Println("╚══════════════════════════════════════════╝")

}
func tampilanCara() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║ 𐙚⋆🍓 dengan cara apa data dicari?? 🍰⋆𐙚  ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Ingin berdasarkan apa data dicari?      ║")
	fmt.Println("║  [1] sequential.                         ║")
	fmt.Println("║  [2] binary.                             ║")
	fmt.Println("╚══════════════════════════════════════════╝")
}
func pencarSequential(n int, resep *tabresep, bahanYgDicari string) {
	var i int
	var found bool

	fmt.Println("               🍓  ≽^• ˕ • ྀི≼ 🍰 ")
	fmt.Printf("         Hasil Pencarian Bahan: %s\n", bahanYgDicari)

	for i = 0; i < n; i++ {
		if toLower(bahanYgDicari) == toLower((*resep)[i].bahanMakanan) {
			found = true
			//══════════════════════════════════════════════
			//══════════════════════════════════════════════
			fmt.Println("══════════════════════════════════════════")
			fmt.Printf("         𐙚⋆🍓 Resep ke-%-2d 🍰⋆𐙚           \n", i+1)
			fmt.Println("══════════════════════════════════════════")
			fmt.Printf("  📖 Judul Resep       : %s\n", (*resep)[i].judulresep)
			fmt.Printf("  🥬 Bahan Makanan     : %s\n", (*resep)[i].bahanMakanan)
			fmt.Printf("  🏷️  Kategori Masakan  : %s\n", (*resep)[i].kategoriMasakan)
			fmt.Printf("  🧄 Komposisi Bahan   : %s\n", (*resep)[i].komposisiBahan)
			fmt.Printf("  📝 Langkah Pembuatan : %s\n", (*resep)[i].langkahPembuatan)
			fmt.Printf("  ⏰ Estimasi Waktu    : %d menit\n", (*resep)[i].estimasiWaktu)
			fmt.Println()
		}
	}
	if found == false {

		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║      𐙚⋆🍓 Resep Tidak Ditemukan 🍰⋆𐙚     ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║    Maaf yaa, datanya tidak tersedia      ║")
		fmt.Println("║                ૮◞ ‸ ◟ ა                  ║")
		fmt.Println("╚══════════════════════════════════════════╝")
	}
}
func judulSequential(n int, resep *tabresep, bahanYgDicari string) {
	var i int
	var found bool

	for i = 0; i < n; i++ {
		if toLower(bahanYgDicari) == toLower((*resep)[i].judulresep) {
			found = true
			fmt.Println("══════════════════════════════════════════")
			fmt.Printf("         𐙚⋆🍓 Resep ke-%-2d 🍰⋆𐙚           \n", i+1)
			fmt.Println("══════════════════════════════════════════")
			fmt.Printf("  📖 Judul Resep       : %s\n", (*resep)[i].judulresep)
			fmt.Printf("  🥬 Bahan Makanan     : %s\n", (*resep)[i].bahanMakanan)
			fmt.Printf("  🏷️  Kategori Masakan  : %s\n", (*resep)[i].kategoriMasakan)
			fmt.Printf("  🧄 Komposisi Bahan   : %s\n", (*resep)[i].komposisiBahan)
			fmt.Printf("  📝 Langkah Pembuatan : %s\n", (*resep)[i].langkahPembuatan)
			fmt.Printf("  ⏰ Estimasi Waktu    : %d menit\n", (*resep)[i].estimasiWaktu)
			fmt.Println()
		}
	}
	if found == false {

		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║      𐙚⋆🍓 Resep Tidak Ditemukan 🍰⋆𐙚     ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║    Maaf yaa, datanya tidak tersedia      ║")
		fmt.Println("║                ૮◞ ‸ ◟ ა                  ║")
		fmt.Println("╚══════════════════════════════════════════╝")
	}
}
func pencarianBinary(n int, resep *tabresep, bahanYgDicari string) {
	var left, right, mid int
	var found int
	var target string
	var i int
	left = 0
	right = n - 1
	found = -1
	target = toLower(bahanYgDicari)
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if target > toLower((*resep)[mid].judulresep) {
			left = mid + 1
		} else if target < toLower((*resep)[mid].judulresep) {
			right = mid - 1
		} else {
			found = mid
		}
	}
	if found != -1 {

	// cari data pertama yang sama
	i = found
	for i > 0 && toLower((*resep)[i-1].bahanMakanan) == target {
		i--
	}

	// tampilkan semua data yang sama
	for i < n && toLower((*resep)[i].bahanMakanan) == target {

		(*resep)[i].jumlahDicari++

		fmt.Println("══════════════════════════════════════════")
		fmt.Printf("         𐙚⋆🍓 Resep ke-%-2d 🍰⋆𐙚           \n",
			(*resep)[i].noAsli)
		fmt.Println("══════════════════════════════════════════")
		fmt.Printf("  📖 Judul Resep       : %s\n", (*resep)[i].judulresep)
		fmt.Printf("  🥬 Bahan Makanan     : %s\n", (*resep)[i].bahanMakanan)
		fmt.Printf("  🏷️  Kategori Masakan  : %s\n", (*resep)[i].kategoriMasakan)
		fmt.Printf("  🧄 Komposisi Bahan   : %s\n", (*resep)[i].komposisiBahan)
		fmt.Printf("  📝 Langkah Pembuatan : %s\n", (*resep)[i].langkahPembuatan)
		fmt.Printf("  ⏰ Estimasi Waktu    : %d menit\n", (*resep)[i].estimasiWaktu)
		fmt.Println()

		i++
	}

} else {

		clearScreen()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║      𐙚⋆🍓 Resep Tidak Ditemukan 🍰⋆𐙚     ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║    Maaf yaa, datanya tidak tersedia      ║")
		fmt.Println("║                ૮◞ ‸ ◟ ა                  ║")
		fmt.Println("╚══════════════════════════════════════════╝")
	}
}
func binaryBahanMakanan(n int, resep *tabresep, bahanYgDicari string) {
	var found int
	var left, right, mid int
	var target string
	var i int
	left = 0
	right = n - 1
	found = -1
	target = toLower(bahanYgDicari)
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if target > toLower((*resep)[mid].bahanMakanan) {
			left = mid + 1
		} else if target < toLower((*resep)[mid].bahanMakanan) {
			right = mid - 1
		} else {
			found = mid
		}
	}
	if found != -1 {

	// cari data pertama yang sama
	i = found
	for i > 0 && toLower((*resep)[i-1].bahanMakanan) == target {
		i--
	}

	// tampilkan semua data yang sama
	for i < n && toLower((*resep)[i].bahanMakanan) == target {

		(*resep)[i].jumlahDicari++

		fmt.Println("══════════════════════════════════════════")
		fmt.Printf("         𐙚⋆🍓 Resep ke-%-2d 🍰⋆𐙚           \n",
			(*resep)[i].noAsli)
		fmt.Println("══════════════════════════════════════════")
		fmt.Printf("  📖 Judul Resep       : %s\n", (*resep)[i].judulresep)
		fmt.Printf("  🥬 Bahan Makanan     : %s\n", (*resep)[i].bahanMakanan)
		fmt.Printf("  🏷️  Kategori Masakan  : %s\n", (*resep)[i].kategoriMasakan)
		fmt.Printf("  🧄 Komposisi Bahan   : %s\n", (*resep)[i].komposisiBahan)
		fmt.Printf("  📝 Langkah Pembuatan : %s\n", (*resep)[i].langkahPembuatan)
		fmt.Printf("  ⏰ Estimasi Waktu    : %d menit\n", (*resep)[i].estimasiWaktu)
		fmt.Println()

		i++
	}

} else {

		clearScreen()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║      𐙚⋆🍓 Resep Tidak Ditemukan 🍰⋆𐙚     ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║    Maaf yaa, datanya tidak tersedia      ║")
		fmt.Println("║                ૮◞ ‸ ◟ ა                  ║")
		fmt.Println("╚══════════════════════════════════════════╝")
	}
}

// ═════════════════════════════════════════
// function case 4 "kelola data"
// ═════════════════════════════════════════
func menukelola() int {
	var pilihanKelola int
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║        𐙚⋆🍓 MENU KELOLA 🍰⋆𐙚             ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  [1] Tambah Data Resep                   ║")
	fmt.Println("║  [2] Ubah Data Resep                     ║")
	fmt.Println("║  [3] Hapus Data Resep                    ║")
	fmt.Println("║  [4] Lihat Statistik                     ║")
	fmt.Println("║  [5] Lanjut ke Pencarian                 ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Print("> masukan angkanya ya!! ₍ᵔ.˛.ᵔ₎ : ")

	fmt.Scanln(&pilihanKelola)
	fmt.Println(" ")

	return pilihanKelola

}

// func tambah data
func tambahData(resep *tabresep, n *int, reader *bufio.Reader) {
	var i int
	if *n >= NMAX {
		fmt.Println(" Maaf ya sayang, kapasitas penyimpanan resep sudah penuh! 𐔌՞ ܸ.ˬ.ܸ՞𐦯")
		return
	}
	clearScreen()
	fmt.Println("══════════════════════════════════════════")
	fmt.Println("      𐙚⋆🍓 Tambah Data Resep 🍰⋆𐙚         ")
	fmt.Println("══════════════════════════════════════════")

	i = *n
	fmt.Print("  📖 Judul Resep       : ")
	resep[i].judulresep = inputString(reader)

	fmt.Print("  🥬 Bahan Makanan     : ")
	resep[i].bahanMakanan = inputString(reader)

	fmt.Print("  🏷️  Kategori Masakan  : ")
	resep[i].kategoriMasakan = inputString(reader)

	fmt.Print("  🧄 Komposisi Bahan   : ")
	resep[i].komposisiBahan = inputString(reader)

	fmt.Print("  📝 Langkah Pembuatan : ")
	resep[i].langkahPembuatan = inputString(reader)

	fmt.Print("  ⏰ Estimasi Waktu    : ")
	fmt.Scan(&resep[i].estimasiWaktu)
	
	fmt.Println(" ")
	clearScreen()

	*n = *n + 1
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║   𐙚⋆🍓 Data Berhasil Ditambahkan 🍰⋆𐙚    ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║   Yeay!! resep baru berhasil disimpan ♡  ║")
	fmt.Println("║          ૮₍ ˶ᵔ ᵕ ᵔ˶ ₎ა                   ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	wait()
	fmt.Println(" ")
}
func tampilanubahdata() int {
	var pilihubah int
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║        𐙚⋆🍓 Mau Ubah Apa? 🍰⋆𐙚           ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║ 𑣲⋆ [1] Judul Resep                       ║")
	fmt.Println("║ 𑣲⋆ [2] Bahan Makanan                     ║")
	fmt.Println("║ 𑣲⋆ [3] Kategori                          ║")
	fmt.Println("║ 𑣲⋆ [4] Komposisi Bahan                   ║")
	fmt.Println("║ 𑣲⋆ [5] Langkah Pembuatan                 ║")
	fmt.Println("║ 𑣲⋆ [6] Estimasi Waktu                    ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Print("> pilih angkanya ya!! ₍ᵔ.˛.ᵔ₎ : ")
	fmt.Scanln(&pilihubah)

	return pilihubah

}

// func ubah data
func ubahData(resep *tabresep, n int, reader *bufio.Reader) {
	var idx int
	var konfirmasi int
	var pilihubah int
	var temp dataresep

	clearScreen()
	tampilanresep(resep, n)

	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║       🍓 Ubah Data Resep 🍰              ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Masukan nomor resep yang mau diubah     ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Print("> masukan angkanya ya!! ₍ᵔ.˛.ᵔ₎ : ")

	fmt.Scanln(&idx)
	idx = idx - 1

	if idx < 0 || idx >= n {
		clearScreen()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║      𐙚⋆🍓 Input Tidak Valid 🍰⋆𐙚         ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║  Maaf yaa, nomor resep tidak ditemukan   ║")
		fmt.Println("║          ૮◞ ﻌ ◟ ა                        ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		wait()
		return
	}
	temp = resep[idx]

	clearScreen()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║        𐙚⋆🍓 Konfirmasi Ubah 🍰⋆𐙚         ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Printf("║  Ubah resep: %-27s ║\n", resep[idx].judulresep)
	fmt.Println("║  [1] Ya           [2] Tidak              ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Print("masukan angkanya ya!! ₍ᵔ.˛.ᵔ₎ :  ")

	fmt.Scanln(&konfirmasi)

	if konfirmasi != 1 {
		clearScreen()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║      𐙚⋆🍓 Update Dibatalkan 🍰⋆𐙚         ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║    Tidak ada perubahan yang disimpan     ║")
		fmt.Println("║                 𐔌՞ ܸ.ˬ.ܸ՞𐦯                  ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		wait()
		return
	} else {
		clearScreen()
		fmt.Println("═════════════════════════════════════════")
		fmt.Println("     ✦ Masukkan data baru yaaa!! ♡ ♡ ♡   ")
		fmt.Println("═════════════════════════════════════════")

		temp = resep[idx]
		pilihubah = tampilanubahdata()
		clearScreen()
		switch pilihubah {
		case 1:
			fmt.Print("  📖 Judul Resep       : ")
			temp.judulresep = inputString(reader)
		case 2:
			fmt.Print("  🥬 Bahan Makanan     : ")
			temp.bahanMakanan = inputString(reader)
		case 3:
			fmt.Print("  🏷️  Kategori Masakan  : ")
			temp.kategoriMasakan = inputString(reader)
		case 4:
			fmt.Print("  🧄 Komposisi Bahan   : ")
			temp.komposisiBahan = inputString(reader)
		case 5:
			fmt.Print("  📝 Langkah Pembuatan : ")
			temp.langkahPembuatan = inputString(reader)
		case 6:
			fmt.Print("  ⏰ Estimasi Waktu    : ")
			fmt.Scanln(&temp.estimasiWaktu)
		}

		fmt.Println(" ")
		clearScreen()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║      𐙚⋆🍓 Konfirmasi Perubahan 🍰⋆𐙚      ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Printf("║  Ubah resep: %-27s ║\n", resep[idx].judulresep)
		fmt.Println("║  [1] Ya      [2] Tidak (Undo)            ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		fmt.Print("> masukan angkanya ya!! ₍ᵔ.˛.ᵔ₎ : ")

		fmt.Scanln(&konfirmasi)

		clearScreen()
		if konfirmasi == 1 {
			resep[idx] = temp // Data asli baru ditimpa di sini kalau setuju
			fmt.Println("╔══════════════════════════════════════════╗")
			fmt.Println("║      𐙚⋆🍓 Update Berhasil! 🍰⋆𐙚          ║")
			fmt.Println("╠══════════════════════════════════════════╣")
			fmt.Println("║     Resep berhasil diperbarui yaa ✨     ║")
			fmt.Println("║                ₍ᐢ. .ᐢ₎♡                  ║")
			fmt.Println("╚══════════════════════════════════════════╝")
			wait()

		} else {
			fmt.Println("╔══════════════════════════════════════════╗")
			fmt.Println("║      𐙚⋆🍓 Update Dibatalkan 🍰⋆𐙚         ║")
			fmt.Println("╠══════════════════════════════════════════╣")
			fmt.Println("║    Tidak ada perubahan yang disimpan     ║")
			fmt.Println("║                 𐔌՞ ܸ.ˬ.ܸ՞𐦯                  ║")
			fmt.Println("╚══════════════════════════════════════════╝")
			wait()
		}
		fmt.Println()
	}
}

// func hapus data
func hapusData(resep *tabresep, n *int) {
	var idx int
	var konfirmasi int
	var i int
	clearScreen()
	tampilanresep(resep, *n)
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║       𐙚⋆🍓 Hapus Data Resep 🍰⋆𐙚         ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Printf("║  Masukan nomor resep yang mau dihapus    ║\n")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Print("> masukan angkanya ya!! ₍ᵔ.˛.ᵔ₎ : ")
	fmt.Print(" ")

	fmt.Scanln(&idx)

	idx = idx - 1
	if idx < 0 || idx >= *n {

		clearScreen()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║      𐙚⋆🍓 Input Tidak Valid 🍰⋆𐙚         ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║  Maaf yaa, nomor resep tidak ditemukan   ║")
		fmt.Println("║          ૮◞ ﻌ ◟ ა                        ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		wait()

		return
	}
	clearScreen()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║       𐙚⋆🍓 Hapus Data Resep 🍰⋆𐙚         ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Printf("║  Nomor Resep: %-26d ║\n", idx+1)
	fmt.Printf("║  Judul      : %-26s ║\n", resep[idx].judulresep)
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Yakin ingin menghapus resep ini?        ║")
	fmt.Println("║  [1] Ya          [2] Tidak               ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Print("> masukan angkanya ya!! ₍ᵔ.˛.ᵔ₎ : ")

	fmt.Scanln(&konfirmasi)

	if konfirmasi != 1 {

		clearScreen()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║     𐙚⋆🍓 Hapus Dibatalkan 🍰⋆𐙚           ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║  Data resep tidak jadi dihapus yaa ♡     ║")
		fmt.Println("║                𐔌՞ ܸ.ˬ.ܸ՞𐦯                  ║")
		fmt.Println("╚━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╝")
		wait()

		return // ← batalkan, tidak jadi hapus
	}
	for i = idx; i < *n-1; i++ {
		resep[i] = resep[i+1]
	}
	*n = *n - 1
	clearScreen()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║    𐙚⋆🍓 Data Berhasil Dihapus 🍰⋆𐙚       ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Resep berhasil dihapus dari daftar ♡    ║")
	fmt.Println("║          ₍ᐢ. .ᐢ₎♡                        ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	wait()

	fmt.Println(" ")
}

// statistik nya
func Statistik(resep tabresep, n int) {
	if n == 0 {
		fmt.Println(" Belum ada data resep")
		return
	}
	var totalWaktu, minWaktu, maxWaktu int
	var rataRata float64
	var maxDicari, idxMax, i int
	minWaktu = resep[0].estimasiWaktu
	maxWaktu = resep[0].estimasiWaktu

	for i = 0; i < n; i++ {
		totalWaktu = totalWaktu + resep[i].estimasiWaktu
		if resep[i].estimasiWaktu < minWaktu {
			minWaktu = resep[i].estimasiWaktu
		}
		if resep[i].estimasiWaktu > maxWaktu {
			maxWaktu = resep[i].estimasiWaktu
		}
	}
	rataRata = float64(totalWaktu) / float64(n)

	fmt.Println("══════════════════════════════════════════")
	fmt.Println("     𐙚⋆🍓 Tabel Statistik Resep 🍰⋆𐙚     ")
	fmt.Println("══════════════════════════════════════════")
	fmt.Printf("  🍓 Total Resep          : %d resep\n", n)
	fmt.Printf("  ⏰ Total Estimasi Waktu : %d menit\n", totalWaktu)
	fmt.Printf("  📊 Rata-rata Waktu      : %.2f menit\n", rataRata)
	fmt.Printf("  ⚡ Waktu Tercepat       : %d menit\n", minWaktu)
	fmt.Printf("  🔥 Waktu Terlama        : %d menit\n", maxWaktu)
	// ← tambah ini sebelum if!
	maxDicari = 0
	idxMax = 0
	for i = 0; i < n; i++ {
		if resep[i].jumlahDicari > maxDicari {
			maxDicari = resep[i].jumlahDicari
			idxMax = i
		}
	}

	if maxDicari > 0 {
		fmt.Printf("  🔍 Paling sering dicari : %s (%dx)\n",
			resep[idxMax].judulresep, maxDicari)
	} else {
		fmt.Println("  🔍 Paling sering dicari : Belum ada yang dicari")
	}

	fmt.Println()
}

func StatistikKategori(resep tabresep, n int) {
	var namaKategori [NMAX]string
	var jumlahKategori [NMAX]int
	var totalKategori int
	var i, j int
	var found bool
	for i = 0; i < n; i++ {
		found = false
		for j = 0; j < totalKategori; j++ {
			if resep[i].kategoriMasakan == namaKategori[j] {
				jumlahKategori[j]++
				found = true
			}
		}
		if found == false {
			namaKategori[totalKategori] = resep[i].kategoriMasakan
			jumlahKategori[totalKategori] = 1
			totalKategori++
		}

	}
	fmt.Println("══════════════════════════════════════════")
	fmt.Println("     𐙚⋆🍓 Statistik Per Kategori 🍰⋆𐙚     ")
	fmt.Println("══════════════════════════════════════════")

	for i = 0; i < totalKategori; i++ {
		fmt.Printf("  🏷️  %-20s : %d resep\n",
			namaKategori[i], jumlahKategori[i])
	}
	fmt.Println()
}

// ═════════════════════════════════════════
// function case 5 "pengurutan"
// ═════════════════════════════════════════
func sortBahanMakanan(resep *tabresep, n int) {
	var i, j int
	var temp dataresep
	for i = 1; i < n; i++ {
		temp = resep[i]
		j = i - 1
		for j >= 0 && toLower(resep[j].bahanMakanan) > toLower(temp.bahanMakanan) {
			resep[j+1] = resep[j]
			j = j - 1
		}
		resep[j+1] = temp
	}
}

func sortWaktu(resep *tabresep, n int) {
	var i, j, min int
	var temp dataresep
	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if resep[j].estimasiWaktu < resep[min].estimasiWaktu {
				min = j

			}
		}
		temp = resep[i]
		resep[i] = resep[min]
		resep[min] = temp
	}
}
func sortWaktuDesc(resep *tabresep, n int) {
	var i, j, max int
	var temp dataresep
	for i = 0; i < n-1; i++ {
		max = i
		for j = i + 1; j < n; j++ {
			if resep[j].estimasiWaktu > resep[max].estimasiWaktu {
				max = j

			}
		}
		temp = resep[i]
		resep[i] = resep[max]
		resep[i] = temp
	}
}

// sorting berdasarkan bahan makanan ascending A ampe Z
func sortAbjad(resep *tabresep, n int) {
	var i, j int
	var temp dataresep
	for i = 1; i < n; i++ {
		temp = resep[i]
		j = i - 1
		for j >= 0 && toLower(resep[j].judulresep) > toLower(temp.judulresep) {
			resep[j+1] = resep[j]
			j = j - 1
		}
		resep[j+1] = temp
	}
}
func sortAbjadDesc(resep *tabresep, n int) {
	var i, j int
	var temp dataresep
	for i = 1; i < n; i++ {
		temp = resep[i]
		j = i - 1
		for j >= 0 && toLower(resep[j].judulresep) < toLower(temp.judulresep) {
			resep[j+1] = resep[j]
			j = j - 1
		}
		resep[j+1] = temp
	}
}
func tampilanSorting(resep *tabresep, n *int) {
	var i int
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║    𐙚⋆🍓 Hasil Setelah Diurutkan 🍰⋆𐙚     ║")
	fmt.Println("╚══════════════════════════════════════════╝")

	for i = 0; i < *n; i++ {

		fmt.Println("══════════════════════════════════════════")
		fmt.Printf("        𐙚⋆🍓 Resep ke-%-2d 🍰⋆𐙚            \n", i+1)
		fmt.Println("══════════════════════════════════════════")
		fmt.Printf("  📖 Judul Resep       : %s\n", resep[i].judulresep)
		fmt.Printf("  🥬 Bahan Makanan     : %s\n", resep[i].bahanMakanan)
		fmt.Printf("  🏷️  Kategori Masakan  : %s\n", resep[i].kategoriMasakan)
		fmt.Printf("  🧄 Komposisi Bahan   : %s\n", resep[i].komposisiBahan)
		fmt.Printf("  📝 Langkah Pembuatan : %s\n", resep[i].langkahPembuatan)
		fmt.Printf("  ⏰ Estimasi Waktu    : %d menit\n", resep[i].estimasiWaktu)
		fmt.Println()
	}
}

// ═════════════════════════════════════════
// function case 6 "favorite"
// ═════════════════════════════════════════
func tampilanfavorit() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║      𐙚⋆🍓 Menu Favorit 🍰⋆𐙚              ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Apa yang ingin anda lakukan?            ║")
	fmt.Println("║  [1] Tambah Resep Favorit                ║")
	fmt.Println("║  [2] Lihat Semua Favorit                 ║")
	fmt.Println("╚══════════════════════════════════════════╝")
}
func favorite(resep *tabresep, n int) {
	var pilihfav int
	tampilanresep(resep, n)
	fmt.Print("> resep nomer berapa yang ingin ditambahkan ke favorit?? : ")
	fmt.Scanln(&pilihfav)

	pilihfav--

	if pilihfav >= 0 && pilihfav < n {
		resep[pilihfav].favorit = true
		clearScreen()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║      𐙚⋆🍓 Resep Favorit! 🍰⋆𐙚            ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║    Yeay! Resep masuk daftar favorit ♡    ║")
		fmt.Println("║               ( ˶ˆᗜˆ˵ )                  ║")
		fmt.Println("╚══════════════════════════════════════════╝")
	} else {
		clearScreen()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║      𐙚⋆🍓 Gagal Menambah Favorit 🍰⋆𐙚    ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║    Nomor resep tidak ditemukan yaa       ║")
		fmt.Println("║               ૮◞ ﻌ ◟ ა                   ║")
		fmt.Println("╚══════════════════════════════════════════╝")
	}
	wait()
}
func lihatfavori(resep tabresep, n int) {
	var i int
	var found bool
	clearScreen()
	for i = 0; i < n; i++ {
		if resep[i].favorit == true {
			found = true
			fmt.Println("             🍓  ≽^• ˕ • ྀི≼ 🍰 ")
			fmt.Printf("         > !! RESEP FAVORITE MUU!! <")
			fmt.Println()
			fmt.Println("══════════════════════════════════════════")
			fmt.Printf("        𐙚⋆🍓 Resep ke-%-2d 🍰⋆𐙚            \n", i+1)
			fmt.Println("══════════════════════════════════════════")
			fmt.Printf("  📖 Judul Resep       : %s\n", resep[i].judulresep)
			fmt.Printf("  🥬 Bahan Makanan     : %s\n", resep[i].bahanMakanan)
			fmt.Printf("  🏷️  Kategori Masakan  : %s\n", resep[i].kategoriMasakan)
			fmt.Printf("  🧄 Komposisi Bahan   : %s\n", resep[i].komposisiBahan)
			fmt.Printf("  📝 Langkah Pembuatan : %s\n", resep[i].langkahPembuatan)
			fmt.Printf("  ⏰ Estimasi Waktu    : %d menit\n", resep[i].estimasiWaktu)
			fmt.Println()
		}

	}
	if !found {
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║      𐙚⋆🍓 Belum Ada Favorit 🍰⋆𐙚         ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║   Kamu belum menambahkan resep favorit   ║")
		fmt.Println("║               ❤️ ૮₍˶ᵔ ᵕ ᵔ˶₎ა             ║")
		fmt.Println("╚══════════════════════════════════════════╝")
	}
}

func inputString(reader *bufio.Reader) string {
	var text string
	text = ""
	for text == "" {
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)
	}
	return text
}

func clearScreen() {
	var cmd *exec.Cmd              //var untuk menyimpan perintah terminal
	if runtime.GOOS == "windows" { //ini kalau user menggunakan windows
		cmd = exec.Command("cmd", "/c", "cls") //jika true maka akan menjalankan perintah "cls"
	} else { //jika osnya Mac atau Linux
		cmd = exec.Command("clear") //maka akan menjalankan perintah "clear"
	}
	cmd.Stdout = os.Stdout //untuk menyambungkan command ke terminal program
	cmd.Run()              //ini agar perintah tadi langsung di eksekusi
}

func wait() {
	fmt.Println("kalo mau lanjut, tekan enter yaa!! ₍ᐢ..ᐢ₎ ")
	fmt.Scanln()
}

func toLower(s string) string {
	var i int
	var res string

	i = 0
	res = ""
	for i < len(s) {
		if s[i] >= 'A' && s[i] <= 'Z' {
			res = res + string(s[i]+32)
		} else {
			res = res + string(s[i])
		}
		i++
	}
	return res
}
func datadummy(resep *tabresep, n *int) {
	*n = 10

	resep[0] = dataresep{
		noAsli: 1,
		judulresep:       "Nasi Goreng Spesial",
		bahanMakanan:     "Nasi",
		kategoriMasakan:  "Makanan Utama",
		komposisiBahan:   "Nasi, telur, kecap, bawang",
		langkahPembuatan: "Tumis bumbu lalu masukkan nasi",
		estimasiWaktu:    20,
	}

	resep[1] = dataresep{
		noAsli: 2,
		judulresep:       "Ayam Goreng Krispi",
		bahanMakanan:     "Ayam",
		kategoriMasakan:  "Makanan Utama",
		komposisiBahan:   "Ayam, tepung, garam",
		langkahPembuatan: "Balur ayam lalu goreng",
		estimasiWaktu:    30,
	}

	resep[2] = dataresep{
		noAsli: 3,
		judulresep:       "Sup Jagung",
		bahanMakanan:     "Jagung",
		kategoriMasakan:  "Sup",
		komposisiBahan:   "Jagung, telur, wortel",
		langkahPembuatan: "Rebus semua bahan hingga matang",
		estimasiWaktu:    25,
	}

	resep[3] = dataresep{
		noAsli: 4,
		judulresep:       "Mie Goreng Jawa",
		bahanMakanan:     "Mie",
		kategoriMasakan:  "Makanan Utama",
		komposisiBahan:   "Mie, kol, kecap, telur",
		langkahPembuatan: "Tumis bumbu lalu masak mie",
		estimasiWaktu:    15,
	}

	resep[4] = dataresep{
		noAsli: 5,
		judulresep:       "Puding Coklat",
		bahanMakanan:     "Coklat",
		kategoriMasakan:  "Dessert",
		komposisiBahan:   "Susu, coklat bubuk, agar-agar",
		langkahPembuatan: "Rebus lalu dinginkan",
		estimasiWaktu:    40,
	}

	resep[5] = dataresep{
		noAsli: 6,
		judulresep:       "Es Buah Segar",
		bahanMakanan:     "Buah",
		kategoriMasakan:  "Minuman",
		komposisiBahan:   "Melon, semangka, sirup",
		langkahPembuatan: "Campur semua bahan",
		estimasiWaktu:    10,
	}

	resep[6] = dataresep{
		noAsli: 7,
		judulresep:       "Sate Ayam Madura",
		bahanMakanan:     "Ayam",
		kategoriMasakan:  "Makanan Utama",
		komposisiBahan:   "Ayam, kecap, bumbu kacang",
		langkahPembuatan: "Bakar ayam lalu sajikan",
		estimasiWaktu:    45,
	}

	resep[7] = dataresep{
		noAsli: 8,
		judulresep:       "Tumis Kangkung",
		bahanMakanan:     "Kangkung",
		kategoriMasakan:  "Sayuran",
		komposisiBahan:   "Kangkung, bawang putih",
		langkahPembuatan: "Tumis hingga layu",
		estimasiWaktu:    12,
	}

	resep[8] = dataresep{
		noAsli: 9,
		judulresep:       "Bakso Kuah",
		bahanMakanan:     "Daging Sapi",
		kategoriMasakan:  "Sup",
		komposisiBahan:   "Bakso, mie, daun bawang",
		langkahPembuatan: "Rebus hingga matang",
		estimasiWaktu:    35,
	}

	resep[9] = dataresep{
		noAsli: 10,
		judulresep:       "Pancake Pisang",
		bahanMakanan:     "Pisang",
		kategoriMasakan:  "Dessert",
		komposisiBahan:   "Pisang, tepung, susu, telur",
		langkahPembuatan: "Campur adonan lalu panggang",
		estimasiWaktu:    18,
	}
}