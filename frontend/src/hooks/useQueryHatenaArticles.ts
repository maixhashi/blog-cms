import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { HatenaArticle } from '../types'
import { useError } from './useError'

export const useQueryHatenaArticles = () => {
  const { switchErrorHandling } = useError()
  
  const getHatenaArticles = async () => {
    const { data } = await axios.get<HatenaArticle[]>(
      `${process.env.REACT_APP_API_URL}/hatena`,
      { withCredentials: true }
    )
    return data
  }

  return useQuery<HatenaArticle[], Error>({
    queryKey: ['hatenaArticles'],
    queryFn: getHatenaArticles,
    staleTime: 0,
    refetchOnWindowFocus: true,
    onError: (err: any) => {
      if (err.response?.data?.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling('Something went wrong')
      }
    },
  })
}
