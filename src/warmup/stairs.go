package warmup

import (
	"fmt"
	"strings"
)

func Staircase(n int32) {
	for i := int32(1); i <= n; i++ {
		spaces := strings.Repeat(" ", int(n-i))
		sharps := strings.Repeat("#", int(i))
		fmt.Println(spaces + sharps)
	}
}
