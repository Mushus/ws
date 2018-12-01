# app

## これはなに？

CMSを作りたい。ただし抽象化されすぎていてそのまま作るのは辛い。
なのでインターフェースごとに設計を行って内部を考える。

## 機能

facade ... 汎用管理画面。BFF
stdapp ... 汎用管理画面用アプリのボイラープレート
api ... マスタ管理APIサーバー
components ... 管理画面用コンポーネント

## 利点

- [x] 複数の画面を統合できる
  - [x] デザインが統一できる(Web component)
  - [x] セッションの管理が簡単。ミドルウェアの準備不要。勝手にやってくれる。
  - [x] ログイン画面を作る必要がない。SSO。
- [x] 設定ファイルに定義すれば管理画面が秒でつくれる
  - [x] 要件を確認しながらすぐ作れる短いサイクルで要件が詰めれる
- [x] サービス全体がダウンしなくなる
- [x] リレーションに基づいたわかりやすいDB設計(MySQL使用時)
  - [x] チューニングが簡単
  - [x] マスタースレーブを最大限に活用する設計