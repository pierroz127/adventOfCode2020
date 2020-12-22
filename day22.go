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
	playRecursiveCardGame(player1, player2, 0)
}

func solveDay22Part1() {
	player1, player2 := getDay22Data()
	playCardGame(player1, player2)
}

func playCardGame(player1 []int, player2 []int) {
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
	playRecursiveCardGame(player1, player2, 0)
}

func playRecursiveCardGame(deck1 []int, deck2 []int, level int) int {
	playedDecks := make(map[string]bool)
	for true {
		decksKey := getDecksKey(deck1, deck2)
		_, ok := playedDecks[decksKey]
		if ok {
			return 1
		}
		playedDecks[decksKey] = true

		top1, top2 := deck1[0], deck2[0]
		if len(deck1)-1 >= top1 && len(deck2)-1 >= top2 {
			subDeck1 := make([]int, top1)
			copy(subDeck1, deck1[1:1+top1])
			subDeck2 := make([]int, top2)
			copy(subDeck2, deck2[1:1+top2])
			winner := playRecursiveCardGame(subDeck1, subDeck2, level+1)
			if winner == 1 {
				deck1 = append(deck1[1:], top1, top2)
				deck2 = deck2[1:]
			} else {
				deck1 = deck1[1:]
				deck2 = append(deck2[1:], top2, top1)
			}
		} else if top1 > top2 {
			deck1 = append(deck1[1:], top1, top2)
			if len(deck2) == 1 {
				if level == 0 {
					fmt.Printf("Player 1 wins with score %d (%v)\n", computeScore(deck1), deck1)
				}
				return 1
			}
			deck2 = deck2[1:]
		} else {
			deck2 = append(deck2[1:], top2, top1)
			if len(deck1) == 1 {
				if level == 0 {
					fmt.Printf("Player 2 wins with score %d (%v)\n", computeScore(deck2), deck2)
				}
				return 2
			}
			deck1 = deck1[1:]
		}
	}
	return 0
}

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
