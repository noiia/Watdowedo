# Docker readme

## Build image 
docker build -t watdowedo:latest .

## Run container
docker run -d -p 8080:8080 --name watdowedo-container watdowedo:latest

## Stop container
docker stop watdowedo-container