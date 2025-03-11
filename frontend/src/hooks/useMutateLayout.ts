import { useQueryClient, useMutation } from '@tanstack/react-query'
import { Layout } from '../types'
import useStore from '../store'
import { useError } from '../hooks/useError'
import { createLayout, updateLayout, deleteLayout } from '../api/layouts'

export const useMutateLayout = () => {
  const queryClient = useQueryClient()
  const { switchErrorHandling } = useError()
  const resetEditedLayout = useStore((state) => state.resetEditedLayout)

  const createLayoutMutation = useMutation({
    mutationFn: (layout: Omit<Layout, 'id' | 'created_at' | 'updated_at'>) => 
      createLayout(layout),
    onSuccess: (res) => {
      const previousLayouts = queryClient.getQueryData<Layout[]>(['layouts'])
      if (previousLayouts) {
        queryClient.setQueryData(['layouts'], [...previousLayouts, res])
      }
      resetEditedLayout()
    },
    onError: (err: any) => {
      if (err.response?.data?.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response?.data || 'Something went wrong')
      }
    },
  })
  
  const updateLayoutMutation = useMutation({
    mutationFn: (layout: Omit<Layout, 'created_at' | 'updated_at'>) => 
      updateLayout(layout.id, { title: layout.title }),
    onSuccess: (res, variables) => {
      const previousLayouts = queryClient.getQueryData<Layout[]>(['layouts'])
      if (previousLayouts) {
        queryClient.setQueryData<Layout[]>(
          ['layouts'],
          previousLayouts.map((layout) =>
            layout.id === variables.id ? res : layout
          )
        )
      }
      resetEditedLayout()
    },
    onError: (err: any) => {
      if (err.response?.data?.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response?.data || 'Something went wrong')
      }
    },
  })
  
  const deleteLayoutMutation = useMutation({
    mutationFn: (id: number) => deleteLayout(id),
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
      if (err.response?.data?.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response?.data || 'Something went wrong')
      }
    },
  })
  
  return {
    createLayoutMutation,
    updateLayoutMutation,
    deleteLayoutMutation,
  }
}