package testSuite

import (
	"fmt"
	"log"
	"projekt2/utils"
	"runtime/debug"
)

func RunFullTests() {
	results := map[string]map[int]float64{}
	percentage := 25
	fmt.Println("Running tests for 25% connected graph")
	log.Println("Running tests for 25% connected graph")
	results["25IMKruskal"], results["25PLKruskal"] = PercentageTestKruskal(percentage)
	results["25IMPrim"], results["25PLPrim"] = PercentageTestPrim(percentage)
	results["25IMDijsktra"], results["25PLDijsktra"] = PercentageTestDijkstra(percentage)
	results["25IMBellmanFord"], results["25PLBellmanFord"] = PercentageTestBellmanFord(percentage)
	fmt.Println("Saving results to CSV files")
	log.Println("Saving results to CSV files")
	for key, value := range results {
		fmt.Println("Saving results for ", key)
		log.Println("Saving results for ", key)
		utils.SaveMapToCSV(value, key)
	}
	debug.FreeOSMemory()
	clear(results)
	results = map[string]map[int]float64{}
	fmt.Println("Running tests for 50% connected graph")
	log.Println("Running tests for 50% connected graph")
	percentage = 50
	results["50IMKruskal"], results["50PLKruskal"] = PercentageTestKruskal(percentage)
	results["50IMPrim"], results["50PLPrim"] = PercentageTestPrim(percentage)
	results["50IMDijsktra"], results["50PLDijsktra"] = PercentageTestDijkstra(percentage)
	results["50IMBellmanFord"], results["50PLBellmanFord"] = PercentageTestBellmanFord(percentage)
	fmt.Println("Saving results to CSV files")
	log.Println("Saving results to CSV files")
	for key, value := range results {
		fmt.Println("Saving results for ", key)
		log.Println("Saving results for ", key)
		utils.SaveMapToCSV(value, key)
	}
	debug.FreeOSMemory()
	clear(results)
	results = map[string]map[int]float64{}
	percentage = 99
	fmt.Println("Running tests for 99% connected graph")
	log.Println("Running tests for 99% connected graph")
	results["99IMKruskal"], results["99PLKruskal"] = PercentageTestKruskal(percentage)
	results["99IMPrim"], results["99PLPrim"] = PercentageTestPrim(percentage)
	results["99IMDijsktra"], results["99PLDijsktra"] = PercentageTestDijkstra(percentage)
	results["99IMBellmanFord"], results["99PLBellmanFord"] = PercentageTestBellmanFord(percentage)
	fmt.Println("Saving results to CSV files")
	log.Println("Saving results to CSV files")
	for key, value := range results {
		fmt.Println("Saving results for ", key)
		log.Println("Saving results for ", key)
		utils.SaveMapToCSV(value, key)
	}

}
