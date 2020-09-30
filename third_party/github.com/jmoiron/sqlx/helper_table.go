// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlx

import (
	"fmt"
	"strings"
)

// shrinkColumns trim empty columns
func shrinkColumns(cols ...string) []string {
	var params []string
	for _, col := range cols {
		if col == "" {
			continue
		}
		params = append(params, col)
	}
	return params
}

// NamedTableColumns returns the []string{table.value1, table.value2 ...}
// query := NamedColumns("table", "foo", "bar")
// // []string{"table.foo", "table.bar"}
func NamedTableColumns(table string, cols ...string) []string {
	cols = shrinkColumns(cols...)
	var namedCols []string
	for _, col := range cols {
		if table == "" {
			namedCols = append(namedCols, col)
		} else {
			namedCols = append(namedCols, fmt.Sprintf("%s.%s", table, col))
		}
	}
	return namedCols
}

// NamedTableValues returns the []string{:value1, :value2 ...}
// query := NamedTableValues("foo", "bar")
// // []string{":foo", ":bar"}
func NamedTableValues(cols ...string) []string {
	cols = shrinkColumns(cols...)

	var namedCols []string
	for _, col := range cols {
		namedCols = append(namedCols, ":"+col)
	}
	return namedCols
}

// NamedColumnsValues returns the []string{table.value1=:value1, table.value2=:value2 ...}
// query := NamedColumnsValues("table", "foo", "bar")
// // []string{"table.foo=:foo", "table.bar=:bar"}
func NamedTableColumnsValues(cmp SqlCompare, table string, cols ...string) []string {
	cols = shrinkColumns(cols...)

	var namedCols []string
	for _, col := range cols {
		if table == "" {
			namedCols = append(namedCols, fmt.Sprintf("%[1]s %[2]s :%[1]s", col, cmp))
		} else {
			namedCols = append(namedCols, fmt.Sprintf("%[1]s.%[2]s %[3]s :%[2]s", table, col, cmp))
		}
	}
	return namedCols
}

// JoinTableColumns concatenates the elements of cols to column1, column2, ...
// query := JoinTableColumns("table", "foo", "bar")
// // "table.foo, table.bar"
func JoinTableColumns(table string, cols ...string) string {
	//cols = shrinkColumns(cols...)
	return strings.Join(NamedTableColumns(table, cols...), ",")
}

// JoinNamedTableValues concatenates the elements of values to :value1, :value2, ...
// query := JoinNamedTableValues("foo", "bar")
// // ":foo,:bar"
// query := JoinNamedTableValues()
// // "DEFAULT"
func JoinNamedTableValues(cols ...string) string {
	cols = shrinkColumns(cols...)
	if len(cols) == 0 {
		// https://dev.mysql.com/doc/refman/5.7/en/data-type-defaults.html
		return "DEFAULT"
	}
	return strings.Join(NamedTableValues(cols...), ",")
}

// JoinNamedTableColumnsValues concatenates the elements of values to table.value1=:value1, table.value2=:value2 ...
// query := JoinNamedTableColumnsValues("table", "foo", "bar")
// // "table.foo=:foo, table.bar=:bar"
func JoinNamedTableColumnsValues(table string, cols ...string) string {
	//cols = shrinkColumns(cols...)
	return strings.Join(NamedTableColumnsValues(SqlCompareEqual, table, cols...), ",")
}

// JoinNamedTableCondition concatenates the elements of values to table.value1=:value1 AND table.value2=:value2 ...
// query := JoinNamedTableCondition(SqlCompareEqual, SqlOperatorAnd, "table", "foo", "bar")
// // "table.foo=:foo AND table.bar=:bar"
func JoinNamedTableCondition(cmp SqlCompare, operator SqlOperator, table string, cols ...string) string {
	//cols = shrinkColumns(cols...)
	return strings.Join(NamedTableColumnsValues(cmp, table, cols...), fmt.Sprintf(" %s ", operator.String()))
}

// JoinNamedTableColumns concatenates the elements of cols in table to column1, column2, ...
// query := JoinNamedTableColumns("table", "foo", "bar")
// // "table.foo, table.bar"
func JoinNamedTableColumns(table string, cols ...string) string {
	//cols = shrinkColumns(cols...)
	return JoinTableColumns(table, cols...)
}