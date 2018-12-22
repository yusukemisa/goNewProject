## 作業録
### circleciのローカル実行
コマンドをダウンロード
```
$ curl -o /usr/local/bin/circleci 
# 下記に入るのでパスを追加する
/usr/local/bin/circleci

# はいったケロ
$ circleci update
Already up-to-date.

# homebrewでも普通にできるようなのでこっちの方が手っ取り早かった模様
$ brew install circleci
```

### Configuring The CLI
下記よりCircleCIのAPIトークンを取得してきてコマンドに設定する
https://circleci.com/account/api

```
circleci setup
```

### ymlのバリデーション 
```
$ circleci config validate -c .circleci/config.yml
```

### 実行
実行したところなんか失敗・・・
```
$ circleci build
Downloading latest CircleCI build agent...
Error: Could not find picard image: failed to pull latest docker image: exit status 1
```
exit status 1ってなんだ

結局circleciコマンド自体Go製だったのでソースを調べてみると、docker pullコマンドの実行が失敗していた模様。

ターミナルで本来実行しようとしていたコマンドを実行すると・・・
```
$ docker pull circleci/picard
Using default tag: latest
Error response from daemon: Get https://registry-1.docker.io/v2/circleci/picard/manifests/latest: unauthorized: incorrect username or password
# まあたしかにexit codeが1だけど・・・
$ echo $?
1
```
結局docker hubの認証に問題、というか認証の有効期限が切れてたっぽいのでログインしなおし
```
$ docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: yusukemisawa
Password: 
Login Succeeded
```
これでもう一度トライしたところうまくいった。

```
$ circleci build
:
:
====>> Checkout code
  #!/bin/bash -eo pipefail
mkdir -p /go/src/github.com/{{ORG_NAME}}/{{REPO_NAME}} && cp -r /tmp/_circleci_local_build_repo/. /go/src/github.com/{{ORG_NAME}}/{{REPO_NAME}}
====>> echo 'Hello world!'
  #!/bin/bash -eo pipefail
echo 'Hello world!'
Hello world!
Success!

```
