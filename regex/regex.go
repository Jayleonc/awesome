package main

import (
	"fmt"
	"regexp"
)

func main() {
	num_str := "$200.49、$1,999.00、$99、50.00美元"

	reg := regexp.MustCompile(`(\d+)`)
	submatch := reg.FindAllStringSubmatch(num_str, -1)
	fmt.Println(submatch)

}
