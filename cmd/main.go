package main

import (
	"myproject/internal/mysql"
	"myproject/internal/nginx"
)

func main() {
	mysql.MysqlErrors()
	nginx.NginxErrors()
}
