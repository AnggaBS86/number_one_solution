/*
// 3. Buatlah fungsi dengan input string jika string tersebut dibalik
//	dan tetap sama dengan string yang diinputkan maka return true jika tidak return false
*/
package main

import (
	"strings"
)

func IsPalindrome(str string, caseSensitive bool) bool {
	parsedStr := str
	length := len(str)

	//Jika non-case sensitive maka ubah ke huruf kecil semua
	if false == caseSensitive {
		parsedStr = strings.ToLower(str)
	}

	for i := 0; i < length/2; i++ {
		if parsedStr[i] != parsedStr[length-i-1] {
			return false
		}
	}

	return true
}
