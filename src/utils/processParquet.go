package utils

import (
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
	"log"
)

type Record struct {
	ID          int32 `parquet:"name=id, type=INT32"`
	NeighborsID int32 `parquet:"name=neighbors_id, type=LIST, of=INT32"`
}

func ReadEmbeddingParquet() []string { //bsonType parquet.BsonType

	///read
	fr, err := local.NewLocalFileReader("test.parquet")
	if err != nil {
		log.Println("Can't open file", err)
		return nil
	}
	pr, err := reader.NewParquetColumnReader(fr, 1)
	if err != nil {
		log.Println("Can't create column reader", err)
		return nil
	}
	num := int64(pr.GetNumRows())
	_ = num

	//
	//schemass := pr.SchemaHandler
	//namesd := schemass.GetRootInName()
	//_ = namesd
	ids, _, _, err := pr.ReadColumnByIndex(0, 1000)
	if err != nil {
		log.Println("err:")
	}

	var idList []int64
	for _, id := range ids {
		idList = append(idList, id.(int64))
	}

	neIds, _, _, err := pr.ReadColumnByIndex(1, 1000)
	if err != nil {
		log.Println("err:")
	}

	vectors := make([]string, 0)

	for i := 0; i < 1000; i++ {
		// 计算当前行的起始索引和结束索引
		start := i * 1536
		end := (i + 1) * 1536

		vector := make([]float64, 1536)

		for j := start; j < end; j++ {
			vector[j-start] = neIds[j].(float64)
		}

		res := ConvertVectorToString(vector)
		vectors = append(vectors, res)
	}

	//log.Println(vectors)

	//log.Println("===============key", neIds, "\n rls", rls, "\n dls", dls, err)

	//for range neIds {
	//
	//}

	pr.ReadStop()
	fr.Close()

	return vectors
	//// 读取数据并打印
	//for {
	//	record := Record{}
	//	if err := pr.Read(&record); err != nil {
	//		break // 读取完成或发生错误时退出循环
	//	}
	//	fmt.Printf("ID: %d, NeighborsID: %v\n", record.ID, record.NeighborsID)
	//}

}

func main() {
	ReadEmbeddingParquet()
}
