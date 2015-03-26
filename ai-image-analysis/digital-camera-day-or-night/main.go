package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_LUMINANCE = 32
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	var total, dark_count int = 0, 0

	for reader.Scan() {
		parts := strings.Split(reader.Text(), " ")

		for ixp := range parts {
			total++ // !!
			pixel := strings.Split(parts[ixp], ",")
			if len(pixel) != 3 {
				continue
			}

			red, _ := strconv.ParseFloat(pixel[0], 32)
			green, _ := strconv.ParseFloat(pixel[1], 32)
			blue, _ := strconv.ParseFloat(pixel[2], 32)
			// 0.2126 R + 0.7152 G + 0.0722 B
			luminance := red*0.2126 + green*0.7152 + blue*0.0722

			if luminance < MAX_LUMINANCE {
				dark_count++
			}
		}
	}
	percent := float64(dark_count) * 100.0 / float64(total)

	if percent < 10 {
		fmt.Println("day")
	} else {
		fmt.Println("night")
	}
}
