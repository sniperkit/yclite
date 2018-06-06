# yclite
hacker news lite web project using `fasthttp` and `fasthttprouter`

## install

`go get -u github.com/shohi/yclite`

## usage

```bash
# build docker
make docker-build

# run image
docker run -d -p 8080:8080 yclite:0.2

# digest HN
http://localhost:8080/list/1?points=100&filter=go
```

## developer

The project uses `golang` team's  `dep` to resolve dependencies. First, install `dep`

`go get -u github.com/golang/dep/cmd/dep`

After that, ensure dependencies and ready to go!

`dep ensure`