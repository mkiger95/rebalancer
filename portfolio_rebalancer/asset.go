package main

//Asset struct
type Asset struct {
	Name				string
	Ticker				string
	CurrentAllocation	float64
	DesiredAllocation	float64
	Value				float64
}

//Assets array
type Assets []Asset

//Change struct
type Change struct {
	Name			string
	Ticker			string
	Action 			string	//buy or sell
	PercentDiff 	string
	Value 			string
}

//Changes array
type Changes []Change