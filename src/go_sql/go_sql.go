package go_sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//! 本身DB 继承sql.DB 调用方法
type DB struct{
sql.DB
}

//小写仅用于本包
type Student struct {
	Id        int
	Name      string
	Dept_name string
	Tot_cred  int
}

type Descdata struct {
	Field     string
	Type      string
	NUll      string
	Key       string
	Default   sql.NullString
	Extra     string 
}


//initialize DB
func InitDB() (db *DB,err error) {
	// DSN:Data Source name
	dsn := "xherror_win:200430@tcp(192.168.0.191:3306)/learn?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	odb, err := sql.Open("mysql", dsn)
	db=&DB{*odb}  //!
	if err != nil {
		return db,err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return db,err
	}
	return db,nil
}

func (db *DB)DescTableDemo(){
	sqlStr:="DESC student;"
	var descdata Descdata
	rows,err:=db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		err := rows.Scan(&descdata.Field,&descdata.Type,&descdata.NUll,&descdata.Key,&descdata.Default,&descdata.Extra)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("student:%v\n", descdata)
	}
}

// 查询单条数据示例
func (db *DB)QueryRowDemo(name string) (student Student,err error) {
	sqlStr := "select Id,name,Dept_name,Tot_cred from student where name=?"
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err = db.QueryRow(sqlStr, name).Scan(&student.Id, &student.Name, &student.Dept_name, &student.Tot_cred)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return student, err
	}
	return student,nil
}

// 查询多条数据示例
func (db *DB)QueryMultiRowDemo() {
	sqlStr := "select * from student;"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.Id, &student.Name, &student.Dept_name, &student.Tot_cred)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("Id:%d\tname:%s\tDept_name:%s\tTot_cred:%d\n", student.Id, student.Name, student.Dept_name, student.Tot_cred)
	}
}

// 插入数据
func (db *DB)InsertRowDemo() {
	sqlStr := "insert into student values (?,?,?,?)"
	ret, err := db.Exec(sqlStr, 66666, "Sammy", "History", 100)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert Id failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the Id is %d.\n", theId)
}

// 预处理查询示例
func (db *DB)PrepareQueryDemo() {
	sqlStr := "select * from student where Id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.Id, &student.Name, &student.Dept_name, &student.Tot_cred)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("Id:%d\tname:%s\tDept_name:%s\tTot_cred:%d\n", student.Id, student.Name, student.Dept_name, student.Tot_cred)
	}
}

// sql注入示例
func (db *DB)SqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select * from student where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var student Student
	err := db.QueryRow(sqlStr).Scan(&student.Id, &student.Name, &student.Dept_name, &student.Tot_cred)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("Id:%d\tname:%s\tDept_name:%s\tTot_cred:%d\n", student.Id, student.Name, student.Dept_name, student.Tot_cred)
	//fmt.Printf("user:%#v\n", student)
}

// 事务操作示例
func (db *DB)TransactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update student set Dept_name='Physics' where Id=?;"
	ret1, err := tx.Exec(sqlStr1, 66666)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "Update student set Tot_cred=120 where Id=?"
	ret2, err := tx.Exec(sqlStr2, 66666)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}