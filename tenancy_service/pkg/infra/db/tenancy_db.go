package tenancy_db

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/LoughlinSpollen/tenancy_service/pkg/domain/model"
	"github.com/LoughlinSpollen/tenancy_service/pkg/infra/env"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	fs       = flag.NewFlagSet("database", flag.ExitOnError)
	host     = fs.String("postgres-host", env.WithDefaultString("PGHOST", "localhost"), "postgres connection string")
	port     = fs.String("postgres-port", env.WithDefaultString("PGPORT", "5432"), "postgres port")
	dbname   = fs.String("postgres-db", env.WithDefaultString("PRIVATUS_DB_NAME", "privatus"), "postgres database name")
	user     = fs.String("postgres-user", env.WithDefaultString("PRIVATUS_DB_USER", "privatus_user"), "postgres user")
	password = fs.String("postgres-password", env.WithDefaultString("PRIVATUS_DB_PASSWORD", ""), "postgres user password")
	ssl      = fs.String("postgres-ssl", env.WithDefaultString("PGSSLMODE", "disable"), "postgres ssl mode")
	debug    = fs.String("postgres-debug", env.WithDefaultString("POSTGRES_DEBUG", "true"), "postgres debug mode")
)

const (
	maxConnections int = 20
)

type tenancyDB struct {
	pool *pgxpool.Pool
}

func NewTenancyDB() *tenancyDB {
	return &tenancyDB{}
}

func (db *tenancyDB) Connect() error {
	log.Debug("Database connect")

	conf, err := pgxpool.ParseConfig(db.getConnectionString())
	if err != nil {
		log.Error(fmt.Printf("Unable to parse database connection string: %v \n", err))
		return err
	}

	// tell the driver to sanitize statement strings
	conf.ConnConfig.RuntimeParams = map[string]string{
		"standard_conforming_strings": "on",
		"timezone":                    "UTC",
	}
	// no statement preparation call, whole queries within a single network call
	conf.ConnConfig.PreferSimpleProtocol = true
	if *debug == "true" {
		conf.ConnConfig.LogLevel = pgx.LogLevelDebug
	} else {
		conf.ConnConfig.LogLevel = pgx.LogLevelError
	}

	ctx := context.Background()
	db.pool, err = pgxpool.ConnectConfig(ctx, conf)
	if err != nil {
		log.Error(fmt.Printf("Unable to create database connection pool: %v \n", err))
	}
	// reconnect is handled by the driver
	return err
}

func (db *tenancyDB) Close() {
	log.Debug("Database Close")
	db.pool.Close()
}

