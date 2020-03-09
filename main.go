package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type data struct {
	targetResponseTime float64
	target             string
	statusCode         int
}

type plotValues struct {
	avg   float64
	p90   float64
	p99   float64
	p999  float64
	p9999 float64
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

	p90, p99, p999, p9999 := findPercentiles(datapoints)
	avg := findAvg(datapoints)
	points := plotValues{
		avg,
		p90,
		p99,
		p999,
		p9999,
	}
	createPlot(points)
}

func createPlot(points plotValues) {
	p, err := plot.New()
	handleErr(err)
	p.Title.Text = "Latencies"
	p.Y.Label.Text = "Latency"

	pointsForAvg := plotter.Values{points.avg}
	pointsForp90 := plotter.Values{points.p90}
	pointsForp99 := plotter.Values{points.p99}
	pointsForp999 := plotter.Values{points.p999}
	pointsForp9999 := plotter.Values{points.p9999}

	w := vg.Points(20)

	barsForAvg, err := plotter.NewBarChart(pointsForAvg, w)
	handleErr(err)
	barsForAvg.Color = plotutil.Color(0)
	barsForAvg.Offset = -2 * w

	barsForp90, err := plotter.NewBarChart(pointsForp90, w)
	handleErr(err)
	barsForp90.Color = plotutil.Color(1)
	barsForp90.Offset = -w

	barsForp99, err := plotter.NewBarChart(pointsForp99, w)
	handleErr(err)
	barsForp99.Color = plotutil.Color(2)

	barsForp999, err := plotter.NewBarChart(pointsForp999, w)
	handleErr(err)
	barsForp999.Color = plotutil.Color(3)
	barsForp999.Offset = w

	barsForp9999, err := plotter.NewBarChart(pointsForp9999, w)
	handleErr(err)
	barsForp9999.Color = plotutil.Color(4)
	barsForp9999.Offset = 2 * w

	p.Add(barsForAvg, barsForp90, barsForp99, barsForp999, barsForp9999)
	p.Legend.Add("Average", barsForAvg)
	p.Legend.Add("p90", barsForp90)
	p.Legend.Add("p99", barsForp99)
	p.Legend.Add("p99.9", barsForp999)
	p.Legend.Add("p99.99", barsForp9999)
	p.Legend.Top = true
	p.NominalX("")

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "plot.png"); err != nil {
		panic(err)
	}
}

func findPercentiles(datapoints []data) (p90, p99, p999, p9999 float64) {
	n := len(datapoints)
	p90 = datapoints[int(math.Ceil(float64(n)*0.90))].targetResponseTime
	p99 = datapoints[int(math.Ceil(float64(n)*0.99))].targetResponseTime
	p999 = datapoints[int(math.Ceil(float64(n)*0.999))].targetResponseTime
	p9999 = datapoints[int(math.Ceil(float64(n)*0.9999))].targetResponseTime
	return
}

func findAvg(datapoints []data) float64 {
	avg := 0.0
	for _, v := range datapoints {
		avg += v.targetResponseTime
	}
	avg /= float64(len(datapoints))
	return avg
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
