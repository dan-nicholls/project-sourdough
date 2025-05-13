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
