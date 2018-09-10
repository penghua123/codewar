package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("12.255.56.1", is_valid_ip("12.255.56.1"))
	fmt.Println("abc.def.ghi.jkl", is_valid_ip("abc.def.ghi.jkl"))
	fmt.Println("123.456.789.0", is_valid_ip("123.456.789.0"))
	fmt.Println("123.045.067.089", is_valid_ip("123.045.067.089"))
}

func is_valid_ip(ip string) bool {
	dots := strings.Split(ip, ".")
	if len(dots) < 4 {
		return false
	}

	for _, dot := range dots {
		dotLen := len(dot)
		dotInt, err := strconv.Atoi(dot)
		if err != nil {
			return false
		}
		if dotInt > 255 || dotInt < 0 || (dotInt < 100 && dotLen > 2) {
			return false
		}
	}
	return true
}
