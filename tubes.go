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
	judulresep string
	bahanMakanan     string
	kategoriMasakan  string
	komposisiBahan   string
	langkahPembuatan string
	estimasiWaktu    int
	jumlahDicari     int
	favorit bool
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
	
	reader = bufio.NewReader(os.Stdin) 

	datadummy(&resep, &n)

	for{
		tampilanmenuUtama()
		fmt.Println("🍓 Apa yang mau anda lakukannn?? 🍓 𐔌՞ ܸ.ˬ.ܸ՞𐦯")
		fmt.Print("Masukan angkanya : ")
		fmt.Scan(&pilihanmenu)
		clearScreen()
		fmt.Println(" ")
		switch pilihanmenu{
		//═════════════════════════════════════════
		// input data
		//═════════════════════════════════════════
		case 1 :
			input(&resep, &n, reader)
			clearScreen()
		//═════════════════════════════════════════
		// tampilan data 
		//═════════════════════════════════════════
		case 2 :
			if n == 0{
				fmt.Println("Belum ada data resep")
			}else{
				tampilanresep(&resep, n )
			}
			wait()
			clearScreen()
			//═════════════════════════════════════════
			// 	searching
			//═════════════════════════════════════════
		case 3 :
			tampilansearching()
			fmt.Print("> masukan angkanya ya sayang ₍ᵔ.˛.ᵔ₎ :  ")
			fmt.Scan(&caraPencarianData)
			switch caraPencarianData{
			case 1 :
				fmt.Println("> Bahan utama apa yang mau anda cari?? 𐔌՞ ܸ.ˬ.ܸ՞𐦯  ")
				bahanYgDicari = inputString(reader)
			case 2 :
				fmt.Println("> judul apa yang mau dicari 𐔌՞ ܸ.ˬ.ܸ՞𐦯  ")
				bahanYgDicari = inputString(reader)
			}
	
			// percabangan sequential dan binary/ judul dan durasi
			if caraPencarianData == 1 {
				fmt.Println(" ")
				(pencarSequential(n, &resep, bahanYgDicari))
			} else {
				sortAbjad(&resep, n )
				fmt.Println(" ")
				(pencarianBinary(n, &resep, bahanYgDicari))
			}
			wait()
			clearScreen()
			//═════════════════════════════════════════
			// 	kelola data
			//═════════════════════════════════════════
		case 4 :
			pilihanKelola = menukelola()

			if pilihanKelola == 1 {
				tambahData(&resep, &n, reader)
			} else if pilihanKelola == 2 {
				ubahData(&resep, n, reader)
			} else if pilihanKelola == 3 {
				hapusData(&resep, &n)
			} else if pilihanKelola == 4 {
				Statistik(resep, n)
				StatistikKategori(resep,n)
			}
			wait()
			clearScreen()
			//═════════════════════════════════════════
			// pengurutan data
			//═════════════════════════════════════════
		case 5 :
			fmt.Println(" Berdasarkan apa resep ingin dikeluarkan??")
			fmt.Println("1.Durasi")
			fmt.Println("2.Abjad")
			fmt.Scan(&caraResepDikeluarkan)
			// kondisi output
			if caraResepDikeluarkan == 1 {
				fmt.Println(" ")
				sortWaktu(&resep, n)
			} else {
				fmt.Println(" ")
				sortAbjad(&resep, n)
				}
			tampilanSorting(&resep, &n)
				wait()
			clearScreen()
			clearScreen()
			//═════════════════════════════════════════
			// 	favorit
			//═════════════════════════════════════════
		case 6 :
			tampilanfavorit()
			fmt.Println("apa yang mau anda lakukan??: ")
			fmt.Scan(&angkafav)
			switch angkafav {
			case 1 :
				favorite(&resep, n )
			case 2 :
				lihatfavori(resep, n )
				wait()
				}
			clearScreen()
			//═════════════════════════════════════════
			// 	keluar
			//═════════════════════════════════════════
		case 7 :
			fmt.Println("terimakasihh telahhh menggunakan aplikasi resepku")
			wait()
			fmt.Println("keluar")
			return

		}
	}
}
//═════════════════════════════════════════
// tampilan menu 
//═════════════════════════════════════════
func tampilanmenuUtama(){
	fmt.Println("╔══════════════════════════════════╗")
    fmt.Println("║   𐙚⋆🍓 Aplikasi Resepku 🍰⋆𐙚     ║")
    fmt.Println("╠══════════════════════════════════╣")
    fmt.Println("║  ⋆𐙚 ̊. MENU UTAMA !!! 🍒 ૮₍•⤙•˶   ║")
    fmt.Println("╠══════════════════════════════════╣")
	fmt.Println("║  [1] input data resep            ║")
    fmt.Println("║  [2] lihat semua resep           ║")
    fmt.Println("║  [3] Cari Resep                  ║")
    fmt.Println("║  [4] kelola data                 ║")
    fmt.Println("║  [5] urutan resep                ║")
	fmt.Println("║  [6] kelola favorite             ║")
    fmt.Println("║  [7] keluar                      ║")
    fmt.Println("╚══════════════════════════════════╝")
}
//═════════════════════════════════════════
// function case 1 "inputan"   
//═════════════════════════════════════════
func input(resep *tabresep, n *int, reader *bufio.Reader){
	var i int
	var konfirmasi int
	var jumlah int
		fmt.Println("══════════════════════════════════════════")
		fmt.Println("           𐙚⋆  🍓  input data 🍰  𐙚⋆ " )
		fmt.Println("══════════════════════════════════════════")
			fmt.Print  (" Masukan jumlah data yang ingin diinput : ")
			fmt.Scan(&jumlah)
			clearScreen()
			fmt.Printf("> Apakah anda yakin ingin menginput data sebanyak %d ya sayang?\n", jumlah)
			fmt.Println("  [1] Ya   [2] Tidak")
			fmt.Print("→ ")
			fmt.Scan(&konfirmasi)
			if konfirmasi != 1 {
				fmt.Println("  Input data dibatalkan 𐔌՞ ܸ.ˬ.ܸ՞𐦯")
				return
			}
			*n = jumlah
			reader.ReadString('\n')
			fmt.Println("> ʚɞ masukan datanya yaaa!! ♡ ♡ ♡   ")
			for i = 0; i < *n; i++ {
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
				fmt.Scan(&resep[i].estimasiWaktu)
				reader.ReadString('\n')

				fmt.Println()
	
				fmt.Println(" ")
			}
}
//═════════════════════════════════════════
// function case 2 "tampilan resep"   
//═════════════════════════════════════════
func tampilanresep(resep *tabresep, n int){
	var i int
	for i = 0; i < n; i++ {
		fmt.Println("══════════════════════════════════════════")
		fmt.Printf("        𐙚⋆🍓 Resep ke-%-2d 🍰⋆𐙚            \n", i+1)
		fmt.Println("══════════════════════════════════════════")
		fmt.Printf ("  📖 Judul Resep       : %s\n", (*resep)[i].judulresep)
		fmt.Printf ("  🥬 Bahan Makanan     : %s\n", (*resep)[i].bahanMakanan)
		fmt.Printf ("  🏷️  Kategori Masakan  : %s\n", (*resep)[i].kategoriMasakan)
		fmt.Printf ("  🧄 Komposisi Bahan   : %s\n", (*resep)[i].komposisiBahan)
		fmt.Printf ("  📝 Langkah Pembuatan : %s\n", (*resep)[i].langkahPembuatan)
		fmt.Printf ("  ⏰ Estimasi Waktu    : %d menit\n", (*resep)[i].estimasiWaktu)
		fmt.Println()
	}
}
//═════════════════════════════════════════
// function case 3 "searching"   
//═════════════════════════════════════════
func tampilansearching(){
	fmt.Println("╔════════════════⋆══════════════════════════╗")
	fmt.Println("║        𐙚⋆🍓 Cari Resep 🍰⋆𐙚              ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Ingin berdasarkan apa data dicari?      ║")
	fmt.Println("║  [1] Bahan makanan utama                 ║")
	fmt.Println("║  [2] Judul                               ║")
	fmt.Println("╚══════════════════════════════════════════╝")
}
func pencarSequential(n int, resep *tabresep, bahanYgDicari string) {
	var i int
	var found bool
	for i = 0; i < n; i++ {
		if toLower(bahanYgDicari) == toLower((*resep)[i].bahanMakanan) {
			found = true
			(*resep)[i].jumlahDicari++
				fmt.Println("══════════════════════════════════════════")
				fmt.Printf ("         𐙚⋆🍓 Resep ke-%-2d 🍰⋆𐙚           \n", i+1)
				fmt.Println("══════════════════════════════════════════")
				fmt.Printf ("  📖 Judul Resep       : %s\n", (*resep)[i].judulresep)
				fmt.Printf ("  🥬 Bahan Makanan     : %s\n", (*resep)[i].bahanMakanan)
				fmt.Printf ("  🏷️  Kategori Masakan  : %s\n", (*resep)[i].kategoriMasakan)
				fmt.Printf ("  🧄 Komposisi Bahan   : %s\n", (*resep)[i].komposisiBahan)
				fmt.Printf ("  📝 Langkah Pembuatan : %s\n", (*resep)[i].langkahPembuatan)
				fmt.Printf ("  ⏰ Estimasi Waktu    : %d menit\n", (*resep)[i].estimasiWaktu)
				fmt.Println()
		}
	}
	if found == false {
			fmt.Printf("maaf yaa datanya tidak tersedia sayang 𐔌՞ ܸ.ˬ.ܸ՞𐦯 ")
		}
}

