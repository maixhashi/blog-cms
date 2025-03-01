import { FC } from 'react'
import { QiitaArticleList } from '../components/QiitaArticleList'
import '../styles/QiitaArticle.css'

export const QiitaPage: FC = () => {
  return (
    <div className="container">
      <QiitaArticleList />
    </div>
  )
}
