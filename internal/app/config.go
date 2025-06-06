package app

type BreadConfigurations struct {
	ShapeTypes   []string
	FlourTypes   []string
	ToppingTypes []string
	ScoringTypes []string
}

var availableConfigs = BreadConfigurations{
	ShapeTypes:   []string{"round", "oval"},
	FlourTypes:   []string{"wholegrain", "white"},
	ToppingTypes: []string{"none", "sesame"},
	ScoringTypes: []string{"none", "plain", "wheat"},
}

type Option struct {
	Name 	string
	Label 	string
	Options []string
} 

var FormOptions []Option = []Option{
	{Name: "shape", Label: "Choose Shape", Options: availableConfigs.ShapeTypes},
	{Name: "flour", Label: "Choose Flour", Options: availableConfigs.FlourTypes},
	{Name: "topping", Label: "Choose Topping", Options: availableConfigs.ToppingTypes},
	{Name: "scoring", Label: "Choose Scoring", Options: availableConfigs.ScoringTypes},
}
