import { useState, useEffect } from 'react';
import { useQueryClient } from '@tanstack/react-query';
import {
  ArrowRightStartOnRectangleIcon,
  ShieldCheckIcon,
} from '@heroicons/react/24/solid';
import useStore from '../store';
import { useQueryFeeds } from '../hooks/useQueryFeeds';
import { useMutateFeed } from '../hooks/useMutateFeed';
import { useMutateAuth } from '../hooks/useMutateAuth';
import { FeedItem } from './FeedItem';
import '../Feed.css';

export const Feed = () => {
  const queryClient = useQueryClient();
  const updateFeed = useStore((state) => state.updateEditedFeed);
  const setSelectedFeedId = useStore((state) => state.setSelectedFeedId);
  const selectedFeedId = useStore((state) => state.selectedFeedId);
  const editedFeed = useStore((state) => state.editedFeed);
  const { data, isLoading } = useQueryFeeds();
  const { createFeedMutation, updateFeedMutation } = useMutateFeed();
  const { logoutMutation } = useMutateAuth();

  const [isEditing, setIsEditing] = useState(false);
  const [isSelecting, setIsSelecting] = useState(false);

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.metaKey && event.key === 'e') {
        event.preventDefault();
        const selectedFeed = data?.find(feed => feed.id === selectedFeedId);
        if (selectedFeed) {
          updateFeed({ ...selectedFeed });
          setIsEditing(true);
        }
      }
      if (event.metaKey && event.key === 's') {
        event.preventDefault();
        handleSubmit();
      }
      if (event.metaKey && event.key === 'n') {
        event.preventDefault();
        updateFeed({ id: 0, title: '', url: '', site_url: '', description: '', last_fetched_at: new Date() });
        setIsEditing(true);
      }
      if (event.metaKey && event.key === 'l') {
        event.preventDefault();
        setIsSelecting(true);
      }
      if (event.key === 'Escape') {
        setIsEditing(false);
        setIsSelecting(false);
      }
    };

    document.addEventListener('keydown', handleKeyDown);
    return () => document.removeEventListener('keydown', handleKeyDown);
  }, [data, selectedFeedId, updateFeed]);

  const handleSubmit = () => {
    if (!editedFeed.title) return;
    if (editedFeed.id === 0) {
      createFeedMutation.mutate({
        title: editedFeed.title,
        url: editedFeed.url || '',
        site_url: editedFeed.site_url || '',
        description: editedFeed.description || '',
        last_fetched_at: new Date(),
      });
    } else {
      updateFeedMutation.mutate(editedFeed);
    }
    setIsEditing(false);
  };

  const logout = async () => {
    await logoutMutation.mutateAsync();
    queryClient.removeQueries(['feeds']);
  };

  return (
    <div className="feed-container">
      <div className="feed-header">
        <ShieldCheckIcon className="header-icon" />
        <span className="header-title">Feed Manager</span>
      </div>
      <ArrowRightStartOnRectangleIcon onClick={logout} className="logout-icon" />
      
      <ul className="feed-list">
        {data?.map((feed, index) => (
          <li
            key={feed.id}
            className={`feed-item ${isSelecting ? 'selecting' : ''}`}
            tabIndex={isSelecting ? 0 : -1}
            onKeyDown={(e) => {
              if (isSelecting && e.key === 'Tab') {
                const nextIndex = (index + (e.shiftKey ? -1 : 1) + data.length) % data.length;
                setSelectedFeedId(data[nextIndex].id);
              }
            }}
          >
            <FeedItem {...feed} onEdit={() => {
              updateFeed({ ...feed });
              setSelectedFeedId(feed.id);
              setIsEditing(true);
            }} />
          </li>
        ))}
      </ul>

      {isEditing && (
        <div className="sidebar">
          <h2>詳細編集</h2>
          <input
            className="form-input"
            placeholder="タイトル"
            type="text"
            onChange={(e) => updateFeed({ ...editedFeed, title: e.target.value })}
            value={editedFeed.title || ''}
          />
          <input
            className="form-input"
            placeholder="URL"
            type="text"
            onChange={(e) => updateFeed({ ...editedFeed, url: e.target.value })}
            value={editedFeed.url || ''}
          />
          <input
            className="form-input"
            placeholder="取得元のサイトURL"
            type="text"
            onChange={(e) => updateFeed({ ...editedFeed, site_url: e.target.value })}
            value={editedFeed.site_url || ''}
          />
          <textarea
            className="form-input"
            placeholder="説明"
            onChange={(e) => updateFeed({ ...editedFeed, description: e.target.value })}
            value={editedFeed.description || ''}
          />
          <div className="sidebar-actions">
            <button className="save-button" onClick={handleSubmit}>保存</button>
            <button className="close-button" onClick={() => setIsEditing(false)}>閉じる</button>
          </div>
        </div>
      )}
    </div>
  );
};
