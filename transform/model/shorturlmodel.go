package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	shorturlFieldNames          = builderx.RawFieldNames(&Shorturl{})
	shorturlRows                = strings.Join(shorturlFieldNames, ",")
	shorturlRowsExpectAutoSet   = strings.Join(stringx.Remove(shorturlFieldNames, "`create_time`", "`update_time`"), ",")
	shorturlRowsWithPlaceHolder = strings.Join(stringx.Remove(shorturlFieldNames, "`shorten`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	ShorturlModel interface {
		Insert(data Shorturl) (sql.Result, error)
		FindOne(shorten string) (*Shorturl, error)
		FindByUrl(url string) (*Shorturl, error)
		Update(data Shorturl) error
		Delete(shorten string) error
	}

	defaultShorturlModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Shorturl struct {
		Shorten string `db:"shorten"` // shorten key
		Url     string `db:"url"`     // original url
	}
)

func NewShorturlModel(conn sqlx.SqlConn) ShorturlModel {
	return &defaultShorturlModel{
		conn:  conn,
		table: "`shorturl`",
	}
}

func (m *defaultShorturlModel) Insert(data Shorturl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, shorturlRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Shorten, data.Url)
	return ret, err
}

func (m *defaultShorturlModel) FindOne(shorten string) (*Shorturl, error) {
	query := fmt.Sprintf("select %s from %s where `shorten` = ? limit 1", shorturlRows, m.table)
	var resp Shorturl
	err := m.conn.QueryRow(&resp, query, shorten)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShorturlModel) FindByUrl(url string) (*Shorturl, error) {
	query := fmt.Sprintf("select %s from %s where `url` = ? limit 1", shorturlRows, m.table)
	var resp Shorturl
	err := m.conn.QueryRow(&resp, query, url)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShorturlModel) Update(data Shorturl) error {
	query := fmt.Sprintf("update %s set %s where `shorten` = ?", m.table, shorturlRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Url, data.Shorten)
	return err
}

func (m *defaultShorturlModel) Delete(shorten string) error {
	query := fmt.Sprintf("delete from %s where `shorten` = ?", m.table)
	_, err := m.conn.Exec(query, shorten)
	return err
}
