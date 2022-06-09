package infrastructure

import (
	"context"
	"database/sql"
	"github.com/fwiedmann/site/backend/internal/opinions/application"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func NewOpinionsRepositorySQLite(dbLocation string) (*OpinionsRepositorySQLite, error) {
	db, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		return &OpinionsRepositorySQLite{}, err
	}

	if err := db.Ping(); err != nil {
		return &OpinionsRepositorySQLite{}, err
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("CREATE TABLE IF NOT EXISTS opinions ( id varchar(255) NOT NULL , userId varchar(255) NOT NULL, creationTime varchar(255) NOT NULL, statement varchar(255) NOT NULL, PRIMARY KEY (id))")
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("CREATE INDEX opinion_id on opinions (id);")
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &OpinionsRepositorySQLite{
		db: db,
	}, nil
}

type OpinionsRepositorySQLite struct {
	db *sql.DB
}

func (o *OpinionsRepositorySQLite) CreateOpinion(ctx context.Context, opinion application.Opinion) error {
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO  opinions (id, userId, creationTime, statement) VALUES (?, ?, ?, ?)", opinion.ID, opinion.Owner, opinion.CreatedAt.Format(time.RFC3339), opinion.Statement)
	if err != nil {
		return err
	}

	return tx.Commit()
}
func (o *OpinionsRepositorySQLite) ListOpinions(ctx context.Context) ([]application.Opinion, error) {
	rows, err := o.db.QueryContext(ctx, "SELECT * FROM opinions")
	if err != nil {
		return nil, err
	}

	opinions := make([]application.Opinion, 0)

	for rows.Next() {

		var id application.OpinionId
		var userId application.UserId
		var date string
		var statement string

		err = rows.Scan(&id, &userId, &date, &statement)
		if err != nil {
			return nil, err
		}

		parsedTime, err := time.Parse(time.RFC3339, date)
		if err != nil {
			return nil, err
		}

		opinions = append(opinions, application.Opinion{
			ID:        id,
			Owner:     userId,
			CreatedAt: parsedTime,
			Statement: statement,
		})
	}

	return opinions, rows.Err()
}

func (o *OpinionsRepositorySQLite) DeleteOpinion(ctx context.Context, id application.OpinionId) error {
	tx, err := o.db.BeginTx(ctx, nil)

	_, err = tx.Exec("DELETE FROM opinions WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (o *OpinionsRepositorySQLite) CreateVote(ctx context.Context, vote application.Vote) error {
	//TODO implement me
	panic("implement me")
}

func (o *OpinionsRepositorySQLite) ListVotes(ctx context.Context) ([]application.Vote, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OpinionsRepositorySQLite) UpdateVote(ctx context.Context, vote application.Vote) error {
	//TODO implement me
	panic("implement me")
}

func (o *OpinionsRepositorySQLite) DeleteVote(ctx context.Context, id application.OpinionId) error {
	//TODO implement me
	panic("implement me")
}
