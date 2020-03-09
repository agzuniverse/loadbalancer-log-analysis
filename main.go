package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type data struct {
	targetResponseTime float64
	target             string
	statusCode         int
}

func main() {
	fmt.Println("Loading file...")
	f, err := os.Open("log.log")
	handleErr(err)
	defer f.Close()

	datapoints := make([]data, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		params := strings.Split(line, " ")
		if len(params) > 1 {
			datapoint := createDataPoint(params)
			datapoints = append(datapoints, datapoint)
		}
	}

	err = scanner.Err()
	handleErr(err)

	// Sort the data points on targetResponseTime for ease of finding percentiles
	sort.Slice(datapoints, func(i, j int) bool {
		return datapoints[i].targetResponseTime < datapoints[j].targetResponseTime
	})
}

func createDataPoint(params []string) data {
	statusCodeAsInt, err := strconv.Atoi(params[8])
	handleErr(err)
	responseTimeAsFloat, err := strconv.ParseFloat(params[6], 32)
	handleErr(err)
	return data{
		targetResponseTime: responseTimeAsFloat,
		target:             params[4],
		statusCode:         statusCodeAsInt,
	}
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
