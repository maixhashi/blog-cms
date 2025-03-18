package main_entry_module

import (
	"go-react-app/controller"
	"gorm.io/gorm"
)

// MainEntryPackage はアプリケーションの主要コンポーネントを保持する構造体
type MainEntryPackage struct {
	UserController            controller.IUserController
	TaskController            controller.ITaskController 
	FeedController            controller.IFeedController
	ExternalAPIController     controller.IExternalAPIController
	ArticleController         controller.IArticleController
	LayoutController          controller.ILayoutController
	LayoutComponentController controller.ILayoutComponentController
	QiitaController           controller.IQiitaController
	HatenaController          controller.IHatenaController
	FeedArticleController     controller.IFeedArticleController
	
	// Swaggerハンドラーを追加（オプション）
	SwaggerEnabled            bool
}

// NewMainEntryPackage は新しいMainEntryPackageインスタンスを作成する
func NewMainEntryPackage(db *gorm.DB) *MainEntryPackage {
	entry := &MainEntryPackage{
		SwaggerEnabled: true, // デフォルトで有効
	}
	
	// 各モジュールの初期化
	entry.initUserModule(db)
	entry.initTaskModule(db)
	entry.initFeedModule(db)
	entry.initExternalAPIModule(db)
	entry.initArticleModule(db)
	entry.initLayoutModule(db)
	entry.initLayoutComponentModule(db)
	entry.initQiitaModule()
	entry.initHatenaModule()
	entry.initFeedArticleModule(db)

	return entry
}