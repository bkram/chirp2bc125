package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFileName := flag.String("input", "", "Path to the input Chirp CSV file")
	outputFileName := flag.String("output", "", "Path to the output bc125 CSV file")
	startNumber := flag.Int("start-number", 1, "Starting number for the first column")
	flag.Parse()

	if *inputFileName == "" || *outputFileName == "" {
		fmt.Println("Both input and output file paths are required.")
		return
	}

	inputFile, err := os.Open(*inputFileName)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(*outputFileName)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	csvReader := csv.NewReader(inputFile)
	csvWriter := csv.NewWriter(outputFile)

	header := []string{"Channel", "Name", "Frequency", "Modulation", "CTCSS/DCS", "Delay", "Lockout", "Priority"}
	csvWriter.Write(header)

	i := *startNumber
	for {
		row, err := csvReader.Read()
		if err != nil {
			break
		}

		location := strconv.Itoa(i)
		ctcssDcs := convertCTCSSDcs(row[5], row[6])
		writeRowToCSV(csvWriter, location, row[1], row[2], row[10], ctcssDcs)

		i++
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
	}

	fmt.Println("Conversion completed.")
}

func convertCTCSSDcs(tone, rToneFreq string) string {
	if (tone == "Tone" || tone == "SQL") && rToneFreq != "" {
		f, err := strconv.ParseFloat(rToneFreq, 64)
		if err == nil {
			return fmt.Sprintf("%.1f Hz", f)
		}
	}
	return "none"
}

func writeRowToCSV(csvWriter *csv.Writer, location, name, frequency, mode, ctcssDcs string) {
	csvWriter.Write([]string{location, name, frequency, mode, ctcssDcs, "2", "no", "no"})
}
