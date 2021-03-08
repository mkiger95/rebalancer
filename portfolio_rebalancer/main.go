package main

func main() {
	assets := readCurrentValues("rebalance.csv")
	changes := calcNewValues(assets)
	writeValues(changes)
}