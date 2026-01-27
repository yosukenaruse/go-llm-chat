https://go-llm-chat.onrender.com/static/login.html

go mod init
go mod tidy
go run main.go 

# go-llm-chat

Go言語で実装されたLLM（Large Language Model）チャットアプリケーション

## 概要

このプロジェクトは、Go言語とVue.jsを使用してDeepSeek APIと対話するチャットアプリケーションを提供します。

## 技術スタック

- **バックエンド**: Go, Gin Framework, LangChain Go
- **フロントエンド**: Vue.js 3, HTML/CSS/JavaScript
- **LLM**: DeepSeek API
- **認証**: Cookie-based認証

## 必要要件

- Go 1.21以上
- DeepSeek APIキー

## プロジェクト構造

```
go-llm-chat/
├── main.go                 # エントリーポイント
├── go.mod                  # Go モジュール定義
├── go.sum                  # 依存関係のチェックサム
├── .env                    # 環境変数（APIキーなど）
├── handlers/               # HTTPハンドラー
│   ├── chat_handler.go    # チャットのエンドポイント処理
│   └── auth_handler.go    # 認証のエンドポイント処理
├── services/               # ビジネスロジック
│   ├── chat_service.go    # DeepSeek/LangChain連携
│   └── auth_service.go    # 認証ロジック
├── models/                 # データ構造
│   └── chat.go            # リクエスト/レスポンスの型定義
└── static/                 # 静的ファイル
    ├── login.html         # ログイン画面
    ├── login.js           # ログイン画面のVue.js
    ├── chat.html          # チャット画面
    ├── chat.js            # チャット画面のVue.js
    └── hello-world.html   # テスト用ページ
```

## ローカル開発

### インストール

```bash
git clone https://github.com/yourusername/go-llm-chat.git
cd go-llm-chat
go mod download
```

### 環境変数の設定

`.env`ファイルを作成して以下の環境変数を設定：

```properties
API_KEY=your-deepseek-api-key
PASSWORD=your-login-password
```

### 起動

```bash
go run main.go
```

アプリケーションは `http://localhost:8080` で起動します。

## エンドポイント

- `GET /` - Hello World（テスト用）
- `GET /static/*` - 静的ファイル配信
- `POST /login` - ログイン（認証不要）
- `POST /logout` - ログアウト（認証不要）
- `POST /chat` - チャット（認証必須）

## 使用方法

1. ブラウザで `http://localhost:8080/static/login.html` にアクセス
2. パスワードを入力してログイン
3. チャット画面でメッセージを入力
4. DeepSeek AIが応答を生成

## Renderへのデプロイ

### 前提条件
- Renderアカウント
- GitHubリポジトリ

### デプロイ手順

1. **GitHubにプッシュ**
   ```bash
   git init
   git add .
   git commit -m "Initial commit"
   git remote add origin https://github.com/yourusername/go-llm-chat.git
   git push -u origin main
   ```

2. **Renderで新しいWeb Serviceを作成**
   - Renderダッシュボードにログイン
   - "New +" → "Web Service" をクリック
   - GitHubリポジトリを接続

3. **設定**
   - **Name**: go-llm-chat（任意の名前）
   - **Environment**: Go
   - **Build Command**: `chmod +x build.sh && ./build.sh`
   - **Start Command**: `./bin/app`
   - **Instance Type**: Free（または必要に応じて選択）

4. **環境変数の設定**
   Renderの環境変数セクションで以下を設定：
   - `API_KEY`: DeepSeekのAPIキー
   - `PASSWORD`: ログイン用パスワード
   - `PORT`: 8080（Renderが自動で設定）

5. **デプロイ**
   - "Create Web Service" をクリック
   - 自動的にビルドとデプロイが開始されます

6. **アクセス**
   - デプロイ完了後、RenderのURLからアクセス可能
   - 例: `https://go-llm-chat-xxxx.onrender.com`

### 注意事項

- Freeプランの場合、15分間アクセスがないとスリープ状態になります
- 初回アクセス時に起動に時間がかかる場合があります
- 環境変数は必ずRenderのダッシュボードで設定してください（.envファイルはGitにコミットしない）

## セキュリティ

- `.env`ファイルは`.gitignore`に追加してください
- 本番環境では必ずHTTPSを使用してください
- パスワードはハッシュ化することを推奨します（現在は平文比較）

## ライセンス

MIT License

## 作者

Yosuke Naruse
