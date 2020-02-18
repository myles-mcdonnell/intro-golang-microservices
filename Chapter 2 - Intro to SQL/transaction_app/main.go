package main

import (
	"context"
	"database/sql"
	"fmt"
	"g2r-api/db"
	"g2r-api/env"
	"github.com/google/uuid"
	"github.com/myles-mcdonnell/logrusx"
	"os"
	"sync"
	"time"
)

func main() {

	conf, err := env.Parse()
	if err != nil {
		panic(err)
	}

	logrusx.Init(
		conf.LogLevel,
		os.Stdout,
		&logrusx.JSONFormatter{Indent: conf.PrettyLogOutput},
		logrusx.NewLogEntryFactory("go_app"))

	sqlDb, err := db.BootstrapDB(conf.DbConfig)
	if err != nil {
		panic(err)
	}

	database := db.NewDatabase(sqlDb)

	counterId := uuid.New()
	fmt.Printf("Incrementing counter %v WITHOUT transaction\r\n", counterId)
	createAndIncrementCounterConcurrently(counterId, database, incrementCounterNoTransaction)

	counterId = uuid.New()
	fmt.Printf("Incrementing counter %v WITH transaction\r\n", counterId)
	createAndIncrementCounterConcurrently(counterId, database, incrementCounterInTransaction)
}

func createAndIncrementCounterConcurrently(counterId uuid.UUID, database *db.Database, incFunc func(id uuid.UUID, database *db.Database)) {

	database.Queries.InsertCounter(context.Background(), db.InsertCounterParams{
		ID:  counterId,
		Val: 0,
	})

	incrementCalledCount := 0
	wg := sync.WaitGroup{}
	var timeElapsed time.Duration
	start := time.Now()
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			incFunc(counterId, database)
			incrementCalledCount++
			wg.Done()
		}()
		timeElapsed = time.Now().Sub(start)
	}
	wg.Wait()

	val, err := database.Queries.SelectCounter(context.Background(), counterId)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Increment function called %v times for counter %v, the counter value is %v, execution time %v\r\n", incrementCalledCount, counterId, val, timeElapsed)
}

func incrementCounterNoTransaction(id uuid.UUID, database *db.Database) {

	incrementCounter(id, database.Queries)
}

func incrementCounterInTransaction(id uuid.UUID, database *db.Database) {

	err := database.WithDefaultTransaction(context.Background(), func(tx *sql.Tx, queries *db.Queries) error {

		tx.Exec("LOCK TABLE counter IN SHARE ROW EXCLUSIVE MODE;")

		incrementCounter(id, queries)

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func incrementCounter(id uuid.UUID, queries *db.Queries) {

	val, err := queries.SelectCounter(context.Background(), id)
	if err != nil {
		panic(err)
	}

	err = queries.UpdateCounter(context.Background(), db.UpdateCounterParams{
		Val: val + 1,
		ID:  id})
	if err != nil {
		panic(err)
	}
}
