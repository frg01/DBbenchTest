package utils

import (
	"fmt"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
)

//type ParquetFile interface {
//	io.Seeker
//	io.Reader
//	io.Writer
//	io.Closer
//	Open(name string) (ParquetFile, error)
//	Create(name string) (ParquetFile, error)
//}

func ReadParquet(bsonType parquet.BsonType) []float32 {
	data := "../result/neighbors.parquet"
	fr, err := local.NewLocalFileReader(data)
	if err != nil {
		fmt.Printf("err:%v", err)
	}

	pr,err := reader.NewParquetReader(fr, nil, nil)
	if err != nil {
		return nil
	}

	u := make([*user])



}
