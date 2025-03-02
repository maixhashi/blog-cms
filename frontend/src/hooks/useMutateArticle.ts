import axios from 'axios'
import { useQueryClient, useMutation } from '@tanstack/react-query'
import { Article } from '../types'
import useStore from '../store'
import { useError } from '../hooks/useError'

export const useMutateArticle = () => {
  const queryClient = useQueryClient()
  const { switchErrorHandling } = useError()
  const resetEditedArticle = useStore((state) => state.resetEditedArticle)

  const createArticleMutation = useMutation(
    (article: Omit<Article, 'id' | 'created_at' | 'updated_at'>) =>
      axios.post<Article>(`${process.env.REACT_APP_API_URL}/articles`, article),
    {
      onSuccess: (res) => {
        const previousArticles = queryClient.getQueryData<Article[]>(['articles'])
        if (previousArticles) {
          queryClient.setQueryData(['articles'], [...previousArticles, res.data])
        }
        resetEditedArticle()
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
  
  const updateArticleMutation = useMutation(
    (article: Omit<Article, 'created_at' | 'updated_at'>) =>
      axios.put<Article>(`${process.env.REACT_APP_API_URL}/articles/${article.id}`, {
        title: article.title,
        content: article.content,
        published: article.published,
        tags: article.tags,
      }),
    {
      onSuccess: (res, variables) => {
        const previousArticles = queryClient.getQueryData<Article[]>(['articles'])
        if (previousArticles) {
          queryClient.setQueryData<Article[]>(
            ['articles'],
            previousArticles.map((article) =>
              article.id === variables.id ? res.data : article
            )
          )
        }
        resetEditedArticle()
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
  
  const deleteArticleMutation = useMutation(
    (id: number) =>
      axios.delete(`${process.env.REACT_APP_API_URL}/articles/${id}`),
    {
      onSuccess: (_, variables) => {
        const previousArticles = queryClient.getQueryData<Article[]>(['articles'])
        if (previousArticles) {
          queryClient.setQueryData<Article[]>(
            ['articles'],
            previousArticles.filter((article) => article.id !== variables)
          )
        }
        resetEditedArticle()
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
    createArticleMutation,
    updateArticleMutation,
    deleteArticleMutation,
  }
}
