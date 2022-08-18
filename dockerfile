FROM golang:1.18

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./
COPY .env.example .env

RUN go mod download

COPY . ./

RUN go mod tidy

RUN go build -o ./backend
RUN chmod +x ./backend

EXPOSE 9000

CMD [ "./backend" ]