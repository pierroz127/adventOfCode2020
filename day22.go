package main

import (
	"fmt"
	"strconv"
)

func getPlayerDeckKey(cards []int) string {
	if len(cards) == 0 {
		return "."
	}
	key := strconv.Itoa(cards[0])
	for i := 1; i < len(cards); i++ {
		key = fmt.Sprintf("%s.%d", key, cards[i])
	}
	return key
}

func getDecksKey(player1 []int, player2 []int) string {
	return fmt.Sprintf("%s#%s", getPlayerDeckKey(player1), getPlayerDeckKey(player2))
}

func solveDay22Example() {
	player1, player2 := parseCards([]string{
		"Player 1:",
		"9",
		"2",
		"6",
		"3",
		"1",
		"",
		"Player 2:",
		"5",
		"8",
		"4",
		"7",
		"10",
	})
	playRecursiveCardGame(player1, player2, false)
}

func solveDay22Part1() {
	player1, player2 := getDay22Data()
	playCardGame(player1, player2)
}

func playCardGame(player1 []int, player2 []int) {
	fmt.Printf("Player 1: %v\n", player1)
	fmt.Printf("Player 2: %v\n", player2)
	fmt.Println("*****")

	if player1[0] > player2[0] {
		player1 = append(player1[1:], player1[0], player2[0])
		if len(player2) == 1 {
			fmt.Printf("Player 1 wins with score %d\n", computeScore(player1))
			return
		}
		playCardGame(player1, player2[1:])
	} else {
		player2 = append(player2[1:], player2[0], player1[0])
		if len(player1) == 1 {
			fmt.Printf("Player 2 wins with score %d (%v)\n", computeScore(player2), player2)
			return
		}
		playCardGame(player1[1:], player2)
	}
}

func computeScore(player []int) int {
	score := 0
	n := len(player)
	for i := 1; i <= n; i++ {
		score += i * player[n-i]
	}
	return score
}

func solveDay22Part2() {
	player1, player2 := getDay22Data()
	playRecursiveCardGame(player1, player2, false)
}

type Game struct {
	player1 []int
	player2 []int
}

func playRecursiveCardGame(deck1 []int, deck2 []int, isSubGame bool) int {
	game := &Game{deck1, deck2}
	playedDecks := make(map[string]bool)
	for true {
		// fmt.Printf("Player 1 (length=%d): %v\n", len(game.player1), game.player1)
		// fmt.Printf("Player 2 (length=%d): %v\n", len(game.player2), game.player2)
		// fmt.Println("*****")
		decksKey := getDecksKey(game.player1, game.player2)
		_, ok := playedDecks[decksKey]
		if ok {
			if len(playedDecks)%100 == 0 {
				fmt.Printf("Entering a recursive loop (%d rounds already played) ! Player 1 wins (key: %s)\n", len(playedDecks), decksKey)
			}
			return 1
		}
		playedDecks[decksKey] = true

		top1, top2 := game.player1[0], game.player2[0]
		if len(game.player1)-1 >= top1 && len(game.player2)-1 >= top2 {
			// fmt.Printf("Starting a sub game...\n")
			subPlayer1, subPlayer2 := game.player1[1:], game.player2[1:]
			winner := playRecursiveCardGame(subPlayer1, subPlayer2, true)
			if winner == 1 {
				game = &Game{append(game.player1[1:], top1, top2), game.player2[1:]}
			} else {
				game = &Game{game.player1[1:], append(game.player2[1:], top2, top1)}
			}
		} else if top1 > top2 {
			nextPlayer1 := append(game.player1[1:], top1, top2)
			if len(game.player2) == 1 {
				if !isSubGame {
					fmt.Printf("Player 1 wins with score %d\n", computeScore(nextPlayer1))
				}
				return 1
			}
			game = &Game{nextPlayer1, game.player2[1:]}
		} else {
			nextPlayer2 := append(game.player2[1:], top2, top1)
			if len(game.player1) == 1 {
				if !isSubGame {
					fmt.Printf("Player 2 wins with score %d (%v)\n", computeScore(nextPlayer2), nextPlayer2)
				}
				return 2
			}
			game = &Game{game.player1[1:], nextPlayer2}
		}
	}
	return 0
}

// func playRecursiveCardGame(player1 []int, player2 []int, isSubGame bool) int {
// 	decksKey := getDecksKey(player1, player2)
// 	_, ok := playedDecks[decksKey]
// 	if ok {
// 		fmt.Printf("Entering a recursive loop! Player 1 wins (key: %s)\n", decksKey)
// 		return 1
// 	}
// 	playedDecks[decksKey] = true

// 	top1, top2 := player1[0], player2[0]
// 	if len(player1)-1 >= top1 && len(player2)-1 >= top2 {
// 		subPlayer1, subPlayer2 := player1[1:], player2[1:]
// 		winner := playRecursiveCardGame(subPlayer1, subPlayer2, true)
// 		if winner == 1 {
// 			return playRecursiveCardGame(append(player1[1:], top1, top2), player2[1:], true)
// 		} else {
// 			return playRecursiveCardGame(player1[1:], append(player2[1:], top2, top1), true)
// 		}
// 	} else if top1 > top2 {
// 		player1 = append(player1[1:], top1, top2)
// 		if len(player2) == 1 {
// 			if !isSubGame {
// 				fmt.Printf("Player 1 wins with score %d\n", computeScore(player1))
// 			}
// 			return 1
// 		}
// 		return playRecursiveCardGame(player1, player2[1:], true)
// 	}
// 	player2 = append(player2[1:], top2, top1)
// 	if len(player1) == 1 {
// 		if !isSubGame {
// 		fmt.Printf("Player 2 wins with score %d (%v)\n", computeScore(player2), player2)
// 		}
// 		return 2
// 	}
// 	return playRecursiveCardGame(player1[1:], player2, true)
// }
// ----------
func getDay22Data() ([]int, []int) {
	return parseCards(getDataFromFile("day22"))
}

func parseCards(lines []string) ([]int, []int) {
	player1, player2 := []int{}, []int{}
	idx := 1
	for lines[idx] != "" {
		card, _ := strconv.Atoi(lines[idx])
		player1 = append(player1, card)
		idx++
	}
	idx += 2 // player 2
	for idx < len(lines) {
		card, _ := strconv.Atoi(lines[idx])
		player2 = append(player2, card)
		idx++
	}
	return player1, player2
}
