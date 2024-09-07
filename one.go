/*
//SOAL NOMOR 1
// 1. Buatlah fungsi dengan parameter berapa banyak angka yg akan ditampilkan secara looping,
// dimana jika angka tersebut habis dibagi 3 cetak "fish" jika habis dibagi 5 cetak "bash"
// dan jika habis dibagi 15 cetak "fish bash", cetak cukup salah satu dari kondisi tersebut,
// cetak menggunakan console log atau sejenisnya. code dalam golang

// Disini saya mempunyai dua (2) solusi, :
// 1. menggunakan modulo atau,
// 2. menggunakan cari nilai terdekat dengan bilangan yang habis dibagi 3 dan 5
*/
package main

import "fmt"

func checkFishBashSolutionUsingModulo(limit int) {
	for i := 1; i <= limit; i++ {
		//jika habis dibagi 15 cetak "fish bash"
		if i%15 == 0 {
			fmt.Println("fish bash")
			//jika habis dibagi 5 cetak "bash"
		} else if i%5 == 0 {
			fmt.Println("bash")
			//jika habis dibagi 3 cetak "fish"
		} else if i%3 == 0 {
			fmt.Println("fish")
		} else {
			fmt.Println(i)
		}
	}
}

func checkFishBashSolutionWithoutModulo(limit int) {
	for i := 1; i <= limit; i++ {
		//menghitung nilai terdekat ketika habis dibagi 3
		divOf3 := i / 3 * 3

		//menghitung nilai terdekat yang habis dibagi 5
		divOf5 := i / 5 * 5

		if divOf3 == i && divOf5 == i {
			fmt.Println("fish bash")
		} else if divOf3 == i {
			fmt.Println("fish")
		} else if divOf5 == i {
			fmt.Println("bash")
		} else {
			fmt.Println(i)
		}
	}
}
