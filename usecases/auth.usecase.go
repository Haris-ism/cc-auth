package usecase

import (
	"cc-auth/controllers/models"
)

func (uc *usecase) Register(req models.Credentials) error{

	if err:=uc.postgre.Register(req); err!=nil{
		return err
	}

	return nil
}