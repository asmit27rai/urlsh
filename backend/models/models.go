package models

import (
	"fmt"
    "context"
    "github.com/go-redis/redis/v8"
    "math/rand"
    "time"
	"os"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	redisHost := os.Getenv("REDIS_HOST")
    if redisHost == "" {
        redisHost = "localhost"
    }
    redisPort := os.Getenv("REDIS_PORT")
    if redisPort == "" {
        redisPort = "6379"
    }

    rdb = redis.NewClient(&redis.Options{
        Addr:     redisHost + ":" + redisPort,
        Password: "",
        DB:       0,
    })
}

type URLRequest struct {
    LongURL string `json:"long_url"`
}

type URL struct {
    ShortCode string `json:"short_code"`
    LongURL   string `json:"long_url"`
    Clicks    int    `json:"clicks"`
}

func GenerateShortCode() string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, 6)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}

func SaveURL(shortCode, longURL string) {
    rdb.HSet(ctx, "urls", shortCode, longURL)
    rdb.HSet(ctx, "clicks", shortCode, 0)
}

func GetURL(shortCode string) (string, error) {
    return rdb.HGet(ctx, "urls", shortCode).Result()
}

func IncrementClicks(shortCode string) {
    rdb.HIncrBy(ctx, "clicks", shortCode, 1)
}

func GetAllURLs() ([]URL, error) {
    urls, err := rdb.HGetAll(ctx, "urls").Result()
    if err != nil {
        return nil, err
    }
    clicks, err := rdb.HGetAll(ctx, "clicks").Result()
    if err != nil {
        return nil, err
    }

    var result []URL
    for shortCode, longURL := range urls {
        result = append(result, URL{
            ShortCode: shortCode,
            LongURL:   longURL,
            Clicks:    parseInt(clicks[shortCode]),
        })
    }
    return result, nil
}

func parseInt(s string) int {
    var n int
    fmt.Sscanf(s, "%d", &n)
    return n
}