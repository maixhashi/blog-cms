import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { BookResponse } from '../types/models/book'
import { useError } from './useError'

export const useQueryBooks = () => {
  const { switchErrorHandling } = useError()
  
  const getBooks = async () => {
    const { data } = await axios.get<BookResponse[]>(
      `${process.env.REACT_APP_API_URL}/books`,
      { withCredentials: true }
    )
    return data
  }
  
  return useQuery<BookResponse[], Error>({
    queryKey: ['books'],
    queryFn: getBooks,
    staleTime: Infinity,
    onError: (err: any) => {
      if (err.response?.data?.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response?.data || 'エラーが発生しました')
      }
    },
  })
}
