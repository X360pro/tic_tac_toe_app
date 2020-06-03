package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tictactoe/components"
	"tictactoe/service"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome")

	//taking size
	fmt.Print("Enter size of the board : ")
AGAIN:
	size_of_board, _ := reader.ReadString('\n')
	size_of_board = strings.TrimSpace(size_of_board)
	size, err := strconv.Atoi(size_of_board)
	if err != nil || size > 5 {
		fmt.Println("value is incorrect. try again")
		goto AGAIN
	}

	//taking names and marks
	fmt.Print("Enter Name of first player : ")
	name1, _ := reader.ReadString('\n')
	name1 = strings.TrimSpace(name1)
	fmt.Print("Enter Name of second player : ")
	name2, _ := reader.ReadString('\n')
	name2 = strings.TrimSpace(name2)
	fmt.Print("For player 1 Enter 1 to take X and anything else to take O : ")
	mark, _ := reader.ReadString('\n')
	mark = strings.TrimSpace(mark)
	var pl [2]*components.Player
	if mark == "1" {
		pl[0] = components.CreatePlayer(name1, "X")
		pl[1] = components.CreatePlayer(name2, "O")
	} else {
		pl[0] = components.CreatePlayer(name1, "O")
		pl[1] = components.CreatePlayer(name2, "X")
	}

	//starting new game
	newGame := service.NewGameService(pl, size)

	//start playing game
	var res string
	play := true
	fmt.Println(pl[0].Name, " your chance ")
	for {
		fmt.Print(newGame.PrintBoard())
		fmt.Print("Enter position : ")
		pos, _ := reader.ReadString('\n')
		pos = strings.TrimSpace(pos)
		position, err := strconv.Atoi(pos)
		if err != nil {
			fmt.Println("Position should be an integer")
			continue
		}
		if play {
			err, res = newGame.Play(uint8(position))
		} else {
			err, res = newGame.Play(uint8(position))
		}
		if err != nil {
			fmt.Println(err)
			continue
		} else if res == "win" {
			if play {
				fmt.Print(newGame.PrintBoard())
				fmt.Print(pl[0].Name, " has won ")
				break
			} else {
				fmt.Print(newGame.PrintBoard())
				fmt.Print(pl[1].Name, " has won ")
				break
			}
		} else if res == "draw" {
			fmt.Print("Match is draw")
			break
		}
		if play {
			fmt.Print(pl[1].Name, " your chance \n")
		} else {
			fmt.Print(pl[0].Name, " your chance \n")
		}
		play = !play
	}
}
