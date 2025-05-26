package app

type BreadConfigurations struct {
	ShapeTypes []string
	FlourTypes []string
	ToppingTypes []string
	ScoringTypes []string
}

var availableConfigs = BreadConfigurations{
	ShapeTypes: []string{"round", "oval"},
	FlourTypes: []string{"wholegrain", "white"},
	ToppingTypes: []string{"none", "sesame"},
	ScoringTypes: []string{"none", "plain", "wheat"},
}

type Step struct {
	Name string
	Label string
	Options []string
} 

var availableSteps []Step = []Step{
	{Name: "shape", Label: "Choose Shape", Options: []string{"round", "oval"}},
	{Name: "flour", Label: "Choose Flour", Options: []string{"white", "wholegrain"}},
	{Name: "topping", Label: "Choose Topping", Options: []string{"none", "sesame"}},
	{Name: "scoring", Label: "Choose Scoring", Options: []string{"none", "plain"}},
}
