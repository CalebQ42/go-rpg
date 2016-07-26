package swrpg

import (
	"math"

	"github.com/CalebQ42/go-rpg"
)

const (
	//Success is a success result on a die
	Success = "Success"
	//Failure is a failure result on a die
	Failure = "Failure"
	//Threat is a threat result on a die
	Threat = "Threat"
	//Advantage is an advantage result on a die
	Advantage = "Advantage"
	//Triumph is a triumph result on a die
	Triumph = "Triumph"
	//Despair is a despair result on a die
	Despair = "Despair"
	//LTPoint is a light side point result on a die
	LTPoint = "Light Side Point"
	//DKPoint is a dark side point result on a die
	DKPoint = "Dark Side Point"
)

var (
	//AbilityDie is an instance of an ability die
	AbilityDie = rpg.CreateStringDie([]string{"", Advantage + Advantage, Success, Advantage, Success + Success, Success + Advantage, Advantage, Success})
	//ProficiencyDie is an instance of a proficiency die
	ProficiencyDie = rpg.CreateStringDie([]string{"", Success + Success, Success + Advantage, Success + Success, Success + Advantage, Success, Success + Advantage, Success, Triumph, Advantage + Advantage, Advantage, Advantage + Advantage})
	//DifficultyDie is an instance of a difficulty die
	DifficultyDie = rpg.CreateStringDie([]string{Failure, Threat + Threat, Failure + Failure, "", Failure, Failure + Threat, Threat, Failure})
	//ChallengeDie is an instance of a difficulty die
	ChallengeDie = rpg.CreateStringDie([]string{"", Threat + Threat, Despair, Threat + Threat, Threat, Failure + Threat, Threat, Failure + Threat, Failure, Failure + Failure, Failure, Failure + Failure})
	//ForceDie is an instance of a force die
	ForceDie = rpg.CreateStringDie([]string{DKPoint + DKPoint, DKPoint, LTPoint, DKPoint, LTPoint, DKPoint, LTPoint + LTPoint, DKPoint, LTPoint + LTPoint, DKPoint, LTPoint + LTPoint, DKPoint})
	//BoostDie is an instance of a boost die
	BoostDie = rpg.CreateStringDie([]string{"", "", Advantage, Success + Advantage, Advantage + Advantage, Success})
	//SetbackDie is an instance of a setback die
	SetbackDie = rpg.CreateStringDie([]string{"", "", Threat, Threat, Failure, Failure})
)

//ProcessDiceResults takes in an array of empty interfaces from rolling Dice and
//processes it into an easy to use map.
func ProcessDiceResults(results []interface{}) (out map[string]int) {
	out = make(map[string]int)
	var strs []string
	for _, v := range results {
		strs = append(strs, v.(string))
	}
	for _, v := range strs {
		switch v {
		case Success:
			out[Success]++
		case Success + Success:
			out[Success] += 2
		case Success + Advantage:
			out[Success]++
			out[Advantage]++
		case Advantage:
			out[Advantage]++
		case Advantage + Advantage:
			out[Advantage] += 2
		case Failure:
			out[Failure]++
		case Failure + Failure:
			out[Failure] += 2
		case Failure + Threat:
			out[Failure]++
			out[Threat]++
		case Threat:
			out[Threat]++
		case Threat + Threat:
			out[Threat] += 2
		case Triumph:
			out[Triumph]++
			out[Success]++
		case Despair:
			out[Despair]++
			out[Failure]++
		}
	}
	return
}

//CancelResults takes in a result map (probably from ProcessDiceResults) and cancels
//out Success and Failure and Advantage and Threat, leaving only the necessary results.
//
//Allows you to do something like CancelResults(ProcessDiceResults(dicePool))
func CancelResults(results map[string]int) (out map[string]int) {
	out = make(map[string]int)
	out[Success] = results[Success] - results[Failure]
	out[Advantage] = results[Advantage] - results[Threat]
	out[Despair] = results[Despair]
	out[Triumph] = results[Triumph]
	if out[Success] < 0 {
		out[Failure] = int(math.Abs(float64(out[Success])))
		delete(out, Success)
	}
	if out[Advantage] < 0 {
		out[Threat] = int(math.Abs(float64(out[Advantage])))
		delete(out, Advantage)
	}
	for i, v := range out {
		if v == 0 {
			delete(out, i)
		}
	}
	return
}
