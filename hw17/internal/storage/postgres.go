package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/sirupsen/logrus"

	"github.com/dark705/otus/hw17/internal/calendar/event"
	"github.com/jmoiron/sqlx"
)

type PostgresConfig struct {
	HostPort       string
	User           string
	Pass           string
	Database       string
	TimeoutConnect int
	TimeoutExecute int
}

type Postgres struct {
	logger  *logrus.Logger
	db      *sqlx.DB
	ctxExec context.Context
}

func NewPG(conf PostgresConfig, logger *logrus.Logger) (pg Postgres, err error) {
	pg = Postgres{logger: logger}
	logger.Infoln("Start connect to Postgres")
	ctxConnect, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(conf.TimeoutConnect))
	pg.db, err = sqlx.ConnectContext(ctxConnect, "pgx", fmt.Sprintf("Postgres://%s:%s@%s/%s", conf.User, conf.Pass, conf.HostPort, conf.Database))
	if err != nil {
		return pg, err
	}
	pg.ctxExec, _ = context.WithCancel(context.Background())
	logger.Infoln("Success connected to Postgres")

	return pg, err
}

func (s *Postgres) Shutdown() {
	s.logger.Infoln("Close Postgres connection...")
	err := s.db.Close()
	if err != nil {
		s.logger.Infoln("Fail to close Postgres connection.")
	}
	s.logger.Infoln("Success close Postgres connection.")
}

func (s *Postgres) Add(e event.Event) (err error) {
	sql := "INSERT INTO events (start_time, end_time, title, description) VALUES (:start_time, :end_time, :title, :description);" // :start_time from `db:"start_time"` and so on
	_, err = s.db.NamedExecContext(s.ctxExec, sql, e)
	return err
}

func (s *Postgres) Del(id int) (err error) {
	sql := "DELETE FROM events WHERE id = :id;"
	res, err := s.db.NamedExecContext(s.ctxExec, sql, struct {
		Id int `db:"id"`
	}{Id: id})
	if err != nil {
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff != 1 {
		return errors.New(fmt.Sprintf("Fail to del, no event with %d:", id))
	}
	return nil
}

func (s *Postgres) Get(id int) (e event.Event, err error) {
	sql := "SELECT * FROM events WHERE id = :id;" // :id from `db:"id"`
	rows, err := s.db.NamedQueryContext(s.ctxExec, sql, struct {
		Id int `db:"id"`
	}{Id: id})
	if err != nil {
		return e, err
	}

	rows.Next()
	if err := rows.StructScan(&e); err != nil {
		return e, err
	}

	return e, err
}

func (s *Postgres) GetAll() (events []event.Event, err error) {
	sql := "SELECT * FROM events;" // :id from `db:"id"`
	rows, err := s.db.QueryxContext(s.ctxExec, sql)
	if err != nil {
		return events, err
	}

	for rows.Next() {
		var e event.Event
		if err = rows.StructScan(&e); err != nil {
			return events, err
		}
		events = append(events, e)
	}

	return events, err
}

func (s *Postgres) GetAllNotScheduled() (events []event.Event, err error) {
	sql := "SELECT * FROM events WHERE is_scheduled is false;"
	rows, err := s.db.QueryxContext(s.ctxExec, sql)
	if err != nil {
		return events, err
	}

	for rows.Next() {
		var e event.Event
		if err = rows.StructScan(&e); err != nil {
			return events, err
		}
		events = append(events, e)
	}

	return events, err
}

func (s *Postgres) Edit(editEvent event.Event) (err error) {
	// :start_time from `db:"start_time"` and so on
	sql := "UPDATE events SET (start_time, end_time, title, description, is_scheduled) = (:start_time, :end_time, :title, :description, :is_scheduled) WHERE id = :id;"
	res, err := s.db.NamedExecContext(s.ctxExec, sql, editEvent)
	if err != nil {
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff != 1 {
		return errors.New(fmt.Sprintf("Fail to edit, no event with %d:", editEvent.Id))
	}

	return nil
}

func (s *Postgres) IntervalIsBusy(checkedEvent event.Event, isNewEvent bool) (exist bool, err error) {
	var rows *sqlx.Rows
	//SELECT * FROM events WHERE start_time < 'checkedEvent_end_time' and end_time > 'checkedEvent_start_time and id != checkedEvent_id'
	//if add new event id = 0, in PG serial start from 1
	sql := "SELECT true FROM events WHERE start_time < :end_time AND end_time > :start_time and id != :id;"
	rows, err = s.db.NamedQueryContext(s.ctxExec, sql, checkedEvent)

	if err != nil {
		return exist, err
	}
	exist = rows.Next()

	return exist, err
}
