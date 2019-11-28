package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var (
	key    string
	isSort bool
)

func init() {
	flag.StringVar(&key, "key", "", "set redis key, gets <key:*>, to csv")
	flag.BoolVar(&isSort, "sort", false, "set redis key, gets <key:*>, to csv")

	flag.Parse()
}

func main() {
	if key == "" {
		logrus.Fatal("find key is nil")
	}

	fmt.Printf("gets %s from redis-server", key)

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err := client.Ping().Err(); err != nil {
		logrus.Fatal(err)
	}

	// data := client.Get(key).String()
	// fmt.Println("this", data)

	keys, err := client.Keys(fmt.Sprintf("%s*", key)).Result()
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("keys print: %+v\n", keys)
	if len(keys) <= 0 {
		logrus.Fatal("keys length is less 0")
	}

	if isSort {
		sort.Strings(keys)
	}

	f, err := os.OpenFile(fmt.Sprintf("redis-keys-%s.csv", key), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		logrus.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	for _, v := range keys {
		row, _ := client.Get(v).Result()
		w.Write(strings.Split(row, ","))
	}

	w.Flush()
}
