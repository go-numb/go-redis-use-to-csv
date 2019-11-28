# go-redis-use

go-redis-use is redis get keys to csv.

## Usage
1. keys <redis_key>*
2. get <keys...>
3. to_csv(keys_data...)

```
$ go run main.go -key <redis_key>
// output to <./redis_keys_***.csv>
```



## TODO
- [ ] LIST
- [ ] HGETALL