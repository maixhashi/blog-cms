import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { Feed } from '../types'
import { useError } from '../hooks/useError'

export const useQueryFeeds = () => {
  const { switchErrorHandling } = useError()
  const getFeeds = async () => {
    const { data } = await axios.get<Feed[]>(
      `${process.env.REACT_APP_API_URL}/feeds`,
      { withCredentials: true }
    )
    return data
  }
  return useQuery<Feed[], Error>({
    queryKey: ['feeds'],
    queryFn: getFeeds,
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