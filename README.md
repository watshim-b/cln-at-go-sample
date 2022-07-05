# cln-at-go-sample
For organizing clean architecture


# 実行方法
1. docker compose up -d
2. http://localhost/{apiの種類}でAPIの実行が可能

# アプリケーションを更新したい場合
1. dockerイメージを再進化する
    docker build ./ -t cln-at-go-sample

2. docker compose up -dでコンテナを再起動する

# TODO
・ アプリケーションの拡張
・ docker-composeにbuildをさせるようにする