import { FC } from 'react'
import { useQueryHatenaArticles } from '../hooks/useQueryHatenaArticles'
import { HatenaArticleItem } from './HatenaArticleItem'

export const HatenaArticleList: FC = () => {
  const { data: articles, isLoading, isError } = useQueryHatenaArticles()

  if (isLoading) return <div>読み込み中...</div>
  if (isError) return <div>エラーが発生しました</div>
  if (!articles || articles.length === 0) return <div>記事がありません</div>

  return (
    <div className="hatena-container">
      <h2 className="hatena-header">はてなブログ記事一覧</h2>
      <ul className="hatena-list">
        {articles.map((article) => (
          <HatenaArticleItem key={article.id} article={article} />
        ))}
      </ul>
    </div>
  )
}
