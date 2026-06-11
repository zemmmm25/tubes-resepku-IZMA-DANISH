package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// =============================================
// KONSTANTA & TIPE BENTUKAN
// =============================================

const MAKS_MENU = 100
const MAKS_BAHAN = 10

type Bahan struct {
	Nama   string
	Jumlah string
}

type Menu struct {
	ID          int
	Nama        string
	Kategori    string // "Coffee", "Non-Coffee", "Food", "Dessert"
	Harga       int
	Bahan       [MAKS_BAHAN]Bahan
	JumlahBahan int
	Tersedia    bool
	Deskripsi   string
}

type KatalogMenu struct {
	Data   [MAKS_MENU]Menu
	Jumlah int
}

type HasilPencarian struct {
	Data   [MAKS_MENU]int
	Jumlah int
}

// =============================================
// VARIABEL GLOBAL
// =============================================

var katalog KatalogMenu
var idCounter int = 1
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// =============================================
// UTILITAS FUNDAMENTAL
// =============================================

func trim(s string) string {
	var start, end, i int
	var res string

	start = 0
	end = len(s) - 1
	res = ""

	for start <= end && (s[start] == ' ' || s[start] == '\n' || s[start] == '\r' || s[start] == '\t') {
		start++
	}
	for end >= start && (s[end] == ' ' || s[end] == '\n' || s[end] == '\r' || s[end] == '\t') {
		end--
	}
	if start <= end {
		i = start
		for i <= end {
			res += string(s[i])
			i++
		}
	}
	return res
}

func toLower(s string) string {
	var i int
	var res string

	i = 0
	res = ""
	for i < len(s) {
		if s[i] >= 'A' && s[i] <= 'Z' {
			res += string(s[i] + 32)
		} else {
			res += string(s[i])
		}
		i++
	}
	return res
}

func contains(s, substr string) bool {
	var found, match bool
	var i, j int

	found = false
	i = 0

	if len(substr) == 0 {
		found = true
	} else if len(substr) <= len(s) {
		for i <= len(s)-len(substr) && !found {
			match = true
			j = 0
			for j < len(substr) && match {
				if s[i+j] != substr[j] {
					match = false
				}
				j++
			}
			if match {
				found = true
			}
			i++
		}
	}
	return found
}

// =============================================
// UTILITAS INPUT & TAMPILAN UI
// =============================================

func inputString(prompt string) string {
	var text string
	fmt.Print(prompt)
	text, _ = reader.ReadString('\n')
	return trim(text)
}

func inputInt(prompt string) int {
	var n int
	var text string

	fmt.Print(prompt)
	text, _ = reader.ReadString('\n')
	text = trim(text)
	fmt.Sscan(text, &n)
	return n
}

func inputBool(prompt string) bool {
	var jawab string
	fmt.Print(prompt + " (y/n): ")
	jawab, _ = reader.ReadString('\n')
	jawab = toLower(trim(jawab))
	return jawab == "y" || jawab == "yes"
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		exec.Command("cmd", "/c", "cls").Run()
	} else {
		exec.Command("clear").Run()
	}
}

func tekanEnter() {
	inputString("\n  Tekan Enter untuk melanjutkan...")
	clearScreen()
}

func garisLurus(panjang int) {
	var i int
	i = 0
	for i < panjang {
		fmt.Print("─")
		i++
	}
}

func headerApp() {
	clearScreen()
	fmt.Println()
	fmt.Print("┌")
	garisLurus(72)
	fmt.Println("┐")
	fmt.Printf("│%-72s│\n", "                  CAFE TELKOM | Katalog Menu Digital")
	fmt.Print("├")
	garisLurus(72)
	fmt.Println("┤")
}

func footerApp() {
	fmt.Print("└")
	garisLurus(72)
	fmt.Println("┘")
}

func headerTabel() {
	var header string
	fmt.Print("┌")
	garisLurus(72)
	fmt.Println("┐")
	header = fmt.Sprintf(" %-3s | %-22s | %-10s | %-11s | %-11s", "No.", "Nama Menu", "Status", "Harga", "Kategori")
	fmt.Printf("│%-72s│\n", header)
	fmt.Print("├")
	garisLurus(72)
	fmt.Println("┤")
}

