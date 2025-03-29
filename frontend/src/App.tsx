import { useEffect } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Auth } from './components/Auth'
import { Todo } from './components/Todo'
import { Feed } from './components/Feed'
import { ExternalAPIManager } from './components/ExternalAPIManager'
import { ArticleManager } from './components/ArticleManager'
import { LayoutManager } from './components/LayoutManager'
import { LayoutComponentManager } from './components/LayoutComponentManager'
import axios from 'axios'
import { CsrfToken } from './types'
import { QiitaPage } from './pages/QiitaPage'
import { ArticleEditorPage } from './pages/ArticleEditorPage'
import { HatenaPage } from './pages/HatenaPage'
import RouteMapPage from './pages/RouteMapPage'
import { FeedArticlesPage } from './pages/FeedArticlesPage'
import LayoutEditorPage from './pages/LayoutEditorPage'
import { LayoutEditor } from './components/LayoutEditor'
import { GoogleBookSearchPage } from './pages/GoogleBookSearchPage'

function App() {
  useEffect(() => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${process.env.REACT_APP_API_URL}/csrf-token`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
  }, [])
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Auth />} />
        <Route path="/task-manager" element={<Todo />} />
        <Route path="/feed-manager" element={<Feed />} />
        <Route path="/article-manager" element={<ArticleManager />} />
        <Route path="/external-api-manager" element={<ExternalAPIManager />} />
        <Route path="/layout-manager" element={<LayoutManager />} /> 
        <Route path="/layout-component-manager" element={<LayoutComponentManager />} />
        
        <Route path="/qiita-articles" element={<QiitaPage />} />
        <Route path="/hatena-articles" element={<HatenaPage />} />
        <Route path="/article-editor-page" element={<ArticleEditorPage />} />
        <Route path="/route-map" element={<RouteMapPage />} />
        <Route path="/feed-articles" element={<FeedArticlesPage />} />
        
        {/* 新しく追加した書籍検索ページのルート */}
        <Route path="/google-book-search" element={<GoogleBookSearchPage />} />
        
        {/* レイアウトエディターのルート */}
        <Route path="/layout-editor/:layoutId" element={<LayoutEditorPage />} />
      </Routes>
    </BrowserRouter>
  )
}
export default App