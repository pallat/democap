# Exercises

## HA Instructions

1. starting
  - `make up`
  - master/slave up
  - go run main.go (master & slave)
2. disconnect master & slave
  - `make disconnect`
  - wait until `Unable to connect to MASTER: Resource temporarily unavailable`
3. test
  - run only master
    - master set 222
  - run only slave to get new value
4. เพิ่ม C

```go
	all := 1
	intCmd := rdb.Wait(context.Background(), all, 3*time.Second)
	if int(intCmd.Val()) != all {
		slog.Error("not guarantee consistency")
	}
```

## MongoDB Instructions

1. docker-compose
1. dev container
2. go run main.go
3. disconnect mongo2
4. เพิ่ม connection option
```go
clientOpts := options.Client().ApplyURI(
		"mongodb://localhost:27017/?connect=direct").
		// "mongodb://localhost:27017,localhost:27018,localhost:27019/?replicaSet=rs0")
		SetWriteConcern(&writeconcern.WriteConcern{W: 3})
```
