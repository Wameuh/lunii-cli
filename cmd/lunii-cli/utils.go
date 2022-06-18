package main

import "fmt"

func logg(datas ...any) {
	for _, data := range datas {
		fmt.Println(data)
	}
	fmt.Println()
}

func boolToShort(boolean bool) int16 {
	if boolean {
		return 1
	} else {
		return 0
	}
}
