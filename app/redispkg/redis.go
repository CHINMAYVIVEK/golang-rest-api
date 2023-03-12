package redispkg

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"go-docker/helper"
	"net/http"
)

type RedisPageResponse struct {
	Value   string `json:"value"`
	Message string `json:"message"`
}

var ctx = context.Background()

func setData(rdb *redis.Client, key, value string) error {

	//  set data to redis for 5 sec
	helper.SugarObj.Info(key, " == ", value)
	err := rdb.Set(ctx, key, value, 5*time.Second).Err()
	helper.SugarObj.Info("Setting value to Redis")
	if err != nil {
		// panic(err)
		helper.SugarObj.Error(err)
		return err
	}
	return nil
}

func fetchData(rdb *redis.Client) (string, error) {

	//  fetch Data from redis
	key := "key"
	value := "value"

	err := setData(rdb, key, value)

	if err != nil {
		helper.SugarObj.Error(err)
		return "", err
	}

	val, err := rdb.Get(ctx, "key").Result()

	if err != nil {
		helper.SugarObj.Error(err)
		return "", err
	}
	helper.SugarObj.Info(val)
	return val, nil

}

func RedisEntrypoint(responseWriter http.ResponseWriter, request *http.Request) {

	helper.SugarObj.Info(request)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := fetchData(rdb)

	if err != nil {
		helper.SugarObj.Error(err)
		val = "No value"
	}
	helper.SugarObj.Info(val)

	resp := &RedisPageResponse{
		Value:   val,
		Message: err.Error(),
	}

	helper.SugarObj.Info(resp)
	fmt.Println(" ==== RedisEntrypoint ==== ", request)

	responseWriter.Header().Set("Content-Type", "application/json")

	json.NewEncoder(responseWriter).Encode(resp)

}
