<div align="center">

![qing-api](/web/public/logo.png)

# Qing API

🍥 **次世代大規模モデルゲートウェイとAI資産管理システム**

<p align="center">
  <a href="./README.zh_CN.md">简体中文</a> |
  <a href="./README.zh_TW.md">繁體中文</a> |
  <a href="./README.md">English</a> |
  <a href="./README.fr.md">Français</a> |
  <strong>日本語</strong>
</p>

<p align="center">
  <a href="https://raw.githubusercontent.com/Isiah998/new-api/main/LICENSE">
    <img src="https://img.shields.io/github/license/Isiah998/new-api?color=brightgreen" alt="license">
  </a><!--
  --><a href="https://github.com/Isiah998/new-api/releases/latest">
    <img src="https://img.shields.io/github/v/release/Isiah998/new-api?color=brightgreen&include_prereleases" alt="release">
  </a><!--
  --><a href="https://github.com/Isiah998/new-api/pkgs/container/qing-api">
    <img src="https://img.shields.io/badge/container-ghcr.io-blue" alt="docker">
  </a><!--
  --><a href="https://goreportcard.com/report/github.com/Isiah998/new-api">
    <img src="https://goreportcard.com/badge/github.com/Isiah998/new-api" alt="GoReportCard">
  </a>
</p>

<p align="center">
  <a href="#-クイックスタート">クイックスタート</a> •
  <a href="#-主な機能">主な機能</a> •
  <a href="#-デプロイ">デプロイ</a> •
  <a href="#-ドキュメント">ドキュメント</a> •
  <a href="#-ヘルプサポート">ヘルプ</a>
</p>

</div>

## 📝 プロジェクト説明

> [!IMPORTANT]
> - 本プロジェクトは、合法的に許可された AI API ゲートウェイ、組織レベルの認証、マルチモデル管理、利用量分析、コスト管理、プライベートデプロイのシナリオのみを対象としています。
> - ユーザーは、上流の API キー、アカウント、モデルサービス、インターフェース権限を合法的に取得し、上流のサービス利用規約および適用される法律法規を遵守する必要があります。
> - ユーザーは、利用方法が上流のサービス利用規約および適用される法律法規に準拠していることを確認してください。
> - 生成 AI サービスを公衆に提供する場合、ユーザーは適用される規制要件を遵守し、管轄区域で求められる届出、ライセンス、コンテンツセキュリティ、本人確認、ログ保持、税務、上流認可などのすべての義務を履行してください。

---

## 🤝 信頼できるパートナー

<p align="center">
  <em>順不同</em>
</p>

<p align="center">
  <a href="https://www.cherry-ai.com/" target="_blank">
    <img src="./docs/images/cherry-studio.png" alt="Cherry Studio" height="80" />
  </a><!--
  --><a href="https://github.com/iOfficeAI/AionUi/" target="_blank">
    <img src="./docs/images/aionui.png" alt="Aion UI" height="80" />
  </a><!--
  --><a href="https://bda.pku.edu.cn/" target="_blank">
    <img src="./docs/images/pku.png" alt="北京大学" height="80" />
  </a><!--
  --><a href="https://www.compshare.cn/?ytag=GPU_yy_gh_qingapi" target="_blank">
    <img src="./docs/images/ucloud.png" alt="UCloud 優刻得" height="80" />
  </a><!--
  --><a href="https://www.aliyun.com/" target="_blank">
    <img src="./docs/images/aliyun.png" alt="Alibaba Cloud" height="80" />
  </a><!--
  --><a href="https://io.net/" target="_blank">
    <img src="./docs/images/io-net.png" alt="IO.NET" height="80" />
  </a>
</p>

---

## 🙏 特別な感謝

<p align="center">
  <a href="https://www.jetbrains.com/?from=qing-api" target="_blank">
    <img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png" alt="JetBrains Logo" width="120" />
  </a>
</p>

<p align="center">
  <strong>感謝 <a href="https://www.jetbrains.com/?from=qing-api">JetBrains</a> が本プロジェクトに無料のオープンソース開発ライセンスを提供してくれたことに感謝します</strong>
