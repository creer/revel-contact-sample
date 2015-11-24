package controllers

import (
	"github.com/revel/revel"
	"myform/app/models"
)

type App struct {
	GorpController
}

/**
	一覧画面
 */
func (c App) Index() revel.Result {

	var contacts []models.Contact
	_, err := Dbm.Select(&contacts, "SELECT * FROM contact ORDER BY Id")

	if(err != nil){
		panic(err)
	}

	return c.Render(contacts)
}