func cetakMenu(m Menu, nomor int) {
	var statusStr, baris string

	statusStr = "[TERSEDIA]"
	if !m.Tersedia {
		statusStr = "[HABIS]"
	}
	baris = fmt.Sprintf(" %3d | %-22s | %-10s | Rp%-8d | %-11s", nomor, m.Nama, statusStr, m.Harga, m.Kategori)
	fmt.Printf("│%-72s│\n", baris)
}

func cetakMenuDetail(m Menu) {
	var i int
	var desk, statusStr, barisBahan string

	i = 0
	fmt.Print("┌")
	garisLurus(72)
	fmt.Println("┐")
	fmt.Printf("│%-72s│\n", " DETAIL MENU")
	fmt.Print("├")
	garisLurus(72)
	fmt.Println("┤")

	fmt.Printf("│%-72s│\n", fmt.Sprintf("  ID        : %d", m.ID))
	fmt.Printf("│%-72s│\n", fmt.Sprintf("  Nama      : %s", m.Nama))
	fmt.Printf("│%-72s│\n", fmt.Sprintf("  Kategori  : %s", m.Kategori))
	fmt.Printf("│%-72s│\n", fmt.Sprintf("  Harga     : Rp%d", m.Harga))

	statusStr = "HABIS"
	if m.Tersedia {
		statusStr = "TERSEDIA"
	}
	fmt.Printf("│%-72s│\n", fmt.Sprintf("  Status    : %s", statusStr))

	desk = m.Deskripsi
	if len(desk) > 50 {
		desk = desk[:47] + "..."
	}
	fmt.Printf("│%-72s│\n", fmt.Sprintf("  Deskripsi : %s", desk))

	if m.JumlahBahan > 0 {
		fmt.Printf("│%-72s│\n", "  Bahan     :")
		for i < m.JumlahBahan {
			barisBahan = fmt.Sprintf("    - %s (%s)", m.Bahan[i].Nama, m.Bahan[i].Jumlah)
			if len(barisBahan) > 65 {
				barisBahan = barisBahan[:62] + "..."
			}
			fmt.Printf("│%-72s│\n", barisBahan)
			i++
		}
	}
	footerApp()
}

// =============================================
// INISIALISASI DATA AWAL 
// =============================================

