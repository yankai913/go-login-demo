package dao

import (
	"github.com/ziutek/mymysql/mysql"
	"os"
	"fmt"
	 _ "github.com/ziutek/mymysql/native" // Native engine
//     _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
	
)


func InitDB() (con mysql.Conn){
	addr := "127.0.0.1:3306"
	user := "root"
	passwd := ""
	dbname := "kaige"
	//db := mysql.New(proto, "", addr, user, pass, dbname)
	con = CreateCon(addr , user , passwd , dbname, )
	if con == nil {
		panic("con is nil!")
	}
	con.Query("set names utf8")
	return
}

func PrintOK() {
	fmt.Println("OK")
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CheckedResult(rows []mysql.Row, res mysql.Result, err error) ([]mysql.Row,
	mysql.Result) {
	CheckError(err)
	return rows, res
}

func CreateCon(addr string, user string, passwd string, dbname string) (con mysql.Conn){
	proto := "tcp"
	con = mysql.New(proto, "", addr, user, passwd, dbname,)
	fmt.Printf("Connect to %s:%s... ", proto, addr)
	CheckError(con.Connect())
	PrintOK()
	return 
}

type Dao interface {
	// QueryForList(sql string) (res []interface{})
	// QueryForOne(sql string) (res interface{})
	 Save(sql string)
	 Update(sql string)
	 Remove(sql string)
}


