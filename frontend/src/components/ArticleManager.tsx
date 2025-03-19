import { useState, useEffect } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import {
  ArrowRightStartOnRectangleIcon,
  DocumentTextIcon,
  PlusCircleIcon,
} from '@heroicons/react/24/solid'
import useStore from '../store'
import { useQueryArticles } from '../hooks/useQueryArticles'
import { useMutateArticle } from '../hooks/useMutateArticle'
import { useMutateAuth } from '../hooks/useMutateAuth'
import { ArticleItem } from './ArticleItem'
import '../ArticleManager.css'

export const ArticleManager = () => {
  const queryClient = useQueryClient()
  const editedArticle = useStore((state) => state.editedArticle)
  const updateArticle = useStore((state) => state.updateEditedArticle)
  const { data, isLoading } = useQueryArticles()
  const { createArticleMutation, updateArticleMutation } = useMutateArticle()
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
        createNewArticle()
      }
      if (event.key === 'Escape') {
        setIsEditing(false)
      }
    }

    document.addEventListener('keydown', handleKeyDown)
    return () => document.removeEventListener('keydown', handleKeyDown)
  }, [editedArticle])

  const createNewArticle = () => {
    updateArticle({ id: 0, title: '', content: '', published: false, tags: '' })
    setIsEditing(true)
  }

  const handleSubmit = () => {
    if (!editedArticle.title) return
    
    if (editedArticle.id === 0) {
      createArticleMutation.mutate({
        title: editedArticle.title,
        content: editedArticle.content,
        published: editedArticle.published,
        tags: editedArticle.tags,
      })
    } else {
      updateArticleMutation.mutate({
        id: editedArticle.id,
        title: editedArticle.title,
        content: editedArticle.content,
        published: editedArticle.published,
        tags: editedArticle.tags,
      })
    }
    setIsEditing(false)
  }

  const logout = async () => {
    await logoutMutation.mutateAsync()
    queryClient.removeQueries(['articles'])
  }

  return (
    <div className="article-container">
      <div className="article-header">
        <DocumentTextIcon className="header-icon" />
        <span className="header-title">記事管理</span>
      </div>
      <ArrowRightStartOnRectangleIcon
        onClick={logout}
        className="logout-icon"
      />
      
      <div className="article-actions-top">
        <button 
          className="add-article-button"
          onClick={createNewArticle}
        >
          <PlusCircleIcon className="add-icon" />
          新規記事作成
        </button>
      </div>
      
      <ul className="article-list">
        {isLoading ? (
          <p>読み込み中...</p>
        ) : data && data.length > 0 ? (
          data.map((article) => (
            <ArticleItem
              key={article.id}
              id={article.id!}
              title={article.title!}
              content={article.content!}
              published={article.published!}
              tags={article.tags!}
              onEdit={() => {
                updateArticle({
                  id: article.id!,
                  title: article.title!,
                  content: article.content!,
                  published: article.published!,
                  tags: article.tags!,
                })
                setIsEditing(true)
              }}
            />
          ))
        ) : (
          <div className="empty-state">
            <p>記事が見つかりません</p>
          </div>
        )}
      </ul>

      {isEditing && (
        <div className="sidebar">
          <h2>{editedArticle.id === 0 ? '新規記事作成' : '記事編集'}</h2>
          <input
            className="form-input"
            placeholder="タイトル"
            type="text"
            onChange={(e) => 
              updateArticle({ 
                ...editedArticle, 
                title: e.target.value 
              })
            }
            value={editedArticle.title || ''}
          />
          
          <textarea
            className="form-textarea"
            placeholder="本文"
            rows={15}
            onChange={(e) => 
              updateArticle({ 
                ...editedArticle, 
                content: e.target.value 
              })
            }
            value={editedArticle.content || ''}
          />
          
          <input
            className="form-input"
            placeholder="タグ（カンマ区切り）"
            type="text"
            onChange={(e) => 
              updateArticle({ 
                ...editedArticle, 
                tags: e.target.value 
              })
            }
            value={editedArticle.tags || ''}
          />
          
          <div className="publish-option">
            <label>
              <input
                type="checkbox"
                checked={editedArticle.published || false}
                onChange={(e) => 
                  updateArticle({ 
                    ...editedArticle, 
                    published: e.target.checked 
                  })
                }
              />
              公開する
            </label>
          </div>
          
          <div className="sidebar-actions">
            <button 
              className="save-button" 
              onClick={handleSubmit}
              disabled={!editedArticle.title}
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