func tambahDataAwal() {
	type DataAwal struct {
		nama      string
		kategori  string
		harga     int
		tersedia  bool
		deskripsi string
		bahan     [MAKS_BAHAN]Bahan
		jmlBahan  int
	}

	var i, j int
	var data [12]DataAwal
	var d DataAwal
	var m Menu

	i = 0
	data = [12]DataAwal{
		{"Espresso", "Coffee", 18000, true, "Kopi murni tanpa susu, pekat dan kuat", [10]Bahan{{"Biji kopi arabika", "18g"}, {"Air panas", "30ml"}}, 2},
		{"Cappuccino", "Coffee", 28000, true, "Espresso dengan busa susu lembut", [10]Bahan{{"Espresso", "1 shot"}, {"Susu full cream", "150ml"}, {"Busa susu", "secukupnya"}}, 3},
		{"Cafe Latte", "Coffee", 30000, true, "Kopi lembut dengan dominasi susu", [10]Bahan{{"Espresso", "1 shot"}, {"Susu steamed", "200ml"}}, 2},
		{"Americano", "Coffee", 22000, true, "Espresso dengan tambahan air panas", [10]Bahan{{"Espresso", "2 shot"}, {"Air panas", "150ml"}}, 2},
		{"Matcha Latte", "Non-Coffee", 32000, true, "Teh hijau premium dengan susu", [10]Bahan{{"Bubuk matcha", "5g"}, {"Susu oat", "200ml"}, {"Madu", "1 sdt"}}, 3},
		{"Teh Tarik", "Non-Coffee", 20000, true, "Teh manis dengan susu kental", [10]Bahan{{"Teh hitam", "1 kantong"}, {"Susu kental manis", "3 sdm"}, {"Air panas", "200ml"}}, 3},
		{"Jus Alpukat", "Non-Coffee", 25000, true, "Jus alpukat segar dengan susu", [10]Bahan{{"Alpukat", "1 buah"}, {"Susu cair", "100ml"}, {"Gula pasir", "1 sdm"}}, 3},
		{"Croissant", "Food", 22000, true, "Croissant mentega panggang renyah", [10]Bahan{{"Tepung terigu", "200g"}, {"Mentega", "100g"}, {"Ragi", "5g"}}, 3},
		{"Roti Bakar", "Food", 18000, true, "Roti tawar panggang dengan selai", [10]Bahan{{"Roti tawar", "2 lembar"}, {"Selai coklat", "secukupnya"}, {"Mentega", "1 sdm"}}, 3},
		{"Waffle", "Dessert", 35000, true, "Waffle crispy dengan es krim vanilla", [10]Bahan{{"Tepung waffle", "150g"}, {"Telur", "1 butir"}, {"Es krim vanilla", "1 scoop"}}, 3},
		{"Cheesecake", "Dessert", 38000, false, "Cheesecake lembut dengan topping blueberry", [10]Bahan{{"Cream cheese", "200g"}, {"Biskuit oreo", "100g"}, {"Selai blueberry", "50g"}}, 3},
		{"Cold Brew", "Coffee", 35000, true, "Kopi cold brew 18 jam brewing", [10]Bahan{{"Biji kopi robusta", "30g"}, {"Air dingin", "300ml"}}, 2},
	}

	for i < len(data) {
		d = data[i]
		m = Menu{}
		m.ID = idCounter
		idCounter++
		m.Nama = d.nama
		m.Kategori = d.kategori
		m.Harga = d.harga
		m.Tersedia = d.tersedia
		m.Deskripsi = d.deskripsi

		j = 0
		for j < d.jmlBahan {
			m.Bahan[j] = d.bahan[j]
			j++
		}
		m.JumlahBahan = d.jmlBahan
		katalog.Data[katalog.Jumlah] = m
		katalog.Jumlah++
		i++
	}
}

// =============================================
// SUBPROGRAM: PENCARIAN
// =============================================

func sequentialSearchNama(k KatalogMenu, keyword string) int {
	var i, hasil int
	var kataKunci string

	i = 0
	hasil = -1
	kataKunci = toLower(keyword)

	for i < k.Jumlah && hasil == -1 {
		if contains(toLower(k.Data[i].Nama), kataKunci) {
			hasil = i
		}
		i++
	}
	return hasil
}

func sequentialSearchKategori(k KatalogMenu, kategori string) HasilPencarian {
	var i int
	var katKunci string
	var hasil HasilPencarian

	i = 0
	katKunci = toLower(kategori)
	hasil.Jumlah = 0

	for i < k.Jumlah {
		if toLower(k.Data[i].Kategori) == katKunci {
			hasil.Data[hasil.Jumlah] = i
			hasil.Jumlah++
		}
		i++
	}
	return hasil
}

func sequentialSearchID(k KatalogMenu, id int) int {
	var i, hasil int

	i = 0
	hasil = -1

	for i < k.Jumlah && hasil == -1 {
		if k.Data[i].ID == id {
			hasil = i
		}
		i++
	}
	return hasil
}

