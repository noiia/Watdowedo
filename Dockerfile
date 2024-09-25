## build tailwind dependencies
FROM node:20 AS build

WORKDIR /usr/src/Watdowedo

COPY package*.json ./

RUN npm install -D tailwindcss

RUN npx tailwindcss init

## run process
## compile tailwind and golang then launch the app
FROM node:20

RUN npx tailwindcss -i ./web/static/css/tailwind.css -o ./web/static/css/output.css 

FROM golang:1.22.4

WORKDIR /usr/src/Watdowedo

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/src/Watdowedo ./...

EXPOSE 8080

CMD ["./watdowedo"]