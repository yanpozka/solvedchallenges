package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_LUMINANCE = 21
)

var nR, nC int

//
func main() {

	letters_map := make(map[string][10][8]bool)

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
		// for col := 5; col < 50; col += 9 { showMatrixCh(extractLetter(&matrix, col)) }

		showMatrix(&matrix)

		if !in_answ_reader.Scan() {
			panic("There is no more answer WHY?")
		}

		text := in_answ_reader.Text()
		{
			fmt.Println("Read = ", text)
			col := 5
			for ix := range text {
				lettr := text[ix : ix+1]
				current_letter_matrix := *extractLetter(&matrix, col)

				if _, ok := letters_map[lettr]; !ok {
					letters_map[lettr] = current_letter_matrix
				}
				col += 9 // !!
			}
		}
		reader = nil
	}

	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	e.Encode(letters_map)

	var decodedMap map[string][10][8]bool
	d := gob.NewDecoder(b)
	// Decoding the serialized data

	if err := d.Decode(&decodedMap); err != nil {
		panic(err)
	}

	file_output, errout := os.Create("mapoutput.go")
	if errout != nil {
		panic(errout)
	}
	defer file_output.Close()
	writer := bufio.NewWriter(file_output)

	// Ta da! It is a map!
	fmt.Fprintf(writer, "package main\n\nfunc LettersMap() map[string][10][8]bool {\n\t return %#v \n}\n", decodedMap)

	writer.Flush()
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
func showMatrix(matrix *[30][60]bool) {
	for i := 10; i < 22; i++ {
		for k := 0; k < 60; k++ {
			if matrix[i][k] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
