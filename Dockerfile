# goの実行環境のイメージの取得
FROM golang:1.18.3 as builder

# dockerfileで利用する作業用ディレクトリ（コンテナ内で作成される）
WORKDIR /app

# ライブラリの関連ファイルをコンテナにコピーして、インストールする
COPY go.mod go.sum ./
RUN go mod download

# プロジェクト全てをコンテナにコピーする
COPY .  ./

# goの実行ファイルを生成する
RUN go build -o /server

# portを開放する。このポートであれば、外部からアクセスできるようになる
EXPOSE 8080

# go buildで作成したバイナリを実行する
CMD [ "/server" ]