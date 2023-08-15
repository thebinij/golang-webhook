# golang-webhook
### Build Docker compose
``` shell
docker-compose up -d --build
```

### Stop Docker Compose

``` shell
docker-compose down      
```

### Deloy to Docker hub
```shell
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t thebinij/go-webhook:latest -t thebinij/go-webhook:1.0  --push .

```