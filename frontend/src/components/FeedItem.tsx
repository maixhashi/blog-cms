import { FC, memo } from 'react'
import { PencilIcon, TrashIcon } from '@heroicons/react/24/solid'
import useStore from '../store'
import { Feed } from '../types'
import { useMutateFeed } from '../hooks/useMutateFeed'
import '../FeedItem.css'

interface FeedItemProps extends Omit<Feed, 'created_at' | 'updated_at'> {
  onEdit: () => void
}

const FeedItemMemo: FC<FeedItemProps> = ({
  id,
  title,
  url,
  site_url,
  description,
  last_fetched_at,
  onEdit, // ← 追加
}) => {
  const updateFeed = useStore((state) => state.updateEditedFeed)
  const { deleteFeedMutation } = useMutateFeed()
  return (
    <li className="feed-item">
      <span className="feed-title">{title}</span>
      <div className="feed-actions">
        <PencilIcon
          className="feed-icon edit"
          onClick={() => {
            updateFeed({
              id: id,
              title: title,
              url: url,
              site_url: site_url,
              description: description,
              last_fetched_at: last_fetched_at,
            })
            onEdit() // ← 追加
          }}
        />
        <TrashIcon
          className="feed-icon delete"
          onClick={() => {
            deleteFeedMutation.mutate(id)
          }}
        />
      </div>
    </li>
  )
}

export const FeedItem = memo(FeedItemMemo)
