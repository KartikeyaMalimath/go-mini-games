package tictactoe

import (
	"fmt"
)

type Matrix struct {
	val string
}

//Global Variables
var Player1 Matrix
var Player2 Matrix

func Gamestatus(curgame [3][3]Matrix, curp int) bool {
	var player Matrix
	if curp%2 == 0 {
		player = Matrix{"X"}
	} else {
		player = Matrix{"O"}
	}

	for i := 0; i < 3; i++ {
		if curgame[i][0] == player && curgame[i][1] == player && curgame[i][2] == player {
			return true
		}
	}
	for i := 0; i < 3; i++ {
		if curgame[0][i] == player && curgame[1][i] == player && curgame[2][i] == player {
			return true
		}
	}
	for i := 0; i < 2; i++ {
		if i == 0 {
			if curgame[0][0] == player && curgame[1][1] == player && curgame[2][2] == player {
				return true
			}
		} else {
			if curgame[0][2] == player && curgame[1][1] == player && curgame[2][0] == player {
				return true
			}
		}

	}
	return false
}

func Printgame(curgame [3][3]Matrix) {
	for _, e := range curgame {
		fmt.Println(e)
	}
}

func Positions(in int) (int, int) {
	ret1 := 999
	ret2 := 999
	switch in {
	case 1:
		ret1 = 0
		ret2 = 0
	case 2:
		ret1 = 0
		ret2 = 1
	case 3:
		ret1 = 0
		ret2 = 2
	case 4:
		ret1 = 1
		ret2 = 0
	case 5:
		ret1 = 1
		ret2 = 1
	case 6:
		ret1 = 1
		ret2 = 2
	case 7:
		ret1 = 2
		ret2 = 0
	case 8:
		ret1 = 2
		ret2 = 1
	case 9:
		ret1 = 2
		ret2 = 2
	}
	return ret1, ret2
}

func Gameover() {
	fmt.Println("===Game Ended===")
}

func TicTacToe() {
	Player1 = Matrix{"X"}
	Player2 = Matrix{"O"}
	var inputstatus [9]int
	gameover := false
	curPlayer := 0
	var input int
	game := [3][3]Matrix{{{"1"}, {"2"}, {"3"}}, {{"4"}, {"5"}, {"6"}}, {{"7"}, {"8"}, {"9"}}}
	fmt.Println("===Welcome to Tic Tac Toe Version 1.0===")
	Printgame(game)
	fmt.Println("===Let the game begin===")
	for !gameover {
		fmt.Println("========================")
		if curPlayer%2 == 0 {
			fmt.Print("Player : 1 Enter Position for X : ")
		} else {
			fmt.Print("Player : 2 Enter Position for O : ")
		}
		fmt.Scan(&input)
		i, j := Positions(input)
		if game[i][j] != Player1 && game[i][j] != Player2 {
			if curPlayer%2 == 0 {
				game[i][j] = Player1
			} else {
				game[i][j] = Player2
			}
			inputstatus[input-1] = input
			gameover = Gamestatus(game, curPlayer)
			curPlayer += 1
		} else {
			fmt.Println("It's already been Marked!")
		}
		fmt.Println("------------------------")
		Printgame(game)
		if gameover {
			fmt.Println("========================")
			fmt.Println("Player :", (((curPlayer - 1) % 2) + 1), " Wins! Yaay!")
			fmt.Println("========================")
		} else {
			temp := 0
			for _, i := range inputstatus {
				temp = temp + i
			}
			if temp == 45 {
				gameover = true
				fmt.Println("=================================")
				fmt.Println("==== It's a Draw! Game Over! ====")
				fmt.Println("=================================")
			}
		}
	}
	Gameover()
}