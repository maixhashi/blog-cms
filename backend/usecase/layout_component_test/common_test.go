package layout_component_test

import (
	"go-react-app/model"
	"testing"
	"time"
)

// レスポンスの検証ヘルパー関数
func validateLayoutComponentResponse(t *testing.T, response model.LayoutComponentResponse, expected model.LayoutComponent) bool {
	if response.ID != expected.ID {
		t.Errorf("ID が一致しません。取得: %d, 期待: %d", response.ID, expected.ID)
		return false
	}
	
	if response.Name != expected.Name {
		t.Errorf("Name が一致しません。取得: %s, 期待: %s", response.Name, expected.Name)
		return false
	}
	
	if response.Type != expected.Type {
		t.Errorf("Type が一致しません。取得: %s, 期待: %s", response.Type, expected.Type)
		return false
	}
	
	if response.Content != expected.Content {
		t.Errorf("Content が一致しません。取得: %s, 期待: %s", response.Content, expected.Content)
		return false
	}
	
	// CreatedAt と UpdatedAt は厳密な一致ではなく、存在確認のみ
	if response.CreatedAt.IsZero() {
		t.Error("CreatedAt が設定されていません")
		return false
	}
	
	if response.UpdatedAt.IsZero() {
		t.Error("UpdatedAt が設定されていません")
		return false
	}
	
	return true
}

// 時間が近いかどうかを確認するヘルパー関数
func isTimeNearby(t1, t2 time.Time, allowedDiff time.Duration) bool {
	diff := t1.Sub(t2)
	if diff < 0 {
		diff = -diff
	}
	return diff <= allowedDiff
}

// データベースからコンポーネントを取得して検証するヘルパー関数
func verifyComponentInDB(t *testing.T, componentID uint, expectedName, expectedType, expectedContent string) bool {
	var component model.LayoutComponent
	if err := db.First(&component, componentID).Error; err != nil {
		t.Errorf("コンポーネントがデータベースに存在しません: %v", err)
		return false
	}
	
	if component.Name != expectedName {
		t.Errorf("データベース内の Name が一致しません。取得: %s, 期待: %s", component.Name, expectedName)
		return false
	}
	
	if component.Type != expectedType {
		t.Errorf("データベース内の Type が一致しません。取得: %s, 期待: %s", component.Type, expectedType)
		return false
	}
	
	if component.Content != expectedContent {
		t.Errorf("データベース内の Content が一致しません。取得: %s, 期待: %s", component.Content, expectedContent)
		return false
	}
	
	return true
}