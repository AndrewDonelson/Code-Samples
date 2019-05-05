package orm_go_pg

import (
	goPg "github.com/AndrewDonelson/go-pg-orm"
	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/model"
)

type GoPgModel struct {
	db *goPg.Model
}

func New(db *goPg.Model) *GoPgModel {
	return &GoPgModel{db}
}

func (gpg *GoPgModel) Create(model *model.Contact) error{
	return gpg.db.SaveModel(model)
}

func (gpg *GoPgModel) View(model *model.Contact) error{
	return gpg.db.GetModel(model)
}

func (gpg *GoPgModel) List(model []*model.Contact) error{
	return gpg.db.GetAllModels(model)
}

func (gpg *GoPgModel) GetWithCondition(model *model.Contact, condition interface{}, args ...interface{}) error{
	return gpg.db.GetAllWithCondition(model, condition, args)
}