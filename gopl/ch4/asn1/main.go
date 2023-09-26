package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

func main() {

	arr := []int{123, 873, 238, 987, 3458}
	bytes, err := asn1.Marshal(arr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d %T\n", bytes, bytes)

	var n []int
	_, err1 := asn1.Unmarshal(bytes, &n)
	checkError(err1)

	fmt.Printf("After bytes/unmarshal: %d", n)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
