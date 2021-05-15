package processes

import "fmt"

//UserMgr实例在服务器有且只有一个，使用的地方很多，因此定义为全局变量

var (
	userMgr *UserMgr
)

type UserMgr struct{
	onlineUsers map[int]*UserProcess
}

//完成对userMgr初始化工作
func init(){
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess,1024),
	}
}

//完成对onlineUsers添加,编辑
func (this *UserMgr) AddOnlineUser(up *UserProcess){
	this.onlineUsers[up.UserId] = up //可以添加也可以实现覆盖功能
}

//删除
func (this *UserMgr) DelOnlineUser(userId int){
	delete(this.onlineUsers,userId)
}

//返回当前所有在线的用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess{
	return this.onlineUsers
}

//返回指定userId的用户
func (this *UserMgr) GetUserById(userId int) (up *UserProcess, err error){
	//如何从map中取出一个值，带检测方式
	up,ok := this.onlineUsers[userId]
	if !ok { //说明，你要查找的这个用户，当前不再线
		err = fmt.Errorf("用户%d 不存在",userId)
		return
	}
	return
}