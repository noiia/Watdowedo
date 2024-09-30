# Docker readme

## Build image 
```bash
docker build -t watdowedo:latest .
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