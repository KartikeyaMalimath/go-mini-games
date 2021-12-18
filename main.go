package main

import (
	"fmt"
	"github.com/KartikeyaMalimath/go-mini-games/tictactoe"
)

func main() {
	gametime := true
	var input int
	fmt.Println("===================================")
	fmt.Println("Welcome to KM's Go-lang Mini Games")
	fmt.Printf("===================================\n\n")
	for gametime {
		fmt.Printf("------------------------\nSelect the game you wanna Play!\n 1.) Tic Tac Toe\n 2.) Exit\n------------------------\nEnter your choice: ")
		fmt.Scan(&input)
		switch(input){
		case 1: tictactoe.TicTacToe()
		case 2: fmt.Println("============================")
		fmt.Println("==== Bye! See you Later ====")
		fmt.Println("============================")
		gametime = false
		default: fmt.Println("***** Please Enter Valid Option *****")
		}
	}
	

}
