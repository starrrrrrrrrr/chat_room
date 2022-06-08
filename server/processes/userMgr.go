package processes

import "fmt"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	OnlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		OnlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (u *UserMgr) AddOnlineUser(up *UserProcess) {
	u.OnlineUsers[up.UserId] = up
}

func (u *UserMgr) DelOnlineUser(userId int) {
	delete(u.OnlineUsers, userId)
}

func (u *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return u.OnlineUsers
}

func (u *UserMgr) GetOnlineUser(userId int) (up *UserProcess, err error) {
	up, ok := u.OnlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}

	return
}
