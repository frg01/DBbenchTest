package src

import (
	"database-benchTest/src/database/pgvector"
	"database-benchTest/src/utils"
)

var DatabaseURL string

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

func (r *Runner) Run() {

	//1. load entire datasets
	vectors := utils.ReadEmbeddingParquet()

	//connection database
	pg := pgvector.NewPgvector(r.DatabaseURL)
	defer pg.Down()

	//2.insert into database
	pg.InsertData(vectors)

	//3. build index
	pg.CreateIndex()

	//4. multi query embedding
	//pg.SingleSearch("[embedding]")

}
