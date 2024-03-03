package lasagna_master

func PreparationTime(layers []string, avgPrepTimeLayer int) int {
    layers_len := len(layers)

	if avgPrepTimeLayer == 0 {
		return layers_len * 2
	} else {
		return layers_len * avgPrepTimeLayer
	}
}

func Quantities(layers []string) (int, float64) {
	noodles_qty := 0
	sauce_qty := 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodles_qty += 50
		}

		if layer == "sauce" {
			sauce_qty += 0.2
		}
	}

	return noodles_qty, sauce_qty
}