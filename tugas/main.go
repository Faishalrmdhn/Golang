//Minggu #1

package main

import (
	"fmt"
)

func cetakGambar(input int) {
	if input % 2 != 1 {
		fmt.Println("Input harus bilangan ganjil")
	
	} else {
		fmt.Println("--- Panjang ---")
		for i := 0 ; i < input; i++ {
			for j := 0; j < input; j++ {
				if j < 1 || j == input-1 || i == input/2 {
					fmt.Print("*  ")
					} else {
						fmt.Print("=  ")
				}
			}
			fmt.Println("\n") 
		}
	}
}

func main() {
	cetakGambar(5)
}