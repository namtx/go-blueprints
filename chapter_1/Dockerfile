FROM golang:1.14.2 AS builder

WORKDIR /go/src/github.com/namtx/go-blueprints/chapter_1

COPY . .

RUN go get -d

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Staging
FROM alpine as staging
ENV GITHUB_APP_CLIENT_ID \
  GITHUB_APP_CLIENT_SECRET

WORKDIR /root/
COPY --from=builder /go/src/github.com/namtx/go-blueprints/chapter_1/app .

COPY ./templates ./templates
RUN mkdir avatars

EXPOSE 3000

CMD ["./app", "-addr", ":3000"]
