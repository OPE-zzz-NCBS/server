server: clean
	go-bindata -pkg="repo" -o mssql/repo/sql.go mssql/repo/sql
	go-bindata -pkg="sql_mssql" -o sql_mssql/sql_mssql.go sql_mssql
	go build .

clean:
	rm -f server
