package pgvector

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

type Pgvector struct {
	pool   *pgxpool.Pool
	ctx    context.Context
	cancel context.CancelFunc
}

func NewPgvector(databaseUrl string) *Pgvector {
	pgVec := &Pgvector{}
	pgVec.init(databaseUrl)
	return pgVec
}

func (p *Pgvector) init(databaseUrl string) {
	p.pool = p.setPool(databaseUrl)
}

func (p *Pgvector) InsertData(vectors []string) {

	for _, value := range vectors {
		_, err := p.pool.Exec(context.Background(), "insert into items (embedding) values ($1);", value)
		if err != nil {
			log.Print(err)
		}
	}
}

func insertDataMutil() {
}

func (p *Pgvector) CreateIndex() {

	_, err := p.pool.Exec(context.Background(), "CREATE INDEX ON items USING spann (embedding vector_cosine_ops) WITH (machine=3,threads=16,assign=1);")
	if err != nil {
		log.Print(err)
	}
}

func (p *Pgvector) SingleSearch(embedding string) (pgx.Rows, error) {

	res, err := p.pool.Query(nil, "select id,embedding from items order by embedding <-> $1", embedding)
	if err != nil {
		log.Print(err)
	}

	return res, err
}

func (p *Pgvector) Down() {
	p.pool.Close()
}

func (p *Pgvector) setPool(databaseURL string) *pgxpool.Pool { //, context.Context, context.CancelFunc

	fmt.Println("databaseURL:", databaseURL)
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "config set wrong ,check input parameters!\\n")
		os.Exit(1)
	}

	////设置取消尝试时间
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//连接数据库
	pool, err := pgxpool.New(context.Background(), config.ConnString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database：%v\n", err)
		return nil
	}

	stat := pool.Stat()
	fmt.Printf("连接池状态：%v", stat)
	return pool
}
