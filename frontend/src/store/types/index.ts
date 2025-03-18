import { TaskState } from './taskTypes';
import { ExternalAPIState } from './apiTypes';
import { FeedState } from './feedTypes';
import { ArticleState } from './articleTypes';
import { LayoutState, LayoutComponentState } from './layoutTypes';

// 各型定義を型として再エクスポート
export type { TaskState } from './taskTypes';
export type { LayoutState, LayoutComponentState } from './layoutTypes';
export type { ExternalAPIState } from './apiTypes';
export type { FeedState } from './feedTypes';
export type { ArticleState } from './articleTypes';

// すべてのスライスの状態を合成した型
export type State = TaskState & ExternalAPIState & FeedState & ArticleState & LayoutState & LayoutComponentState;