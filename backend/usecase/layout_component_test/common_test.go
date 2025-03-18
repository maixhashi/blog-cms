package layout_component_test

import (
	"fmt"
	"github.com/google/uuid"
)

// 一意のコンポーネント名を生成
func generateUniqueName() string {
	// UUIDを生成（ほぼ確実に一意）
	uuid := uuid.New().String()
	// UUIDの最初の8文字だけを使用
	shortUUID := uuid[:8]
	
	return fmt.Sprintf("Test Component %s", shortUUID)
}
