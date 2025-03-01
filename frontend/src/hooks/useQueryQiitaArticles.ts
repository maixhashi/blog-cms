import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { QiitaArticle } from '../types'

export const useQueryQiitaArticles = () => {
  const getQiitaArticles = async () => {
    const { data } = await axios.get<QiitaArticle[]>(
      `${process.env.REACT_APP_API_URL}/qiita/articles`,
      { withCredentials: true }
    )
    return data
  }

  return useQuery<QiitaArticle[], Error>({
    queryKey: ['qiitaArticles'],
    queryFn: getQiitaArticles,
    staleTime: 0,
    refetchOnWindowFocus: true,
  })
}
