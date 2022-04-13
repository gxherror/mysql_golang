module myproject

go 1.18

require github.com/go-sql-driver/mysql v1.6.0

require github.com/gxherror/my_utils v1.0.0

replace github.com/gxherror/my_utils v1.0.0 => ../local/my_utils
