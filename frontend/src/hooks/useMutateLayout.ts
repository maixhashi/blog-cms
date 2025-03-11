import axios from 'axios'
import { useQueryClient, useMutation } from '@tanstack/react-query'
import { Layout } from '../types'
import useStore from '../store'
import { useError } from '../hooks/useError'

export const useMutateLayout = () => {
  const queryClient = useQueryClient()
  const { switchErrorHandling } = useError()
  const resetEditedLayout = useStore((state) => state.resetEditedLayout)

  const createLayoutMutation = useMutation(
    (layout: Omit<Layout, 'id' | 'created_at' | 'updated_at'>) =>
      axios.post<Layout>(`${process.env.REACT_APP_API_URL}/layouts`, layout),
    {
      onSuccess: (res) => {
        const previousLayouts = queryClient.getQueryData<Layout[]>(['layouts'])
        if (previousLayouts) {
          queryClient.setQueryData(['layouts'], [...previousLayouts, res.data])
        }
        resetEditedLayout()
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
  
  const updateLayoutMutation = useMutation(
    (layout: Omit<Layout, 'created_at' | 'updated_at'>) =>
      axios.put<Layout>(`${process.env.REACT_APP_API_URL}/layouts/${layout.id}`, {
        title: layout.title,
      }),
    {
      onSuccess: (res, variables) => {
        const previousLayouts = queryClient.getQueryData<Layout[]>(['layouts'])
        if (previousLayouts) {
          queryClient.setQueryData<Layout[]>(
            ['layouts'],
            previousLayouts.map((layout) =>
              layout.id === variables.id ? res.data : layout
            )
          )
        }
        resetEditedLayout()
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
  
  const deleteLayoutMutation = useMutation(
    (id: number) =>
      axios.delete(`${process.env.REACT_APP_API_URL}/layouts/${id}`),
    {
      onSuccess: (_, variables) => {
        const previousLayouts = queryClient.getQueryData<Layout[]>(['layouts'])
        if (previousLayouts) {
          queryClient.setQueryData<Layout[]>(
            ['layouts'],
            previousLayouts.filter((layout) => layout.id !== variables)
          )
        }
        resetEditedLayout()
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
    createLayoutMutation,
    updateLayoutMutation,
    deleteLayoutMutation,
  }
}
