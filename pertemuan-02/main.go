package main

import (
	"errors"
	"fmt"
)

// TODO(Level 1): lihat SOAL.md untuk kontrak lengkap tiap fungsi di bawah.
// Ganti setiap "panic" dengan implementasi yang benar.

func HitungSubtotal(qty int, hargaSatuan float64) float64 {
	return float64(qty) * hargaSatuan
}

func HitungTotalPesanan(qty []int, hargaSatuan []float64) float64 {
	if len(qty) != len(hargaSatuan) {
		return 0
	}

	total := 0.0
	for i := 0; i < len(qty); i++ {
		total += HitungSubtotal(qty[i], hargaSatuan[i])
	}
	return total
}

func TerapkanPajak(total float64, tarifPajak float64) float64 {
	return (total * tarifPajak) + total
}

func HitungDiskon(total float64) float64 {
	if total >= 1000000 {
		return total * 0.10
	} else if total >= 500000 {
		return total * 0.05
	}
	return 0
}

func TotalSetelahDiskon(qty []int, hargaSatuan []float64, tarifPajak float64) float64 {
	subtotal := HitungTotalPesanan(qty, hargaSatuan)
	diskon := HitungDiskon(subtotal)
	setelahDiskon := subtotal - diskon
	return TerapkanPajak(setelahDiskon, tarifPajak)
}

func ValidasiPesanan(qty []int, hargaSatuan []float64) (bool, string) {
	if len(qty) != len(hargaSatuan) {
		return false, "jumlah barang dan harga satuan tidak sama banyak"
	}

	for i := 0; i < len(qty); i++ {
		if qty[i] <= 0 {
			return false, "ada jumlah barang yang bukan bilangan positif"
		}
		if hargaSatuan[i] <= 0 {
			return false, "ada harga satuan yang bukan bilangan positif"
		}
	}

	return true, ""
}

func TentukanStatus(total float64) string {
	if total > 1000000 {
		return "Prioritas"
	} else if total > 100000 {
		return "Reguler"
	}
	return "Hemat"
}

func RingkasanPesanan(qty []int, hargaSatuan []float64, tarifPajak float64) string {
	subtotal := HitungTotalPesanan(qty, hargaSatuan)
	diskon := HitungDiskon(subtotal)
	totalAkhir := TotalSetelahDiskon(qty, hargaSatuan, tarifPajak)
	status := TentukanStatus(totalAkhir)

	return fmt.Sprintf(
		"Subtotal: %.2f\nDiskon: %.2f\nTotal Akhir: %.2f\nStatus: %s",
		subtotal, diskon, totalAkhir, status,
	)
}

// TODO(Level 9): signature ini SUDAH benar (cari tahu sendiri kenapa
// bentuknya begini - lihat SOAL.md) - tinggal implementasikan isinya.
func Total(harga ...float64) float64 {
	total := 0.0
	for _, h := range harga {
		total += h
	}
	return total
}

// TODO(Level 10, bonus): signature ini SUDAH benar (cari tahu sendiri
// kenapa ada dua nilai balik - lihat SOAL.md) - tinggal implementasikan isinya.
func HitungOngkosKirim(beratKg float64, jarakKm float64) (float64, error) {
	if beratKg <= 0 || jarakKm < 0 {
		return 0, errors.New("berat harus lebih dari 0 dan jarak tidak boleh negatif")
	}
	return (beratKg * 2000) + (jarakKm * 3000), nil
}

func main() {
	fmt.Println("Sales Order Processor - pertemuan 2")
}
