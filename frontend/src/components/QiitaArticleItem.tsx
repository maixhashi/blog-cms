import { FC } from 'react'
import { QiitaArticle } from '../types'

type Props = {
  article: QiitaArticle
}

export const QiitaArticleItem: FC<Props> = ({ article }) => {
  const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return date.toLocaleDateString('ja-JP')
  }

  return (
    <li className="qiita-item">
      <div className="qiita-content">
        <h3 className="qiita-title">
          <a href={article.url} target="_blank" rel="noopener noreferrer">
            {article.title}
          </a>
        </h3>
        <div className="qiita-meta">
          <div className="qiita-user">
            <img 
              src={article.user.profile_image_url} 
              alt={article.user.id} 
              className="qiita-user-image" 
            />
            <span>{article.user.id}</span>
          </div>
          <div className="qiita-info">
            <span className="qiita-date">{formatDate(article.created_at)}</span>
            <span className="qiita-likes">♥ {article.likes_count}</span>
          </div>
        </div>
        <div className="qiita-tags">
          {article.tags.map((tag) => (
            <span key={tag.name} className="qiita-tag">
              {tag.name}
            </span>
          ))}
        </div>
      </div>
    </li>
  )
}
