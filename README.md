
##### Folder Structure:
```
cmd/api -> for main.go
api -> for yaml spec files
internal -> for most of the code
```

##### Commands:
```
go mod init github.com/rajpgr8/goapi
go mod tidy
go run cmd/api/main.go

Test:
curl --location 'http://localhost:8080/account/coins?username=alex' \
--header 'Authorization: 123ABC'


docker build -t goapi . --no-cache   --progress plain
docker -p 8081:8081 run goapi 

```

##### CGO_ENABLED 
```
https://medium.com/@pengcheng1222/exploring-cgo-enabled-in-go-23cf5cf2fe88#:~:text=What%20is%20CGO_ENABLED%20%3F,calling%20C%20code%20from%20Go.
```
