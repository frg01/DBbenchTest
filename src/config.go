package src

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	databaseUrl string
}

func GetUrl() string {
	// --help usages of tips
	help := flag.Bool("help", false, "显示帮助信息")
	if *help {
		fmt.Printf("这是一个帮助信息")
		os.Exit(0)
	}

	//defined command line params
	var (
		ip       = "192.168.50.60"
		port     = 5433
		username = "bigmath"
		pwd      = "bigmath"
		database = "bigmath"
		maxConn  = 1
	)

	// 解析命令行参数
	flag.StringVar(&ip, "ip", "", "连接地址")

	flag.IntVar(&port, "port", 5433, "端口")
	flag.StringVar(&username, "username", "", "用户名")
	flag.StringVar(&pwd, "password", "", "密码")
	flag.StringVar(&database, "bigmath", "", "数据库")
	flag.IntVar(&maxConn, "max_conn", 1, "最大连接数")

	flag.Parse()

	// 打印解析结果
	fmt.Printf("用户名: %s\n", username)

	//databaseUrl := "postgres://" + username + ":" + pwd + "@" + ip + ":" + string(port) + "/" + database + "?sslmode=disable&pool_max_conns=" + string(maxConn)
	databaseUrl := "postgres://bigmath:bigmath@192.168.50.60:5433/bigmath?sslmode=disable&pool_max_conns=100"
	fmt.Printf("url=", databaseUrl)

	return databaseUrl

}
