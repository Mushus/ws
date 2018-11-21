# components

コンポーネントのビルド

## 環境構築

### 準備

以下が入っていたらここスキップしてください
Windows は WSL を使うかインストーラーで node 入れたほうがいいと思います。

* node なるべく最新版
* yarn

```bash
# NVM をダウンロード
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.11/install.sh | bash
# もしくは
# wget -qO- https://raw.githubusercontent.com/creationix/nvm/v0.33.11/install.sh | bash

# 起動時の設定を読み直す
source ~/.bashrc

# 一番下をコピペ
nvm ls-remote

# インストールする
nvm install [バージョン番号]

# yarn をインストールする
npm install -g yarn
```

### 本番

```bash
# カレントディレクトリをこのディレクトリに設定する
cd ./components

# 必要なライブラリ等を集める
yarn
```

## 作業方法

### ディレクトリ構成

```
src
 + assets
   + scss 全体の共通するスタイルシートを書くところです
     + variable.scss scss の変数定義
   + template 開発用の確認用テンプレートです(使い方等はこちらに記載)
     + index.html 開発用テンプレートのサンプル
 + [any dir] コンポーネントのディレクトリです
 + index.js エントリポイント
 + util.js 便利メソッド定義するところです
```

### 規約

* BEM

TODO: 足りてないので追加

### 開発用ローカルサーバー起動

立ったサーバーの情報はコンソールに出力されるのですが、だいたい `http://localhost:8080/` に立ちます。

```bash
yarn dev
```

### ビルド

```bash
yarn build
```

## TODO:

* アセットのバンドル
