# gocron-example
Testing gocron

#### Run the server
```shell
go run server.go
```

#### Run these tests and monitor changes in logs
```shell
curl --location --request POST 'http://localhost:1234/update/3?timezone=Africa/Windhoek'

curl --location --request POST 'http://localhost:1234/update/3?timezone=Africa/Windhoek2'

```
