import { useState, useEffect } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import {
  ArrowRightStartOnRectangleIcon,
  LinkIcon,
  PlusCircleIcon,
} from '@heroicons/react/24/solid'
import useStore from '../store'
import { useQueryExternalAPIs } from '../hooks/useQueryExternalApis'
import { useMutateExternalAPI } from '../hooks/useMutateExternalApi'
import { useMutateAuth } from '../hooks/useMutateAuth'
import { ExternalAPIItem } from './ExternalApiItem'
import '../ExternalAPIManager.css'

export const ExternalAPIManager = () => {
  const queryClient = useQueryClient()
  const editedExternalAPI = useStore((state) => state.editedExternalAPI)
  const updateExternalAPI = useStore((state) => state.updateEditedExternalAPI)
  const { data, isLoading } = useQueryExternalAPIs()
  const { createExternalAPIMutation, updateExternalAPIMutation } = useMutateExternalAPI()
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
        createNewAPI()
      }
      if (event.key === 'Escape') {
        setIsEditing(false)
      }
    }

    document.addEventListener('keydown', handleKeyDown)
    return () => document.removeEventListener('keydown', handleKeyDown)
  }, [updateExternalAPI])

  const createNewAPI = () => {
    updateExternalAPI({ id: 0, name: '', base_url: '', description: '' })
    setIsEditing(true)
  }

  const handleSubmit = () => {
    // 修正点1: すべてのフィールドが空の場合のみ早期リターン
    if (!editedExternalAPI.name && !editedExternalAPI.base_url && !editedExternalAPI.description) return
    
    if (editedExternalAPI.id === 0) {
      createExternalAPIMutation.mutate({
        name: editedExternalAPI.name,
        base_url: editedExternalAPI.base_url,
        description: editedExternalAPI.description,
      })
    } else {
      updateExternalAPIMutation.mutate(editedExternalAPI)
    }
    setIsEditing(false)
  }

  const logout = async () => {
    await logoutMutation.mutateAsync()
    queryClient.removeQueries(['externalAPIs'])
  }

  return (
    <div className="api-container">
      <div className="api-header">
        <LinkIcon className="header-icon" />
        <span className="header-title">External API Manager</span>
      </div>
      <ArrowRightStartOnRectangleIcon
        onClick={logout}
        className="logout-icon"
      />
      
      <ul className="api-list">
        {isLoading ? (
          <p>Loading...</p>
        ) : data && data.length > 0 ? (
          data.map((api) => (
            <ExternalAPIItem
              key={api.id}
              id={api.id}
              name={api.name}
              base_url={api.base_url}
              description={api.description}
              onEdit={() => {
                updateExternalAPI({
                  id: api.id,
                  name: api.name,
                  base_url: api.base_url,
                  description: api.description,
                })
                setIsEditing(true)
              }}
            />
          ))
        ) : (
          <div className="empty-state">
            <p>No External APIs found</p>
            <button 
              className="add-button"
              onClick={createNewAPI}
            >
              <PlusCircleIcon className="add-icon" />
              Add New API
            </button>
          </div>
        )}
      </ul>

      {isEditing && (
        <div className="sidebar">
          <h2>API詳細編集</h2>
          <input
            className="form-input"
            placeholder="API Name"
            type="text"
            onChange={(e) => 
              updateExternalAPI({ 
                ...editedExternalAPI, 
                name: e.target.value 
              })
            }
            value={editedExternalAPI.name || ''}
          />
          
          <input
            className="form-input"
            placeholder="Base URL"
            type="text"
            onChange={(e) => 
              updateExternalAPI({ 
                ...editedExternalAPI, 
                base_url: e.target.value 
              })
            }
            value={editedExternalAPI.base_url || ''}
          />
          
          <textarea
            className="form-textarea"
            placeholder="Description"
            onChange={(e) => 
              updateExternalAPI({ 
                ...editedExternalAPI, 
                description: e.target.value 
              })
            }
            value={editedExternalAPI.description || ''}
          />
          
          <div className="sidebar-actions">
            <button 
              className="save-button" 
              onClick={handleSubmit}
              // 修正点2: すべてのフィールドが空の場合のみボタンを無効化
              disabled={!editedExternalAPI.name && !editedExternalAPI.base_url && !editedExternalAPI.description}
            >
              保存
            </button>
            <button className="close-button" onClick={() => setIsEditing(false)}>閉じる</button>
          </div>
        </div>
      )}
    </div>
  )
}