import { FC, FormEvent, useState } from 'react'
import { useQueryLayoutComponents } from '../hooks/useQueryLayoutComponents'
import { useMutateLayoutComponent } from '../hooks/useMutateLayoutComponent'
import useStore from '../store'
import { LayoutComponent } from '../types'
import { 
  PencilIcon, 
  TrashIcon, 
  PlusCircleIcon, 
  ArrowPathIcon, 
  Squares2X2Icon 
} from '@heroicons/react/24/solid'
import '../LayoutComponent.css'

export const LayoutComponentManager: FC = () => {
  const [selectedComponentId, setSelectedComponentId] = useState<number | null>(null)
  const { data: components, isLoading, refetch } = useQueryLayoutComponents()
  const { createLayoutComponentMutation, updateLayoutComponentMutation, deleteLayoutComponentMutation } = useMutateLayoutComponent()
  const { editedLayoutComponent, updateEditedLayoutComponent, resetEditedLayoutComponent } = useStore((state) => state)

  const submitHandler = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (editedLayoutComponent.id === 0) {
      createLayoutComponentMutation.mutate({
        name: editedLayoutComponent.name,
        type: editedLayoutComponent.type,
        content: editedLayoutComponent.content
      })
    } else {
      updateLayoutComponentMutation.mutate(editedLayoutComponent)
    }
  }

  const resetForm = () => {
    resetEditedLayoutComponent()
    setSelectedComponentId(null)
  }

  const getTypeClassName = (type: string): string => {
    switch (type) {
      case 'header': return 'type-header'
      case 'footer': return 'type-footer'
      case 'sidebar': return 'type-sidebar'
      case 'main': return 'type-main'
      default: return 'type-custom'
    }
  }

  if (isLoading) return (
    <div className="loading-spinner">
      <div className="spinner"></div>
    </div>
  )

  return (
    <div className="layout-component-container">
      <div className="layout-component-header">
        <Squares2X2Icon className="header-icon" />
        <h1 className="header-title">レイアウトコンポーネント管理</h1>
      </div>

      <div className="flex justify-end mb-4">
        <button 
          onClick={() => refetch()}
          className="refresh-button"
        >
          <ArrowPathIcon className="refresh-icon" />
          <span>更新</span>
        </button>
      </div>
      
      <div className="layout-component-grid">
        {/* フォーム部分 */}
        <div className="component-card">
          <div className="component-card-header">
            <div className={`card-indicator ${editedLayoutComponent.id === 0 ? 'indicator-create' : 'indicator-edit'}`}></div>
            <h2 className="card-title">
              {editedLayoutComponent.id === 0 ? '新規コンポーネント作成' : 'コンポーネント編集'}
            </h2>
          </div>
          
          <form onSubmit={submitHandler} className="component-form">
            <div className="form-group">
              <label className="form-label">
                名前
              </label>
              <input
                type="text"
                className="form-input"
                value={editedLayoutComponent.name}
                onChange={(e) => 
                  updateEditedLayoutComponent({
                    ...editedLayoutComponent,
                    name: e.target.value
                  })
                }
                placeholder="コンポーネント名を入力"
                required
              />
            </div>
            
            <div className="form-group">
              <label className="form-label">
                タイプ
              </label>
              <select
                className="form-select"
                value={editedLayoutComponent.type}
                onChange={(e) => 
                  updateEditedLayoutComponent({
                    ...editedLayoutComponent,
                    type: e.target.value
                  })
                }
                required
              >
                <option value="">選択してください</option>
                <option value="header">ヘッダー</option>
                <option value="footer">フッター</option>
                <option value="sidebar">サイドバー</option>
                <option value="main">メインコンテンツ</option>
                <option value="custom">カスタム</option>
              </select>
            </div>
            
            <div className="form-group">
              <label className="form-label">
                コンテンツ
              </label>
              <textarea
                className="form-textarea"
                value={editedLayoutComponent.content}
                onChange={(e) => 
                  updateEditedLayoutComponent({
                    ...editedLayoutComponent,
                    content: e.target.value
                  })
                }
                placeholder="HTML、CSS、またはJSONコンテンツを入力"
              />
            </div>
            
            <div className="form-actions">
              {editedLayoutComponent.id !== 0 ? (
                <button
                  type="button"
                  onClick={resetForm}
                  className="cancel-button"
                >
                  キャンセル
                </button>
              ) : (
                <div></div>
              )}
              <button
                type="submit"
                className={`submit-button ${editedLayoutComponent.id === 0 ? 'create' : 'edit'}`}
              >
                {editedLayoutComponent.id === 0 ? (
                  <>
                    <PlusCircleIcon className="button-icon" />
                    <span>作成</span>
                  </>
                ) : (
                  <>
                    <PencilIcon className="button-icon" />
                    <span>更新</span>
                  </>
                )}
              </button>
            </div>
          </form>
        </div>
        
        {/* コンポーネント一覧 */}
        <div className="component-card">
          <h2 className="card-title mb-4">コンポーネント一覧</h2>
          
          {components && components.length > 0 ? (
            <div className="component-list">
              {components.map((component) => (
                <div 
                  key={component.id}
                  className={`component-item ${selectedComponentId === component.id ? 'selected' : ''}`}
                >
                  <div className="component-item-header">
                    <div className="component-info">
                      <h3 className="component-name">{component.name}</h3>
                      <span className={`component-type ${getTypeClassName(component.type)}`}>
                        {component.type}
                      </span>
                    </div>
                    <div className="component-actions">
                      <button
                        onClick={() => {
                          updateEditedLayoutComponent({
                            id: component.id,
                            name: component.name,
                            type: component.type,
                            content: component.content
                          })
                          setSelectedComponentId(component.id)
                        }}
                        className="action-button edit-button"
                        title="編集"
                      >
                        <PencilIcon className="action-icon" />
                      </button>
                      <button
                        onClick={() => {
                          if (window.confirm(`"${component.name}"を削除してもよろしいですか？`)) {
                            deleteLayoutComponentMutation.mutate(component.id)
                            if (selectedComponentId === component.id) {
                              resetForm()
                            }
                          }
                        }}
                        className="action-button delete-button"
                        title="削除"
                      >
                        <TrashIcon className="action-icon" />
                      </button>
                    </div>
                  </div>
                  
                  {component.content && (
                    <div className="content-preview">
                      <p className="preview-label">コンテンツプレビュー:</p>
                      <div className="preview-content">
                        {component.content}
                      </div>
                    </div>
                  )}
                </div>
              ))}
            </div>
          ) : (
            <div className="empty-state">
              <svg 
                className="empty-icon" 
                fill="none" 
                viewBox="0 0 24 24" 
                stroke="currentColor" 
                aria-hidden="true"
              >
                <path 
                  strokeLinecap="round" 
                  strokeLinejoin="round" 
                  strokeWidth={1} 
                  d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" 
                />
              </svg>
              <p className="empty-text">コンポーネントがまだありません</p>
              <p className="empty-subtext">新しいコンポーネントを作成してください</p>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}
