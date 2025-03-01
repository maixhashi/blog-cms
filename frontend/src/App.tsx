import { useEffect } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Auth } from './components/Auth'
import { Todo } from './components/Todo'
import { Feed } from './components/Feed'
import { ExternalAPIManager } from './components/ExternalAPIManager'
import axios from 'axios'
import { CsrfToken } from './types'
import { QiitaPage } from './pages/QiitaPage'

function App() {
  useEffect(() => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${process.env.REACT_APP_API_URL}/csrf`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
  }, [])
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Auth />} />
        <Route path="/todo" element={<Todo />} />
        <Route path="/feed" element={<Feed />} />
        <Route path="/external-api-manager" element={<ExternalAPIManager />} />
        <Route path="/qiita-articles" element={<QiitaPage />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App