// Package rocketc is fast, simple and lightweight library for CSV data manipulation and mathematical computation involving 2D Matrices.
package rocketc

import (
	"fmt"
	"strconv"
)

// DataFrame : Basic data container, stores data in form of 2D slices of string.
type DataFrame [][]string

// Rows : Returns number of rows in DataFrame.
func (d DataFrame) Rows() int {
	return len(d)
}

// Cols : Returns number of columns in DataFrame, DataFrame must be uniform for accurate result.
func (d DataFrame) Cols() int {
	if len(d) != 0 {
		return len(d[0])
	}
	return 0
}

// Shape : Returns shape of DataFrame, slice of length 2.
func (d DataFrame) Shape() []int {
	var size = make([]int, 2, 2)
	size[0] = d.Rows()
	size[1] = d.Cols()
	return size
}

// Headers : Returns header of the dataframe i.e row 0.
func (d DataFrame) Headers() []string {
	return d[0]
}

// Head : Returns first n rows of DataFrame including headers.
func (d DataFrame) Head(n int) DataFrame {
	if n <= len(d) {
		return d[0:n]
	}
	return d[0:]
}

// SetHeaders : Set custom column names to a DataFrame.
// Takes a slice of string containing name of columns.
func (d *DataFrame) SetHeaders(header []string) {
	newDataFrame := Allocate(d.Rows()+1, len(header))
	newDataFrame[0] = header
	r := newDataFrame.Rows()
	for i := 1; i < r; i++ {
		newDataFrame[i] = (*d)[i-1]
	}
	*d = newDataFrame
}

// Allocate : Allocate a blank DataFrame of given size.
func Allocate(row, col int) DataFrame {
	var d = make(DataFrame, row)
	for i := 0; i < row; i++ {
		d[i] = make([]string, col)
	}
	return d
}

// WipeDown : Returns unifom DataFrame by only including rows of length l
// in returned DataFrame. Takes a DataFrame and a integer as arguments.
func WipeDown(m DataFrame, l int) DataFrame {
	var r DataFrame
	n := m.Rows()
	for i := 0; i < n; i++ {
		value := m[i]
		if len(value) == l {
			r = append(r, value)
		}
	}
	return r
}

// DropColumn : Drops columns from a DataFrame, takes a DataFrame and variable number of arguments
// which are indexes of columns to be droped.
func DropColumn(d DataFrame, i ...int) DataFrame {
	f := func(arr []int) int {
		var max = arr[0]
		for _, value := range arr {
			if value > max {
				max = value
			}
		}
		return max
	}
	var result = make(DataFrame, len(d))
	var arr = make([]int, f(i)+1)
	for _, value := range i {
		arr[value]++
	}
	for j := 0; j < len(d[0]); j++ {
		for i := 0; i < len(d); i++ {
			if arr[j] > 0 {
				break
			}
			result[i] = append(result[i], d[i][j])
		}
	}
	return result
}

// ConvMatrix : Converts numerical DataFrame into Matrix, returns err if
// dataframe contains values that cannot be converted into a float64.
func ConvMatrix(d DataFrame) (Matrix, error) {
	var m = Zeros(d.Rows(), d.Cols())
	var r = d.Rows()
	var c = d.Cols()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			temp, err := strconv.ParseFloat(d[i][j], 64)
			if err != nil {
				return nil, err
			}
			m[i][j] = float32(temp)
		}
	}
	return m, nil
}

// PrintDataframe : for pretty printing of DataFrame.
func PrintDataframe(d ...DataFrame) {
	lambda := func(d DataFrame) {
		for row := range d {
			fmt.Printf("%3d |", row)
			for col := range d[row] {
				if col < len(d[row])-1 {
					fmt.Printf("%-15s, ", d[row][col])
				} else {
					fmt.Printf("%-15s \n", d[row][col])
				}
			}
		}
	}
	for _, value := range d {
		lambda(value)
		fmt.Println()
	}
}

// GetColumnsDataFrame : Returns a DataFrame by only including specific columns
// whose column indexs are passed as argument. Take a DataFrame and variadic number
// integers which are column indexes.
func GetColumnsDataFrame(d DataFrame, i ...int) DataFrame {
	var c = Allocate(d.Rows(), len(i))
	for row, value := range d {
		for index, v := range i {
			c[row][index] = value[v]
		}
	}
	return c
}
