## build tailwind dependencies
FROM --platform=$BUILDPLATFORM node:latest AS build

WORKDIR /usr/src/Watdowedo

COPY package*.json ./

RUN npm install -D tailwindcss && npx tailwindcss init

## run process
## compile tailwind and golang then launch the app
FROM --platform=$BUILDPLATFORM node:latest

RUN npx tailwindcss -i ./web/static/css/tailwind.css -o ./web/static/css/output.css 

FROM --platform=$BUILDPLATFORM golang:1.22.4

RUN apt-get update && apt-get install -y

WORKDIR /usr/src/Watdowedo

COPY go.mod go.sum ./

RUN go mod tidy && go mod verify

COPY . .

RUN go build -v -o /usr/src/Watdowedo ./...

EXPOSE 8080