import { FormEvent, useState, useEffect } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import {
  ArrowRightStartOnRectangleIcon,
  Squares2X2Icon,
  PlusCircleIcon,
} from '@heroicons/react/24/solid'
import useStore from '../store'
import { useQueryLayouts } from '../hooks/useQueryLayouts'
import { useMutateLayout } from '../hooks/useMutateLayout'
import { useMutateAuth } from '../hooks/useMutateAuth'
import { LayoutItem } from './LayoutItem'
import '../Layout.css'

export const LayoutManager = () => {
  const queryClient = useQueryClient()
  const { editedLayout } = useStore()
  const updateLayout = useStore((state) => state.updateEditedLayout)
  const { data, isLoading } = useQueryLayouts()
  const { createLayoutMutation, updateLayoutMutation } = useMutateLayout()
  const { logoutMutation } = useMutateAuth()
  
  const [isEditing, setIsEditing] = useState(false)

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.metaKey && event.key === 'e') {
        event.preventDefault()
        setIsEditing(true)
      }
      if (event.metaKey && event.key === 's') {
        event.preventDefault()
        handleSubmit()
      }
      if (event.metaKey && event.key === 'n') {
        event.preventDefault()
        createNewLayout()
      }
      if (event.key === 'Escape') {
        setIsEditing(false)
      }
    }

    document.addEventListener('keydown', handleKeyDown)
    return () => document.removeEventListener('keydown', handleKeyDown)
  }, [updateLayout])

  const createNewLayout = () => {
    updateLayout({ id: 0, title: '' })
    setIsEditing(true)
  }

  const handleSubmit = () => {
    if (!editedLayout.title) return
    
    if (editedLayout.id === 0) {
      createLayoutMutation.mutate({
        title: editedLayout.title,
      })
    } else {
      updateLayoutMutation.mutate(editedLayout)
    }
    setIsEditing(false)
  }

  const submitLayoutHandler = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    handleSubmit()
  }

  const logout = async () => {
    await logoutMutation.mutateAsync()
    queryClient.removeQueries(['layouts'])
  }

  return (
    <div>
      <h1>レイアウト管理</h1>
      <div className="layout-container">
        <div className="layout-header">
          <Squares2X2Icon className="header-icon" />
          <span className="header-title">Layout Manager</span>
        </div>
        <ArrowRightStartOnRectangleIcon
          onClick={logout}
          className="logout-icon"
        />
        
        <div className="layout-actions-top">
          <button 
            className="add-layout-button"
            onClick={createNewLayout}
          >
            <PlusCircleIcon className="add-icon" />
            新規レイアウト作成
          </button>
        </div>

        {isLoading ? (
          <p className="loading-text">Loading...</p>
        ) : (
          <ul className="layout-list">
            {data?.map((layout) => (
              <LayoutItem 
                key={layout.id}
                id={layout.id} 
                title={layout.title}
                onEdit={() => {
                  updateLayout({
                    id: layout.id,
                    title: layout.title,
                  })
                  setIsEditing(true)
                }}
              />
            ))}
          </ul>
        )}

        {isEditing && (
          <div className="sidebar">
            <h2>{editedLayout.id === 0 ? '新規レイアウト作成' : 'レイアウト編集'}</h2>
            <form onSubmit={submitLayoutHandler}>
              <input
                className="form-input"
                placeholder="レイアウトタイトル"
                type="text"
                onChange={(e) => updateLayout({ ...editedLayout, title: e.target.value })}
                value={editedLayout.title || ''}
              />
              
              <div className="sidebar-actions">
                <button 
                  className="save-button" 
                  type="submit"
                  disabled={!editedLayout.title}
                >
                  保存
                </button>
                <button 
                  className="close-button" 
                  type="button" 
                  onClick={() => setIsEditing(false)}
                >
                  閉じる
                </button>
              </div>
            </form>
          </div>
        )}
      </div>
    </div>
  )
}