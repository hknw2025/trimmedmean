package trimmedmean

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Average(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func stringsToFloats(strs []string) ([]float64, error) {
	var numbers []float64
	for _, str := range strs {
		str = strings.TrimPrefix(str, "\ufeff") // Remove BOM if present
		num, err := strconv.ParseFloat(str, 64) // Convert string to float
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to float: %v", str, err)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

// flatten takes a slice of slices of integers and returns a single slice
func Flatten(nestedSlices [][]string) []string {
	var flatSlice []string
	for _, subSlice := range nestedSlices {
		flatSlice = append(flatSlice, subSlice...) // Append all elements of subSlice
	}
	return flatSlice
}

func TrimmedMean() {

	// open file from the first user argument
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// Read CSV file using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Flatten the slice
	flat_data := Flatten(data)

	final_data, err := stringsToFloats(flat_data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// read second user inputted argument
	trim1, err := strconv.Atoi(os.Args[2])
	//check if error occured
	if err != nil {
		//executes if there is any error
		fmt.Println(err)
	}

	trim2 := trim1

	arg_len := len(os.Args)
	if arg_len == 4 {
		trim2, err = strconv.Atoi(os.Args[3])
		//check if error occured
		if err != nil {
			//executes if there is any error
			fmt.Println(err)
		}
	}

	filtered := []float64{}

	sort.Float64s(final_data)

	for i, value := range final_data {
		percentile := float64(i) / float64(len(final_data)-1) * 100
		if percentile > float64(trim1) && percentile < float64(100-trim2) {
			filtered = append(filtered, value)
		}
	}

	fmt.Println(Average(filtered))

}
