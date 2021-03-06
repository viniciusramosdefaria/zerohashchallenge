package runner

import (
	"crypto/rand"
	"log"
	"math/big"
)

var (
	cardPack = cards{sliceOfCards: []string{
		"cat",
		"corolla",
		"crayon",
		"carp",
		"cadillac",
		"corn",
		"camping",
		"condo",
		"camera",
		"cardinal",
		"casino",
	}}
	userGroup = users{sliceOfUsers: map[string]user{
		"justin":
		{
			name:       "justin",
			multiplier: 2,
		},
		"erik":
		{
			name:       "erik",
			multiplier: 3,
		},
		"alex":
		{
			name:       "alex",
			multiplier: 2,
		},
		"brian":
		{
			name:       "brian",
			multiplier: 2.5,
		},
		"edward":
		{
			name:       "edward",
			multiplier: 1.22,
		},
		"matt":
		{
			name:       "matt",
			multiplier: 1.66,
		},
	}}
	teamFormation = teams{teams: []team{
		{
			name: "red",
			users: []string{
				"justin",
				"alex",
				"erik",
			},
		},
		{
			name: "blue",
			users: []string{
				"brian",
				"edward",
				"matt",
			},
		},
	}}

	turnOrder = []string{
		"edward",
		"justin",
		"matt",
		"erik",
		"brian",
		"alex",
	}
)

type score map[string]string

type cards struct {
	sliceOfCards []string
}

type user struct {
	name       string
	multiplier float64
}

type users struct {
	sliceOfUsers map[string]user
}

type team struct {
	name  string
	users []string
}

type teams struct {
	teams []team
}

func randomInt(max int32) int64 {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max)+1))
	if err != nil {
		log.Fatal()
	}
	return r.Int64()
}

func generateRandomGameData(c cards, turnOrder []string, n int) []score {

	var turnOrderPosition int
	var remainingCards cards
	var s score = make(map[string]string)
	var scores []score
	var tempCardSlice = make([]string, len(c.sliceOfCards))

	for i := 0; i < n; i += 1 {
		remainingCards.sliceOfCards = c.sliceOfCards
		for ; ; {
			tempCardSlice = []string{}
			if turnOrderPosition > len(turnOrder)-1 {
				turnOrderPosition = 0
			}

			userTerm := randomInt(3)

			if userTerm != 0 {
				nextCardPosition := randomInt(int32(len(remainingCards.sliceOfCards) - 1))
				nextCard := remainingCards.sliceOfCards[nextCardPosition]
				s[nextCard] = turnOrder[turnOrderPosition]

				for position := 0; position < len(remainingCards.sliceOfCards)-1; position += 1 {
					if position != int(nextCardPosition) {
						tempCardSlice = append(tempCardSlice, remainingCards.sliceOfCards[position])
					}
				}
				remainingCards.sliceOfCards = tempCardSlice
			}
			if len(remainingCards.sliceOfCards) == 0 {
				break
			}
			turnOrderPosition += 1
		}
		scores = append(scores, s)
		s = score{}
	}
	return scores
}

func runGame(s []score) string{

	var maxPoints float64
	var winner string

	playerFinalScores := map[string]float64{}
	teamFinalScores := map[string]float64{}

	for key, round := range s {
		log.Printf("ROUND %d\n", key+1)
		playerScores := map[string]float64{}
		teamScores := map[string]float64{}
		for _, value := range turnOrder {
			playerScores[value] = 0
		}
		for _, value := range round {
			playerScores[value] = playerScores[value] + (1 * userGroup.sliceOfUsers[value].multiplier)
		}
		for key, playerScore := range playerScores {
			for _, team := range teamFormation.teams {
				for _, participant := range team.users {
					if participant == key {
						teamScores[team.name] = teamScores[team.name] + playerScore
					}
				}
			}
		}

		for key, score := range playerScores {
			log.Printf("player name: %s, score: %.2f\n", key, score)
			playerFinalScores[key] += score
		}

		for key, score := range teamScores {
			log.Printf("team name: %s, team score: %.2f\n", key, score)
			teamFinalScores[key] += score
		}
	}
	for key, score := range playerFinalScores {
		log.Printf("player name: %s, final score: %.2f\n", key, score)
	}

	for key, score := range teamFinalScores {
		if score > maxPoints {
			maxPoints = score
			winner = key

		}
		log.Printf("team name: %s, team final score: %.2f\n", key, score)
	}
	return winner
}
