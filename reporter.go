package main

import (
	"time"
	"fmt"
	"strings"
	"github.com/bmizerany/perks/quantile"
)

func report(results []Result) {
	quants := quantile.NewTargeted(0.95, 0.99)

	total := len(results)
	totalRtt := time.Duration(0)
	min := time.Hour
	max := time.Duration(0)

	for _, result := range results {
		quants.Insert(float64(result.RTT))
		totalRtt += result.RTT

		if result.RTT > max {
			max = result.RTT
		}

		if result.RTT < min {
			min = result.RTT
		}
	}

	fmt.Println("Results")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Requests:", total)

	fmt.Println("Latencies")
	fmt.Printf("\tTotal:\t%s\t\n\n", totalRtt)

	fmt.Printf("\t0.95:\t%s\n", time.Duration(quants.Query(0.95)))
	fmt.Printf("\t0.99:\t%s\n", time.Duration(quants.Query(0.99)))
	fmt.Printf("\t0.999:\t%s\n\n", time.Duration(quants.Query(0.999)))

	fmt.Printf("\tmean\t%s\n", time.Duration(float64(totalRtt) / float64(total)))
	fmt.Printf("\tmin:\t%s\n", min)
	fmt.Printf("\tmax:\t%s\n", max)
}