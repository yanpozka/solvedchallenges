package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	EXIT  = 101 // 101 == 'e'
	EMPTY = 45  // -
	WALL  = 35  // #
)

var Matrix [3][3]uint8

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for i := 0; i < 3 && scanner.Scan(); i++ {
		var line string = scanner.Text()
		for k := range line {
			Matrix[i][k] = line[k]
		}
	}
	fmt.Println(nextMove())
}

//
func nextMove() string {
	if ok, d := _checkExit(); ok {
		return d
	}

	if Matrix[1][2] != WALL {
		return "RIGHT"
	}
	if Matrix[0][1] != WALL {
		return "UP"
	}
	if Matrix[2][1] != WALL {
		return "DOWN"
	}
	if Matrix[1][0] != WALL {
		return "LEFT"
	}
	panic("Impossible to move the bot!")
	return ""
}

//
func _checkExit() (bool, string) {
	if Matrix[0][1] == EXIT {
		return true, "UP"
	}
	if Matrix[1][0] == EXIT {
		return true, "LEFT"
	}
	if Matrix[1][2] == EXIT {
		return true, "RIGHT"
	}
	if Matrix[2][1] == EXIT {
		return true, "DOWN"
	}
	return false, ""
}
