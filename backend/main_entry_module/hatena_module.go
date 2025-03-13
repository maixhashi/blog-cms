package main_entry_module

import (
	"go-react-app/controller"
	"go-react-app/repository"
	"go-react-app/usecase"
)

func (m *MainEntryPackage) initHatenaModule() {
	hatenaRepository := repository.NewHatenaRepository("https://tech.smarthr.jp/feed?exclude_body=1")
	hatenaUsecase := usecase.NewHatenaUsecase(hatenaRepository)
	m.HatenaController = controller.NewHatenaController(hatenaUsecase)
}
