package usecase

import "cc-auth/models"

func (uc *usecase) InsertDB(req models.ItemList) error {
	err := uc.postgre.Insert(req)
	if err != nil {
		return err
	}

	return nil
}
func (uc *usecase) QueryDB() ([]models.ItemList, error) {
	result, err := uc.postgre.Query()
	if err != nil {
		return result, err
	}

	return result, nil
}
