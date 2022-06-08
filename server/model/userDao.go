package model

import (
	"encoding/json"
	"exercise/chatroom/common/message"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (u *UserDao) GetUserById(conn redis.Conn, id int) (user *User, err error) {
	str, err := redis.String(conn.Do("Hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(str), user)
	if err != nil {
		fmt.Println("json.Unmashal err =", err)
		return
	}

	return
}

func (u *UserDao) Login(id int, pwd string) (user *User, err error) {
	conn := u.pool.Get()
	defer conn.Close()

	user, err = u.GetUserById(conn, id)
	if err != nil {
		return
	}

	if pwd != user.UserPwd {
		err = ERROR_USER_PWD
	}

	return
}

func (u *UserDao) Register(user *message.User) (err error) {
	conn := u.pool.Get()
	defer conn.Close()

	_, err = u.GetUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("marshal err", err)
		return
	}
	_, err = conn.Do("hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("存入用户数据出错...")
	}
	return
}
