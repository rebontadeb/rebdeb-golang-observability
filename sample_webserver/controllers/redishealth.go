package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"sample_webserver/selflogger"

	"github.com/redis/go-redis/v9"
)

var ctx = context.TODO()

func MyRedisClient() *redis.Client {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

type redisHealth struct {
	Status string
}

func RedisHealthz(w http.ResponseWriter, req *http.Request) {
	_, err := MyRedisClient().Ping(ctx).Result()
	if err != nil {
		healthstatus := redisHealth{Status: "Cannot Connect to Redis "}
		jsonResponse, jsonError := json.Marshal(healthstatus)
		if jsonError != nil {
			selflogger.ErrorLogger.Println("Unable to Marshall JSON")
		} else {
			selflogger.ErrorLogger.Println("Cannot Connect to Redis ", err)
			FailureResponseWriter(w, req, jsonResponse)
		}
	} else {
		healthstatus := redisHealth{Status: "Healthy "}
		jsonResponse, jsonError := json.Marshal(healthstatus)
		if jsonError != nil {
			selflogger.ErrorLogger.Println("Unable to Marshall JSON")
		} else {
			SuccessResponseWriter(w, req, jsonResponse)
			selflogger.InfoLogger.Println("Redis Healthz Endpoint Accessed")
		}
	}
}
