# Docker readme

## Create a new builder instance that enables for multiple platforms
```bash
docker buildx create --name mybuilder --use
```
## Start it
```bash
docker buildx inspect --bootstrap
```

## Build image 
Choose your building version between: `linux/amd64,linux/arm64,windows/amd64`
```bash
docker buildx build --platform linux/amd64 -t watdowedo:latest --load .
```

## Check if it has been built 
```bash
docker manifest inspect
```

## Run container
```bash
docker run -d -p 8080:8080 --name watdowedo-container watdowedo:latest
```

## Stop container
```bash
docker stop watdowedo-container
```

## Delete a container
```bash
docker rm watdowedo-container
```