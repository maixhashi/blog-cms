import { FC, memo } from 'react'
import { PencilIcon, TrashIcon } from '@heroicons/react/24/solid'
import { useMutateArticle } from '../hooks/useMutateArticle'
import useStore from '../store'
import { Article } from '../types/models/article'
import '../ArticleItem.css'

// Articleから必要なプロパティのみを取り出す
type ArticleItemProps = Required<Pick<Article, 'id' | 'title' | 'content' | 'published' | 'tags'>> & {
  onEdit: () => void
}

const ArticleItemMemo: FC<ArticleItemProps> = ({
  id,
  title,
  content,
  published,
  tags,
  onEdit,
}) => {
  const { deleteArticleMutation } = useMutateArticle()
  const updateArticle = useStore((state) => state.updateEditedArticle)
  
  return (
    <li className="article-item">
      <div className="article-content">
        <h3 className="article-title">{title}</h3>
        <div className="article-meta">
          <span className={`article-status ${published ? 'published' : 'draft'}`}>
            {published ? '公開中' : '下書き'}
          </span>
          {tags && <span className="article-tags">{tags}</span>}
        </div>
        <p className="article-excerpt">{content.substring(0, 150)}...</p>
      </div>
      <div className="article-actions">
        <PencilIcon
          className="article-icon edit"
          onClick={() => {
            updateArticle({
              id,
              title,
              content,
              published,
              tags,
            })
            onEdit()
          }}
        />
        <TrashIcon
          className="article-icon delete"
          onClick={() => {
            deleteArticleMutation.mutate(id)
          }}
        />
      </div>
    </li>
  )
}

export const ArticleItem = memo(ArticleItemMemo)
