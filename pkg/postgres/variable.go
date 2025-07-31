package postgres

import (
	"fmt"
	"strings"
)

func Array(start int, len int) string {
	if len == 0 {
		return ""
	}

	var builder strings.Builder
	// Pre-allocate approximate capacity: each variable is roughly "$N," format
	builder.Grow(len * 4)

	for i := range len {
		if i > 0 {
			builder.WriteByte(',')
		}
		builder.WriteByte('$')
		builder.WriteString(fmt.Sprintf("%d", start+i))
	}

	return builder.String()
}

func BatchInsert(nVariables int, nItems int) string {
	if nItems == 0 || nVariables == 0 {
		return ""
	}

	var builder strings.Builder
	// Pre-allocate approximate capacity: each item is roughly "($1,$2,...$N)," format
	builder.Grow(nItems * (nVariables*4 + 3))

	for i := range nItems {
		if i > 0 {
			builder.WriteByte(',')
		}
		builder.WriteByte('(')

		for j := range nVariables {
			if j > 0 {
				builder.WriteByte(',')
			}
			builder.WriteByte('$')
			builder.WriteString(fmt.Sprintf("%d", i*nVariables+j+1))
		}

		builder.WriteByte(')')
	}

	return builder.String()
}

func With(s string, name string) string {
	if s == "" {
		return ""
	}

	b := strings.Builder{}
	b.WriteString("WITH ")
	b.WriteString(name)
	b.WriteString(" AS (")
	b.WriteString(s)
	b.WriteString(")")

	return b.String()
}
