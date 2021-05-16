package dao

import (
	"context"
	"database/sql"
	"go-learning-server/internal/model"
	"go-learning-server/internal/service"
	"go-learning-server/pkg/config"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Dao dao.
type Dao struct {
	db *sql.DB
}

func newMySQL(c *config.MySQLConfig) (*sql.DB, error) {
	cfg, err := mysql.ParseDSN(c.DSN)
	if err != nil {
		return nil, err
	}

	if cfg.Params == nil {
		cfg.Params = make(map[string]string)
	}

	cfg.Loc = time.UTC
	cfg.ParseTime = true
	cfg.InterpolateParams = true
	cfg.Params["time_zone"] = "'+00:00'"

	conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// New new a dao and return.
func NewDao() (dao *Dao) {
	conf := config.Get()
	db, err := newMySQL(&conf.MySQL)
	if err != nil {
		panic(err)
	}
	return &Dao{db: db}
}

// Close close the resource.
func (d *Dao) Close() {
	d.db.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(context.Context) (err error) {
	return d.db.Ping()
}

func (d *Dao) GetUser(ctx context.Context, r *service.GetUserReq) (*service.GetUserRes, error) {
	// DO to PO
	var m model.UserModel
	err := d.db.QueryRowContext(ctx, "select id, name from user where id = ?;", r.ID).Scan(&m)
	if err != nil {
		return nil, err
	}

	return &service.GetUserRes{
		Name: m.Name,
	}, nil
}
