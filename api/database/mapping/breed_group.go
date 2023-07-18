package mapping

import "example/swift-comply/models"

func BreedGroupJson(bg models.BreedGroup) models.BreedGroupJson {
	return models.BreedGroupJson{
		Id:   bg.Id,
		Name: bg.Name,
	}
}
