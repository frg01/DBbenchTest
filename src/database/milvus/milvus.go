package milvus

import (
	"context"
	"github.com/milvus-io/milvus-proto/go-api/v2/client"
	"github.com/milvus-io/milvus-proto/go-api/v2/entity"
	"log"

	"time"
)

type Milvus struct {

	//MilvusAddr = `localhost:19530`
	//collectionName = 'test'
	//dim = 128
	//idCol,embeddingCol = "ID","embeddings"
}

func NewPgvector(databaseUrl string) *Milvus {
	milvus := &Milvus{}
	milvus.init(databaseUrl)
	return milvus
}

func (m *Milvus) init(databaseUrl string) {

	collectionName := "test"
	milvusAddr := "localhost:19530"
	idCol := ""
	embeddingCol := ""
	var dim int64 = 128

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	c, err := client.NewClient(ctx, client.Config{
		Address: milvusAddr,
	})
	if err != nil {
		// handling error and exit, to make example simple here
		log.Fatal("failed to connect to milvus:", err.Error())
	}

	// first, lets check the collection exists
	collExists, err := c.HasCollection(ctx, collectionName)
	if err != nil {
		log.Fatal("failed to check collection exists:", err.Error())
	}
	if collExists {
		// let's say the example collection is only for sampling the API
		// drop old one in case early crash or something
		_ = c.DropCollection(ctx, collectionName)
	}

	// define collection schema
	schema := entity.NewSchema().WithName(collectionName).WithDescription("this is the basic example collection").
		// currently primary key field is compulsory, and only int64 is allowed
		WithField(entity.NewField().WithName(idCol).WithDataType(entity.FieldTypeInt64).WithIsPrimaryKey(true).WithIsAutoID(false)).
		// also the vector field is needed
		WithField(entity.NewField().WithName(embeddingCol).WithDataType(entity.FieldTypeFloatVector).WithDim(dim))

	err = c.CreateCollection(ctx, schema, entity.DefaultShardNumber)
	if err != nil {
		log.Fatal("failed to create collection:", err.Error())
	}

	//collections, err := c.ListCollections(ctx)
	//if err != nil {
	//	log.Fatal("failed to list collections:", err.Error())
	//}
	//for _, collection := range collections {
	//	// print all the collections, id & name
	//	log.Printf("Collection id: %d, name: %s\n", collection.ID, collection.Name)
	//
	//}
	//
	//// show collection partitions
	//partitions, err := c.ShowPartitions(ctx, collectionName)
	//if err != nil {
	//	log.Fatal("failed to show partitions:", err.Error())
	//}
	//for _, partition := range partitions {
	//	// print partition info, the shall be a default partition for out collection
	//	log.Printf("partition id: %d, name: %s\n", partition.ID, partition.Name)
	//}
	//
	//partitionName := "new_partition"
	//// now let's try to create a partition
	//err = c.CreatePartition(ctx, collectionName, partitionName)
	//if err != nil {
	//	log.Fatal("failed to create partition:", err.Error())
	//}
	//
	//log.Println("After create partition")
	//// show collection partitions, check creation
	//partitions, err = c.ShowPartitions(ctx, collectionName)
	//if err != nil {
	//	log.Fatal("failed to show partitions:", err.Error())
	//}
	//for _, partition := range partitions {
	//	log.Printf("partition id: %d, name: %s\n", partition.ID, partition.Name)
	//}

	// clean up our mess
	_ = c.DropCollection(ctx, collectionName)
	c.Close()
}

func (m *Milvus) InsertData(vectors []string) {

}

func (m *Milvus) insertDataMutil() {

}

func (m *Milvus) CreateIndex() {

}

func (m *Milvus) SingleSearch(embedding string) {

}

func (m *Milvus) Down() {

}

func (m *Milvus) setPool(databaseURL string) { //, context.Context, context.CancelFunc

}
