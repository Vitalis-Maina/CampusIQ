package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Vitalis-Maina/internal/data"

	_ "github.com/lib/pq"
)

type Config struct {
	Port int
	db   struct {
		dns          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

// struct to handler our server depandancies
type application struct {
	config Config
	models data.UmsModel
}

func main() {
	var cfg Config

	flag.IntVar(&cfg.Port, "port", 9009, "api server port")
	flag.StringVar(&cfg.db.dns, "db-dns", os.Getenv("UMS_DB_DNS"), "Postgres DNS")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.Parse()

	db, err := connDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("starting server on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}

}

func connDB(cfg Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.db.dns)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxIdleTime(duration)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
