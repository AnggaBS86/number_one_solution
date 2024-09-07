/*
*
// 2. Buatlah fungsi dengan input array integer untuk
// mengurutkan angka dari kecil kebesar dan besar kekecil dengan manual tanpa fungsi bawaan javascript
//Disini saya menggunakan algoritma Bubble Sort : https://www-geeksforgeeks-org.translate.goog/bubble-sort-algorithm/?_x_tr_sl=en&_x_tr_tl=id&_x_tr_hl=id&_x_tr_pto=tc
*/
package main

import (
	"reflect"
	"testing"
)

func TestIntegerArraySort(t *testing.T) {
	type args struct {
		arr  []int
		mode string
	}

	input := []int{3, 1, 12, 8}
	expectedAscSort := []int{1, 3, 8, 12}
	expectedDescSort := []int{12, 8, 3, 1}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test ASC sort", args{input, "ASC"}, expectedAscSort},
		{"Test DESC sort", args{input, "DESC"}, expectedDescSort},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntegerArraySort(tt.args.arr, tt.args.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntegerArraySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
