import { TaskState } from './taskTypes';
import { ExternalAPIState } from './apiTypes';
import { FeedState } from './feedTypes';
import { ArticleState } from './articleTypes';

// 各型定義を型として再エクスポート
export type { TaskState } from './taskTypes';
export type { ExternalAPIState } from './apiTypes';
export type { FeedState } from './feedTypes';
export type { ArticleState } from './articleTypes';

// 全体のアプリケーション状態の型
// 再エクスポートした型を使用
export type State = TaskState & ExternalAPIState & FeedState & ArticleState;
