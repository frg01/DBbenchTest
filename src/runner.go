package src

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"time"
)

var DatabaseURL string

type Runner struct {
}

func init() {

	DatabaseURL = getConfig()
}

func getConfig() string {
	config := Config{GetUrl()}
	return config.databaseUrl
}

func SetPool() (*pgxpool.Pool, context.Context, context.CancelFunc) {

	fmt.Println("DatabaseURL:", DatabaseURL)
	config, err := pgxpool.ParseConfig(DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "config set wrong ,check input parameters!\\n")
		os.Exit(1)
	}

	////设置取消尝试时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//连接数据库
	pool, err := pgxpool.New(ctx, config.ConnString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database：%v\n", err)
		os.Exit(1)
	}

	stat := pool.Stat()
	fmt.Printf("连接池状态：%v", stat)

	return pool, ctx, cancel
}

func Run() {

}
