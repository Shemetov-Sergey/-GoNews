package dbconfig

func ConnString() string {
	return ("host=172.22.0.2 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
}
