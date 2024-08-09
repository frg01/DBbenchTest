package src

import (
	"DBbenchTest/src/database/pgvector"
	"DBbenchTest/src/utils"
	"fmt"
	"sync"
	"time"
)

type Runner struct {
	DatabaseURL string
}

func NewRunner() *Runner {
	r := &Runner{}
	r.init()
	return r
}

func (r *Runner) init() {

	r.DatabaseURL = r.getConfig()
}

func (r *Runner) getConfig() string {
	config := Config{GetUrl()}
	return config.databaseUrl
}

func (r *Runner) run() {

	//1. load entire datasets
	vectors := utils.ReadEmbeddingParquet()

	//connection database
	pg := new(pgvector.Pgvector)
	pg.init(r.DatabaseURL)
	defer pg.down()

	////2.insert into database
	//pg.insertData(vectors)
	//
	////3. build index
	//pg.createIndex()

	//4. multi query embedding
	r.mulConcurrency(pg, 1, vectors)

}

func (r *Runner) mulConcurrency(pg *pgvector.Pgvector, concurrency int, vectors []string) {

	// 定义一个等待组
	var wg sync.WaitGroup

	timeout := 30 * time.Second
	timer := time.NewTimer(timeout)

	// 记录已经执行的任务数量
	var counter int
	var failed int
	// 定义一个通道用于接收任务执行的时间
	results := make(chan time.Duration, concurrency)

	// 记录开始时间1
	start := time.Now()

	for i := 0; i < concurrency; i++ {
		index := i
		for {
			select {
			case <-timer.C:
				//如果定时器已触发，打印超时信息，并台跳出循环
				fmt.Println("duration=", timeout, "s search finished")
				return
			default:
				wg.Add(1) // 增加等待组计数
				go func() {
					defer wg.Done()                             // 减少等待组计数
					queryStart := time.Now()                    // 记录任务开始时间
					res, err := pg.singleSearch(vectors[index]) // 执行SQL查询
					queryDuration := time.Since(queryStart)     // 计算任务执行时间
					_ = res
					if err != nil {
						fmt.Println(err)
						failed += 1
					}
					results <- queryDuration // 将任务执行时间发送到通道中
					counter += 1
				}()
			}
		}
	}
	defer timer.Stop()
	//total durations
	totalDuration := time.Since(start)

	go func() {
		wg.Wait()      // 等待所有任务执行完毕
		close(results) // 关闭结果通道
	}()

	//qps
	qps := float64(counter) / float64(totalDuration)
	fmt.Printf("QPS: %.2f\n", qps)

	// 主goroutine 统计所有任务的执行时间
	var queryDurations []time.Duration
	for result := range results {
		queryDurations = append(queryDurations, result)
	}
}

// 计算 QPS 和 P99
//totalQueries := len(queryDurations) // 总执行任务数量
// 总执行时间
//qps := float64(totalQueries) / totalDuration.Seconds() // 每秒执行的任务数量
//fmt.Printf("QPS: %.2f\n", qps)            // 打印QPS
//p99 := utils.CalculateP99(queryDurations) // 计算P99延迟
//fmt.Printf("P99: %v\n", p99)              // 打印P99延迟
//
//fmt.Printf("counter: %v\n", counter)
