package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetUserMove(scanner *bufio.Scanner) Bitboard {
	var square = ""

	for len(square) != 1 {
		fmt.Print("Enter a square to place a piece [1-9]: ")
		scanner.Scan()
		fmt.Println("")
		square = scanner.Text()
	}

	index, err := strconv.ParseInt(square, 10, 32)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return (1 << (index - 1))
}

func main() {
	var p Position
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println()

	result, winner := p.GetResult()
	for result == Ongoing {
		if p.active_side == Crosses {
			p.Print()
			p.MakeMove(GetUserMove(scanner))
		} else {
			p.MakeMove(GetBestMove(p))
		}

		result, winner = p.GetResult()
	}

	p.Print()

	if result == Cat {
		fmt.Println("Cat game")
	} else if winner == Noughts {
		fmt.Println("Noughts won")
	} else {
		fmt.Println("Crosses won")
	}
}
