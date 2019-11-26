FROM golang:latest as builder

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go/src/gin-blog
COPY . /go/src/gin-blog
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin-example .

FROM scratch

WORKDIR /gin-blog/
COPY --from=builder /go/src/gin-blog/go-gin-example .
COPY --from=builder /go/src/gin-blog/docs ./docs/
COPY --from=builder /go/src/gin-blog/conf ./conf/
COPY --from=builder /go/src/gin-blog/runtime ./runtime

EXPOSE 8000
ENTRYPOINT ["/gin-blog/go-gin-example"]