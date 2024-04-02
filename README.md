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

milvus thread pool
https://golang.0voice.com/?id=388#:~:text=Milvus%E6%98%AF%E4%B8%80%E4%B8%AA%E5%BC%80%E6%BA%90%E7%9A%84%E5%90%91%E9%87%8F%E6%95%B0%E6%8D%AE%E5%BA%93%EF%BC%8C%E6%94%AF%E6%8C%81%E9%AB%98%E7%BB%B4%E5%90%91%E9%87%8F%E7%9B%B8%E4%BC%BC%E5%BA%A6%E6%90%9C%E7%B4%A2%E3%80%82,%E5%9C%A8%E4%BD%BF%E7%94%A8Milvus%E6%97%B6%EF%BC%8C%E5%8F%AF%E4%BB%A5%E4%BD%BF%E7%94%A8%E8%BF%9E%E6%8E%A5%E6%B1%A0%E6%9D%A5%E7%AE%A1%E7%90%86%E8%BF%9E%E6%8E%A5%EF%BC%8C%E6%8F%90%E9%AB%98%E6%80%A7%E8%83%BD%E5%92%8C%E6%95%88%E7%8E%87%E3%80%82%20%E5%9C%A8Go%E8%AF%AD%E8%A8%80%E4%B8%AD%EF%BC%8C%E5%8F%AF%E4%BB%A5%E4%BD%BF%E7%94%A8%E7%AC%AC%E4%B8%89%E6%96%B9%E5%BA%93%E5%A6%82go-redis%E6%88%96%E8%80%85%E8%87%AA%E5%B7%B1%E5%AE%9E%E7%8E%B0%E8%BF%9E%E6%8E%A5%E6%B1%A0%E6%9D%A5%E7%AE%A1%E7%90%86Milvus%E7%9A%84%E8%BF%9E%E6%8E%A5%E3%80%82