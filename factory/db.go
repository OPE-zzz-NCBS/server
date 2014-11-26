package factory

import (
	"fmt"
	"path"
	"log"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/OPENCBS/server/config"
	"github.com/OPENCBS/server/repo"
	"github.com/OPENCBS/server/mssql"
)

var db *sql.DB

func getDb() *sql.DB {
	if db != nil {
		return db
	}

	var conf *config.Configuration
	conf, err := config.Get()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	template := "server=%s;user id=%s;password=%s;database=%s;connection timeout=5"
	connString := fmt.Sprintf(template,
		conf.Database.Host,
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Name)
	db, err = sql.Open("mssql", connString)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func getSql(name string) string {
	path := path.Join("mssql", name)
	sql, err := mssql.Asset(path)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(sql)
}

func NewUserRepo() *repo.UserRepo {
	repo := new(repo.UserRepo)
	repo.GetSql = getSql
	repo.Db = getDb()
	return repo
}

func NewClientRepo() *repo.ClientRepo {
	repo := new(repo.ClientRepo)
	repo.GetSql = getSql
	repo.Db = getDb()
	return repo
}

func NewPersonRepo() *repo.PersonRepo {
	repo := new(repo.PersonRepo)
	repo.GetSql = getSql
	repo.Db = getDb()
	return repo
}

func NewActivityRepo() *repo.ActivityRepo {
	repo := new(repo.ActivityRepo)
	repo.GetSql = getSql
	repo.Db = getDb()
	return repo
}

func NewBranchRepo() *repo.BranchRepo {
	repo := new(repo.BranchRepo)
	repo.GetSql = getSql
	repo.Db = getDb()
	return repo
}

func NewCityRepo() *repo.CityRepo {
	repo := new(repo.CityRepo)
	repo.GetSql = getSql
	repo.Db = getDb()
	return repo
}

func NewDistrictRepo() *repo.DistrictRepo {
	repo := new(repo.DistrictRepo)
	repo.GetSql = getSql
	repo.Db = getDb()
	return repo
}

