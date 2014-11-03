server: clean
	go-bindata -pkg="repo" -o mssql/repo/sql.go mssql/repo/sql
	go build .

clean:
	rm -f server