func pencarianBinary(n int, resep *tabresep, bahanYgDicari string) {
	var left, right, mid int
	var found int
	var target string
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
		(*resep)[found].jumlahDicari++
			fmt.Println("══════════════════════════════════════════")
			fmt.Printf ("         𐙚⋆🍓 Resep ke-%-2d 🍰⋆𐙚           \n", found+1)
			fmt.Println("══════════════════════════════════════════")
			fmt.Printf ("  📖 Judul Resep       : %s\n", resep[found].judulresep)
			fmt.Printf ("  🥬 Bahan Makanan     : %s\n", resep[found].bahanMakanan)
			fmt.Printf ("  🏷️  Kategori Masakan  : %s\n", resep[found].kategoriMasakan)
			fmt.Printf ("  🧄 Komposisi Bahan   : %s\n", resep[found].komposisiBahan)
			fmt.Printf ("  📝 Langkah Pembuatan : %s\n", resep[found].langkahPembuatan)
			fmt.Printf ("  ⏰ Estimasi Waktu    : %d menit\n", resep[found].estimasiWaktu)
			fmt.Println()
	} else {
		fmt.Printf("maaf yaa datanya tidak tersedia sayang 𐔌՞ ܸ.ˬ.ܸ՞𐦯 ")
	}
}
//═════════════════════════════════════════
// function case 4 "kelola data"   
//═════════════════════════════════════════
func menukelola()int{
	var pilihanKelola int
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║        𐙚⋆🍓 MENU KELOLA 🍰⋆𐙚             ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  [1] Tambah Data Resep                    ║")
	fmt.Println("║  [2] Ubah Data Resep                      ║")
	fmt.Println("║  [3] Hapus Data Resep                     ║")
	fmt.Println("║  [4] Lihat Statistik                      ║")
	fmt.Println("║  [5] Lanjut ke Pencarian                  ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Masukan angkanya ya sayang ₍ᵔ.˛.ᵔ₎ :    ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Print("→ ")
	fmt.Scan(&pilihanKelola)
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
	reader.ReadString('\n')
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
	fmt.Println("            yeay!! data berhasil ditambahkan ♡ ♡ ♡      ")
	fmt.Println(" ")
	wait()
	}

// func ubah data
func ubahData(resep *tabresep, n int, reader *bufio.Reader) {
	var idx int
	var konfirmasi int

	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║        𐙚⋆🍓 Ubah Data Resep 🍰⋆𐙚         ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Printf ("║  Masukan nomor resep yang mau diubah     ║\n")
	fmt.Printf ("║  (1 - %d) :                              ║\n", n)
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Print("→ ")
	fmt.Scan(&idx)
	idx = idx - 1
	if idx < 0 || idx >= n {
		fmt.Println(" maaf yaa nomor resepnya tidak valid     ")
		return
	}
	clearScreen()
    fmt.Printf("  Apakah anda yakin ingin mengubah resep \"%s\"?\n", resep[idx].judulresep)
    fmt.Println("  [1] Ya, lanjut ubah   [2] Tidak, batalkan")
    fmt.Print("→ ")
    fmt.Scan(&konfirmasi)
    clearScreen()

	if konfirmasi == 1 {
		reader.ReadString('\n')
		fmt.Println("═════════════════════════════════════════")
		fmt.Println("     ✦ Masukkan data baru yaaa!! ♡ ♡ ♡   ")
		fmt.Println("═════════════════════════════════════════")

		var temp dataresep

		fmt.Print("  📖 Judul Resep       : ")
		temp.judulresep = inputString(reader)

		fmt.Print("  🥬 Bahan Makanan     : ")
		temp.bahanMakanan = inputString(reader)

		fmt.Print("  🏷️  Kategori Masakan  : ")
		temp.kategoriMasakan = inputString(reader)

		fmt.Print("  🧄 Komposisi Bahan   : ")
		temp.komposisiBahan = inputString(reader)

		fmt.Print("  📝 Langkah Pembuatan : ")
		temp.langkahPembuatan = inputString(reader)

		fmt.Print("  ⏰ Estimasi Waktu    : ")
		fmt.Scan(&temp.estimasiWaktu)
		fmt.Println(" ")

		fmt.Printf("  Apakah anda yakin ingin mengubah resep \"%s\" menjadi data di atas?\n", resep[idx].judulresep)
		fmt.Println("  [1] Ya   [2] Tidak (Undo)")
		fmt.Print("→ ")
		fmt.Scan(&konfirmasi)
		clearScreen()
		if konfirmasi == 1 {
			resep[idx] = temp // Data asli baru ditimpa di sini kalau setuju
			fmt.Println("  Yeay! Data resep berhasil diperbarui... ✨")
		} else {
			fmt.Println("  Ubah data dibatalkan! Data lama kamu tetap aman 𐔌՞ ܸ.ˬ.ܸ՞𐦯")
		}
	} else {
		fmt.Println("  Ubah data dibatalkan! Data lama kamu tetap aman 𐔌՞ ܸ.ˬ.ܸ՞𐦯")
	}
	fmt.Println()
	wait()
}

// func hapus data
func hapusData(resep *tabresep, n *int) {
	var idx int
	var konfirmasi int
	var i int
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║       𐙚⋆🍓 Hapus Data Resep 🍰⋆𐙚         ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Printf ("║  Masukan nomor resep yang mau dihapus    ║\n")
	fmt.Printf ("║  (1 - %d) :                               ║\n", *n)
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Print("→ ")
	fmt.Scan(&idx)
	idx = idx - 1
	if idx < 0 || idx >= *n {
		fmt.Println("  maaf yaa nomor resepnya tidak valid     ")
		
		return
	}
	fmt.Println("══════════════════════════════════════════")
	fmt.Printf("  Yakin mau hapus \"%s\"?\n", (*resep)[idx].judulresep) // ← tampilkan nama resepnya
	fmt.Println("  [1] Ya   [2] Tidak")
	fmt.Println("══════════════════════════════════════════")
	fmt.Print("→ ")
	fmt.Scan(&konfirmasi)

if konfirmasi != 1 {
    fmt.Println("  oke, penghapusan dibatalkan 𐔌՞ ܸ.ˬ.ܸ՞𐦯")
    return  // ← batalkan, tidak jadi hapus
}
	for i = idx; i < *n-1; i++ {
		resep[i] = resep[i+1]
	}
	*n = *n - 1
	fmt.Println("        data berhasil dihapus ♡ ♡ ♡          ")
	
	fmt.Println(" ")
	wait()
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
	fmt.Printf ("  🍓 Total Resep          : %d resep\n", n)
	fmt.Printf ("  ⏰ Total Estimasi Waktu : %d menit\n", totalWaktu)
	fmt.Printf ("  📊 Rata-rata Waktu      : %.2f menit\n", rataRata)
	fmt.Printf ("  ⚡ Waktu Tercepat       : %d menit\n", minWaktu)
	fmt.Printf ("  🔥 Waktu Terlama        : %d menit\n", maxWaktu)
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

func StatistikKategori(resep tabresep, n int){
	var namaKategori [NMAX]string
	var jumlahKategori [NMAX]int
	var totalKategori int
	var i,j int
	var found bool
	for i = 0; i < n; i++{
		found = false
		for j = 0; j < totalKategori; j++{
			if resep[i].kategoriMasakan == namaKategori[j]{
				jumlahKategori[j] ++
				found = true
			}
		}
		if found == false {
			namaKategori[totalKategori]= resep[i].kategoriMasakan
			jumlahKategori[totalKategori] = 1
			totalKategori ++
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
//═════════════════════════════════════════
// function case 5 "pengurutan"   
//═════════════════════════════════════════

func sortWaktu(resep *tabresep, n int) {
	var i,j,min int
	var temp dataresep
	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if urutan == 1{
				resep[j].estimasiWaktu < resep[min].estimasiWaktu 
			}else{
				resep[j].estimasiWaktu > resep[ekstrem].estimasiWaktu

			}
		}
		temp = resep[i]
		resep[i] = resep[min]
		resep[min] = temp
	}
}

// sorting berdasarkan bahan makanan ascending A ampe Z
func sortAbjad(resep *tabresep, n int) {
	var i,j int
	var temp dataresep
	for i = 1; i < n; i++ {
		temp = resep[i]
		j = i - 1
		for j >= 0 && toLower(resep[j].judulresep) >toLower(temp.judulresep) {
			resep[j+1] = resep[j]
			j = j-1
		}
		 resep[j+1] = temp
	}
}
func tampilanSorting(resep *tabresep, n *int){
	var i int
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║    𐙚⋆🍓 Hasil Setelah Diurutkan 🍰⋆𐙚     ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		for i = 0; i < *n; i++ {
			fmt.Println("══════════════════════════════════════════")
			fmt.Printf("        𐙚⋆🍓 Resep ke-%-2d 🍰⋆𐙚            \n", i+1)
			fmt.Println("══════════════════════════════════════════")
			fmt.Printf ("  📖 Judul Resep       : %s\n", resep[i].judulresep)
			fmt.Printf ("  🥬 Bahan Makanan     : %s\n", resep[i].bahanMakanan)
			fmt.Printf ("  🏷️  Kategori Masakan  : %s\n", resep[i].kategoriMasakan)
			fmt.Printf ("  🧄 Komposisi Bahan   : %s\n", resep[i].komposisiBahan)
			fmt.Printf ("  📝 Langkah Pembuatan : %s\n", resep[i].langkahPembuatan)
			fmt.Printf ("  ⏰ Estimasi Waktu    : %d menit\n", resep[i].estimasiWaktu)
			fmt.Println()
		}
}
//═════════════════════════════════════════
// function case 6 "favorite"   
//═════════════════════════════════════════
func tampilanfavorit(){
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║      𐙚⋆🍓 Menu Favorit 🍰⋆𐙚             ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Apa yang ingin anda lakukan?            ║")
	fmt.Println("║  [1] Tambah Resep Favorit                ║")
	fmt.Println("║  [2] Lihat Semua Favorit                 ║")
	fmt.Println("╚══════════════════════════════════════════╝")
}
func favorite(resep *tabresep, n int){
	var pilihfav int
	tampilanresep(resep, n )
	fmt.Println(" resep nomer berapa yang ingin ditambahkan ke favorit??")
	fmt.Scan(&pilihfav)

	pilihfav--

	if pilihfav >= 0 && pilihfav < n {
		resep[pilihfav].favorit = true
		fmt.Println("♡ Resep berhasil ditambahkan ke favorit!")
	} else {
		fmt.Println("Nomor tidak valid.")
	}
}
func lihatfavori(resep tabresep, n int ){
	var i int
	var found bool
	for i = 0; i < n; i++{
		if resep[i].favorit== true {
			found = true 

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
		fmt.Println("Belum ada resep favorit ❤️")
	}
}
func inputString(reader *bufio.Reader) string {
	var text string
	text, _ = reader.ReadString('\n')
	return strings.TrimSpace(text)
}
func clearScreen() {
    var cmd *exec.Cmd //var untuk menyimpan perintah terminal
    if runtime.GOOS == "windows" { //ini kalau user menggunakan windows
        cmd = exec.Command("cmd", "/c", "cls") //jika true maka akan menjalankan perintah "cls"
    } else { //jika osnya Mac atau Linux
        cmd = exec.Command("clear") //maka akan menjalankan perintah "clear"
    }
    cmd.Stdout = os.Stdout //untuk menyambungkan command ke terminal program
    cmd.Run() //ini agar perintah tadi langsung di eksekusi
}

func wait() {
	var buang string
	fmt.Scanln(&buang)
	fmt.Print("kalo mau lanjut, tekan enter yaa!! 𐔌՞ ܸ.ˬ.ܸ՞")
	fmt.Scanln()
}

func toLower(s string) string {
	var i int
	var res string

	i = 0
	res = ""
	for i < len(s) {
		if s[i] >= 'A' && s[i] <= 'Z' {
			res = res + string(s[i] + 32)
		} else {
			res = res + string(s[i])
		}
		i++
	}
	return res
}


func datadummy(resep *tabresep, n *int){
	*n = 10

resep[0] = dataresep{
	judulresep: "Nasi Goreng Spesial",
	bahanMakanan: "Nasi",
	kategoriMasakan: "Makanan Utama",
	komposisiBahan: "Nasi, telur, kecap, bawang",
	langkahPembuatan: "Tumis bumbu lalu masukkan nasi",
	estimasiWaktu: 20,
}

resep[1] = dataresep{
	judulresep: "Ayam Goreng Krispi",
	bahanMakanan: "Ayam",
	kategoriMasakan: "Makanan Utama",
	komposisiBahan: "Ayam, tepung, garam",
	langkahPembuatan: "Balur ayam lalu goreng",
	estimasiWaktu: 30,
}

resep[2] = dataresep{
	judulresep: "Sup Jagung",
	bahanMakanan: "Jagung",
	kategoriMasakan: "Sup",
	komposisiBahan: "Jagung, telur, wortel",
	langkahPembuatan: "Rebus semua bahan hingga matang",
	estimasiWaktu: 25,
}

resep[3] = dataresep{
	judulresep: "Mie Goreng Jawa",
	bahanMakanan: "Mie",
	kategoriMasakan: "Makanan Utama",
	komposisiBahan: "Mie, kol, kecap, telur",
	langkahPembuatan: "Tumis bumbu lalu masak mie",
	estimasiWaktu: 15,
}

resep[4] = dataresep{
	judulresep: "Puding Coklat",
	bahanMakanan: "Coklat",
	kategoriMasakan: "Dessert",
	komposisiBahan: "Susu, coklat bubuk, agar-agar",
	langkahPembuatan: "Rebus lalu dinginkan",
	estimasiWaktu: 40,
}

resep[5] = dataresep{
	judulresep: "Es Buah Segar",
	bahanMakanan: "Buah",
	kategoriMasakan: "Minuman",
	komposisiBahan: "Melon, semangka, sirup",
	langkahPembuatan: "Campur semua bahan",
	estimasiWaktu: 10,
}

resep[6] = dataresep{
	judulresep: "Sate Ayam Madura",
	bahanMakanan: "Ayam",
	kategoriMasakan: "Makanan Utama",
	komposisiBahan: "Ayam, kecap, bumbu kacang",
	langkahPembuatan: "Bakar ayam lalu sajikan",
	estimasiWaktu: 45,
}

resep[7] = dataresep{
	judulresep: "Tumis Kangkung",
	bahanMakanan: "Kangkung",
	kategoriMasakan: "Sayuran",
	komposisiBahan: "Kangkung, bawang putih",
	langkahPembuatan: "Tumis hingga layu",
	estimasiWaktu: 12,
}

resep[8] = dataresep{
	judulresep: "Bakso Kuah",
	bahanMakanan: "Daging Sapi",
	kategoriMasakan: "Sup",
	komposisiBahan: "Bakso, mie, daun bawang",
	langkahPembuatan: "Rebus hingga matang",
	estimasiWaktu: 35,
}

resep[9] = dataresep{
	judulresep: "Pancake Pisang",
	bahanMakanan: "Pisang",
	kategoriMasakan: "Dessert",
	komposisiBahan: "Pisang, tepung, susu, telur",
	langkahPembuatan: "Campur adonan lalu panggang",
	estimasiWaktu: 18,
}
}