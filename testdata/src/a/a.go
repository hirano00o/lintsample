package a

import (
	"fmt"
	"strings"
)

func a() {
	s := []string{"h", "e", "l", "l", "o"}
	for i, v := range s {
		fmt.Println(s[i])
		fmt.Println(v)                    // want "should use identifier of slice or array"
		if strings.Compare(v, "h") == 0 { // want "should use identifier of slice or array"
			fmt.Println("value is h")
		}
	}
}
