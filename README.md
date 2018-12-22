## goNewProject
goの空プロジェクトを作成するコマンド

## 使い方
空プロジェクトを作りたいディレクトリ上で
```
$ goNewProject {projectName}
```

## 機能
### 空プロジェクト作成機能
- {projectName}のディレクトリを作成する
- {projectName}配下にREADME.mdを作成する
  - README.md内はプロジェクト名とinstallationが自動で記載される
- {projectName}配下にHello World表示するだけのmain.goを作成する

TODO:コマンド化した場合のテンプレートファイルのパスが正しくならない

### github連携機能(設計中)
githubのアカウントを登録することで、からプロジェクト作成後自動で
リモートにプッシュしておいてくれる。

登録は以下のように行う
```
$ goNewProject init
Enter your github account name:xxx

```

##　開発中・・・
### テストの雛形
- circle ciの定義ファイルを作る
- 環境変数で何かいい感じに指定したらいい感じにプッシュ後にジョブが動くようにする
- トリガーはプッシュ、タグ作成、定期実行を定義する
