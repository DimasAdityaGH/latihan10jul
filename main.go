package main

import "fmt"

func main () {
	//hello world
	fmt.Println("hello, world!")

	//var const string int & type datanumber
	var siswa string = "dimas"
	siswa = "aditya" //var bisa diganti
	fmt.Println(siswa)

	const nomorSiswa int8 = 10
	//nomorSiswa = 11 const nilai tetap
	fmt.Println(nomorSiswa)

	var nilai8 int8 = 127
	var nilai16 int16 = 32767
	var nilai32 int32 = 40000
	var nilai64 int64 = 220202000
	fmt.Println(nilai8)
	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	//perbandingan
	var ten = 10
	var one = 1
	fmt.Println(ten == one)

	//operasi matematika
	var i = 9
	i++
	i += 10
	fmt.Println(i)

	var hasil = 10 * 100
	fmt.Println(hasil)

	//conversion
	var merk = "asus"
	var m = merk[0]
	var mMerk = string(m)
	fmt.Println(merk)
	fmt.Println(mMerk)

	//type dec
	type str string
	type num int8

	var produk str = "susu"
	var harga num = 10
	fmt.Println(produk)
	fmt.Println(harga)

	//boolean

	var married bool = false
	fmt.Println(married)

	//operasi boolean dan array

	var rendi [2]int
	rendi[0] = 75 //0 nilai teori
	rendi[1] = 80 //1 nilai praktek

	var kkmTeori = 75
	var kkmPraktek = 80

	var hasilTeori = rendi[0] >= kkmTeori
	var hasilPraktek = rendi[1] >= kkmPraktek

	var lulus = hasilTeori && hasilPraktek
	fmt.Println(lulus)
}
