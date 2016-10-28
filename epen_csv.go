package epen_csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
)

type EpenCSV struct {
	DataRows         [][]string
	Header           []string
	comma            rune
	trimLeadingSpace bool
}

func NewCSV(filePath string, settings map[string]interface{}) (*EpenCSV, error) {

	defaultSetting := map[string]interface{}{
		"comma":            ',',
		"trimLeadingSpace": true,
	}

	if settings["comma"] == nil {
		settings["comma"] = defaultSetting["comma"]
	}

	if settings["trimLeadingSpace"] == nil {
		settings["trimLeadingSpace"] = defaultSetting["trimLeadingSpace"]
	}

	if settings["comma"] != nil && reflect.TypeOf(settings["comma"]) != reflect.TypeOf(defaultSetting["comma"]) {
		return nil, errors.New("comma has to be rune type")
	}

	if settings["trimLeadingSpace"] != nil && reflect.TypeOf(settings["trimLeadingSpace"]) != reflect.TypeOf(defaultSetting["trimLeadingSpace"]) {
		return nil, errors.New("trimLeadingSpace has to be bool type")
	}

	r := &EpenCSV{
		comma:            settings["comma"].(rune),
		trimLeadingSpace: settings["trimLeadingSpace"].(bool),
	}

	// Read CSV
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = r.comma
	reader.TrimLeadingSpace = r.trimLeadingSpace
	rows, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	r.Header = rows[0]

	r.DataRows = rows[1:]

	return r, nil
}

func checkColumnIndex(r *EpenCSV, column int) error {
	if column > len(r.Header) {
		return errors.New("Index is out of bound")
	}
	return nil
}

func (r *EpenCSV) GetMean(column int) (float64, error) {

	column--
	e := checkColumnIndex(r, column)

	if e != nil {
		return .0, e
	}

	var data_sum, counter float64
	for _, row := range r.DataRows {
		data, err := strconv.ParseFloat(row[column], 64)

		if err != nil {
			return .0, err
			break
		}

		data_sum += data
		counter++
	}
	return data_sum / counter, nil
}

func (r *EpenCSV) GetMedian(column int) (float64, error) {

	column--
	var data_holder []float64

	e := checkColumnIndex(r, column)

	if e != nil {
		return .0, e
	}

	for _, row := range r.DataRows {
		data, err := strconv.ParseFloat(row[column], 64)

		if err != nil {
			return .0, err
			break
		}

		data_holder = append(data_holder, data)
	}

	sort.Float64s(data_holder) // sort slice

	middle := len(data_holder) / 2

	if len(data_holder)%2 != 0 {
		return data_holder[middle], nil
	} else {
		higher := data_holder[middle]
		lower := data_holder[middle-1]
		return (higher + lower) / 2, nil
	}
}

func (r *EpenCSV) Print() {
	// Print the header
	for index, row := range r.Header {
		if len(r.Header)-1 == index {
			fmt.Printf("%s\n", row)
		} else {
			fmt.Printf("%s\t", row)
		}
	}

	//Print the data
	for _, rows := range r.DataRows {
		for i, row := range rows {
			if len(rows)-1 == i {
				fmt.Printf("%s\n", row)
			} else {
				fmt.Printf("%s\t", row)
			}
		}
	}
}
