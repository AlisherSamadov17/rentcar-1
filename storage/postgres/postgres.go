package postgres

import (
	"context"
	"fmt"
	"rent-car/config"
	"rent-car/storage"
	"time"
	
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Store struct {
	Pool *pgxpool.Pool
}


func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	pgPoolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	pgPoolConfig.MaxConns = 100
	pgPoolConfig.MaxConnLifetime = time.Hour

	ctx,cancel:= context.WithTimeout(ctx,config.TimewithContex)
	defer cancel()
	
	newPool, err := pgxpool.NewWithConfig(ctx, pgPoolConfig)
	if err != nil {
		fmt.Println("error while connecting to db", err.Error())
		return nil, err
	}

	return Store{
		Pool: newPool,
	}, nil

}

func (s Store) CloseDB() {
	s.Pool.Close()
}

func (s Store) Car() storage.ICarStorage {
	newCar := NewCar(s.Pool)

	return &newCar
}

func (s Store) Customer() storage.ICustomerStorage {
	newCustomer := NewCustomer(s.Pool)

	return &newCustomer
}
func (s Store) Order() storage.IOrderStorage {
	newOrder := NewOrder(s.Pool)

	return &newOrder
}
