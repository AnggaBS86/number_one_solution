/*
// 3. Buatlah fungsi dengan input string jika string tersebut dibalik
//	dan tetap sama dengan string yang diinputkan maka return true jika tidak return false
*/
package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		str           string
		caseSensitive bool
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test IsPalindrome with case sensitive false", args{"Malam", false}, true},
		{"Test IsPalindrome with case sensitive true", args{"Malam", true}, false},
		{"Test IsPalindrome with case sensitive true", args{"MalaM", false}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.str, tt.args.caseSensitive); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
