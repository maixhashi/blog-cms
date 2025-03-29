import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { GoogleBookSearchResponse } from '../types/models/googleBook'
import { useError } from './useError'
import useStore from '../store'

export const useQueryGoogleBooks = (enabled = false) => {
  const { switchErrorHandling } = useError()
  const { searchQuery } = useStore((state) => state)
  
  const searchBooks = async () => {
    if (!searchQuery.trim()) return { items: [], totalItems: 0 }
    
    try {
      // POSTリクエストでクエリを送信
      const { data } = await axios.post<GoogleBookSearchResponse>(
        `${process.env.REACT_APP_API_URL}/google-books/search`,
        { query: searchQuery },
        { withCredentials: true }
      )
      return data
    } catch (error: any) {
      console.error('Google Books検索エラー:', error)
      if (error.response) {
        console.error('エラーレスポンス:', error.response.data)
      }
      throw error
    }
  }
  
  // @tanstack/react-queryのバージョンに合わせた型定義
  return useQuery({
    queryKey: ['googleBooks', searchQuery],
    queryFn: searchBooks,
    enabled: enabled && !!searchQuery.trim(),
    staleTime: 60000, // 1分間キャッシュ
    onError: (err: any) => {
      if (err.response?.data?.message) {
        switchErrorHandling(err.response.data.message)
      } else if (err.response?.data?.error) {
        switchErrorHandling(err.response.data.error)
      } else {
        switchErrorHandling(err.message || '書籍検索中にエラーが発生しました')
      }
    },
  })
}
