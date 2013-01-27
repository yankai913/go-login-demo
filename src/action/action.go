package action

import (
	"net/http"
	"dao"
	"log"
	"html/template"
	"model"
)

func StartUp() {
	log.Println("start...")
	Register()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("ListenAndServe", err.Error())
	}
}

var  userDao *dao.UserDaoImpl = new(dao.UserDaoImpl)

func Register() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/doLogin.do", doLoginHandler)
	http.HandleFunc("/register.do", registerHandler)
	http.HandleFunc("/doRegister.do", doRegisterHandler)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//index
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("indexHandler...")
	tmpl, err := template.ParseFiles("src/action/index.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

//doLogin
func doLoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	
	log.Print("loginHandler...  " + name + "  " + password)
	
	sql := "SELECT * FROM userinfo where name='" + name +"' and password ='" + password + "'"
	log.Print("sql : " + sql)
	
	objs := userDao.QueryForList(sql)
	tmpl, err := template.ParseFiles("src/action/logined.html")
	checkErr(err)
	
	if len(objs) == 0 {
		err = tmpl.Execute(w, map[string] string {"err_msg": "username or password is invalid!"})
		checkErr(err)
		return
	}
	
	var  user *model.User =  objs[0]
	err = tmpl.Execute(w, map[string] interface{} {"user": user})
	checkErr(err)
	
	
}

//register
func registerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("src/action/register.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

//doRegister
func doRegisterHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	address := r.FormValue("address")
	age := r.FormValue("age")
	
	sql := "INSERT INTO userinfo(NAME, PASSWORD, address, age) VALUES ('" + name + "' ,'" + password + "', '" + address + "'," + age + ")"
	userDao.Save(sql)
	tmpl, err := template.ParseFiles("src/action/registered.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}