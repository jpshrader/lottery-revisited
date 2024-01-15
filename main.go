package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	data, err := os.ReadFile("simulation.yml")
	if err != nil {
		log.Fatal(err)
	}

	var config simulationConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

	lot := generateLottery(config)

	winningDistribution := make(map[int]int)

	for i := 0; i < config.NumDrawings; i++ {
		if item, won := lot.draw(); won {
			winningDistribution[item.Priority]++
		}
	}

	lotteryResults := make([]int, 0, config.MaxPriorities)
	for i := 0; i < config.MaxPriorities; i++ {
		res, ok := winningDistribution[i]
		if !ok {
			res = 0
		}
		lotteryResults = append(lotteryResults, res)
	}

	fmt.Println("wins by priority:")
	for priority, count := range lotteryResults {
		population := len(lot.MultilevelFeedbackQueue[priority].Queue)
		fmt.Printf("%d (pop: %d):\t%d\n", priority, population, count)
	}
}