func (db *tenancyDB) Save(tenancy *model.Tenancy) error {
	log.Debug("TenancyDB Save")

	ctx := context.Background()

	tx, err := db.pool.Begin(ctx)
	if err != nil {
		log.Error(fmt.Printf("datatabase error. Failed transaction begin: %v \n", err))
		return err
	}

	var sqlStatementTenancy = `
	INSERT INTO ` + *dbname + `.tenancy (id, ml_model)
	VALUES (uuid_generate_v4(), $1)
	RETURNING id
	`
	err = tx.QueryRow(ctx,
		sqlStatementTenancy,
		tenancy.MlModel,
	).Scan(&tenancy.ID)
	if err != nil {
		log.Error(fmt.Printf("datatabase error. Failed to insert tenancy: %v \n", err))
		return err
	}

	if tenancy.Federation != nil {
		tenancy.Federation.TenancyID = tenancy.ID
		var sqlStatementFederation = `
		INSERT INTO ` + *dbname + `.federation
			(id,
			tenancy_id,
			threshold, 
			epochs,
			rate,
			rounds,
			batch)
		VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6)
		RETURNING id
		`
		err = tx.QueryRow(ctx,
			sqlStatementFederation,
			tenancy.ID,
			tenancy.Federation.Threshold,
			tenancy.Federation.Epochs,
			tenancy.Federation.Rate,
			tenancy.Federation.Rounds,
			tenancy.Federation.Batch,
		).Scan(&tenancy.Federation.ID)
		if err != nil {
			log.Error(fmt.Printf("datatabase error. Failed to insert federation: %v \n", err))
			if err2 := tx.Rollback(ctx); err2 != nil {
				log.Error(fmt.Printf("datatabase error. Failed to rollback transaction: %v \n", err2))
			}
			return err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		log.Error(fmt.Printf("datatabase error. Failed tenancy save transaction commit: %v \n", err))
	}
	return err
}

func (db *tenancyDB) AddFederation(federation *model.Federation) error {
	log.Debug("TenancyDB AddFederation")

	ctx := context.Background()
	var sqlStatementFederation = `
	INSERT INTO ` + *dbname + `.federation
		(id,
		tenancy_id,
		threshold, 
		epochs,
		rate,
		rounds,
		batch)
	VALUES (uuid_generate_v4(), $1)
	RETURNING id
	`
	err := db.pool.QueryRow(ctx,
		sqlStatementFederation,
		federation.TenancyID,
		federation.Threshold,
		federation.Epochs,
		federation.Rate,
		federation.Rounds,
		federation.Batch,
	).Scan(&federation.ID)
	if err != nil {
		log.Error(fmt.Printf("datatabase error. Failed to insert federation: %v \n", err))
	}
	return err
}

func (db *tenancyDB) UpdateFederation(federation *model.Federation) error {
	log.Debug("TenancyDB UpdateFederation")

	ctx := context.Background()
	var sqlStatementFederation = `
	UPDATE ` + *dbname + `.federation SET
		threshold = $1,
		epochs = $2,
		rate = $3,
		rounds = $4,
		batch = $5
		WHERE id = $6
		`
	_, err := db.pool.Exec(ctx,
		sqlStatementFederation,
		federation.Threshold,
		federation.Epochs,
		federation.Rate,
		federation.Rounds,
		federation.Batch,
		federation.ID,
	)
	if err == pgx.ErrNoRows {
		log.Error(fmt.Printf("datatabase error. Failed to find federation: %v \n", err))
		return errors.New("not found")
	} else if err != nil {
		log.Error(fmt.Printf("datatabase error. Failed to update federation: %v \n", err))
	}
	return err
}

func (db *tenancyDB) AddTraining(training *model.Training) error {
	log.Debug("TenancyDB AddTraining")

	ctx := context.Background()
	var sqlStatementTraining = `
	INSERT INTO ` + *dbname + `.training
		(id,
		tenancy_id)
	VALUES (uuid_generate_v4(), $1)
	RETURNING id
	`
	err := db.pool.QueryRow(ctx,
		sqlStatementTraining,
		training.TenancyID,
	).Scan(&training.ID)
	if err != nil {
		log.Error(fmt.Printf("datatabase error. Failed to insert training: %v \n", err))
	}
	return err
}

func (db *tenancyDB) UpdateTraining(training *model.Training) error {
	log.Debug("TenancyDB UpdateTraining")

	ctx := context.Background()
	var sqlStatementTraining = `
	UPDATE ` + *dbname + `.training SET
		WHERE id = $4
		`
	_, err := db.pool.Exec(ctx,
		sqlStatementTraining,
	)
	if err == pgx.ErrNoRows {
		log.Error(fmt.Printf("datatabase error. Failed to find training: %v \n", err))
		return errors.New("not found")
	} else if err != nil {
		log.Error(fmt.Printf("datatabase error. Failed to update training: %v \n", err))
	}
	return err
}

func (db *tenancyDB) ReadFederation(federationID uuid.UUID) (*model.Federation, error) {
	log.Debug("TenancyDB UpdateTraining")

	ctx := context.Background()
	var sqlStatementFederation = `
	SELECT id, tenancy_id, threshold, epochs, rate, rounds, batch FROM ` + *dbname + `.federation WHERE id=$1;`
	row := db.pool.QueryRow(ctx, sqlStatementFederation, federationID)
	var federation model.Federation
	switch err := row.Scan(&federation.ID, &federation.TenancyID, &federation.Threshold, &federation.Epochs,
		&federation.Rate, &federation.Rounds, &federation.Batch); err {
	case pgx.ErrNoRows:
		log.Error(fmt.Printf("No federation dao returned from database for id: %v", federationID))
		return nil, errors.New("not found")
	case nil:
		return &federation, nil
	default:
		log.Error(fmt.Printf("datatabase error. Failed to retrieve federation dao for id %v: %v", federationID, err))
		return nil, err
	}
}

func (db *tenancyDB) ReadTenancy(tenancyID uuid.UUID) (*model.Tenancy, error) {
	log.Debug("TenancyDB ReadTenancy")

	ctx := context.Background()
	var sqlStatementTenancy = `SELECT
		tenancy.id, tenancy.ml_model,
		federation.id, federation.tenancy_id, federation.threshold, federation.epochs, 
		federation.rate, federation.rounds, federation.batch 
		FROM ` + *dbname + `.tenancy
		INNER JOIN ` + *dbname + `.federation ON tenancy.id = federation.tenancy_id
		WHERE tenancy.id=$1;`

	row := db.pool.QueryRow(ctx, sqlStatementTenancy, tenancyID)
	var tenancy model.Tenancy
	var federation model.Federation
	tenancy.Federation = &federation
	switch err := row.Scan(&tenancy.ID, &tenancy.MlModel, &federation.ID, &federation.TenancyID,
		&federation.Threshold, &federation.Epochs, &federation.Rate, &federation.Rounds, &federation.Batch); err {
	case pgx.ErrNoRows:
		log.Error(fmt.Printf("No tenancy dao returned from database for id: %v", tenancy.ID))
		return nil, errors.New("not found")
	case nil:
		return &tenancy, nil
	default:
		log.Error(fmt.Printf("datatabase error. Failed to retrieve tenancy dao for id %v: %v", tenancyID, err))
		return nil, err
	}
}

func (db *tenancyDB) getConnectionString() string {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&pool_max_conns=%d", *user, *password, *host, *port, *dbname, *ssl, maxConnections)
	return connString
}
