package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

//Matrix type holding data
type Matrix struct {
	data [][]int
}

//New creates new matrix of int
func New(input string) (*Matrix, error) {
	m := &Matrix{}
	s := strings.Split(input, "\n")
	b := len(strings.Split(s[0], " "))
	for _, t := range s {
		t = strings.TrimSpace(t)
		in := strings.Split(t, " ")
		if b != len(in) {
			return nil, fmt.Errorf("error")
		}
		data := make([]int, len(in))
		for i, t := range in {
			a, err := strconv.Atoi(t)
			if err != nil {
				return nil, fmt.Errorf("error")
			}
			data[i] = a
		}
		m.data = append(m.data, data)
	}
	return m, nil
}

//Rows returns the rows of a matrix
func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m.data))
	for i, t := range m.data {
		rows[i] = make([]int, len(t))
		for j, value := range t {
			rows[i][j] = value
		}
	}
	return rows
}

//Cols returns the columns of a matrix
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m.data[0]))
	for i := 0; i < len(m.data[0]); i++ {
		cols[i] = make([]int, len(m.data))
	}
	for i, t := range m.data {
		for j, value := range t {
			cols[j][i] = value
		}
	}
	return cols
}

//Set sets a value in the matrix
func (m Matrix) Set(a int, b int, c int) bool {
	if a < 0 || a >= len(m.data) || b < 0 || b >= len(m.data[0]) {
		return false
	}
	m.data[a][b] = c
	return true
}
