# Go Messenger

Simple client-server message interaction through UDP

## Build

```
go build .
```

## Run server

```
./go-messenger server
```

## Send message from client

```
./go-messenger client --p 8888 --m "This is a message!"
```

## Example using NetCat

nc -u 127.0.0.1 8888