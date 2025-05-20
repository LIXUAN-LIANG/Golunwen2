package App

import (
	"Golunwen2/App/model"
	"Golunwen2/App/router"
)

func InitStart() {
	model.InitMysql()
	defer func() {
		model.Close()
	}()
	router.InitGin()
}
