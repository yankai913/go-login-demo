package model

import (
	"time"
)

/**
CREATE TABLE `userinfo` (
  `id` INT(10) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(64) DEFAULT NULL,
  `password` VARCHAR(64) DEFAULT NULL,
  `createtime` DATE DEFAULT NULL,
  `address` VARCHAR(64) DEFAULT NULL,
  `age` INT(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8

INSERT INTO userinfo(NAME, PASSWORD, address, age) VALUES ('admin' ,'admin', 'china',25)
**/

type User struct {
	Id int
	Name string
	Password string
	CreateTime time.Time
	Address string
	Age int
}

func (user *User) ToString() (str string) {
	str = "name:" + user.Name
	return str
}