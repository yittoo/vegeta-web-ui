package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
	vegetaPlot "github.com/tsenart/vegeta/lib/plot"
)

/*
	BODY KEYS
	freq: String. the frequency of requests per second. REQUIRED
	duration: String. the duration of requests. default "5"
	method: String. method of requests. possible options: "GET", "POST"
	target: String. the target url that will be load tested. REQUIRED
	reportType: String. type of expected result. possible options: "graph", "json". Default: "json"
*/
var options = struct {
	freq     string
	duration string
	method   string
	target   string

	reportType string
}{
	freq:       "freq",
	duration:   "duration",
	method:     "method",
	target:     "target",
	reportType: "reportType",
}

// SuccessResult is struct definition for sending json and graph responses together
type SuccessResult struct {
	AsGraphHTML string
	AsGraphJSON string
}

func mapVegetaOptions(j []byte) (map[string]string, error) {
	v := make(map[string]string)
	err := json.Unmarshal(j, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func execVegetaCall(o map[string]string) (string, string, error) {
	freqAsStr, err := checkMapKeyExists(&o, options.freq)
	if err != nil {
		return "", "", err
	}
	freq, err := strconv.Atoi(freqAsStr)
	if err != nil {
		return "", "", err
	}
	rate := vegeta.Rate{Freq: freq, Per: time.Second}

	durationAsStr, err := checkMapKeyExists(&o, options.duration)
	if err != nil {
		// default to 5 second
		durationAsStr = "5"
	}
	dur, err := strconv.Atoi(durationAsStr)
	duration := time.Duration(dur) * time.Second

	method, err := checkMapKeyExists(&o, options.method)
	if err != nil {
		// default to 5 second
		method = "GET"
	}

	target, err := checkMapKeyExists(&o, options.target)
	if err != nil {
		return "", "", err
	}

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: method,
		URL:    target,
	})
	attacker := vegeta.NewAttacker()
	var reporter vegeta.Reporter
	// var bf bytes.Buffer
	var contentType string

	// reportType, err := checkMapKeyExists(&o, options.reportType)
	// if err != nil {
	// 	// default to json
	// 	reportType = "json"
	// }

	// if reportType == "json" {
	// 	contentType = "application/json"
	// 	var metrics vegeta.Metrics
	// 	for res := range attacker.Attack(targeter, rate, duration, "Boom") {
	// 		metrics.Add(res)
	// 	}
	// 	metrics.Close()
	// 	reporter = vegeta.NewJSONReporter(&metrics)
	// 	err = reporter.Report(&bf)
	// 	if err != nil {
	// 		return "", "", err
	// 	}
	// } else if reportType == "graph" {
	// 	contentType = "text/html"
	// 	plot := vegetaPlot.New()
	// 	for res := range attacker.Attack(targeter, rate, duration, "Boom") {
	// 		plot.Add(res)
	// 	}
	// 	plot.Close()
	// 	_, err = plot.WriteTo(&bf)
	// 	if err != nil {
	// 		return "", "", err
	// 	}
	// }

	/* both together */
	contentType = "application/json"
	var jsonBuffer bytes.Buffer
	var graphBuffer bytes.Buffer

	var metrics vegeta.Metrics
	plot := vegetaPlot.New()

	for res := range attacker.Attack(targeter, rate, duration, "Boom") {
		metrics.Add(res)
		plot.Add(res)
	}
	metrics.Close()
	plot.Close()

	reporter = vegeta.NewJSONReporter(&metrics)

	err = reporter.Report(&jsonBuffer)
	if err != nil {
		return "", "", err
	}
	_, err = plot.WriteTo(&graphBuffer)
	if err != nil {
		return "", "", err
	}

	// convert responses into struct
	sr := SuccessResult{
		AsGraphHTML: graphBuffer.String(),
		AsGraphJSON: jsonBuffer.String(),
	}

	// serialize struct into json
	s, err := json.Marshal(sr)
	if err != nil {
		return "", "", err
	}
	/***/

	return string(s), contentType, nil
}

func checkMapKeyExists(options *map[string]string, key string) (string, error) {
	if val, ok := (*options)[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("The key %v does not exist", key)
}
