package postgres

import (
	"fmt"
	"strings"
)

func Array(start int, len int) string {
	n := make([]string, len)

	for i := 0; i < len; i++ {
		n[i] = fmt.Sprintf("$%d", start+i)
	}

	return strings.Join(n, ",")
}
