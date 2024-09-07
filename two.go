/*
*
// 2. Buatlah fungsi dengan input array integer untuk
// mengurutkan angka dari kecil kebesar dan besar kekecil dengan manual tanpa fungsi bawaan javascript
//Disini saya menggunakan algoritma Bubble Sort : https://www-geeksforgeeks-org.translate.goog/bubble-sort-algorithm/?_x_tr_sl=en&_x_tr_tl=id&_x_tr_hl=id&_x_tr_pto=tc
*/
package main

func IntegerArraySort(arr []int, mode string) []int {
	lengthArr := len(arr)

	switch mode {
	case "ASC":
		//Dalam double loop bersarang, bandingkan tiap elemen dengan elemen yang berdekatan.
		for i := 0; i < lengthArr-1; i++ {
			for j := 0; j < lengthArr-i-1; j++ {
				if arr[j] > arr[j+1] {
					// Tukar supaya urut dari yang lebih besar, OK!
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	case "DESC":
		//dalam 2 looping bandingkan tiap element dengan element yang berdekatan
		for i := 0; i < lengthArr-1; i++ {
			for j := 0; j < lengthArr-i-1; j++ {
				if arr[j] < arr[j+1] {
					// Tukar supaya urut dari yang lebih kecil
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}

	return arr
}

// func main() {
// 	data := []int{3, 9, 45, 1, 88}

// 	fmt.Println("ASC Sort")
// 	fmt.Println(IntegerArraySort(data, "ASC"))
// 	fmt.Println("====================================")
// 	fmt.Println("DESC Sort")
// 	fmt.Println(IntegerArraySort(data, "DESC"))
// }
