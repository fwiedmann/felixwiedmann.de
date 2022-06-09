package infrastructure_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fwiedmann/site/backend/internal/opinions/application"
	"github.com/fwiedmann/site/backend/internal/opinions/infrastructure"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestNewOpinionsRepositorySQLite_create_client_with_new_db_instance(t *testing.T) {
	t.Parallel()
	const testDBInstance = "testInstance.db"
	dbAbsolutePath := fmt.Sprintf("%s/%s", t.TempDir(), testDBInstance)

	_, err := infrastructure.NewOpinionsRepositorySQLite(dbAbsolutePath)
	if err != nil {
		t.Errorf("NewOpinionsRepositorySQLite() retunred error %s, but no error is expected", err)
	}

	_, err = os.Stat(dbAbsolutePath)
	if err != nil {
		t.Errorf("NewOpinionsRepositorySQLite() retunred error %s, but no error is expected", err)
	}
}

func TestNewOpinionsRepositorySQLite_create_client_with_exising_db_instance(t *testing.T) {
	t.Parallel()
	const testDBInstance = "testInstance.db"
	dbAbsolutePath := fmt.Sprintf("%s/%s", t.TempDir(), testDBInstance)

	_, err := sql.Open("sqlite3", dbAbsolutePath)
	if err != nil {
		t.Errorf("Could not create db for test: error %s", err)
	}

	_, err = infrastructure.NewOpinionsRepositorySQLite(dbAbsolutePath)
	if err != nil {
		t.Errorf("NewOpinionsRepositorySQLite() retunred error %s, but no error is expected", err)
	}

	_, err = os.Stat(dbAbsolutePath)
	if err != nil {
		t.Errorf("os.Stat() retunred error %s, but no error is expected", err)
	}
}

func TestNewOpinionsRepositorySQLite_error_invalid_path(t *testing.T) {
	t.Parallel()
	const testDBInstance = "testInstance.db"
	dbAbsolutePath := fmt.Sprintf("%s/does-not-exist/%s", t.TempDir(), testDBInstance)

	_, err := infrastructure.NewOpinionsRepositorySQLite(dbAbsolutePath)
	if err == nil {
		t.Errorf("NewOpinionsRepositorySQLite() retunred no error, but is expected")
	}
}

func TestOpinionsRepositorySQLite_CreateOpinion_successfully(t *testing.T) {
	t.Parallel()
	const testDBInstance = "testInstance.db"
	dbAbsolutePath := fmt.Sprintf("%s/%s", t.TempDir(), testDBInstance)

	repo, err := infrastructure.NewOpinionsRepositorySQLite(dbAbsolutePath)
	if err != nil {
		t.Errorf("NewOpinionsRepositorySQLite() retunred error %s, but no error is expected", err)
	}

	testTime := time.Now()
	var testUserId application.UserId = "123"
	var testId application.OpinionId = "1"
	testStatement := "copy and pasta is fine"

	err = repo.CreateOpinion(context.Background(), application.Opinion{
		ID:        testId,
		Owner:     testUserId,
		CreatedAt: testTime,
		Statement: testStatement,
	})

	if err != nil {
		t.Errorf("CreateOpinion() retunred error %s, but no error is expected", err)
	}

	db, err := sql.Open("sqlite3", dbAbsolutePath)

	row := db.QueryRow("SELECT * FROM opinions WHERE id = ?", testId)
	if row.Err() != nil {
		t.Errorf("could not exec query satement: %q", err)
	}

	var userId application.UserId
	var id application.OpinionId
	var date string
	var statement string

	err = row.Scan(&id, &userId, &date, &statement)
	if err != nil {
		t.Errorf("could not scan: %q", err)
	}

	assert.Equal(t, testId, id)
	assert.Equal(t, testUserId, userId)
	assert.Equal(t, testTime.Format(time.RFC3339), date)
	assert.Equal(t, testStatement, statement)
}

func TestOpinionsRepositorySQLite_ListOpinionsOpinion_successfully(t *testing.T) {
	t.Parallel()
	const testDBInstance = "testInstance.db"
	dbAbsolutePath := fmt.Sprintf("%s/%s", t.TempDir(), testDBInstance)

	repo, err := infrastructure.NewOpinionsRepositorySQLite(dbAbsolutePath)
	if err != nil {
		t.Errorf("NewOpinionsRepositorySQLite() retunred error %s, but no error is expected", err)
	}

	testTime := time.Now()
	var testUserId application.UserId = "123"
	var testId application.OpinionId = "1"
	testStatement := "copy and pasta is fine"

	db, err := sql.Open("sqlite3", dbAbsolutePath)

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Errorf("could not create transaction: %q", err)
	}

	_, err = tx.Exec("INSERT INTO  opinions (id, userId, creationTime, statement) VALUES (?, ?, ?, ?)", testId, testUserId, testTime.Format(time.RFC3339), testStatement)
	if err != nil {
		t.Errorf("could exec transaction: %q", err)
	}
	if err := tx.Commit(); err != nil {
		t.Errorf("could commit transaction: %q", err)

	}

	list, err := repo.ListOpinions(context.Background())
	if err != nil {
		t.Errorf("ListOpinions() returned error: %q", err)
	}

	assert.Len(t, list, 1)

	assert.Equal(t, testId, list[0].ID)
	assert.Equal(t, testUserId, list[0].Owner)
	assert.Equal(t, testTime.Format(time.RFC3339), list[0].CreatedAt.Format(time.RFC3339))
	assert.Equal(t, testStatement, list[0].Statement)
}

func TestOpinionsRepositorySQLite_DeleteOpinion(t *testing.T) {
	t.Parallel()
	const testDBInstance = "testInstance.db"
	dbAbsolutePath := fmt.Sprintf("%s/%s", t.TempDir(), testDBInstance)

	repo, err := infrastructure.NewOpinionsRepositorySQLite(dbAbsolutePath)
	if err != nil {
		t.Errorf("NewOpinionsRepositorySQLite() retunred error %s, but no error is expected", err)
	}

	testTime := time.Now()
	var testUserId application.UserId = "123"
	var testId application.OpinionId = "1"
	testStatement := "copy and pasta is fine"

	db, err := sql.Open("sqlite3", dbAbsolutePath)

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Errorf("could not create transaction: %q", err)
	}

	_, err = tx.Exec("INSERT INTO  opinions (id, userId, creationTime, statement) VALUES (?, ?, ?, ?)", testId, testUserId, testTime.Format(time.RFC3339), testStatement)
	if err != nil {
		t.Errorf("could exec transaction: %q", err)
	}
	if err := tx.Commit(); err != nil {
		t.Errorf("could commit transaction: %q", err)
	}

	if err := repo.DeleteOpinion(context.Background(), testId); err != nil {
		t.Errorf("could not delete opinion: %q", err)
	}

	row := db.QueryRow("SELECT * FROM opinions WHERE id = ?", testId)

	if row.Scan() == nil {
		t.Errorf("Scan() should return an error because opinion could not be found, but no error received")
	}
}
