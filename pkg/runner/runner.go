package runner

import (
	"crypto/rand"
	"log"
	"math/big"
)

var (
	CardPack = cards{sliceOfCards: []string{
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
	UserGroup = users{sliceOfUsers: []user{
		{
			name:       "justin",
			multiplier: 2,
		},
		{
			name:       "erik",
			multiplier: 3,
		},
		{
			name:       "alex",
			multiplier: 2,
		},
		{
			name:       "brian",
			multiplier: 2.5,
		},
		{
			name:       "edward",
			multiplier: 1.22,
		},
		{
			name:       "matt",
			multiplier: 1.66,
		},
	}}
	TeamFormation = teams{teams: []team{
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

	TurnOrder = []string{
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
	sliceOfUsers []user
}

type team struct {
	name  string
	users []string
}

type teams struct {
	teams []team
}

func RandomInt(max int32) int64 {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max)+1))
	if err != nil {
		log.Fatal()
	}
	return r.Int64()
}

func GenerateRandomGameData(c cards, u users, t teams, turnOrder []string, n int) []score {

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

			userTerm := RandomInt(3)

			if userTerm != 0 {
				nextCardPosition := RandomInt(int32(len(remainingCards.sliceOfCards) - 1))
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
	}
	return scores
}
