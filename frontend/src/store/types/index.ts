import { TaskState } from './taskTypes';
import { LayoutState, LayoutComponentState } from './layoutTypes';
import { ExternalAPIState } from './apiTypes';
import { FeedState } from './feedTypes';
import { ArticleState } from './articleTypes';

// 各型定義を型として再エクスポート
export type { TaskState } from './taskTypes';
export type { LayoutState, LayoutComponentState } from './layoutTypes';
export type { ExternalAPIState } from './apiTypes';
export type { FeedState } from './feedTypes';
export type { ArticleState } from './articleTypes';

// 修正点: 単純に型を組み合わせる
export type State = 
  & TaskState 
  & LayoutState 
  & LayoutComponentState 
  & ExternalAPIState
  & FeedState 
  & ArticleState;