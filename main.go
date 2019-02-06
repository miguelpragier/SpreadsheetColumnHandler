package main

import (
	"fmt"
	"strings"
)

type ColumnHandler struct {
	headers []string
}

func(h*ColumnHandler)lastColumn() string {
	return "XFD"
}

func New() ColumnHandler {
	var h ColumnHandler

	lastColumn:=h.lastColumn()

	a := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

	columns := make([][]string, 3)

	for _, s1 := range a {
		columns[0] = append(columns[0], s1)

		for _, s2 := range a {
			t2 := fmt.Sprintf("%s%s", s1, s2)
			columns[1] = append(columns[1], t2)

			for _, s3 := range a {
				t3 := fmt.Sprintf("%s%s%s", s1, s2, s3)

				if len(columns[2]) > 1 && columns[2][len(columns[2])-1] == lastColumn {
					break
				}

				columns[2] = append(columns[2], t3)
			}
		}
	}

	h.headers = columns[0]

	for _, col := range columns[1] {
		h.headers = append(h.headers, col)
	}

	for _, col := range columns[2] {
		h.headers = append(h.headers, col)
	}

	return h
}

func(h ColumnHandler)nextColumn(currentColumn string, skipColumns uint) string{
	currentColumn=strings.ToUpper(currentColumn)

	if currentColumn==""||currentColumn==h.lastColumn() {
		return ""
	}

	c:=""

	for i,col:=range h.headers{
		if col==currentColumn && i<len(h.headers){
			c = h.headers[i+1]

			if skipColumns>0 {
				skipColumns--

				return h.nextColumn(c,skipColumns)
			}

			break
		}
	}

	return c
}

func(h ColumnHandler)Columns(initialColumn string, howManyColumns uint) []string {
	var (
		a []string
		c= initialColumn
	)

	a = append(a, c)

	for i := uint(0); i < howManyColumns-1; i++ {
		c = h.nextColumn(c, 0)

		if c!="" {
			a = append(a, c)
		}
	}

	return a
}

func main(){
	h:= New()

	fmt.Println(h.nextColumn("Z",0))
	fmt.Println(h.nextColumn("A",3))
	fmt.Println(h.nextColumn("AA",0))
	fmt.Println(h.nextColumn("AA",3))
	fmt.Println(h.nextColumn("AZ",0))
	fmt.Println(h.nextColumn("ZZ",0))
	fmt.Println(h.nextColumn("XFD",0))
	fmt.Println(h.nextColumn("XFD",2))
	fmt.Println(h.nextColumn("XFB",0))
	fmt.Println(h.nextColumn("XFB",1))
	fmt.Println(h.nextColumn("XFB",2))
	fmt.Println(h.nextColumn("XFB",3))
	fmt.Println(h.nextColumn("323",0))
	fmt.Println(h.nextColumn("fsdfsd",0))
	fmt.Println(h.nextColumn("fsdfsd",0))
	fmt.Println(h.Columns("A",5))
	fmt.Println(h.Columns("XFA",6))
}
