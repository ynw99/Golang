package main

import (
	"fmt"
	"strings"
)

func main() {
	// Package strings
	fmt.Println("Is 'Nur' in string?", strings.Contains("Yusuf Nur Wahid", "Nur"))
	fmt.Println("Split from string", strings.Split("Yusuf Nur Wahid", " "))
	fmt.Println("To lower strings:", strings.ToLower("Yusuf Nur Wahid"))
	fmt.Println("To upper strings:", strings.ToUpper("Yusuf Nur Wahid"))
	fmt.Println("To title strings:", strings.ToTitle("Yusuf Nur Wahid"))
	fmt.Println("Trim the left and right whitelines strings:", strings.Trim("   Yusuf    ", " "))
	fmt.Println("Replace all:", strings.ReplaceAll("Yusuf Nur Yusuf", "Yusuf", "Ucup"))
}
