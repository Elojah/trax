package postgres

import (
	"fmt"
	"strings"
)

func Array(start int, len int) string {
	n := make([]string, len)

	for i := range len {
		n[i] = fmt.Sprintf("$%d", start+i)
	}

	return strings.Join(n, ",")
}

func BatchInsert(nVariables int, nItems int) string {
	n := make([]string, nItems)

	for i := range nVariables {
		vars := make([]string, nItems)
		for j := range nItems {
			vars[j] = fmt.Sprintf("$%d", i*nItems+j+1)
		}
		n[i] = fmt.Sprintf("(%s)", strings.Join(vars, ","))
	}

	return strings.Join(n, ",")
}
