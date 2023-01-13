package lasagna

const (
	defaultPreparationTimePerLayer = 2
	noodlesQntPerLayer             = 50
	sauceQntPerLayer               = 0.2
	defaultServingRecipe           = 2.0
)

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, AvgPreparationTimePerLayer int) int {
	if AvgPreparationTimePerLayer == 0 {
		AvgPreparationTimePerLayer = defaultPreparationTimePerLayer
	}

	return len(layers) * AvgPreparationTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodlesLayers := 0
	sauceLayers := 0.0

	for _, v := range layers {
		if v == "sauce" {
			sauceLayers++
		}
		if v == "noodles" {
			noodlesLayers++
		}
	}

	return (noodlesQntPerLayer * noodlesLayers), (sauceQntPerLayer * sauceLayers)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(portions []float64, number int) []float64 {
	newPortions := make([]float64, len(portions))
	scale := float64(number) / defaultServingRecipe

	for k, portion := range portions {
		newPortions[k] = portion * scale
	}

	return newPortions
}
