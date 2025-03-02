export interface RouteNode {
  id: string;
  name: string;
  path: string;
  children?: string[]; // 遷移可能な画面のID
}

export const routes: RouteNode[] = [
  {
    id: 'home',
    name: 'ホーム画面',
    path: '/',
    children: ['external-apis', 'qiita-articles', 'hatena-articles']
  },
  {
    id: 'external-apis',
    name: '外部API管理',
    path: '/external-apis',
    children: ['external-api-edit']
  },
  {
    id: 'external-api-edit',
    name: '外部API編集',
    path: '/external-apis/edit',
    children: ['external-apis']
  },
  {
    id: 'qiita-articles',
    name: 'Qiita記事一覧',
    path: '/qiita',
    children: ['article-detail']
  },
  {
    id: 'hatena-articles',
    name: 'はてな記事一覧',
    path: '/hatena',
    children: ['article-detail']
  },
  {
    id: 'article-detail',
    name: '記事詳細',
    path: '/articles/:id',
    children: ['qiita-articles', 'hatena-articles']
  }
  // 他の画面も同様に追加
];