func binarySearchNama(k KatalogMenu, keyword string) int {
	var kiri, kanan, hasil, tengah int
	var namaTengah, kataKunci string
	
	kiri = 0
	kanan = k.Jumlah - 1
	hasil = -1
	kataKunci = toLower(keyword)

	for kiri <= kanan && hasil == -1 {
		tengah = (kiri + kanan) / 2
		namaTengah = toLower(k.Data[tengah].Nama)

		if namaTengah == kataKunci {
			hasil = tengah
		} else if namaTengah < kataKunci {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return hasil
}

// =============================================
// SUBPROGRAM: PENGURUTAN (KONSOLIDASI OPTIMAL)
// =============================================

func selectionSort(k *KatalogMenu, byNama bool, ascending bool) {
	var i, j, idxEkstrem int
	var kondisi bool

	i = 0
	for i < k.Jumlah-1 {
		idxEkstrem = i
		j = i + 1
		for j < k.Jumlah {
			kondisi = false
			if byNama {
				if ascending {
					kondisi = toLower(k.Data[j].Nama) < toLower(k.Data[idxEkstrem].Nama)
				} else {
					kondisi = toLower(k.Data[j].Nama) > toLower(k.Data[idxEkstrem].Nama)
				}
			} else {
				if ascending {
					kondisi = k.Data[j].Harga < k.Data[idxEkstrem].Harga
				} else {
					kondisi = k.Data[j].Harga > k.Data[idxEkstrem].Harga
				}
			}

			if kondisi {
				idxEkstrem = j
			}
			j++
		}
		if idxEkstrem != i {
			k.Data[i], k.Data[idxEkstrem] = k.Data[idxEkstrem], k.Data[i]
		}
		i++
	}
}

func insertionSort(k *KatalogMenu, byNama bool, ascending bool) {
	var i, j, posisi int
	var kondisi bool
	var kunci Menu

	i = 1
	for i < k.Jumlah {
		kunci = k.Data[i]
		j = i - 1
		posisi = 0
		for j >= 0 {
			kondisi = false
			if byNama {
				if ascending {
					kondisi = toLower(k.Data[j].Nama) > toLower(kunci.Nama)
				} else {
					kondisi = toLower(k.Data[j].Nama) < toLower(kunci.Nama)
				}
			} else {
				if ascending {
					kondisi = k.Data[j].Harga > kunci.Harga
				} else {
					kondisi = k.Data[j].Harga < kunci.Harga
				}
			}

			if kondisi {
				k.Data[j+1] = k.Data[j]
				j--
				posisi = j + 1
			} else {
				posisi = j + 1
				j = -1
			}
		}
		k.Data[posisi] = kunci
		i++
	}
}

// =============================================
// SUBPROGRAM: CRUD
// =============================================

func pilihKategori(pil int, defaultKategori string) string {
	var hasil string

	if pil == 1 {
		hasil = "Coffee"
	} else if pil == 2 {
		hasil = "Non-Coffee"
	} else if pil == 3 {
		hasil = "Food"
	} else if pil == 4 {
		hasil = "Dessert"
	} else {
		hasil = defaultKategori
	}
	return hasil
}

func tambahMenu(k *KatalogMenu) {
	var pil, jmlBahan, i int
	var m Menu

	i = 0

	if k.Jumlah >= MAKS_MENU {
		fmt.Println("\n  [!] Kapasitas katalog penuh!")
		return
	}

	fmt.Println()
	fmt.Println("  -- TAMBAH MENU BARU --")
	
	m.ID = idCounter
	idCounter++

	m.Nama = inputString("  Nama Menu    : ")
	fmt.Println("  Kategori     : 1.Coffee  2.Non-Coffee  3.Food  4.Dessert")

	pil = inputInt("  Pilih        : ")
	m.Kategori = pilihKategori(pil, "Other")

	m.Harga = inputInt("  Harga (Rp)   : ")
	m.Deskripsi = inputString("  Deskripsi    : ")
	m.Tersedia = inputBool("  Tersedia?    ")

	fmt.Print("  Jumlah bahan (maks 10): ")
	jmlBahan = inputInt("")
	if jmlBahan > MAKS_BAHAN {
		jmlBahan = MAKS_BAHAN
	}

	for i < jmlBahan {
		fmt.Printf("  Bahan ke-%d:\n", i+1)
		m.Bahan[i].Nama = inputString("    Nama bahan  : ")
		m.Bahan[i].Jumlah = inputString("    Jumlah/takaran: ")
		i++
	}
	m.JumlahBahan = jmlBahan

	k.Data[k.Jumlah] = m
	k.Jumlah++
	fmt.Printf("\n  [OK] Menu '%s' berhasil ditambahkan! (ID: %d)\n", m.Nama, m.ID)
}

func ubahMenu(k *KatalogMenu) {
	var id, idx, pil, hargaB int
	var namaB, deskB, statusStr string
	var m Menu

	fmt.Println()
	fmt.Println("  -- UBAH DATA MENU --")

	id = inputInt("  Masukkan ID menu yang ingin diubah: ")
	idx = sequentialSearchID(*k, id)

	if idx == -1 {
		fmt.Println("  [!] Menu dengan ID tersebut tidak ditemukan.")
		return
	}

	m = k.Data[idx]
	fmt.Printf("\n  Menu ditemukan: %s (ID: %d)\n", m.Nama, m.ID)
	fmt.Println("  Kosongkan input untuk tidak mengubah field tersebut.")
	fmt.Println()

	namaB = inputString(fmt.Sprintf("  Nama [%s]: ", m.Nama))
	if namaB != "" {
		m.Nama = namaB
	}

	fmt.Printf("  Kategori saat ini: %s\n", m.Kategori)
	fmt.Println("  1.Coffee  2.Non-Coffee  3.Food  4.Dessert  0.Tidak ubah")

	pil = inputInt("  Pilih: ")
	m.Kategori = pilihKategori(pil, m.Kategori)

	fmt.Printf("  Harga saat ini: Rp%d\n", m.Harga)
	hargaB = inputInt("  Harga baru (0 = tidak ubah): ")
	if hargaB > 0 {
		m.Harga = hargaB
	}

	deskB = inputString(fmt.Sprintf("  Deskripsi [%s]: ", m.Deskripsi))
	if deskB != "" {
		m.Deskripsi = deskB
	}

	statusStr = "tersedia"
	if !m.Tersedia {
		statusStr = "habis"
	}
	fmt.Printf("  Status saat ini: %s\n", statusStr)
	m.Tersedia = inputBool("  Tersedia?")

	k.Data[idx] = m
	fmt.Printf("\n  [OK] Menu '%s' berhasil diubah!\n", m.Nama)
}

func hapusMenu(k *KatalogMenu) {
	var id, idx, i int
	var konfirmasi bool
	var namaHapus string

	fmt.Println()
	fmt.Println("  -- HAPUS MENU --")
	id = inputInt("  Masukkan ID menu yang ingin dihapus: ")
	idx = sequentialSearchID(*k, id)

	if idx == -1 {
		fmt.Println("  [!] Menu dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Printf("\n  Menu: %s - Rp%d [%s]\n", k.Data[idx].Nama, k.Data[idx].Harga, k.Data[idx].Kategori)
	konfirmasi = inputBool("  Yakin ingin menghapus?")

	if konfirmasi {
		namaHapus = k.Data[idx].Nama
		i = idx
		for i < k.Jumlah-1 {
			k.Data[i] = k.Data[i+1]
			i++
		}
		k.Jumlah--
		fmt.Printf("\n  [OK] Menu '%s' berhasil dihapus!\n", namaHapus)
	} else {
		fmt.Println("  [!] Penghapusan dibatalkan.")
	}
}

// =============================================
// SUBPROGRAM: TAMPILKAN
// =============================================

func tampilkanSemuaMenu(k KatalogMenu) {
	var i int
	
	i = 0

	if k.Jumlah == 0 {
		fmt.Println("  [!] Katalog masih kosong.")
		return
	}

	fmt.Println()
	fmt.Printf("  Total menu di sistem: %d\n", k.Jumlah)
	
	headerTabel()
	for i < k.Jumlah {
		cetakMenu(k.Data[i], i+1)
		i++
	}
	footerApp()
}

func tampilkanDetailMenu(k KatalogMenu) {
	var id, idx int

	fmt.Println()
	id = inputInt("  Masukkan ID menu untuk melihat detail: ")
	idx = sequentialSearchID(k, id)

	if idx == -1 {
		fmt.Println("  [!] Menu tidak ditemukan.")
		return
	}

	fmt.Println()
	cetakMenuDetail(k.Data[idx])
}

// =============================================
// SUBPROGRAM: INTERFACE MENU LAINNYA
// =============================================

func menuPencarian(k KatalogMenu) {
	var pilihan, idx, i int
	var keyword string
	var hasil HasilPencarian
	var salinan KatalogMenu

	fmt.Println()
	fmt.Println("  -- PENCARIAN MENU --")
	fmt.Println("  1. Berdasarkan nama (Sequential)")
	fmt.Println("  2. Berdasarkan kategori (Sequential)")
	fmt.Println("  3. Berdasarkan nama (Binary - Diurutkan Otomatis)")
	fmt.Println("  0. Batal")

	pilihan = inputInt("  Pilih: ")

	if pilihan == 1 {
		keyword = inputString("  Masukkan nama menu: ")
		idx = sequentialSearchNama(k, keyword)
		if idx == -1 {
			fmt.Println("  [!] Menu tidak ditemukan.")
		} else {
			fmt.Println("\n  [OK] Menu ditemukan:")
			cetakMenuDetail(k.Data[idx])
		}
	} else if pilihan == 2 {
		fmt.Println("  Kategori: Coffee | Non-Coffee | Food | Dessert")
		keyword = inputString("  Masukkan kategori: ")
		hasil = sequentialSearchKategori(k, keyword)

		if hasil.Jumlah == 0 {
			fmt.Println("  [!] Tidak ada menu dengan kategori tersebut.")
		} else {
			fmt.Printf("\n  Ditemukan %d menu dengan kategori '%s':\n", hasil.Jumlah, keyword)
			headerTabel()
			i = 0
			for i < hasil.Jumlah {
				cetakMenu(k.Data[hasil.Data[i]], i+1)
				i++
			}
			footerApp()
		}
	} else if pilihan == 3 {
		salinan = k
		selectionSort(&salinan, true, true)
		fmt.Println("  [INFO] Data disalin & diurutkan untuk Binary Search.")
		keyword = inputString("  Masukkan nama menu (spesifik): ")
		idx = binarySearchNama(salinan, keyword)

		if idx == -1 {
			fmt.Println("  [!] Menu tidak ditemukan.")
		} else {
			fmt.Println("\n  [OK] Menu ditemukan:")
			cetakMenuDetail(salinan.Data[idx])
		}
	}
}

func menuPengurutan(k *KatalogMenu) {
	var algoPil, fieldPil, urutanPil int
	var ascending, byNama bool
	var algoNama, fieldNama, urutanStr string

	algoNama = "Selection Sort"
	fieldNama = "Harga"
	urutanStr = "Ascending"

	fmt.Println()
	fmt.Println("  -- PENGURUTAN MENU --")
	fmt.Println("  Algoritma : 1. Selection Sort   2. Insertion Sort")
	algoPil = inputInt("  Pilih algoritma: ")

	fmt.Println("  Field     : 1. Harga   2. Nama")
	fieldPil = inputInt("  Pilih field: ")

	fmt.Println("  Urutan    : 1. Ascending (A-Z / Termurah)   2. Descending (Z-A / Termahal)")
	urutanPil = inputInt("  Pilih urutan: ")

	ascending = (urutanPil == 1)
	byNama = (fieldPil == 2)

	if algoPil == 2 {
		algoNama = "Insertion Sort"
		insertionSort(k, byNama, ascending)
	} else {
		selectionSort(k, byNama, ascending)
	}

	if byNama {
		fieldNama = "Nama"
	}
	if !ascending {
		urutanStr = "Descending"
	}

	fmt.Printf("\n  [OK] Data diurutkan dengan %s berdasarkan %s (%s)!\n", algoNama, fieldNama, urutanStr)
	fmt.Println()
	tampilkanSemuaMenu(*k)
}

func tampilkanStatistik(k KatalogMenu) {
	var totalMenu, totalHarga, totalTersedia, i, j, jumlah, totalHargaKat, tersediaKat, rataKat, rataTotal int
	var kat, barisStat string
	var kategoriList [4]string
	var salinan KatalogMenu

	kategoriList = [4]string{"Coffee", "Non-Coffee", "Food", "Dessert"}
	totalMenu = k.Jumlah
	totalHarga = 0
	totalTersedia = 0
	i = 0

	fmt.Println()
	fmt.Print("┌")
	garisLurus(72)
	fmt.Println("┐")
	fmt.Printf("│%-72s│\n", " STATISTIK KESELURUHAN")
	fmt.Print("├")
	garisLurus(72)
	fmt.Println("┤")

	for i < len(kategoriList) {
		kat = kategoriList[i]
		jumlah = 0
		totalHargaKat = 0
		tersediaKat = 0
		j = 0

		for j < k.Jumlah {
			if k.Data[j].Kategori == kat {
				jumlah++
				totalHargaKat += k.Data[j].Harga
				if k.Data[j].Tersedia {
					tersediaKat++
				}
			}
			j++
		}

		rataKat = 0
		if jumlah > 0 {
			rataKat = totalHargaKat / jumlah
		}
		
		barisStat = fmt.Sprintf(" %-12s : %2d menu | Tersedia: %2d | Rata-rata: Rp%d", kat, jumlah, tersediaKat, rataKat)
		fmt.Printf("│%-72s│\n", barisStat)
		totalHarga += totalHargaKat
		i++
	}

	i = 0
	for i < k.Jumlah {
		if k.Data[i].Tersedia {
			totalTersedia++
		}
		i++
	}

	rataTotal = 0
	if totalMenu > 0 {
		rataTotal = totalHarga / totalMenu
	}

	fmt.Print("├")
	garisLurus(72)
	fmt.Println("┤")
	
	fmt.Printf("│%-72s│\n", fmt.Sprintf(" Total Menu   : %d", totalMenu))
	fmt.Printf("│%-72s│\n", fmt.Sprintf(" Total Tersedia: %d", totalTersedia))
	fmt.Printf("│%-72s│\n", fmt.Sprintf(" Total Habis  : %d", totalMenu-totalTersedia))
	fmt.Printf("│%-72s│\n", fmt.Sprintf(" Rata-rata Harga Keseluruhan: Rp%d", rataTotal))

	if totalMenu > 0 {
		salinan = k
		selectionSort(&salinan, false, true)
		fmt.Print("├")
		garisLurus(72)
		fmt.Println("┤")
		fmt.Printf("│%-72s│\n", fmt.Sprintf(" Menu Termurah: %s - Rp%d", salinan.Data[0].Nama, salinan.Data[0].Harga))
		fmt.Printf("│%-72s│\n", fmt.Sprintf(" Menu Termahal: %s - Rp%d", salinan.Data[salinan.Jumlah-1].Nama, salinan.Data[salinan.Jumlah-1].Harga))
	}
	footerApp()
}

func filterKetersediaan(k KatalogMenu) {
	var pil, count, i int
	var cariBool bool
	var label string

	label = "TERSEDIA"
	count = 0
	i = 0

	fmt.Println()
	fmt.Println("  Filter: 1. Tersedia   2. Habis")

	pil = inputInt("  Pilih: ")
	cariBool = (pil == 1)

	if !cariBool {
		label = "HABIS"
	}

	fmt.Printf("\n  Menu dengan status [%s]:\n", label)
	headerTabel()

	for i < k.Jumlah {
		if k.Data[i].Tersedia == cariBool {
			count++
			cetakMenu(k.Data[i], count)
		}
		i++
	}
	footerApp()

	if count == 0 {
		fmt.Println("  Tidak ada menu dengan status tersebut.")
	} else {
		fmt.Printf("  Total: %d menu\n", count)
	}
}

// =============================================
// MENU ADMIN & PELANGGAN
// =============================================

func menuAdmin() {
	var pilihan int
	var selesai bool 

	selesai = false

	for !selesai {
		headerApp()
		fmt.Printf("│%-72s│\n", " MODE: ADMIN")
		fmt.Print("├")
		garisLurus(72)
		fmt.Println("┤")
		fmt.Printf("│%-72s│\n", " [1] Tampilkan Semua Menu")
		fmt.Printf("│%-72s│\n", " [2] Lihat Detail Menu")
		fmt.Printf("│%-72s│\n", " [3] Tambah Menu")
		fmt.Printf("│%-72s│\n", " [4] Ubah Menu")
		fmt.Printf("│%-72s│\n", " [5] Hapus Menu")
		fmt.Printf("│%-72s│\n", " [6] Cari Menu")
		fmt.Printf("│%-72s│\n", " [7] Urutkan Menu")
		fmt.Printf("│%-72s│\n", " [8] Filter Ketersediaan")
		fmt.Printf("│%-72s│\n", " [9] Statistik Katalog")
		fmt.Printf("│%-72s│\n", " [0] Kembali ke Menu Utama")
		footerApp()

		pilihan = inputInt("  Pilih menu: ")

		if pilihan == 1 {
			tampilkanSemuaMenu(katalog)
		} else if pilihan == 2 {
			tampilkanDetailMenu(katalog)
		} else if pilihan == 3 {
			tambahMenu(&katalog)
		} else if pilihan == 4 {
			ubahMenu(&katalog)
		} else if pilihan == 5 {
			hapusMenu(&katalog)
		} else if pilihan == 6 {
			menuPencarian(katalog)
		} else if pilihan == 7 {
			menuPengurutan(&katalog)
		} else if pilihan == 8 {
			filterKetersediaan(katalog)
		} else if pilihan == 9 {
			tampilkanStatistik(katalog)
		} else if pilihan == 0 {
			selesai = true
		} else {
			fmt.Println("  [!] Pilihan tidak valid.")
		}

		if !selesai {
			tekanEnter()
		}
	}
}

func menuPelanggan() {
	var pilihan int
	var selesai bool 

	selesai = false

	for !selesai {
		headerApp()
		fmt.Printf("│%-72s│\n", " MODE: PELANGGAN")
		fmt.Print("├")
		garisLurus(72)
		fmt.Println("┤")
		fmt.Printf("│%-72s│\n", " [1] Lihat Semua Menu")
		fmt.Printf("│%-72s│\n", " [2] Lihat Detail Menu")
		fmt.Printf("│%-72s│\n", " [3] Cari Menu")
		fmt.Printf("│%-72s│\n", " [4] Urutkan Menu") 
		fmt.Printf("│%-72s│\n", " [5] Filter Menu Tersedia")
		fmt.Printf("│%-72s│\n", " [6] Statistik Cafe")
		fmt.Printf("│%-72s│\n", " [0] Kembali ke Menu Utama")
		footerApp()

		pilihan = inputInt("  Pilih menu: ")

		if pilihan == 1 {
			tampilkanSemuaMenu(katalog)
		} else if pilihan == 2 {
			tampilkanDetailMenu(katalog)
		} else if pilihan == 3 {
			menuPencarian(katalog)
		} else if pilihan == 4 {
			menuPengurutan(&katalog) 
		} else if pilihan == 5 {
			filterKetersediaan(katalog)
		} else if pilihan == 6 {
			tampilkanStatistik(katalog)
		} else if pilihan == 0 {
			selesai = true
		} else {
			fmt.Println("  [!] Pilihan tidak valid.")
		}

		if !selesai {
			tekanEnter()
		}
	}
}

// =============================================
// MAIN
// =============================================

func main() {
	var pilihan int
	var selesai bool 

	selesai = false

	tambahDataAwal()

	for !selesai {
		headerApp()
		fmt.Printf("│%-72s│\n", " Silakan pilih mode akses:")
		fmt.Print("├")
		garisLurus(72)
		fmt.Println("┤")
		fmt.Printf("│%-72s│\n", " [1] Admin     - Kelola katalog menu")
		fmt.Printf("│%-72s│\n", " [2] Pelanggan - Lihat & cari menu")
		fmt.Printf("│%-72s│\n", " [0] Keluar")
		footerApp()
		
		fmt.Println()
		pilihan = inputInt("  Pilih: ")

		if pilihan == 1 {
			menuAdmin()
		} else if pilihan == 2 {
			menuPelanggan()
		} else if pilihan == 0 {
			selesai = true
			fmt.Println()
			fmt.Println("  Terima kasih telah menggunakan Cafe-Menu!")
			fmt.Println("  Sampai jumpa :)")
			fmt.Println()
		} else {
			fmt.Println("  [!] Pilihan tidak valid.")
			tekanEnter()
		}
	}
}