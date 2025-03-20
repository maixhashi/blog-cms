/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
  "/articles": {
    /** ログインユーザーのすべての記事を取得する */
    get: {
      responses: {
        /** OK */
        200: {
          schema: definitions["model.ArticleResponse"][];
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
    /** ユーザーの新しい記事を作成する */
    post: {
      parameters: {
        body: {
          /** 記事情報 */
          article: definitions["model.ArticleRequest"];
        };
      };
      responses: {
        /** Created */
        201: {
          schema: definitions["model.ArticleResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
  "/articles/{articleId}": {
    /** 指定されたIDの記事を取得する */
    get: {
      parameters: {
        path: {
          /** 記事ID */
          articleId: number;
        };
      };
      responses: {
        /** OK */
        200: {
          schema: definitions["model.ArticleResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
    /** 指定されたIDの記事を更新する */
    put: {
      parameters: {
        path: {
          /** 記事ID */
          articleId: number;
        };
        body: {
          /** 更新する記事情報 */
          article: definitions["model.ArticleRequest"];
        };
      };
      responses: {
        /** OK */
        200: {
          schema: definitions["model.ArticleResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
    /** 指定されたIDの記事を削除する */
    delete: {
      parameters: {
        path: {
          /** 記事ID */
          articleId: number;
        };
      };
      responses: {
        /** No Content */
        204: never;
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
  "/csrf-token": {
    /** CSRFトークンを取得する */
    get: {
      responses: {
        /** OK */
        200: {
          schema: definitions["model.CsrfTokenResponse"];
        };
      };
    };
  };
  "/layout-components": {
    /** ログインユーザーのすべてのレイアウトコンポーネントを取得する */
    get: {
      responses: {
        /** OK */
        200: {
          schema: definitions["model.LayoutComponentResponse"][];
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
    /** ユーザーの新しいレイアウトコンポーネントを作成する */
    post: {
      parameters: {
        body: {
          /** コンポーネント情報 */
          component: definitions["model.LayoutComponentRequest"];
        };
      };
      responses: {
        /** Created */
        201: {
          schema: definitions["model.LayoutComponentResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
  "/layout-components/{componentId}": {
    /** 指定されたIDのレイアウトコンポーネントを取得する */
    get: {
      parameters: {
        path: {
          /** コンポーネントID */
          componentId: number;
        };
      };
      responses: {
        /** OK */
        200: {
          schema: definitions["model.LayoutComponentResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
    /** 指定されたIDのレイアウトコンポーネントを更新する */
    put: {
      parameters: {
        path: {
          /** コンポーネントID */
          componentId: number;
        };
        body: {
          /** 更新するコンポーネント情報 */
          component: definitions["model.LayoutComponentRequest"];
        };
      };
      responses: {
        /** OK */
        200: {
          schema: definitions["model.LayoutComponentResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
    /** 指定されたIDのレイアウトコンポーネントを削除する */
    delete: {
      parameters: {
        path: {
          /** コンポーネントID */
          componentId: number;
        };
      };
      responses: {
        /** No Content */
        204: never;
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
  "/layout-components/{componentId}/assign": {
    /** 指定されたコンポーネントをレイアウトから削除する */
    delete: {
      parameters: {
        path: {
          /** コンポーネントID */
          componentId: number;
        };
      };
      responses: {
        /** OK */
        200: unknown;
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
  "/layout-components/{componentId}/assign/{layoutId}": {
    /** 指定されたコンポーネントを特定のレイアウトに割り当てる */
    post: {
      parameters: {
        path: {
          /** コンポーネントID */
          componentId: number;
          /** レイアウトID */
          layoutId: number;
        };
        body: {
          /** 割り当て情報 */
          position: definitions["model.AssignLayoutRequest"];
        };
      };
      responses: {
        /** OK */
        200: unknown;
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
  "/layout-components/{componentId}/position": {
    /** 指定されたコンポーネントの位置情報を更新する */
    put: {
      parameters: {
        path: {
          /** コンポーネントID */
          componentId: number;
        };
        body: {
          /** 位置情報 */
          position: definitions["model.PositionRequest"];
        };
      };
      responses: {
        /** OK */
        200: unknown;
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
  "/layouts": {
    /** ログインユーザーのすべてのレイアウトを取得する */
    get: {
      responses: {
        /** OK */
        200: {
          schema: definitions["model.LayoutResponse"][];
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
    /** ユーザーの新しいレイアウトを作成する */
    post: {
      parameters: {
        body: {
          /** レイアウト情報 */
          layout: definitions["model.LayoutRequest"];
        };
      };
      responses: {
        /** Created */
        201: {
          schema: definitions["model.LayoutResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
  "/layouts/{layoutId}": {
    /** 指定されたIDのレイアウトを取得する */
    get: {
      parameters: {
        path: {
          /** レイアウトID */
          layoutId: number;
        };
      };
      responses: {
        /** OK */
        200: {
          schema: definitions["model.LayoutResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
    /** 指定されたIDのレイアウトを更新する */
    put: {
      parameters: {
        path: {
          /** レイアウトID */
          layoutId: number;
        };
        body: {
          /** 更新するレイアウト情報 */
          layout: definitions["model.LayoutRequest"];
        };
      };
      responses: {
        /** OK */
        200: {
          schema: definitions["model.LayoutResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
    /** 指定されたIDのレイアウトを削除する */
    delete: {
      parameters: {
        path: {
          /** レイアウトID */
          layoutId: number;
        };
      };
      responses: {
        /** No Content */
        204: never;
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
  "/login": {
    /** 既存ユーザーのログイン処理 */
    post: {
      parameters: {
        body: {
          /** ログイン情報 */
          user: definitions["model.UserLoginRequest"];
        };
      };
      responses: {
        /** OK */
        200: {
          schema: string;
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
  "/logout": {
    /** ユーザーのログアウト処理 */
    post: {
      responses: {
        /** OK */
        200: {
          schema: string;
        };
      };
    };
  };
  "/signup": {
    /** 新しいユーザーアカウントを作成する */
    post: {
      parameters: {
        body: {
          /** ユーザー登録情報 */
          user: definitions["model.UserSignupRequest"];
        };
      };
      responses: {
        /** Created */
        201: {
          schema: definitions["model.UserResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
  "/tasks": {
    /** ログインユーザーのすべてのタスクを取得する */
    get: {
      responses: {
        /** OK */
        200: {
          schema: definitions["model.TaskResponse"][];
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
    /** ユーザーの新しいタスクを作成する */
    post: {
      parameters: {
        body: {
          /** タスク情報 */
          task: definitions["model.TaskRequest"];
        };
      };
      responses: {
        /** Created */
        201: {
          schema: definitions["model.TaskResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
  "/tasks/{taskId}": {
    /** 指定されたIDのタスクを取得する */
    get: {
      parameters: {
        path: {
          /** タスクID */
          taskId: number;
        };
      };
      responses: {
        /** OK */
        200: {
          schema: definitions["model.TaskResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
    /** 指定されたIDのタスクを更新する */
    put: {
      parameters: {
        path: {
          /** タスクID */
          taskId: number;
        };
        body: {
          /** 更新するタスク情報 */
          task: definitions["model.TaskRequest"];
        };
      };
      responses: {
        /** OK */
        200: {
          schema: definitions["model.TaskResponse"];
        };
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
    /** 指定されたIDのタスクを削除する */
    delete: {
      parameters: {
        path: {
          /** タスクID */
          taskId: number;
        };
      };
      responses: {
        /** No Content */
        204: never;
        /** Bad Request */
        400: {
          schema: { [key: string]: string };
        };
        /** Internal Server Error */
        500: {
          schema: { [key: string]: string };
        };
      };
    };
  };
}

export interface definitions {
  "model.ArticleRequest": {
    /** @example Goは静的型付け言語です... */
    content?: string;
    /** @example true */
    published?: boolean;
    /** @example Go,プログラミング,チュートリアル */
    tags?: string;
    /** @example Goプログラミングの基礎 */
    title: string;
  };
  "model.ArticleResponse": {
    /** @example Goは静的型付け言語です... */
    content?: string;
    /** @example 2023-01-01T00:00:00Z */
    created_at?: string;
    /** @example 1 */
    id?: number;
    /** @example true */
    published?: boolean;
    /** @example Go,プログラミング,チュートリアル */
    tags?: string;
    /** @example Goプログラミングの基礎 */
    title?: string;
    /** @example 2023-01-01T00:00:00Z */
    updated_at?: string;
  };
  "model.AssignLayoutRequest": {
    /** @example 1 */
    layout_id: number;
    position?: definitions["model.PositionRequest"];
  };
  "model.CsrfTokenResponse": {
    /** @example token-string-here */
    csrf_token?: string;
  };
  "model.LayoutComponentRequest": {
    /** @example <h1>ブログタイトル</h1> */
    content?: string;
    /** @example 50 */
    height?: number;
    /** @example ヘッダーコンポーネント */
    name: string;
    /** @example header */
    type: string;
    /** @example 100 */
    width?: number;
    /** @example 0 */
    x?: number;
    /** @example 0 */
    y?: number;
  };
  "model.LayoutComponentResponse": {
    /** @example <h1>ブログタイトル</h1> */
    content?: string;
    /** @example 2023-01-01T00:00:00Z */
    created_at?: string;
    /** @example 50 */
    height?: number;
    /** @example 1 */
    id?: number;
    /** @example 1 */
    layout_id?: number;
    /** @example ヘッダーコンポーネント */
    name?: string;
    /** @example header */
    type?: string;
    /** @example 2023-01-01T00:00:00Z */
    updated_at?: string;
    /** @example 100 */
    width?: number;
    /** @example 0 */
    x?: number;
    /** @example 0 */
    y?: number;
  };
  "model.LayoutRequest": {
    /** @example ブログのメインレイアウト */
    title: string;
  };
  "model.LayoutResponse": {
    components?: definitions["model.LayoutComponentResponse"][];
    /** @example 2023-01-01T00:00:00Z */
    created_at?: string;
    /** @example 1 */
    id?: number;
    /** @example ブログのメインレイアウト */
    title?: string;
    /** @example 2023-01-01T00:00:00Z */
    updated_at?: string;
  };
  "model.PositionRequest": {
    /** @example 75 */
    height?: number;
    /** @example 150 */
    width?: number;
    /** @example 10 */
    x?: number;
    /** @example 20 */
    y?: number;
  };
  "model.TaskRequest": {
    /** @example 買い物に行く */
    title: string;
  };
  "model.TaskResponse": {
    /** @example 2023-01-01T00:00:00Z */
    created_at?: string;
    /** @example 1 */
    id?: number;
    /** @example 買い物に行く */
    title?: string;
    /** @example 2023-01-01T00:00:00Z */
    updated_at?: string;
  };
  "model.UserLoginRequest": {
    /** @example user@example.com */
    email: string;
    /** @example password123 */
    password: string;
  };
  "model.UserResponse": {
    /** @example user@example.com */
    email?: string;
    /** @example 1 */
    id?: number;
  };
  "model.UserSignupRequest": {
    /** @example user@example.com */
    email: string;
    /** @example password123 */
    password: string;
  };
}

export interface operations {}

export interface external {}
