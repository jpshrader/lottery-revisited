package main

import "github.com/google/uuid"

type simulationConfig struct {
	NumDrawings  int `yaml:"num_drawings"`
	Distribution struct {
		InitialPopulation float64 `yaml:"initial_population"`
		Multiplier        float64 `yaml:"multiplier"`
	} `yaml:"distribution"`
	MaxPriorities int `yaml:"max_priorities"`
	DrawingConfig struct {
		Constant   float64 `yaml:"constant"`
		Multiplier float64 `yaml:"multiplier"`
	} `yaml:"drawing_config"`
}

type participant struct {
	Id       uuid.UUID
	Priority int
}

func generateLottery(config simulationConfig) lottery {
	lot := lottery{
		Constant:                config.DrawingConfig.Constant,
		Multiplier:              config.DrawingConfig.Multiplier,
		MultilevelFeedbackQueue: make(multilevelFeedbackQueue, 0, config.MaxPriorities),
	}
	for lvl := 0; lvl < config.MaxPriorities; lvl++ {
		levelPopulation := int(config.Distribution.InitialPopulation + float64(lvl) * config.Distribution.Multiplier)
		queue := make([]participant, 0, levelPopulation)
		for j := 0; j < levelPopulation; j++ {
			queue = append(queue, participant{
				Id:       uuid.New(),
				Priority: lvl,
			})
		}
		lot.MultilevelFeedbackQueue = append(lot.MultilevelFeedbackQueue, level{
			Queue: queue,
		})
	}
	return lot
}
