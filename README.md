# BlogSystemEcho



## 概要

* 誰でも投稿できるようにする
  * ログインしている人もしていない人も！
* ユーザーのログいは名前＋パスワード
* フォロー機能なし
* コメント機能なし
* カテゴリ(タグ機能)あり→タグ除外機能
* 指定日付範囲での記事の絞り込み
* **とりあえずゲスト投稿はなしで作成**
* ユーザー作成はできない→ユーザー＝管理者、増やす予定はなし
* adminのログインはベーシック認証で良さそう？



## 実装手順

- 各テーブル作成、モデル定義
- 管理者側の実装　ベーシック認証でのログイン、投稿ページを作成
- 閲覧画面の実装　とりあえず表示のみできるようにする





## DB

* dbの作成はdocker build時に行うが各テーブル、カラムはgormに任せる



| admin    |                           |
| -------- | ------------------------- |
| id       | int,primary,autoincrement |
| name     | string                    |
| password | string                    |

* ベーシック認証の場合は必要なし



| articles   |                 |
| ---------- | --------------- |
| id         | id              |
| title      | char(20),unique |
| body       | text            |
| created_at | date(Y-m-d)     |



| article_tags |      |
| ------------ | ---- |
| id           | id   |
| article_id   | int  |
| tag_id       | int  |



| tag  |          |
| ---- | -------- |
| id   | id       |
| name | char(20) |



## Route

| url          | method | about              | group  |
| ------------ | ------ | ------------------ | ------ |
| /login       | GET    | ログインフォーム   | /admin |
| /login       | POST   | ログイン処理       | /admin |
| /post        | GET    | 投稿フォーム       | /admin |
| /post        | POST   | 投稿処理           | /admin |
| /article     | GET    | 記事取得           |        |
| /article/:id | GET    | 記事閲覧           |        |
| /tag         | GET    | タグ取得           |        |
| /tag/:id     | GET    | そのタグの記事取得 |        |

* ベーシック認証の場合ログインフォームは必要なし
* 管理者側で記事の更新、削除もつける
* 日での絞り込み用のルートも考慮する