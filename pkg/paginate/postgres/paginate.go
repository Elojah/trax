package postgres

import (
	"fmt"
	"strings"

	"github.com/elojah/trax/pkg/paginate"
)

type Paginate paginate.Paginate

func (p Paginate) Row(sortCols map[string]string) string {
	sort, ok := sortCols[p.Sort]

	b := strings.Builder{}
	b.WriteString(` ,ROW_NUMBER() OVER (`)
	if ok {
		b.WriteString(` ORDER BY `)
		b.WriteString(sort)
		if p.Order {
			b.WriteString(` ASC `)
		} else {
			b.WriteString(` DESC `)
		}
	}
	b.WriteString(` )`)

	return b.String()
}

func (p Paginate) CTE(subquery string) string {
	b := strings.Builder{}
	b.WriteString(`WITH cte as ( `)
	b.WriteString(subquery)
	b.WriteString(fmt.Sprintf(` ) SELECT * FROM cte WHERE ROW_NUMBER BETWEEN %d AND %d`, p.Start, p.End))

	return b.String()
}
