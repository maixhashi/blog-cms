import { FormEvent } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import {
  ArrowRightStartOnRectangleIcon,
  ShieldCheckIcon,
} from '@heroicons/react/24/solid'
import useStore from '../store'
import { useQueryFeeds } from '../hooks/useQueryFeeds'
import { useMutateFeed } from '../hooks/useMutateFeed'
import { useMutateAuth } from '../hooks/useMutateAuth'
import { FeedItem } from './FeedItem'
import '../Todo.css' // ← CSS をインポート

export const Feed = () => {
  const queryClient = useQueryClient()
  const { editedFeed } = useStore()
  const updateFeed = useStore((state) => state.updateEditedFeed)
  const { data, isLoading } = useQueryFeeds()
  const { createFeedMutation, updateFeedMutation } = useMutateFeed()
  const { logoutMutation } = useMutateAuth()

  const submitFeedHandler = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (editedFeed.id === 0)
      createFeedMutation.mutate({
        title: editedFeed.title,
        url: editedFeed.url || "",
        site_url: editedFeed.site_url || "",
        description: editedFeed.description || "",
        last_fetched_at: editedFeed.last_fetched_at || new Date().toISOString(),
      })
    else {
      updateFeedMutation.mutate(editedFeed)
    }
  }

  const logout = async () => {
    await logoutMutation.mutateAsync()
    queryClient.removeQueries(['feeds'])
  }

  return (
    <div className="todo-container">
      <div className="todo-header">
        <ShieldCheckIcon className="header-icon" />
        <span className="header-title">Feed Manager</span>
      </div>
      <ArrowRightStartOnRectangleIcon
        onClick={logout}
        className="logout-icon"
      />
      <form onSubmit={submitFeedHandler} className="todo-form">
        <input
          className="form-input"
          placeholder="title ?"
          type="text"
          onChange={(e) => updateFeed({ ...editedFeed, title: e.target.value })}
          value={editedFeed.title || ''}
        />
        <button className="submit-button" disabled={!editedFeed.title}>
          {editedFeed.id === 0 ? 'Create' : 'Update'}
        </button>
      </form>
      {isLoading ? (
        <p className="loading-text">Loading...</p>
      ) : (
        <ul className="feed-list">
          {data?.map((feed) => (
            <li key={feed.id} className="feed-item">
              <FeedItem
                id={feed.id}
                title={feed.title}
                url={feed.url}
                site_url={feed.site_url}
                description={feed.description}
                last_fetched_at={feed.last_fetched_at}
              />
            </li>
          ))}
        </ul>
      )}
    </div>
  )
}
