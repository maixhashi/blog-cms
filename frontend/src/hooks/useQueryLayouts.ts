import { useQuery } from '@tanstack/react-query'
import { Layout } from '../types/'
import { useError } from '../hooks/useError'
import { fetchLayouts } from '../api/layouts'

export const useQueryLayouts = () => {
  const { switchErrorHandling } = useError()
  
  return useQuery<Layout[], Error>({
    queryKey: ['layouts'],
    queryFn: fetchLayouts,
    staleTime: Infinity,
    onError: (err: any) => {
      if (err.response?.data?.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response?.data || 'Something went wrong')
      }
    },
  })
}