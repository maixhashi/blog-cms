import { DefaultHeader } from './DefaultHeader';
import { DefaultCalendar } from './DefaultCalendar';
import { DefaultSidebar } from './DefaultSidebar';
import { DefaultFooter } from './DefaultFooter';
import { DefaultArticleCard } from './DefaultArticleCard';

export { DefaultHeader } from './DefaultHeader';
export { DefaultCalendar } from './DefaultCalendar';
export { DefaultSidebar } from './DefaultSidebar';
export { DefaultFooter } from './DefaultFooter';
export { DefaultArticleCard } from './DefaultArticleCard';

// コンポーネントタイプの定義（重複を削除）
export type ComponentType = 'header' | 'calendar' | 'sidebar' | 'footer' | 'main';

// デフォルトコンポーネントの定義
export interface DefaultComponent {
  id: string;
  name: string;
  type: ComponentType;
  component: React.FC<any>;
  defaultProps: any;
  thumbnail: string; // サムネイル画像のURL
  description: string;
}

// デフォルトコンポーネントのレジストリ
export const defaultComponents: DefaultComponent[] = [
  {
    id: 'default-header',
    name: 'デフォルトヘッダー',
    type: 'header',
    component: DefaultHeader,
    defaultProps: {
      title: 'ブログタイトル',
      links: [
        { label: 'ホーム', url: '/' },
        { label: '記事一覧', url: '/articles' },
        { label: 'カテゴリー', url: '/categories' },
        { label: '問い合わせ', url: '/contact' },
      ],
      backgroundColor: '#2196f3',
      textColor: '#ffffff',
    },
    thumbnail: '/thumbnails/default-header.png',
    description: 'シンプルなナビゲーションリンク付きヘッダー',
  },
  {
    id: 'default-calendar',
    name: 'カレンダー',
    type: 'calendar',
    component: DefaultCalendar,
    defaultProps: {
      highlightedDates: [new Date()],
      primaryColor: '#2196f3',
      secondaryColor: '#bbdefb',
    },
    thumbnail: '/thumbnails/default-calendar.png',
    description: 'ブログ投稿日をハイライト表示できるカレンダー',
  },
  {
    id: 'default-sidebar',
    name: 'サイドバー',
    type: 'sidebar',
    component: DefaultSidebar,
    defaultProps: {
      title: 'サイドバー',
      items: [
        { label: '最新の記事', url: '/latest' },
        { label: '人気の記事', url: '/popular' },
        { label: 'カテゴリー1', url: '/category/1' },
        { label: 'カテゴリー2', url: '/category/2' },
        { label: 'アーカイブ', url: '/archive' },
      ],
      backgroundColor: '#f5f5f5',
      textColor: '#333333',
    },
    thumbnail: '/thumbnails/default-sidebar.png',
    description: 'ブログナビゲーション用サイドバー',
  },
  {
    id: 'default-footer',
    name: 'フッター',
    type: 'footer',
    component: DefaultFooter,
    defaultProps: {
      title: 'ブログタイトル',
      description: 'ブログの説明文をここに入力します。',
      links: [
        {
          section: 'カテゴリー',
          items: [
            { label: 'テクノロジー', url: '/category/tech' },
            { label: 'ライフスタイル', url: '/category/lifestyle' },
            { label: '旅行', url: '/category/travel' },
          ],
        },
        {
          section: 'リンク',
          items: [
            { label: 'ホーム', url: '/' },
            { label: '記事一覧', url: '/articles' },
            { label: '問い合わせ', url: '/contact' },
          ],
        },
      ],
      copyright: '© 2023 ブログタイトル. All rights reserved.',
      backgroundColor: '#333333',
      textColor: '#ffffff',
    },
    thumbnail: '/thumbnails/default-footer.png',
    description: 'リンクセクション付きのフッター',
  },
  {
    id: 'default-article-card',
    name: '記事カード',
    type: 'main',
    component: DefaultArticleCard,
    defaultProps: {
      title: '記事タイトル',
      excerpt: 'これは記事の抜粋です。ここに記事の概要が表示されます。',
      imageUrl: 'https://source.unsplash.com/random/300x200/?blog',
      date: '2023年1月1日',
      author: '著者名',
      tags: ['タグ1', 'タグ2'],
      primaryColor: '#2196f3',
      secondaryColor: '#bbdefb',
      showImage: true,
      showTags: true,
    },
    thumbnail: '/thumbnails/default-article-card.png',
    description: 'ブログ記事を表示するカード',
  },
];
// コンポーネントを取得する関数
export const getComponentById = (id: string): DefaultComponent | undefined => {
  return defaultComponents.find(comp => comp.id === id);
};
