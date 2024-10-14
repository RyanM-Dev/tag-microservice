FROM golang:1.22.3-alpine AS build

WORKDIR /app

ENV GOPROXY=https://goproxy.io,direct
ENV GOPRIVATE=git.mycompany.com,github.com/my/private

COPY go.mod go.sum ./
RUN go mod download &&  go mod tidy

COPY . .

WORKDIR /app/cmd


RUN go build -o main .

FROM alpine:3.20.3
WORKDIR /app

COPY --from=build /app/cmd/main .

FROM alpine:3.20.3
WORKDIR /app

COPY --from=build /app/cmd/main .

EXPOSE 8080

CMD ["./main"]