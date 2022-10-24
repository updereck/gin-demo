# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work
# GoApp起動
CMD ["go","run","main.go"]
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD src/ /go/src/work
