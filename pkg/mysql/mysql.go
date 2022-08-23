package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Host string
	Port int32
	User string
	Pwd  string
	DB   string
}

type MysqlConn struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string

	db *sql.DB
}

// If you do not want to preselect a database, leave dbName empty
// thus, caller need select database before operation on table
func NewMysqlConn(host, port, user, pwd, dbName string, ext ...string) (*MysqlConn, error) {
	conn := &MysqlConn{
		Host:   host,
		Port:   port,
		User:   user,
		Pwd:    pwd,
		DbName: dbName,
	}

	charSet := "UTF8"
	if len(ext) > 0 {
		charSet = ext[0]
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, pwd, host, port, dbName, charSet)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	conn.db = db
	return conn, nil
}

func (conn *MysqlConn) Close() error {
	return conn.db.Close()
}

func (conn *MysqlConn) Query(query string, args ...interface{}) ([]map[string]string, error) {
	rows, err := conn.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return rowsToRecords(rows)
}

func (conn *MysqlConn) Exec(query string, args ...interface{}) (int64, int64, error) {
	res, err := conn.db.Exec(query, args...)
	if err != nil {
		return 0, 0, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, 0, err
	}
	return num, lastId, nil
}

func (conn *MysqlConn) NewTx() (*MysqlTx, error) {
	tx, err := conn.db.Begin()
	if err != nil {
		return nil, err
	}

	return &MysqlTx{Tx: tx}, nil
}

func (conn *MysqlConn) PrepareExec(query string, args ...interface{}) (int64, int64, error) {
	//prepare the statement
	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return 0, 0, err
	}
	//format all vals at once
	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, 0, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, 0, err
	}
	return num, lastId, nil
}

// transaction

type MysqlTx struct {
	Tx *sql.Tx
}

func (tx *MysqlTx) Exec(query string, args ...interface{}) (num int64, lastId int64, err error) {
	res, err := tx.Tx.Exec(query, args...)
	if err != nil {
		return 0, 0, err
	}

	num, err = res.RowsAffected()
	if err != nil {
		tx.Tx.Rollback()
		return 0, 0, err
	}

	lastId, err = res.LastInsertId()
	if err != nil {
		tx.Tx.Rollback()
		return 0, 0, err
	}

	return num, lastId, nil
}

func (tx *MysqlTx) Query(query string, args ...interface{}) ([]map[string]string, error) {
	rows, err := tx.Tx.Query(query, args...)
	if err != nil {
		tx.Tx.Rollback()
		return nil, err
	}

	defer rows.Close()
	return rowsToRecords(rows)
}

func (tx *MysqlTx) Commit() error {
	return tx.Tx.Commit()
}

func (tx *MysqlTx) RollBack() {
	tx.Tx.Rollback()
}

func (tx *MysqlTx) PrepareExec(query string, args ...interface{}) (int64, int64, error) {
	//prepare the statement
	stmt, err := tx.Tx.Prepare(query)
	if err != nil {
		return 0, 0, err
	}
	//format all vals at once
	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, 0, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, 0, err
	}
	return num, lastId, nil
}

// util

func rowsToRecords(rows *sql.Rows) ([]map[string]string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	records := []map[string]string{}
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}

		one := make(map[string]string)
		for i, col := range values {
			if col != nil {
				one[columns[i]] = string(col)
			}
		}
		records = append(records, one)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return records, nil
}
