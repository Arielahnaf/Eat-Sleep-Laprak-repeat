package main

import "fmt"

func main() {
	var nama, nim, kelas string
	fmt.Print("Masukan nama anda: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukan NIM anda : ")
	fmt.Scanln(&nim)

	fmt.Print("Masukan kelas anda : ")
	fmt.Scanln(&kelas)

	fmt.Println("\n---Perkenalan---")

	fmt.Print("\nPerkenalkan nama saya ", nama)

	fmt.Printf("\nsalah satu mahasiswa Prodi S1-IF dari kelas %s ", kelas)

	fmt.Printf("\ndengan NIM %s ", nim)

	fmt.Println("\nSaya sangat bersemangat untuk menjalani kehidupan kuliah saya dan lulus dengan nilai yang bagus ")
}
