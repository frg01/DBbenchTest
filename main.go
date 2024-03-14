package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"time"
)

func main() {

	//获取运行参数，检查参数
	//args := os.Args[1:]
	//fmt.Printf("参数：%v,%v\n", args[0], args[1])
	//for i := 0; i < len(args); i++ {
	//	values := strings.Split(args[i], "=")
	//	fmt.Printf("\tkey:%s", values[0])
	//	fmt.Printf("\tvalue:%s", values[1])
	//}

	config, err := pgxpool.ParseConfig("postgres://bigmath:bigmath@192.168.50.60:5433/bigmath?sslmode=disable&pool_max_conns=20")
	if err != nil {
		fmt.Fprintf(os.Stderr, "config set wrong ,check input parameters!\\n")
		os.Exit(1)
	}

	////设置取消尝试时间
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//连接数据库
	pool, err := pgxpool.New(ctx, config.ConnString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database：%v\n", err)
		os.Exit(1)
	}

	stat := pool.Stat()
	fmt.Printf("连接池状态：%v", stat)

	//执行建表语句
	tag, err := pool.Exec(ctx, "create table if not exists go_test6 (id bigserial,embedding vector(3));")
	if err != nil {
		fmt.Fprintf(os.Stderr, "exec sql failed:%v\n", err)
	}
	_ = tag

	defer pool.Close()
}
