package main

import (
	"bufio"
	"fmt"
	"image/jpeg"
	"os"
	"strconv"
)

func main() {

	for index := 12; index <= 12; index++ {
		name := "img"
		if index < 10 {
			name += ("0" + strconv.Itoa(index))
		} else {
			name += strconv.Itoa(index)
		}
		file_img1, errfi := os.Open("./pics/" + name + ".jpg")
		if errfi != nil {
			panic("[-] Error fatal. os.Open")
		}
		defer file_img1.Close()

		file_input, errfin := os.Create("./ins/" + name + ".txt")
		if errfin != nil {
			panic("[-] Error fatal. os.Create")
		}
		writer := bufio.NewWriter(file_input)

		defer func() {
			file_input.Close()
			writer.Flush()
		}()

		img, err := jpeg.Decode(bufio.NewReader(file_img1))
		if err != nil {
			panic("[-] Error trying to create a jpeg.Decode.")
		}
		s := img.Bounds().Size()
		var W, H int = s.X, s.Y

		for i := 0; i < H; i++ {
			for k := 0; k < W; k++ {
				pixel := img.At(i, k)
				var r, g, b, _ uint32 = pixel.RGBA()

				fmt.Fprintf(writer, "%d,%d,%d", uint(r>>8), uint(g>>8), uint(b>>8))
				if k < W-1 {
					fmt.Fprintf(writer, " ")
				}
			}
			fmt.Fprintln(writer)
		}
	}
}
