import { useQueryClient, useMutation } from '@tanstack/react-query'
import { LayoutComponent } from '../types'
import useStore from '../store'
import { useError } from '../hooks/useError'
import { createLayoutComponent, updateLayoutComponent, deleteLayoutComponent } from '../api/layoutComponents'

export const useMutateLayoutComponent = () => {
  const queryClient = useQueryClient()
  const { switchErrorHandling } = useError()
  const resetEditedLayoutComponent = useStore((state) => state.resetEditedLayoutComponent)

  const createLayoutComponentMutation = useMutation({
    mutationFn: (component: Omit<LayoutComponent, 'id' | 'created_at' | 'updated_at'>) => 
      createLayoutComponent(component),
    onSuccess: (res) => {
      const previousComponents = queryClient.getQueryData<LayoutComponent[]>(['layoutComponents'])
      if (previousComponents) {
        queryClient.setQueryData(['layoutComponents'], [...previousComponents, res])
      }
      resetEditedLayoutComponent()
    },
    onError: (err: any) => {
      if (err.response?.data?.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response?.data || 'Something went wrong')
      }
    },
  })
  
  const updateLayoutComponentMutation = useMutation({
    mutationFn: (component: Omit<LayoutComponent, 'created_at' | 'updated_at'>) => 
      updateLayoutComponent(component.id, { name: component.name, type: component.type, content: component.content }),
    onSuccess: (res, variables) => {
      const previousComponents = queryClient.getQueryData<LayoutComponent[]>(['layoutComponents'])
      if (previousComponents) {
        queryClient.setQueryData<LayoutComponent[]>(
          ['layoutComponents'],
          previousComponents.map((component) =>
            component.id === variables.id ? res : component
          )
        )
      }
      resetEditedLayoutComponent()
    },
    onError: (err: any) => {
      if (err.response?.data?.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response?.data || 'Something went wrong')
      }
    },
  })
  
  const deleteLayoutComponentMutation = useMutation({
    mutationFn: (id: number) => deleteLayoutComponent(id),
    onSuccess: (_, variables) => {
      const previousComponents = queryClient.getQueryData<LayoutComponent[]>(['layoutComponents'])
      if (previousComponents) {
        queryClient.setQueryData<LayoutComponent[]>(
          ['layoutComponents'],
          previousComponents.filter((component) => component.id !== variables)
        )
      }
      resetEditedLayoutComponent()
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
    createLayoutComponentMutation,
    updateLayoutComponentMutation,
    deleteLayoutComponentMutation,
  }
}
