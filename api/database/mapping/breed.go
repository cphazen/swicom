package mapping

import "example/swift-comply/models"

func BreedJson(b models.Breed) models.BreedJson {
	result := models.BreedJson{
		Id:   b.Id,
		Name: b.Name,
	}
	if b.BreedGroup != nil {
		breedGroup := BreedGroupJson(*b.BreedGroup)
		result.BreedGroup = &breedGroup
	}
	return result
}
