package main

var region map[string]string

// makeMap makes a map for each region
func makeMap() {
	// initialize the map
	region = make(map[string]string)
	// start creating the map keys
	// US Regions
	region["E"], region["L"], region["O"], region["T"] = "US", "US", "US", "US"
	// European
	region["D"], region["W"], region["X"], region["Y"],
		region["Z"], region["C"], region["H"], region["S"], region["N"],
		region["M"], region["Q"], region["U"], region["V"], region["I"],
		region["P"] =
		"EN", "EN", "EN", "EN", "EN", "EN", "EN", "EN", "EN", "EN", "EN",
		"EN", "EN", "EN", "EN"
	// Korean
	region["K"] = "KO"
	// Japanese
	region["J"] = "JA"
	// French
	region["F"] = "FR"
	// Russian
	region["R"] = "RU"

}
