package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/user/go-exchange-rate/internal/domain"
	"github.com/user/go-exchange-rate/internal/vo"
)

type ExchangeRateRepository interface {
	Save(ctx context.Context, e *domain.CurrentExchangeRate) error
	Find(ctx context.Context, c1, c2 vo.Currency, from, to *time.Time) ([]domain.CurrentExchangeRate, error)
}

type PostgresRepository struct {
	Pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{Pool: pool}
}

func (r *PostgresRepository) Save(ctx context.Context, e *domain.CurrentExchangeRate) error {
	_, err := r.Pool.Exec(ctx,
		`INSERT INTO current_exchange_rate (id, currency1, currency2, rate, created_at) VALUES ($1,$2,$3,$4,$5)`,
		e.ID, string(e.Currency1), string(e.Currency2), e.Rate.String(), e.CreatedAt,
	)
	return err
}

func (r *PostgresRepository) Find(ctx context.Context, c1, c2 vo.Currency, from, to *time.Time) ([]domain.CurrentExchangeRate, error) {
	q := `SELECT id, currency1, currency2, rate, created_at FROM current_exchange_rate WHERE currency1=$1 AND currency2=$2`
	args := []any{string(c1), string(c2)}
	idx := 3
	if from != nil {
		q += fmt.Sprintf(" AND created_at >= $%d", idx)
		args = append(args, *from)
		idx++
	}
	if to != nil {
		q += fmt.Sprintf(" AND created_at < $%d", idx)
		args = append(args, *to)
		idx++
	}
	q += " ORDER BY created_at DESC"
	rows, err := r.Pool.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []domain.CurrentExchangeRate
	for rows.Next() {
		var id string
		var c1s, c2s, rate string
		var ca time.Time
		if err := rows.Scan(&id, &c1s, &c2s, &rate, &ca); err != nil {
			return nil, err
		}
		uid, _ := uuid.Parse(id)
		c1v, _ := vo.NewCurrency(c1s)
		c2v, _ := vo.NewCurrency(c2s)
		n, _ := vo.NewNumber(rate)
		res = append(res, domain.CurrentExchangeRate{
			ID:        uid,
			Currency1: c1v,
			Currency2: c2v,
			Rate:      n,
			CreatedAt: ca,
		})
	}
	return res, nil
}
