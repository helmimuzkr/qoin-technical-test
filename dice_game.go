package main

import (
	"fmt"
	"math/rand"
	"time"
)

func diceGame(players int, dices int) {
	// Init seed to set the default source
	// if want to get the different behavior/value each run
	rand.Seed(time.Now().UnixNano())

	playersPoint := make(map[int]int)
	dicePlayers := make(map[int][]int)
	for i := 1; i <= players; i++ {
		for j := 1; j <= dices; j++ {
			dicePlayers[i] = append(dicePlayers[i], j)
		}
	}

	count := 1
	for {
		fmt.Println("==================")
		// Rolling the dice.
		fmt.Println("Giliran ", count)
		count++
		for i := 1; i <= len(dicePlayers); i++ {
			for j := 0; j < len(dicePlayers[i]); j++ {
				roll := rand.Intn(6) + 1
				dicePlayers[i][j] = roll
			}
			fmt.Printf("\tPemain #%d (%d): %v\n", i, playersPoint[i], dicePlayers[i])
		}

		// Evaluating the result.
		fmt.Println("Setelah evaluasi:")
		for i := 1; i <= len(dicePlayers); i++ {
			for j := 0; j < len(dicePlayers[i]); j++ {
				if dicePlayers[i][j] == 6 {
					playersPoint[i]++
					dicePlayers[i] = append(dicePlayers[i][:j], dicePlayers[i][j+1:]...)
				} else if dicePlayers[i][j] == 1 {
					// Di sini saya memberikan -1 dikarenakan program akan mengoper angka 1 ketika diterima.
					// Misal, player #1 mendapatkan dadu angka 1, lalu diberikan ke player #2 pada giliran pertama.
					// Pada giliran kedua, player #2 akan memberikan kembali dadu angka 1 ke player #3.
					// Ini terjadi dikarenakan setiap iterasi jika menemukan dadu angka 1 akan dioper ke player selanjutnya.
					if i == players {
						// Ada bug ketika player #3 mengoper dadu angka 1 ke player #1. disaat program dijalankan,
						// value -1 tidak ditambahkan ke array dicePlayers[1].
						dicePlayers[1] = append(dicePlayers[1], -1)
					} else {
						dicePlayers[i+1] = append(dicePlayers[i+1], -1)
					}
					dicePlayers[i] = append(dicePlayers[i][:j], dicePlayers[i][j+1:]...)
				}
			}
			fmt.Printf("\tPemain #%d (%d): %v\n", i, playersPoint[i], dicePlayers[i])
		}

		// Check the number of players who have completed the game.
		var finishedPlayer int
		for _, v := range dicePlayers {
			if len(v) == 0 {
				finishedPlayer++
			}
		}
		if finishedPlayer == len(dicePlayers)-1 {
			break
		}

	}

	// Determine the winner and loser.
	var winner, loser, max int
	for player, point := range playersPoint {
		if point > max && len(dicePlayers[player]) == 0 {
			max = point
			winner = player
		}
		if len(dicePlayers[player]) > 0 {
			loser = player
		}
	}
	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", loser)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya\n", winner)
}

func main() {
	var players, dices int
	fmt.Printf("Masukan jumlah player: ")
	fmt.Scanf("%d\n", &players)
	fmt.Printf("Masukan jumlah dadu: ")
	fmt.Scanf("%d\n", &dices)
	diceGame(players, dices)
}
