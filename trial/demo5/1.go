package main

import "fmt"

func main() {
	slice := make([]string, 5, 8)
	slice[0] = "A"
	slice[1] = "B"
	slice[2] = "C"
	slice[3] = "D"
	slice[4] = "E"
	fmt.Println("================")
	slice3 := make([]string, len(slice), cap(slice))
	copy(slice3, slice)
	slice3[0] = "00000"
	inspectSlice(slice)
	inspectSlice(slice3)
}

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
	}
}
