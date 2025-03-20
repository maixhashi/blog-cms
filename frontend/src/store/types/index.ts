import { TaskState } from './taskTypes';
import { LayoutComponentState } from './layoutComponentTypes';
import { ExternalAPIState } from './apiTypes';
import { FeedState } from './feedTypes';
import { ArticleState } from './articleTypes';
import { LayoutState } from './layoutTypes';

// 各型定義を型として再エクスポート
export type { TaskState } from './taskTypes';
export type { LayoutComponentState } from './layoutComponentTypes';
export type { LayoutState } from './layoutTypes';
export type { ExternalAPIState } from './apiTypes';
export type { FeedState } from './feedTypes';
export type { ArticleState } from './articleTypes';

// すべてのスライスの状態を合成した型
export interface State extends 
  TaskState, 
  LayoutComponentState,
  ExternalAPIState,
  FeedState,
  ArticleState,
  LayoutState
{}