# go-em-twapi-receiver

## TODO

* [x] 複数コネクションで正しく対象データが取得出来ているか
    * 取得できている。rate limiting errorsなどに達しなければ問題なく取得出来そう。
* [x] 自社アカウント以外にも対象にするべきアカウント・データを把握・設計・実装する
    * stream filterで取得すべきは自社のアカウント。一部、効果測定の投稿一覧/インプレッションで取得が必要の旨が記載されているので要確認。

* [x] アカウント一覧のファイルを変更して動的に変更出来る機能の実装
* [ ] APIサーバ化して、start&stop&restartをAPIコールで行えるようにする
* [ ] 出力ファイルの単位を考える1日・1時間・5分など(扱い易い単位とは？)

* [ ] 最初のrelease タグをどの段階で付与するか考える

* [x] Anacondaライブラリの実装で420エラーや再接続のロジックを理解する
    * Anacondaライブラリにおいて、rate limiting errors のハンドリングをしている為、利用側がハンドリングする必要がない。
    * ThrottlingにおいてはSetDelayメソッドで指定出来るが、現在はoffになっている。将来的に、デフォルトでonになるかもしれない。
