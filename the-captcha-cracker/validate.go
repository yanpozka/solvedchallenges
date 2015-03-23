package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_LUMINANCE   = 21
	MAX_DIFFERENCES = 7
)

var nR, nC int

//
func main() {

	for f_ix := 0; f_ix < 25; f_ix++ {

		file_name := "0" + strconv.Itoa(f_ix) + ".txt"
		if f_ix >= 10 {
			file_name = strconv.Itoa(f_ix) + ".txt"
		}
		fmt.Println(file_name)

		file, errf := os.Open("./input/input" + file_name)
		if errf != nil {
			fmt.Println(errf)
			return
		}
		defer file.Close()

		file_in_answ, errfilein := os.Open("./output/output" + file_name)
		if errfilein != nil {
			panic("[-] Impossible read output file")
		}
		defer file_in_answ.Close()

		in_answ_reader := bufio.NewScanner(file_in_answ)

		//
		fmt.Fscanf(file, "%d%d", &nR, &nC)

		reader := bufio.NewScanner(file)
		var matrix [30][60]bool

		for ix_row := 0; reader.Scan(); ix_row++ {
			row := strings.Split(reader.Text(), " ")

			for ix_col, p := range row {
				pixel := strings.Split(p, ",")
				red, _ := strconv.ParseFloat(pixel[0], 32)
				green, _ := strconv.ParseFloat(pixel[1], 32)
				blue, _ := strconv.ParseFloat(pixel[2], 32)

				// 0.2126 R + 0.7152 G + 0.0722 B
				luminance := red*0.2126 + green*0.7152 + blue*0.0722

				if luminance < MAX_LUMINANCE { // if red < 3 && green < 3 && blue < 3 {
					matrix[ix_row][ix_col] = true
				} else {
					matrix[ix_row][ix_col] = false
				}
			}
		}

		if !in_answ_reader.Scan() {
			panic("There is no more answer WHY?")
		}

		lmap := LettersMap()
		letter_formed := ""
		expected := in_answ_reader.Text()
		ix := 0

		for col := 5; ix < 5; col += 9 {
			current_letter_matrix := extractLetter(&matrix, col)

			val := lmap[expected[ix:ix+1]]
			if compareMatrices(&val, current_letter_matrix) {
				letter_formed += expected[ix : ix+1]
			} else {
				fmt.Println("Something is failing: ")

				showMatrix(&val)
				showMatrix(current_letter_matrix)
			}
			ix++ // !!
		}

		fmt.Println(expected == letter_formed)
		fmt.Println("Expected =", expected, "letter_formed =", letter_formed)

		reader = nil
	}
}

//
func compareMatrices(mtxA *[10][8]bool, mtxB *[10][8]bool) bool {
	diffs_count := 0

	for i := 0; i < 10; i++ {
		for k := 0; k < 8; k++ {
			if mtxA[i][k] != mtxB[i][k] {
				diffs_count++
				if diffs_count >= MAX_DIFFERENCES {
					return false
				}
			}
		}
	}
	return diffs_count < MAX_DIFFERENCES // should be just true, explicit is better than implicit
}

//
func extractLetter(matrix *[30][60]bool, left_top_col int) *[10][8]bool {
	var letter_matrix [10][8]bool
	i := 0

	for ix_row := 11; i < 10; ix_row++ {
		ix_col := left_top_col

		for k := 0; k < 8; ix_col++ {
			letter_matrix[i][k] = matrix[ix_row][ix_col]
			k++ // k !!
		}

		i++ // i !!
	}

	return &letter_matrix
}

//
func showMatrix(matrix *[10][8]bool) {
	for i := 0; i < 10; i++ {
		for k := 0; k < 8; k++ {
			if matrix[i][k] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
