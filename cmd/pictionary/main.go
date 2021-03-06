package main

import (
	"fmt"
	"github.com/viniciusramosdefaria/zerohashchallenge/pkg/runner"
)

func main()  {
	fmt.Println(runner.GenerateRandomGameData(runner.CardPack,runner.UserGroup, runner.TeamFormation,runner.TurnOrder, 3))
}
