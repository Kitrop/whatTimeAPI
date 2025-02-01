package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Start redis server
	rdb := StartRedisServer()

  r := chi.NewRouter()

  // Middleware for logging and revive after panic
  // r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

    // 404 Not Found
  r.NotFound(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(404)
    w.Write([]byte("route does not exist"))
  })

  r.Get("/weekdate/{date}", func(w http.ResponseWriter, r *http.Request) {
		GetWeekDate(w, r, rdb)
	})

  http.ListenAndServe(os.Getenv("PORT"), r)
}

// Start redis server
func StartRedisServer() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0,
	})

	return rdb
}

// Checking and counting day of week
func GetWeekDate(w http.ResponseWriter, r *http.Request, rdb *redis.Client) {
  date := chi.URLParam(r, "date")

	// Get date week from redis
	val, _ := rdb.Get(ctx, date).Result()

	if val != "" {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(val))
		return
	}

  // Check format date is correct
	dayWeek, error := isValidDate(date)

  if error != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("invalid date format"))
    return
  }
	
	// Set date week to redis
	rdb.Set(ctx, date, dayWeek, 0)

  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte(dayWeek))
}

// Check is correct format date
func isValidDate(date string) (string, error) {
	// Correct formats
  formats := []string{
    "02.01.2006",
    "02-01-2006",
    "02:01:2006",
  }

  // Parcing date
  for _, format := range formats {
    if day, err := time.Parse(format, date); err == nil {
      return day.Weekday().String(), nil
    }
  }

  return "", errors.New("invalid date format")
}