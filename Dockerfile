FROM golang:1.22-alpine 

WORKDIR /app

ENV GOPROXY=https://goproxy.io,direct
ENV GOPRIVATE=git.mycompany.com,github.com/my/private

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o main .


EXPOSE 8080

CMD ["./main"]