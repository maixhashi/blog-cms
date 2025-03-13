package main_entry_module

import (
	"go-react-app/controller"
	"go-react-app/repository"
	"go-react-app/usecase"
)

func (m *MainEntryPackage) initQiitaModule() {
	qiitaRepository := repository.NewQiitaRepository()
	qiitaUsecase := usecase.NewQiitaUsecase(qiitaRepository)
	m.QiitaController = controller.NewQiitaController(qiitaUsecase)
}
