package main

import (
	"log"
	"os"

	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

func main() {
	data, err := os.ReadFile("output.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var config simulationConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

}

type simulationConfig struct {
	NumDrawings  int `yaml:"num_drawings"`
	Distribution struct {
		InitialPopulation int `yaml:"initial_population"`
		Multiplier        int `yaml:"multiplier"`
	} `yaml:"distribution"`
	MaxPriorities int `yaml:"max_priorities"`
	DrawingConfig struct {
		Constant   int `yaml:"constant"`
		Multiplier int `yaml:"multiplier"`
	} `yaml:"drawing_config"`
}

type participant struct {
	Id       uuid.UUID
	Priority int
}

func generateLottery() (lottery[participant], error) {
	return lottery[participant]{}, nil
}
