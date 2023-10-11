package main

import (
	"database/sql"
	"log"
	"net"

	db "github.com/jakoo13/simplebank/db/sqlc"
	"github.com/jakoo13/simplebank/gapi"
	"github.com/jakoo13/simplebank/pb"
	"github.com/jakoo13/simplebank/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	// Run db migrations
	runDBMigrations(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)
	runGrpcServer(config, store)

}

func runDBMigrations(migrationURL string, dbSource string){
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("Cannot create migration:", err)
	}

	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Cannot run migrate up:", err)
	}

	log.Println("DB migrated successfully")
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener", err)
	}

	log.Printf("gRPC server is listening on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}

// func runGinServer (config util.Config, store db.Store) {
// 	server,err := api.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal("Cannot create server", err)
// 	}

// 	err = server.Start(config.HTTPServerAddress)
// 	if err != nil {
// 		log.Fatal("Cannot start server", err)
// 	}
// }
