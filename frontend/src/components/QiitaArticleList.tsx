import { FC } from 'react'
import { useQueryQiitaArticles } from '../hooks/useQueryQiitaArticles'
import { QiitaArticleItem } from './QiitaArticleItem'

export const QiitaArticleList: FC = () => {
  const { data: articles, isLoading, isError } = useQueryQiitaArticles()

  if (isLoading) return <div>読み込み中...</div>
  if (isError) return <div>エラーが発生しました</div>
  if (!articles || articles.length === 0) return <div>記事がありません</div>

  return (
    <div className="qiita-container">
      <h2 className="qiita-header">Qiita記事一覧</h2>
      <ul className="qiita-list">
        {articles.map((article) => (
          <QiitaArticleItem key={article.id} article={article} />
        ))}
      </ul>
    </div>
  )
}