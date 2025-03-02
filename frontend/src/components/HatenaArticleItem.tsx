import { FC } from 'react'
import { HatenaArticle } from '../types'

interface Props {
  article: HatenaArticle
}

export const HatenaArticleItem: FC<Props> = ({ article }) => {
  // 日付をフォーマット
  const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return new Intl.DateTimeFormat('ja-JP', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
    }).format(date)
  }

  return (
    <li className="hatena-item">
      <h3 className="hatena-title">
        <a href={article.url} target="_blank" rel="noopener noreferrer">
          {article.title}
        </a>
      </h3>
      <div className="hatena-meta">
        <div className="hatena-info">
          <span className="hatena-date">
            {formatDate(article.published_at)}
          </span>
          <span className="hatena-author">作者: {article.author}</span>
        </div>
      </div>
      <p>{article.summary}</p>
      <div className="hatena-tags">
        {article.categories.map((category, index) => (
          <span key={index} className="hatena-tag">
            {category}
          </span>
        ))}
      </div>
    </li>
  )
}
