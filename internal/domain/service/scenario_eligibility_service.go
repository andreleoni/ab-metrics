package service

import (
	"ab-metrics/internal/domain/entity"
	"crypto/sha256"
	"fmt"
)

type ScenarioEligibilityService interface {
	GetVariation(ScenarioEligibilityServiceInput) ScenarioEligibilityServiceOutput
}

type ScenarioEligibilityServiceInput struct {
	Identifier string
	UniqueKey  string
	Variations []entity.Variation
}

type ScenarioEligibilityServiceOutput struct {
	Variation entity.Variation
}

type scenarioEligibilityService struct{}

func NewScenarioEligibilityService() scenarioEligibilityService {
	return scenarioEligibilityService{}
}

func (scenarioEligibilityService) GetVariation(
	sesi ScenarioEligibilityServiceInput) ScenarioEligibilityServiceOutput {

	// Calculate total percentage across all variations in all experiments
	totalPercentage := 0

	for _, variation := range sesi.Variations {
		totalPercentage += variation.Percentage
	}

	keyForAbTest := fmt.Sprint(sesi.Identifier, sesi.UniqueKey)

	// Generate hash value from the device UUID
	hash := sha256.New()
	hash.Write([]byte(keyForAbTest))
	hashValue := hash.Sum(nil)

	// Convert first 2 bytes of hexadecimal hash to integer
	hashInt := int(hashValue[0])<<8 + int(hashValue[1])

	// Map hash value to a range (0 to totalPercentage)
	hashRange := hashInt % totalPercentage

	// Assign user to a testing group based on the range and probabilities
	cumulativeProb := 0
	for _, variation := range sesi.Variations {
		cumulativeProb += variation.Percentage
		if hashRange < cumulativeProb {
			return ScenarioEligibilityServiceOutput{Variation: variation}
		}
	}

	// Default to control group if no group is assigned
	return ScenarioEligibilityServiceOutput{}
}
