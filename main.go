package main

import (
		"fmt"
		) 

func main()  { // Seleksi If, Else, dan Else if 
	var point = 12 

	if point == 12 {
		fmt.Println("lulusan dengan nilai sempurna")
	} else if point > 10 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("Kurang Baik")
	} else {
		fmt.Println("tidak lulus", point)
	}
}