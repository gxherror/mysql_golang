module main

go 1.18

require github.com/go-sql-driver/mysql v1.6.0 // indirect

require go_sql v1.0.0

replace go_sql => ../go_sql

require gee v1.0.0

replace gee => ../gee

require my_utils v1.0.0

replace my_utils => ../my_utils
