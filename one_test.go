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

import (
	"testing"
)

func Test_checkFishBashSolutionUsingModulo(t *testing.T) {
	type args struct {
		limit int
	}

	tests := []struct {
		name string
		args args
	}{
		{"Test checkFishBashSolutionUsingModulo 1", args{limit: 15}},
		{"Test checkFishBashSolutionUsingModulo 2", args{limit: 20}},
		{"Test checkFishBashSolutionUsingModulo 3", args{limit: 25}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkFishBashSolutionUsingModulo(tt.args.limit)
		})
	}
}

func Test_checkFishBashSolutionWithoutModulo(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
	}{
		{"Test Test_checkFishBashSolutionWithoutModulo 1", args{n: 15}},
		{"Test Test_checkFishBashSolutionWithoutModulo 2", args{n: 20}},
		{"Test Test_checkFishBashSolutionWithoutModulo 3", args{n: 25}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkFishBashSolutionWithoutModulo(tt.args.n)
		})
	}
}
