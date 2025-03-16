package layout_test

import (
	"fmt"
	"github.com/google/uuid"
)

// 一意のレイアウトタイトルを生成
func generateUniqueTitle() string {
	// UUIDを生成（ほぼ確実に一意）
	uuid := uuid.New().String()
	// UUIDの最初の8文字だけを使用
	shortUUID := uuid[:8]
	
	return fmt.Sprintf("Test Layout %s", shortUUID)
}