</p>

---

## 🚀 クイックスタート

### Docker Composeを使用（推奨）

```bash
# プロジェクトをクローン
git clone https://github.com/Isiah998/new-api.git
cd qing-api

# docker-compose.yml 設定を編集
nano docker-compose.yml

# サービスを起動
docker-compose up -d
```

<details>
<summary><strong>Dockerコマンドを使用</strong></summary>

```bash
# 最新のイメージをプル
docker pull ghcr.io/isiah998/qing-api:latest

# SQLiteを使用（デフォルト）
docker run --name qing-api -d --restart always \
  -p 3000:3000 \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  ghcr.io/isiah998/qing-api:latest

# MySQLを使用
docker run --name qing-api -d --restart always \
  -p 3000:3000 \
  -e SQL_DSN="root:123456@tcp(localhost:3306)/oneapi" \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  ghcr.io/isiah998/qing-api:latest
```

> **💡 ヒント:** `-v ./data:/data` は現在のディレクトリの `data` フォルダにデータを保存します。絶対パスに変更することもできます：`-v /your/custom/path:/data`

</details>

---

🎉 デプロイが完了したら、`http://localhost:3000` にアクセスして使用を開始してください！

> [!WARNING]
> 本プロジェクトを公衆向け生成 AI サービスまたは API 再販サービスとして運営する場合、ユーザーは届出、コンテンツセキュリティ、本人確認、ログ保持、税務、決済、上流認可などの必要なコンプライアンス義務を先に完了してください。

