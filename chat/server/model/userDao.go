package model
import(
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个UserDao的结构体,完成对user结构体的各种操作
type UserDao struct{
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao){
	userDao = &UserDao{
		pool:pool,
	}
	return
}

//思考一下UserDao应该提供
//根据userid，返回一个user实例+err
func (this *UserDao) getUserById(conn redis.Conn,id int) (user *User,err error){
	//通过给定id去redis查询这个用户
	res,err := redis.String(conn.Do("HGet","users",id))
	if err != nil{
		//错误
		if err == redis.ErrNil{ //表示在users哈希中，没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	//这里需要把res反序列化成User实例
	err = json.Unmarshal([]byte(res),user)
	if err != nil{
		fmt.Println("json.Unmarshal err=",err) //json反序列化出错
		return
	}

	return
}

//完成登录的校验 Login
// 1. Login完成对用户的验证
//2. 如果用户的id和pwd都正确，则返回一个user实例
//3.如果用户的id或pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int,userPwd string) (user *User,err error){
	//先从UserDao的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user,err = this.getUserById(conn,userId)
	if err != nil{
		return
	}
	//证明用户获取到了
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}


