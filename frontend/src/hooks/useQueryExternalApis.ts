import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { ExternalAPI } from '../types'
import { useError } from '../hooks/useError'

export const useQueryExternalAPIs = () => {
  const { switchErrorHandling } = useError()
  const getExternalAPIs = async () => {
    const { data } = await axios.get<ExternalAPI[]>(
      `${process.env.REACT_APP_API_URL}/external-apis`,
      { withCredentials: true }
    )
    return data
  }
  return useQuery<ExternalAPI[], Error>({
    queryKey: ['externalAPIs'],
    queryFn: getExternalAPIs,
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
