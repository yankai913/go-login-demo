package dao

import 
( 
	"model"
	"log"
)

type UserDao interface {
	QueryForList(sql string) (objs []*model.User)
	
	Dao
}

type UserDaoImpl struct {

}

func (ud *UserDaoImpl) QueryForList(sql string) (objs []*model.User) {
	db := InitDB()
	defer db.Close()
	
	rows, res := CheckedResult(db.Query(sql))
	id := res.Map("id")
	name := res.Map("name")
	password := res.Map("password")
	createtime := res.Map("createtime")
	address := res.Map("address")
	age := res.Map("age")
	
	var length int = len(rows)
	log.Println("rows length:" , length)
	log.Println("objs length:" , len(objs))
	
	objs = make([]*model.User, length)
	
	for i, row := range rows {
		log.Println("i=", i)
		objs[i] = &model.User{
						row.Int(id), 
						row.Str(name), 
						row.Str(password), 
						row.Localtime(createtime),
						row.Str(address), 
						row.Int(age)}
						
		log.Print(objs[i].CreateTime)
	}
	return objs
}

func (ud *UserDaoImpl) Save(sql string) {
	db := InitDB()
	defer db.Close()
	
	CheckedResult(db.Query(sql))
}

func (ud *UserDaoImpl) Update(sql string) {

}

func (ud *UserDaoImpl) Remove(sql string) {

}

var  userDao *UserDaoImpl = new(UserDaoImpl)

func Test_select() {
	log.Println("select start")
	objs := userDao.QueryForList("SELECT * FROM userinfo")
	for ii, obj := range objs {
		log.Println("row : " , ii, " value : ", obj.ToString())
	}
	log.Println("select end")	
}

func Test_save() {
	log.Println("save start")
	sql := "INSERT INTO userinfo(NAME, PASSWORD, address, age) VALUES ('小黑' ,'admin', '中国',25)"
	userDao.Save(sql)
	log.Println("save end")	
}