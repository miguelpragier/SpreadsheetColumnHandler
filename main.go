package main

import (
	"fmt"
	"strings"
)

type ColumnHandler struct {
	headers []string
}

// Returns Excel's standard last column
func(h*ColumnHandler)lastColumn() string {
	return "XFD"
}

// New Init the headers/columns buffer
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


// NextColumn returns the next column, after the given one.
// The second param "skipColumns" jumps over N columns after the given one, and then return the expected identifier
func(h ColumnHandler) NextColumn(currentColumn string, skipColumns uint) string{
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

				return h.NextColumn(c,skipColumns)
			}

			break
		}
	}

	return c
}

// Columns return a slice of "howManyColumns" columns, including the given one
func(h ColumnHandler)Columns(initialColumn string, howManyColumns uint) []string {
	var (
		a []string
		c= initialColumn
	)

	a = append(a, c)

	for i := uint(0); i < howManyColumns-1; i++ {
		c = h.NextColumn(c, 0)

		if c!="" {
			a = append(a, c)
		}
	}

	return a
}

func main(){
	h:= New()

	fmt.Println(h.NextColumn("A",0))      // Prints "B"
	fmt.Println(h.NextColumn("A",5))      // Prints "G"
	fmt.Println(h.NextColumn("Z",0))      // Prints "AA"
	fmt.Println(h.NextColumn("A",3))      // Prints "E"
	fmt.Println(h.NextColumn("AA",0))     // Prints "AB"
	fmt.Println(h.NextColumn("AA",3))     // Prints "AE"
	fmt.Println(h.NextColumn("AZ",0))     // Prints "BA"
	fmt.Println(h.NextColumn("ZZ",0))     // Prints "AAA"
	fmt.Println(h.NextColumn("XFD",0))    // Prints ""
	fmt.Println(h.NextColumn("XFD",2))    // Prints ""
	fmt.Println(h.NextColumn("XFB",0))    // Prints "XFC"
	fmt.Println(h.NextColumn("XFB",1))    // Prints "XFD"
	fmt.Println(h.NextColumn("XFB",2))    // Prints ""
	fmt.Println(h.NextColumn("XFB",3))    // Prints ""
	fmt.Println(h.NextColumn("323",0))    // Prints ""
	fmt.Println(h.NextColumn("fsdfsd",0)) // Prints ""
	fmt.Println(h.NextColumn("fsdfsd",0)) // Prints ""
	fmt.Println(h.Columns("A",5))         // Prints "[A B C D E]"
	fmt.Println(h.Columns("XFA",6))       // Prints "[XFA XFB XFC XFD]"
}
