package builder

import (
	"fmt"
	"strings"
)

type SqlBuilder struct {
	buffer          strings.Builder
	tableName       string
	selectedColumns []string
	conditions      []condition
	orders          []OrderBy
}

func New() *SqlBuilder {
	return &SqlBuilder{}
}

func (builder *SqlBuilder) Table(name string) *SqlBuilder {
	builder.tableName = name
	return builder
}

func (builder *SqlBuilder) Select(column ...string) *SqlBuilder {
	builder.selectedColumns = column
	return builder
}

func (builder *SqlBuilder) Where(cond string, args ...any) *SqlBuilder {
	cond = fmt.Sprintf(cond, args...)
	builder.conditions = append(builder.conditions, condition(cond))
	return builder
}

func (builder *SqlBuilder) ToString() string {
	return builder.buffer.String()
}

func (builder *SqlBuilder) BuildSQL() string {
	selectFrom := builder.selectedColumnsToString()
	table := builder.tableToString()
	where := conditions(builder.conditions).ToString()
	orderBy := builder.ordersToString()

	return selectFrom + table + where + orderBy + endSQL()
}

func (builder *SqlBuilder) selectedColumnsToString() string {
	buff := strings.Builder{}
	buff.WriteString("SELECT ")
	for _, column := range builder.selectedColumns[0 : len(builder.selectedColumns)-1] {
		buff.WriteString(column + ", ")
	}
	buff.WriteString(builder.selectedColumns[len(builder.selectedColumns)-1])

	buff.WriteString(" FROM ")

	return buff.String()
}

func (builder *SqlBuilder) tableToString() string {
	tableName := "*"
	if len(builder.tableName) != 0 {
		tableName = builder.tableName
	}
	return " " + tableName + " "
}

func endSQL() string {
	return ";"
}
