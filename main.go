package main

import (
	"database-benchTest/src"
	"sort"
	"time"
)

type Record struct {
	ID          float32 `parquet:"name=id, type=FLOAT"`
	NeighborsID float32 `parquet:"name=neighbors_id, type=FLOAT"`
}

func main() {

	runner := src.NewRunner()
	runner.Run()

	//concurrency := 1000          // 并发数，即同时执行的goroutine数量
	//duration := 10 * time.Second // 运行时长，即测试运行的总时长
	//
	//_ = duration // 忽略duration，暂未使用
	//
	//// 获取连接池、上下文和取消函数
	//pool, ctx, cancel := src.SetPool()
	//
	//// 在函数结束时调用取消函数和关闭连接池
	//defer cancel()
	//defer pool.Close()
	//
	//// 定义一个等待组
	//var wg sync.WaitGroup
	//
	//// 记录开始时间
	//start := time.Now()
	//
	//// 记录已经执行的任务数量
	//var counter int
	//
	//// 定义一个通道用于接收任务执行的时间
	//results := make(chan time.Duration, concurrency)
	//
	//// 循环启动并发数量的goroutine执行任务
	//for i := 0; i < concurrency; i++ {
	//	wg.Add(1) // 增加等待组计数
	//	go func() {
	//		defer wg.Done()                                                                             // 减少等待组计数
	//		queryStart := time.Now()                                                                    // 记录任务开始时间
	//		_, err := pool.Exec(ctx, "SELECT * FROM items order by embedding <-> $1 limit 10;", res[i]) // 执行SQL查询
	//		if err != nil {
	//			fmt.Fprintf(os.Stderr, "exec sql failed: %v\n", err) // 打印错误信息
	//			return
	//		}
	//		queryDuration := time.Since(queryStart) // 计算任务执行时间
	//		results <- queryDuration                // 将任务执行时间发送到通道中
	//		counter += 1
	//	}()
	//}
	//
	//// 启动一个goroutine等待所有任务执行完毕
	//go func() {
	//	wg.Wait()      // 等待所有任务执行完毕
	//	close(results) // 关闭结果通道
	//}()
	//
	//// 主goroutine 统计所有任务的执行时间
	//var queryDurations []time.Duration
	//for result := range results {
	//	queryDurations = append(queryDurations, result)
	//}
	//
	//// 计算 QPS 和 P99
	//totalQueries := len(queryDurations)                    // 总执行任务数量
	//totalDuration := time.Since(start)                     // 总执行时间
	//qps := float64(totalQueries) / totalDuration.Seconds() // 每秒执行的任务数量
	//fmt.Printf("QPS: %.2f\n", qps)                         // 打印QPS
	//p99 := calculateP99(queryDurations)                    // 计算P99延迟
	//fmt.Printf("P99: %v\n", p99)                           // 打印P99延迟
	//
	//fmt.Printf("counter: %v\n", counter)

}

func calculateP99(durations []time.Duration) time.Duration {
	if len(durations) == 0 {
		return 0
	}

	var totalDuration time.Duration
	for _, duration := range durations {
		totalDuration += duration
	}
	averageDuration := totalDuration / time.Duration(len(durations))
	_ = averageDuration

	var p99Index = int(float64(len(durations)) * 0.99)
	var sortedDurations = append([]time.Duration(nil), durations...)
	sort.Slice(sortedDurations, func(i, j int) bool { return sortedDurations[i] < sortedDurations[j] })

	return sortedDurations[p99Index]
}
