package runner

import (
	"crypto/rand"
	"log"
	"math/big"
)

var (
	CardPack = cards{SliceOfCards: []string{
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
	UserGroup = Users{SliceOfUsers: map[string]User{
		"justin":
		{
			Name:       "justin",
			Multiplier: 2,
		},
		"erik":
		{
			Name:       "erik",
			Multiplier: 3,
		},
		"alex":
		{
			Name:       "alex",
			Multiplier: 2,
		},
		"brian":
		{
			Name:       "brian",
			Multiplier: 2.5,
		},
		"edward":
		{
			Name:       "edward",
			Multiplier: 1.22,
		},
		"matt":
		{
			Name:       "matt",
			Multiplier: 1.66,
		},
	}}
	TeamFormation = Teams{Teams: []Team{
		{
			Name: "red",
			Users: []string{
				"justin",
				"alex",
				"erik",
			},
		},
		{
			Name: "blue",
			Users: []string{
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

type Score map[string]string

type cards struct {
	SliceOfCards []string
}

type User struct {
	Name       string
	Multiplier float64
}

type Users struct {
	SliceOfUsers map[string]User
}

type Team struct {
	Name  string
	Users []string
}

type Teams struct {
	Teams []Team
}

func randomInt(max int32) int64 {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max)+1))
	if err != nil {
		log.Fatal()
	}
	return r.Int64()
}

func GenerateRandomGameData(c cards, turnOrder []string, n int) []Score {

	var turnOrderPosition int
	var remainingCards cards
	var s Score = make(map[string]string)
	var scores []Score
	var tempCardSlice = make([]string, len(c.SliceOfCards))

	for i := 0; i < n; i += 1 {
		remainingCards.SliceOfCards = c.SliceOfCards
		for ; ; {
			tempCardSlice = []string{}
			if turnOrderPosition > len(turnOrder)-1 {
				turnOrderPosition = 0
			}

			userTerm := randomInt(3)

			if userTerm != 0 {
				nextCardPosition := randomInt(int32(len(remainingCards.SliceOfCards) - 1))
				nextCard := remainingCards.SliceOfCards[nextCardPosition]
				s[nextCard] = turnOrder[turnOrderPosition]

				for position := 0; position < len(remainingCards.SliceOfCards)-1; position += 1 {
					if position != int(nextCardPosition) {
						tempCardSlice = append(tempCardSlice, remainingCards.SliceOfCards[position])
					}
				}
				remainingCards.SliceOfCards = tempCardSlice
			}
			if len(remainingCards.SliceOfCards) == 0 {
				break
			}
			turnOrderPosition += 1
		}
		scores = append(scores, s)
		s = Score{}
	}
	return scores
}

func RunGame(s []Score) string{

	var maxPoints float64
	var winner string

	playerFinalScores := map[string]float64{}
	teamFinalScores := map[string]float64{}

	for key, round := range s {
		log.Printf("ROUND %d\n", key+1)
		playerScores := map[string]float64{}
		teamScores := map[string]float64{}
		for _, value := range TurnOrder {
			playerScores[value] = 0
		}
		for _, value := range round {
			playerScores[value] = playerScores[value] + (1 * UserGroup.SliceOfUsers[value].Multiplier)
		}
		for key, playerScore := range playerScores {
			for _, team := range TeamFormation.Teams {
				for _, participant := range team.Users {
					if participant == key {
						teamScores[team.Name] = teamScores[team.Name] + playerScore
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
