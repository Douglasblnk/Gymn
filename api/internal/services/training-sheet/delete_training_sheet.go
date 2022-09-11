package trainingSheetService

import (
	trainingSheetRepository "gymn/internal/repository/training-sheet"
	"gymn/internal/utils"
)

func DeleteTrainingSheet(userID int, trainingSheetID string) (string, *utils.Error) {
	trainingSheet, err := trainingSheetRepository.GetTrainingSheetByUID(userID, trainingSheetID)

	if err != nil {
		return "", err
	}

	if err := trainingSheetRepository.DeleteTrainingSheet(trainingSheet); err != nil {
		return "", err
	}

	return "deleted", nil
}
