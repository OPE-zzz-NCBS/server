server: clean
	go-bindata -pkg="mssql" -o mssql/mssql.go mssql
	go build .

clean:
	rm -f server
