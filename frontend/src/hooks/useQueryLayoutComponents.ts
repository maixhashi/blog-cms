import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { LayoutComponent } from '../types'
import { useError } from '../hooks/useError'

export const useQueryLayoutComponents = (layoutId: number) => {
  const { switchErrorHandling } = useError()
  const getLayoutComponents = async () => {
    const { data } = await axios.get<LayoutComponent[]>(
      `${process.env.REACT_APP_API_URL}/layouts/${layoutId}/components`,
      { withCredentials: true }
    )
    return data
  }
  return useQuery<LayoutComponent[], Error>({
    queryKey: ['layoutComponents', layoutId],
    queryFn: getLayoutComponents,
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
