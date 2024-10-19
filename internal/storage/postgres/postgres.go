package postgres

import (
	"context"
	"fmt"

	config "github.com/Jofich/URL-Shoritng-Service/internal/config"
	"github.com/Jofich/URL-Shoritng-Service/internal/storage"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	db *pgx.Conn
}

func New(storageConfig config.StorageCfg) (*Storage, error) {
	const fn = "storage.postgres.New"

	connUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		storageConfig.Login, storageConfig.Password, storageConfig.Host, storageConfig.Port, storageConfig.DB_name)

	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &Storage{db: conn}, nil
}

func (s *Storage) SaveUrl(URLtoSave string, alias string) error {

	const fn = "storage.postgres.SaveUrl"

	query := `INSERT INTO urls(alias, url) VALUES (@alias, @url)`
	args := pgx.NamedArgs{
		"alias": alias,
		"url":   URLtoSave,
	}
	_, err := s.db.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetUrl(alias string) (string, error) {

	query := `SELECT * FROM urls WHERE alias = (@alias)`
	args := pgx.NamedArgs{
		"alias": alias,
	}
	var id int
	var url string

	err := s.db.QueryRow(context.Background(), query, args).Scan(&id, &alias, &url)
	fmt.Print(err)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return "", storage.ErrURLNotExists
		}
		return "", storage.ErrFailedToGetUrl
	}
	return url, nil
}
