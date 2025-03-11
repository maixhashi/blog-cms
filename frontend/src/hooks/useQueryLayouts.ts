import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { Layout } from '../types/'
import { useError } from '../hooks/useError'

export const useQueryLayouts = () => {
  const { switchErrorHandling } = useError()
  const getLayouts = async () => {
    const { data } = await axios.get<Layout[]>(
      `${process.env.REACT_APP_API_URL}/layouts`,
      { withCredentials: true }
    )
    return data
  }
  return useQuery<Layout[], Error>({
    queryKey: ['layouts'],
    queryFn: getLayouts,
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
