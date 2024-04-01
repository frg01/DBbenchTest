# database-benchTest
bench test for postgresql ,support vector now 




#解决go引用包暴红的问题
https://blog.csdn.net/lizarel/article/details/108401578

parquet-go 存储测试集
Parquet：Parquet 是一种列式存储格式，通常用于大规模数据分析和数据仓库场景。它可以有效地压缩数据，提供了高性能的列式读取操作，并支持复杂数据结构和嵌套数据类型。Parquet 适合于需要进行高效数据分析和查询的场景，特别是大规模数据集的存储和处理。
go get github.com/xitongsys/parquet-go


连接postgresql的工具
go get github.com/jackc/pgx/v5
https://github.com/lib/pq.git


project has problem to fix, 
pgvector.go import database-benchTest/src and main import database-benchTest/src
PS D:\project\goProjecet\databaseBench\database-benchTest> go run .\main.go
package command-line-arguments
imports database-benchTest/src
imports database-benchTest/src/database/pgvector
imports database-benchTest/src: import cycle not allowed


关于qps的计算

请求成功和请求失败的记录
public double requestsPerSecondThroughput() {
return (double) measuredRequests / (double) nanoseconds * 1e9;
}

public double requestsPerSecondGoodput() {
return (double) success.getSampleCount() / (double) nanoseconds * 1e9;
}
