package main

import (
	"fmt"
)

type Mahasiswa struct {
	nama  string
	nilai [5]float64
	rata  float64
	grade string
}

func main() {
	var jumlahMahasiswa int

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlahMahasiswa)

	mahasiswa := make([]Mahasiswa, jumlahMahasiswa)

	for i := 0; i < jumlahMahasiswa; i++ {
		fmt.Printf("Masukkan nama mahasiswa %d: ", i+1)
		fmt.Scan(&mahasiswa[i].nama)

		for j := 0; j < 5; j++ {
			fmt.Printf("Masukkan nilai 5 mata kuliah %d: ", j+1)
			fmt.Scan(&mahasiswa[i].nilai[j])
		}

		mahasiswa[i].rata = hitungRataRata(mahasiswa[i].nilai)
		mahasiswa[i].grade = tentukanNilaiHuruf(mahasiswa[i].rata)
	}

	fmt.Println("\nHasil Data Mahasiswa:")
	for _, m := range mahasiswa {
		fmt.Printf("Nama: %s, Rata-rata: %.2f, Nilai Huruf: %s\n", m.nama, m.rata, m.grade)
	}

	cariMahasiswaTertinggi(mahasiswa)
}

func hitungRataRata(nilai [5]float64) float64 {
	var total float64
	for _, n := range nilai {
		total += n
	}
	return total / 5
}

func tentukanNilaiHuruf(rata float64) string {
	switch {
	case rata >= 80:
		return "A"
	case rata >= 70 && rata < 80:
		return "B"
	case rata >= 60 && rata < 70:
		return "C"
	case rata >= 50 && rata < 60:
		return "D"
	default:
		return "E"
	}
}

func cariMahasiswaTertinggi(mahasiswa []Mahasiswa) {
	tertinggi := mahasiswa[0]
	for _, m := range mahasiswa {
		if m.rata > tertinggi.rata {
			tertinggi = m
		}
	}
	fmt.Printf("\nMahasiswa dengan nilai rata-rata tertinggi:\nNama: %s, Rata-rata: %.2f\n", tertinggi.nama, tertinggi.rata)
}
