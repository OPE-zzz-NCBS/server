package mssql

import (
	"path"
)

func GetSql(name string) (string, error) {
	path := path.Join("mssql", name)
	sql, err := Asset(path)
	if err != nil {
		return "", err
	}
	return string(sql), nil
}

