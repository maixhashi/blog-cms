import axios from 'axios'
import { useQueryClient, useMutation } from '@tanstack/react-query'
import { ExternalAPI } from '../types'
import useStore from '../store'
import { useError } from '../hooks/useError'

export const useMutateExternalAPI = () => {
  const queryClient = useQueryClient()
  const { switchErrorHandling } = useError()
  const resetEditedExternalAPI = useStore((state) => state.resetEditedExternalAPI)

  const createExternalAPIMutation = useMutation(
    (api: Omit<ExternalAPI, 'id' | 'created_at' | 'updated_at'>) =>
      axios.post<ExternalAPI>(`${process.env.REACT_APP_API_URL}/external-apis`, api),
    {
      onSuccess: (res) => {
        const previousAPIs = queryClient.getQueryData<ExternalAPI[]>(['externalAPIs'])
        if (previousAPIs) {
          queryClient.setQueryData(['externalAPIs'], [...previousAPIs, res.data])
        }
        resetEditedExternalAPI()
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )
  
  const updateExternalAPIMutation = useMutation(
    (api: Omit<ExternalAPI, 'created_at' | 'updated_at'>) =>
      axios.put<ExternalAPI>(`${process.env.REACT_APP_API_URL}/external-apis/${api.id}`, {
        name: api.name,
        base_url: api.base_url,
        description: api.description,
      }),
    {
      onSuccess: (res, variables) => {
        const previousAPIs = queryClient.getQueryData<ExternalAPI[]>(['externalAPIs'])
        if (previousAPIs) {
          queryClient.setQueryData<ExternalAPI[]>(
            ['externalAPIs'],
            previousAPIs.map((api) =>
              api.id === variables.id ? res.data : api
            )
          )
        }
        resetEditedExternalAPI()
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )
  
  const deleteExternalAPIMutation = useMutation(
    (id: number) =>
      axios.delete(`${process.env.REACT_APP_API_URL}/external-apis/${id}`),
    {
      onSuccess: (_, variables) => {
        const previousAPIs = queryClient.getQueryData<ExternalAPI[]>(['externalAPIs'])
        if (previousAPIs) {
          queryClient.setQueryData<ExternalAPI[]>(
            ['externalAPIs'],
            previousAPIs.filter((api) => api.id !== variables)
          )
        }
        resetEditedExternalAPI()
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )
  
  return {
    createExternalAPIMutation,
    updateExternalAPIMutation,
    deleteExternalAPIMutation,
  }
}