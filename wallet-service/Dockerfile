FROM golang

LABEL maintainer="Sai Ravi Teja Karanam <sairavitejaiitkgp@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o grabathon main.go

EXPOSE 4000

CMD ["./grabathon"]