FROM golang:1.19

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./
COPY .env.docker.example .env

RUN go build -o backend
RUN chmod +x ./backend

EXPOSE 9000

CMD [ "./backend" ]