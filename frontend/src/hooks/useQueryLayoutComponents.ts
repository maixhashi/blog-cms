import { useQuery } from '@tanstack/react-query'
import { LayoutComponent } from '../types/'
import { useError } from '../hooks/useError'
import { fetchLayoutComponents } from '../api/layoutComponents'

export const useQueryLayoutComponents = () => {
  const { switchErrorHandling } = useError()
  
  return useQuery<LayoutComponent[], Error>({
    queryKey: ['layoutComponents'],
    queryFn: fetchLayoutComponents,
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