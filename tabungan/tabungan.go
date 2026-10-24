package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tabungan struct {
	nama              string
	target, terkumpul float64
	status            string
}

var reader = bufio.NewReader(os.Stdin)

func masukan(msg string) string {
	fmt.Print(msg)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// Format angka jadi rupiah, contoh: Rp1.500.000
func formatRp(n float64) string {
	s := strconv.FormatFloat(n, 'f', 0, 64)
	negatif := strings.HasPrefix(s, "-")
	if negatif {
		s = s[1:]
	}

	var b strings.Builder
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			b.WriteByte('.')
		}
		b.WriteRune(c)
	}

	hasil := "Rp" + b.String()
	if negatif {
		hasil = "-" + hasil
	}
	return hasil
}

// Perbarui status: "tercapai" jika terkumpul >= target, selain itu "proses"
func perbaruiStatus(t *tabungan) {
	if t.terkumpul >= t.target {
		t.status = "tercapai"
	} else {
		t.status = "proses"
	}
}

func daftarMenu() {
	fmt.Println("1. Tampilkan daftar tabungan")
	fmt.Println("2. Buat tabungan baru")
	fmt.Println("3. Isi tabungan")
	fmt.Println("4. Tarik tabungan")
	fmt.Println("5. Edit tabungan")
	fmt.Println("6. Hapus tabungan")
	fmt.Println("0. Keluar")
}

// Buat tabungan baru
func buatTabunganBaru(daftar []tabungan, namaTabungan string, targetTabungan float64) []tabungan {
	tabunganBaru := tabungan{
		nama:      namaTabungan,
		target:    targetTabungan,
		terkumpul: 0.0,
		status:    "proses",
	}
	fmt.Printf("\nTabungan berhasil dibuat\n")
	return append(daftar, tabunganBaru)
}

// Tampilkan tabungan
func tampilkanTabungan(daftar []tabungan) {
	if len(daftar) == 0 {
		fmt.Println("Tabungan belum pernah dibuat")
		return
	}

	fmt.Printf("%-4s %-25s %-18s %-18s %-18s %-10s\n", "No.", "Nama", "Target", "Terkumpul", "Kekurangan", "Status")
	fmt.Println(strings.Repeat("-", 98))

	for i, v := range daftar {
		fmt.Printf("%-4d %-25s %-18s %-18s %-18s %-10s\n",
			i+1, v.nama, formatRp(v.target), formatRp(v.terkumpul), formatRp(v.target-v.terkumpul), v.status)
	}
}

// Isi tabungan yang sudah dibuat
func isiTabungan(daftar []tabungan, namaTabungan string) {
	for i, v := range daftar {
		if strings.EqualFold(v.nama, namaTabungan) {
			nominal, err := strconv.ParseFloat(masukan("Masukkan nominal: Rp."), 64)

			if err != nil {
				fmt.Println("Nominal harus berupa angka!")
				return
			}

			daftar[i].terkumpul += nominal
			perbaruiStatus(&daftar[i])
			fmt.Printf("\nTabungan %q berhasil diisi\n", v.nama)
			if daftar[i].status == "tercapai" {
				fmt.Printf("Selamat! Target tabungan %q sudah tercapai\n", v.nama)
			}
		} else {
			fmt.Printf("\nNama tabungan tidak ditemukan\n")
		}
	}
}

// Ambil isi Tabungan
func ambilIsiTabungan(daftar []tabungan, namaTabungan string) {
	for i, v := range daftar {
		if strings.EqualFold(v.nama, namaTabungan) {
			nominal, err := strconv.ParseFloat(masukan("Nominal penarikan: Rp."), 64)

			if err != nil {
				fmt.Println("Nominal harus berupa angka!")
				return
			}

			if nominal > v.terkumpul {
				fmt.Printf("\nSaldo tidak mencukupi\n")
				return
			}

			daftar[i].terkumpul -= nominal
			perbaruiStatus(&daftar[i])
			fmt.Printf("\nTabungan berhasil diambil\n")
			break
		} else {
			fmt.Printf("\nNama tabungan tidak ditemukan\n")
		}
	}
}

// Edit tabungan
func editTabungan(daftar []tabungan, namaTabungan string) {
	for i, v := range daftar {
		if strings.EqualFold(v.nama, namaTabungan) {
			namaBaru := masukan("Masukkan nama tabungan baru: ")
			targetBaru, err := strconv.ParseFloat(masukan("Masukkan target tabungan baru: Rp."), 64)

			if err != nil {
				fmt.Printf("\nNominal harus berupa angka\n")
				return
			}

			daftar[i].nama = namaBaru
			daftar[i].target = targetBaru
			perbaruiStatus(&daftar[i])
			fmt.Printf("\nTabungan berhasil diedit\n")

			return
		} else {
			fmt.Printf("\nNama tabungan tidak ditemukan\n")
		}
	}
}

// Hapus tabungan
func hapusTabungan(daftar []tabungan, namaTabungan string) []tabungan {
	for i, v := range daftar {
		if strings.EqualFold(v.nama, namaTabungan) {
			fmt.Printf("\nTabungan %q berhasil dihapus\n", v.nama)
			return append(daftar[:i], daftar[i+1:]...)
		} else {
			fmt.Printf("\nNama tabungan tidak ditemukan\n")
		}
	}
	return daftar
}

func main() {
	daftarTabungan := []tabungan{}
	fmt.Println("Tabungan Golang")
	fmt.Println()

	for {
		daftarMenu()

		pilihan, _ := strconv.Atoi(masukan("Masukkan nomor perintah: "))

		if pilihan == 0 {
			break
		}

		switch pilihan {
		case 1:
			tampilkanTabungan(daftarTabungan)
		case 2:
			nama := masukan("Masukkan nama: ")
			target, err := strconv.ParseFloat(masukan("Masukan target: Rp."), 64)

			if err != nil {
				fmt.Println("Target tabungan harus berupa angka")
			} else {
				daftarTabungan = buatTabunganBaru(daftarTabungan, nama, target)
			}
		case 3:
			nama := masukan("Masukkan nama tabungan: ")
			isiTabungan(daftarTabungan, nama)
		case 4:
			nama := masukan("Masukkan nama tabungan yang ingin diambil: ")
			ambilIsiTabungan(daftarTabungan, nama)
		case 5:
			nama := masukan("Masukkan nama tabungan yang ingin diedit: ")
			editTabungan(daftarTabungan, nama)
		case 6:
			nama := masukan("Masukkan nama tabungan yang ingin dihapus: ")
			daftarTabungan = hapusTabungan(daftarTabungan, nama)
		}
	}
}