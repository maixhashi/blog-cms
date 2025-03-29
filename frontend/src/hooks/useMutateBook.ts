import axios from 'axios'
import { useQueryClient, useMutation } from '@tanstack/react-query'
import { BookResponse, BookRequest } from '../types/models/book'
import useStore from '../store'
import { useError } from './useError'

export const useMutateBook = () => {
  const queryClient = useQueryClient()
  const { switchErrorHandling } = useError()
  const resetEditedBook = useStore((state) => state.resetEditedBook)

  const createBookMutation = useMutation(
    (book: BookRequest) =>
      axios.post<BookResponse>(`${process.env.REACT_APP_API_URL}/books`, book),
    {
      onSuccess: (res) => {
        const previousBooks = queryClient.getQueryData<BookResponse[]>(['books'])
        if (previousBooks) {
          queryClient.setQueryData(['books'], [...previousBooks, res.data])
        }
        resetEditedBook()
      },
      onError: (err: any) => {
        if (err.response?.data?.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response?.data || '書籍の作成に失敗しました')
        }
      },
    }
  )

  const updateBookMutation = useMutation(
    (book: BookRequest & { id: number }) =>
      axios.put<BookResponse>(`${process.env.REACT_APP_API_URL}/books/${book.id}`, book),
    {
      onSuccess: (res, variables) => {
        const previousBooks = queryClient.getQueryData<BookResponse[]>(['books'])
        if (previousBooks) {
          queryClient.setQueryData<BookResponse[]>(
            ['books'],
            previousBooks.map((book) =>
              book.id === variables.id ? res.data : book
            )
          )
        }
        resetEditedBook()
      },
      onError: (err: any) => {
        if (err.response?.data?.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response?.data || '書籍の更新に失敗しました')
        }
      },
    }
  )

  const deleteBookMutation = useMutation(
    (id: number) =>
      axios.delete(`${process.env.REACT_APP_API_URL}/books/${id}`),
    {
      onSuccess: (_, variables) => {
        const previousBooks = queryClient.getQueryData<BookResponse[]>(['books'])
        if (previousBooks) {
          queryClient.setQueryData<BookResponse[]>(
            ['books'],
            previousBooks.filter((book) => book.id !== variables)
          )
        }
        resetEditedBook()
      },
      onError: (err: any) => {
        if (err.response?.data?.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response?.data || '書籍の削除に失敗しました')
        }
      },
    }
  )

  const importGoogleBookMutation = useMutation(
    (googleBookId: string) =>
      axios.post<BookResponse>(`${process.env.REACT_APP_API_URL}/google-books/${googleBookId}/import`),
    {
      onSuccess: (res) => {
        const previousBooks = queryClient.getQueryData<BookResponse[]>(['books'])
        if (previousBooks) {
          queryClient.setQueryData(['books'], [...previousBooks, res.data])
        }
        resetEditedBook()
      },
      onError: (err: any) => {
        if (err.response?.data?.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response?.data || 'Google Booksからの書籍インポートに失敗しました')
        }
      },
    }
  )

  return {
    createBookMutation,
    updateBookMutation,
    deleteBookMutation,
    importGoogleBookMutation,
  }
}
