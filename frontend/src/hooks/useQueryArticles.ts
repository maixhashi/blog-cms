import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { Article } from '../types/index'
import { useError } from '../hooks/useError'

export const useQueryArticles = () => {
  const { switchErrorHandling } = useError()
  const getArticles = async () => {
    const { data } = await axios.get<Article[]>(
      `${process.env.REACT_APP_API_URL}/articles`,
      { withCredentials: true }
    )
    return data
  }
  return useQuery<Article[], Error>({
    queryKey: ['articles'],
    queryFn: getArticles,
    staleTime: Infinity,
    onError: (err: any) => {
      if (err.response.data.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response.data)
      }
    },
  })
}
