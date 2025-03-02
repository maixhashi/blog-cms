import { FC } from 'react'
import { HatenaArticleList } from '../components/HatenaArticleList'
import '../styles/HatenaArticle.css'

export const HatenaPage: FC = () => {
  return (
    <div className="container">
      <HatenaArticleList />
    </div>
  )
}
