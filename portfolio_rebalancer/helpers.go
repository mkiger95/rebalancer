package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
	"math"
)

func readCurrentValues(fileName string) (assets Assets) {
	
	fmt.Println("reading csv file..")

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for _, coin := range records {

		if coin[0] == "Name" {
			continue
		}

		curAlloFloat, _ := strconv.ParseFloat(coin[2], 64)
		desAlloFloat, _ := strconv.ParseFloat(coin[3], 64)
		valFloat, _ := strconv.ParseFloat(coin[4], 64)

		asset := Asset{
			Name: 				coin[0],
			Ticker: 			coin[1],
			CurrentAllocation: curAlloFloat,
			DesiredAllocation: 	desAlloFloat,
			Value: 				valFloat,
		}
		assets = append(assets, asset)
	}

	fmt.Println(assets)

	return
}

func calcNewValues(assets Assets) (changes Changes) {
	
	for _, asset := range assets {

		//total invested currently
		total := asset.Value / (asset.CurrentAllocation/100)
		//desired amount to be invested
		desiredValue := total * (asset.DesiredAllocation/100)
		//difference between current investment and desired investment
		diff := asset.CurrentAllocation - asset.DesiredAllocation
		//convert to a percentage string
		diffString := strconv.FormatFloat(diff, 'f', 2, 64)
		diffString = diffString + "%"
		
		
		var action string
		//determine if buy or sell
		if asset.Value > desiredValue {
			action = "sell"
		} else {
			action = "buy"
		}

		//find value needed to buy or sell
		newValue := math.Abs(asset.Value - desiredValue)
		newValueString := strconv.FormatFloat(newValue, 'f', 8, 64)

		change := Change{
			Name:	asset.Name,
			Ticker: asset.Ticker,
			Action: action,
			PercentDiff: diffString,
			Value: newValueString,
		}
		changes = append(changes, change)
	}
	return
}

func writeValues(changes Changes) {
	var rows [][]string
	headers := []string{"Name", "Ticker", "Action", "Difference", "Value"}

	rows = append(rows, headers)

	for _, change := range changes {
		row := []string{change.Name, change.Ticker, change.Action, change.PercentDiff, change.Value}
		rows = append(rows, row)
	}

	csvfile, err := os.Create("rebalance_new.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	for _, row := range rows {
		err := writer.Write(row)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
}