📖 その他のデプロイ方法については[デプロイガイド](https://github.com/Isiah998/new-api/tree/main/docs/installation)を参照してください。

---

## 📚 ドキュメント

<div align="center">

### 📖 [公式ドキュメント](https://github.com/Isiah998/new-api/tree/main/docs) | [![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/Isiah998/new-api)

</div>

**クイックナビゲーション:**

| カテゴリ | リンク |
|------|------|
| 🚀 デプロイガイド | [インストールドキュメント](https://github.com/Isiah998/new-api/tree/main/docs/installation) |
| ⚙️ 環境設定 | [環境変数](https://github.com/Isiah998/new-api/tree/main/docs/installation) |
| 📡 APIドキュメント | [APIドキュメント](https://github.com/Isiah998/new-api/tree/main/docs) |
| ❓ よくある質問 | [FAQ](https://github.com/Isiah998/new-api/issues) |
| 💬 コミュニティ交流 | [交流チャネル](https://github.com/Isiah998/new-api/issues) |

---

## ✨ 主な機能

> 詳細な機能については[機能説明](https://github.com/Isiah998/new-api/tree/main/docs)を参照してください。

### 🎨 コア機能

| 機能 | 説明 |
|------|------|
| 🎨 新しいUI | モダンなユーザーインターフェースデザイン |
| 🌍 多言語 | 簡体字中国語、繁体字中国語、英語、フランス語、日本語をサポート |
| 🔄 データ互換性 | オリジナルのOne APIデータベースと完全に互換性あり |
| 📈 データダッシュボード | ビジュアルコンソールと統計分析 |
| 🔒 権限管理 | トークングループ化、モデル制限、ユーザー管理 |

### 💰 認可済み利用量とコスト管理

- ✅ 合法的に許可されたシナリオでの内部チャージとクォータ割り当て（EPay、Stripe）
- ✅ 組織レベルのリクエスト単位、使用量ベース、キャッシュヒットのコスト会計
- ✅ OpenAI、Azure、DeepSeek、Claude、Qwen などのモデルのキャッシュ課金統計
- ✅ 内部管理または認可済み企業顧客向けの柔軟な課金ポリシー

### 🔐 認証とセキュリティ

- 😈 Discord認証ログイン
- 🤖 LinuxDO認証ログイン
- 📱 Telegram認証ログイン
- 🔑 OIDC統一認証
- 🔍 Key使用量クォータ照会（[new-api-key-tool](https://github.com/Calcium-Ion/new-api-key-tool)と併用）



### 🚀 高度な機能

**APIフォーマットサポート:**
- ⚡ [OpenAI Responses](https://github.com/Isiah998/new-api/tree/main/docs)
- ⚡ [OpenAI Realtime API](https://github.com/Isiah998/new-api/tree/main/docs)（Azureを含む）
- ⚡ [Claude Messages](https://github.com/Isiah998/new-api/tree/main/docs)
- ⚡ [Google Gemini](https://github.com/Isiah998/new-api/tree/main/docs)
- 🔄 [Rerankモデル](https://github.com/Isiah998/new-api/tree/main/docs)（Cohere、Jina）

**インテリジェントルーティング:**
- ⚖️ チャネル重み付けランダム
- 🔄 失敗自動リトライ
- 🚦 ユーザーレベルモデルレート制限

**フォーマット変換:**
- 🔄 **OpenAI Compatible ⇄ Claude Messages**
- 🔄 **OpenAI Compatible → Google Gemini**
- 🔄 **Google Gemini → OpenAI Compatible** - テキストのみ、関数呼び出しはまだサポートされていません
- 🚧 **OpenAI Compatible ⇄ OpenAI Responses** - 開発中
- 🔄 **思考からコンテンツへの機能**

**Reasoning Effort サポート:**

<details>
<summary>詳細設定を表示</summary>

**OpenAIシリーズモデル:**
- `o3-mini-high` - 高思考努力
- `o3-mini-medium` - 中思考努力
- `o3-mini-low` - 低思考努力
- `gpt-5-high` - 高思考努力
- `gpt-5-medium` - 中思考努力
- `gpt-5-low` - 低思考努力

**Claude思考モデル:**
- `claude-3-7-sonnet-20250219-thinking` - 思考モードを有効にする

**Google Geminiシリーズモデル:**
- `gemini-2.5-flash-thinking` - 思考モードを有効にする
- `gemini-2.5-flash-nothinking` - 思考モードを無効にする
- `gemini-2.5-pro-thinking` - 思考モードを有効にする
- `gemini-2.5-pro-thinking-128` - 思考モードを有効にし、思考予算を128トークンに設定する
- Gemini モデル名の末尾に `-low` / `-medium` / `-high` を付けることで推論強度を直接指定できます（追加の思考予算サフィックスは不要です）。

</details>

---

## 🤖 モデルサポート

> 詳細については[APIドキュメント - ゲートウェイインターフェース](https://github.com/Isiah998/new-api/tree/main/docs)

| モデルタイプ | 説明 | ドキュメント |
|---------|------|------|
| 🤖 OpenAI-Compatible | OpenAI互換モデル | [ドキュメント](https://github.com/Isiah998/new-api/tree/main/docs) |
| 🤖 OpenAI Responses | OpenAI Responsesフォーマット | [ドキュメント](https://github.com/Isiah998/new-api/tree/main/docs) |
| 🎨 Midjourney-Proxy | [Midjourney-Proxy(Plus)](https://github.com/novicezk/midjourney-proxy) | [ドキュメント](https://github.com/Isiah998/new-api/tree/main/docs) |
| 🎵 Suno-API | [Suno API](https://github.com/Suno-API/Suno-API) | [ドキュメント](https://github.com/Isiah998/new-api/tree/main/docs) |
| 🔄 Rerank | Cohere、Jina | [ドキュメント](https://github.com/Isiah998/new-api/tree/main/docs) |
| 💬 Claude | Messagesフォーマット | [ドキュメント](https://github.com/Isiah998/new-api/tree/main/docs) |
| 🌐 Gemini | Google Geminiフォーマット | [ドキュメント](https://github.com/Isiah998/new-api/tree/main/docs) |
| 🔧 Dify | ChatFlowモード | - |
| 🎯 カスタム上流 | 合法的に許可された上流エンドポイントの設定をサポート | - |

### 📡 サポートされているインターフェース

<details>
<summary>完全なインターフェースリストを表示</summary>

- [チャットインターフェース (Chat Completions)](https://github.com/Isiah998/new-api/tree/main/docs)
- [レスポンスインターフェース (Responses)](https://github.com/Isiah998/new-api/tree/main/docs)
- [イメージインターフェース (Image)](https://github.com/Isiah998/new-api/tree/main/docs)
- [オーディオインターフェース (Audio)](https://github.com/Isiah998/new-api/tree/main/docs)
- [ビデオインターフェース (Video)](https://github.com/Isiah998/new-api/tree/main/docs)
- [エンベッドインターフェース (Embeddings)](https://github.com/Isiah998/new-api/tree/main/docs)
- [再ランク付けインターフェース (Rerank)](https://github.com/Isiah998/new-api/tree/main/docs)
- [リアルタイム対話インターフェース (Realtime)](https://github.com/Isiah998/new-api/tree/main/docs)
- [Claudeチャット](https://github.com/Isiah998/new-api/tree/main/docs)
- [Google Geminiチャット](https://github.com/Isiah998/new-api/tree/main/docs)

</details>

---

## 🚢 デプロイ

> [!TIP]
> **最新のDockerイメージ:** `ghcr.io/isiah998/qing-api:latest`

### 📋 デプロイ要件

| コンポーネント | 要件 |
|------|------|
| **ローカルデータベース** | SQLite（Dockerは `/data` ディレクトリをマウントする必要があります）|
| **リモートデータベース** | MySQL ≥ 5.7.8 または PostgreSQL ≥ 9.6 |
| **コンテナエンジン** | Docker / Docker Compose |
| **システムアーキテクチャ** | 64ビットのみ対応（amd64 / arm64）。32ビットシステムは非対応 |

### ⚙️ 環境変数設定

<details>
<summary>一般的な環境変数設定</summary>

| 変数名 | 説明 | デフォルト値 |
|--------|------|--------|
| `SESSION_SECRET` | 認証署名シークレット。すべてのノードで同じ値が必要 | - |
| `SESSION_COOKIE_SECURE` | `false`/未設定ではローカル HTTP 開発プロキシ向けに refresh/logout の OriginGuard を無効化し、`true` では Secure Cookie と厳格な Origin 検証を有効化 | `false` |
| `SESSION_COOKIE_TRUSTED_URL` | Secure モードでは必須。refresh/logout を許可する完全一致の HTTPS Origin をカンマ区切りで指定。relay CORS 設定ではありません | - |
| `TRUSTED_PROXIES` | 未設定/空ではループバック、RFC 1918、IPv6 ULA を信頼して起動時に警告し、`none` ではすべて無効、明示的なプロキシ IP/CIDR リストは既定値を完全に置き換えます | `127.0.0.0/8, ::1, 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, fc00::/7` |
| `USER_SESSION_ACTIVE_LIMIT` | 1 ユーザーあたりの有効なログイン Session 上限 | `50` |
| `USER_SESSION_ISSUANCE_LIMIT` | カウント期間内に作成できる Session 数の上限（取り消し済みを含む） | `100` |
| `USER_SESSION_ISSUANCE_WINDOW_SECONDS` | Session 発行のカウント期間（秒）。取り消し済み Session の保持期間を超える場合は自動的に制限 | `86400` |
| `USER_SESSION_REVOKED_RETENTION_DAYS` | 監査と発行数計算のため取り消し済み Session を保持する日数 | `7` |
| `USER_SESSION_HOURLY_ALERT_THRESHOLD` | 1 時間あたりのグローバル Session 発行数の警告閾値。ログインは拒否しません | `5000` |
| `CRYPTO_SECRET` | キャッシュキー用 HMAC シークレット。Redis を共有するノードでは同じ実効値が必要 | デフォルトは `SESSION_SECRET` |
| `SQL_DSN** | データベース接続文字列 | - |
| `REDIS_CONN_STRING` | Redis接続文字列 | - |
| `STREAMING_TIMEOUT` | ストリーミング応答のタイムアウト時間（秒） | `300` |
| `STREAM_SCANNER_MAX_BUFFER_MB` | ストリームスキャナの1行あたりバッファ上限（MB）。4K画像など巨大なbase64 `data:` ペイロードを扱う場合は値を増加させてください | `64` |
| `MAX_REQUEST_BODY_MB` | リクエストボディ最大サイズ（MB、**解凍後**に計測。巨大リクエスト/zip bomb によるメモリ枯渇を防止）。超過時は `413` | `32` |
| `AZURE_DEFAULT_API_VERSION` | Azure APIバージョン | `2025-04-01-preview` |
| `ERROR_LOG_ENABLED` | エラーログスイッチ | `false` |
| `PYROSCOPE_URL` | Pyroscopeサーバーのアドレス | - |
| `PYROSCOPE_APP_NAME` | Pyroscopeアプリ名 | `qing-api` |
| `PYROSCOPE_BASIC_AUTH_USER` | Pyroscope Basic Authユーザー | - |
| `PYROSCOPE_BASIC_AUTH_PASSWORD` | Pyroscope Basic Authパスワード | - |
| `PYROSCOPE_MUTEX_RATE` | Pyroscope mutexサンプリング率 | `5` |
| `PYROSCOPE_BLOCK_RATE` | Pyroscope blockサンプリング率 | `5` |
| `HOSTNAME` | Pyroscope用のホスト名タグ | `qing-api` |

📖 **完全な設定:** [環境変数ドキュメント](https://github.com/Isiah998/new-api/tree/main/docs/installation)

</details>

### 🔧 デプロイ方法

<details>
<summary><strong>方法 1: Docker Compose（推奨）</strong></summary>

```bash
# プロジェクトをクローン
git clone https://github.com/Isiah998/new-api.git
cd qing-api

# 設定を編集
nano docker-compose.yml

# サービスを起動
docker-compose up -d
```

</details>

<details>
<summary><strong>方法 2: Dockerコマンド</strong></summary>

**SQLiteを使用:**
```bash
docker run --name qing-api -d --restart always \
  -p 3000:3000 \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  ghcr.io/isiah998/qing-api:latest
```

**MySQLを使用:**
```bash
docker run --name qing-api -d --restart always \
  -p 3000:3000 \
  -e SQL_DSN="root:123456@tcp(localhost:3306)/oneapi" \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  ghcr.io/isiah998/qing-api:latest
```

> **💡 パス説明:**
> - `./data:/data` - 相対パス、データは現在のディレクトリのdataフォルダに保存されます
> - 絶対パスを使用することもできます：`/your/custom/path:/data`

</details>

<details>
<summary><strong>方法 3: 宝塔パネル</strong></summary>

1. 宝塔パネル（**9.2.0バージョン**以上）をインストールし、アプリケーションストアで**Qing API**を検索してインストールします。

📖 [画像付きチュートリアル](./docs/BT.md)

</details>

### ⚠️ マルチマシンデプロイの注意事項

> [!WARNING]
> - すべてのノードで同じプライマリデータベースと同じ `SESSION_SECRET` を使用してください。異なる場合、Access Token、Refresh セッション、一時認証フローを一貫して検証できません。
> - 同じ Redis に接続するノードでは同じ `CRYPTO_SECRET` も設定してください。異なる場合、キャッシュキーのダイジェストが一致せず、共有エントリを正しく再利用できません。

ログイン Session とユーザー単位の有効数／発行数制限では、データベースが信頼できる唯一の情報源です。Redis の Session エントリは短期キャッシュであり、TTL は `SYNC_FREQUENCY`（デフォルト 60 秒）に従い、Session の残り有効期間を超えません。

| Redis トポロジー | Session 状態の伝播 | レート制限 |
| --- | --- | --- |
| すべてのノードで Redis を共有 | 取り消しとバージョン更新は通常即時に伝播 | Redis の制限枠はノード間で共有 |
| ノードごとに独立した Redis | 有効な `SYNC_FREQUENCY` 以内にデータベースへフォールバックして収束。バージョンローテーション直後の新しい Token は、古いキャッシュを持つノードで一時的に 401 になる場合があります | ノードごとに独立して計数するため、クラスター全体では設定値の約ノード数倍まで許可される可能性があります |
| Redis なし | Session の検証ごとにデータベースを直接参照 | メモリ内の制限枠はノードごとに独立 |

`SYNC_FREQUENCY` を短くすると独立 Redis のキャッシュ陳腐化時間は短くなりますが、有効な SID ごと、ノードごと、TTL ごとにデータベースへの主キー照会が 1 回増えます。この保証は Session 認証の陳腐化時間を限定するものです。レート制限や Redis を使うその他のコントロールプレーンキャッシュは、引き続きトポロジーに依存します。

Token、Origin 検証、PAT の契約については[ユーザー認証とログインセッション](./docs/authentication.md)を参照してください。

### 🔄 チャネルリトライとキャッシュ

**リトライ設定:** `設定 → 運営設定 → 一般設定 → 失敗リトライ回数`

**キャッシュ設定:**
- `REDIS_CONN_STRING`：Redisキャッシュ（推奨）
- `MEMORY_CACHE_ENABLED`：メモリキャッシュ

---

## 🔗 関連プロジェクト

### 上流プロジェクト

| プロジェクト | 説明 |
|------|------|
| [One API](https://github.com/songquanpeng/one-api) | オリジナルプロジェクトベース |
| [Midjourney-Proxy](https://github.com/novicezk/midjourney-proxy) | Midjourneyインターフェースサポート |

### 補助ツール

| プロジェクト | 説明 |
|------|------|
| [new-api-key-tool](https://github.com/Calcium-Ion/new-api-key-tool) | キー使用量クォータ照会ツール |
| [new-api-horizon](https://github.com/Calcium-Ion/new-api-horizon) | New API高性能最適化版 |

---

## 💬 ヘルプサポート

### 📖 ドキュメントリソース

| リソース | リンク |
|------|------|
| 📘 よくある質問 | [FAQ](https://github.com/Isiah998/new-api/issues) |
| 💬 コミュニティ交流 | [交流チャネル](https://github.com/Isiah998/new-api/issues) |
| 🐛 問題のフィードバック | [問題フィードバック](https://github.com/Isiah998/new-api/issues) |
| 📚 完全なドキュメント | [公式ドキュメント](https://github.com/Isiah998/new-api/tree/main/docs) |

### 🤝 貢献ガイド

あらゆる形の貢献を歓迎します！

- 🐛 バグを報告する
- 💡 新しい機能を提案する
- 📝 ドキュメントを改善する
- 🔧 コードを提出する

---

## 📜 ライセンス

Qing API は [New API](https://github.com/QuantumNous/new-api) の改変版であり、New API は [One API](https://github.com/songquanpeng/one-api) を基に開発されています。本バージョンも [GNU Affero General Public License v3.0（AGPLv3）](./LICENSE)で公開されます。

- この Qing API バージョンの対応ソースコードは [Isiah998/new-api](https://github.com/Isiah998/new-api) で公開されています。
- 既存の著作権表示および New API の追加帰属条件は [NOTICE](./NOTICE) に保持されています。
- 改変版をネットワークサービスとして運用する場合、AGPLv3 第13条に従い、ネットワーク経由で利用するすべてのユーザーに対応ソースコードの取得手段を提示する必要があります。
- Qing API はこの改変ディストリビューションの製品ブランドであり、ブランド変更によって上流貢献者の著作権帰属が置き換わることはありません。

---

## 🌟 スター履歴

<div align="center">

[![スター履歴チャート](https://api.star-history.com/svg?repos=Isiah998/new-api&type=Date)](https://star-history.com/#Isiah998/new-api&Date)

</div>

---

<div align="center">

### 💖 Qing APIをご利用いただきありがとうございます

このプロジェクトがあなたのお役に立てたなら、ぜひ ⭐️ スターをください！

**[公式ドキュメント](https://github.com/Isiah998/new-api/tree/main/docs)** • **[問題フィードバック](https://github.com/Isiah998/new-api/issues)** • **[最新リリース](https://github.com/Isiah998/new-api/releases)**

<sub>❤️ で構築された QingFlow</sub>

</div>
