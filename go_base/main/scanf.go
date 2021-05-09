package main

import "fmt"

func main(){
	//var i int =5
	//fmt.Printf("%b\r\n",i)
	//
	////八进制： 0-7 满8进一，艺术字0开头
	//
	//var j int = 011 //八进制
	//fmt.Println("j=",j)
	//
	//var k int = 0x11 //16 bitsize
	//fmt.Println("k=",k)


//进制的转换

	//fmt.Println(2&3)
	//fmt.Println(2|3)
	//fmt.Println(2^3)
	//fmt.Println(-2^2 )
	//fmt.Println(1>>2 )
	//fmt.Println(1<<2 )
	//fmt.Println(-2<<2 )
	//fmt.Println(-2>>2 ) //补码形式运算


//var age int
//fmt.Println("请输入年龄：")
//fmt.Scanln(&age)
//
//if age > 18{ //不要括号
//	fmt.Println("你必须为自己行为负责")
//}

	//age :=20;
	//if (age > 18) { //不要括号
	//	fmt.Println("你必须为自己行为负责")
	//} else{
	//	fmt.Println("小伙子有前途")
	//}


	//age :=20;
	//if (age > 18) { //不要括号
	//	fmt.Println("你必须为自己行为负责")
	//} else{
	//	fmt.Println("小伙子有前途")
	//}

	//age := 1996
  //if (age % 4==0 && age %100 != 0 ) || (age %400 ==0 ){
  //	fmt.Println("你是闰年生的")
  //}


	//var a float64 = 2.0
	//var b float64 = 4.0
	//var c float64 = 2.0
	//
	//m := b*b -4*a*c
	////多分支判断
	//if m > 0{
	//	x1 := (-b + math.Sqrt(m)) / 2*a
	//	x2 := (-b - math.Sqrt(m)) / 2*a
	//	fmt.Printf("x1=%v x2=%v",x1,x2)
	//}else if m ==0 {
	//	x1 := (-b + math.Sqrt(m)) / 2*a
	//	fmt.Printf("x1=%v",x1)
	//}else {
	//	fmt.Println("无解。。。")
	//}

	//var n1 int32 = 5
	//var n2 int32 = 15
	//
	//switch n1{
	//case n2,n1,5,16,19:
	//	fmt.Println("ok")
	//}

//case中也可以对范围进行判断

//switch score := 90; {
//case score > 90:
//	fmt.Println("考的很好")
//case score >80 && score <= 90:
//	fmt.Println("考的不错")
//	fallthrough //加了穿透就会继续执行下一个default分支
//default:
//	fmt.Println("不及格")
//}

//for i := 1;i<=10;i++{
//	fmt.Println("nihaofwfmwefmweimwei",i,"vitas好帅")
//}

//if true{ //判断条件只能是bool值
//	fmt.Println("对了")
//}

//还可以类似while
//j := 1
//for j <= 10{
//fmt.Println("nihao1",j)
//j++
//}

//k := 1
//for ;;{
//	if k <= 10{
//		fmt.Println("vitas好帅",k)
//	}else{
//		break;
//	}
//	k++
//}
//

//遍历字符串
//var str string = "hello ,world!"
//for i := 0;i< len(str);i++{
//	fmt.Printf("%c \n", str[i]) //使用下标
//}


//for range遍历字符串
//for index,val := range str{
//	fmt.Printf("%d => %c \n",index,val)
//}

	//str := "vitas好帅呀"
	//str2 := []rune(str)
	//for i := 0;i< len(str2);i++{
	//	fmt.Printf("%c \n",str2[i])
	//}

	//range是按字符遍历的
	//for index,val := range str{
	//	fmt.Printf("%d => %c \n",index,val)
	//}

	//for i := 1;i<10;i++{
	//	fmt.Printf("%v + %v = 10 \n",i,10-i)
	//}


	//var i int = 1
	//for{
	//	fmt.Println("wefwefweqwe",i)
	//	i++
	//	if i > 10{
	//		break;
	//	}
	//}

//var count int = 0
//for {
//	rand.Seed(time.Now().UnixNano()) //纳秒做随机数种子
//	n := rand.Intn(100)+1
//	fmt.Println("n=",n)
//	count++
//	if(n==99){
//		break; //表示跳出for循环
//	}
//}
//
//fmt.Println("生成99 一共用了",count)

//u.Echo("亲爱的宝宝")
	//var res1 int
	//var res2 int
	//res1 , res2 = u.Test(1,2)

//	var res1 int
//	//var res2 int
//	res1 , _ = u.Test(1,2)
//	fmt.Println(res1)
//var num int
//fmt.Scanln(&num)
//fmt.Println(num)

//utils.Test2(10)

	//fmt.Println(utils.Fbn(10))
	//fmt.Println(utils.Hanshu(2))
	//fmt.Println(utils.Monkey(10))
	//fmt.Println(utils.Sum(4))


	a :=getSum
	fmt.Printf("a 的类型%T, getSum类型是%T \n",a,getSum)

	//res := a(10,40)

	res := myFun(getSum,10,40)  //将函数作为参数传
	fmt.Println(res)

	type myInt int //给int取别名

	var num myInt
	var num2 int
	num = 40
	num2 = int(num) //在go中还是认为是两个不同的类型，不等价
	fmt.Println(num,num2)

	res3 := myFun2(getSum,13,12)
	fmt.Println(res3)

	sum,sub := myFun3(4,2)
	fmt.Printf("sum=%v sub=%v",sum,sub)
}

func getSum(n1 int ,n2 int ) int {
	return n1+n2
}

func myFun(funvar func(int,int) int,num1 int,num2 int) int{
	return funvar(num1,num2)
}

type myFunType func(int ,int) int

func myFun2(funvar myFunType,num1 int,num2 int) int{
	return funvar(num1,num2)
}

func myFun3(n1 int ,n2 int) (sum int ,sub int){
	sum = n1 + n2
	sub = n1 - n2
	return
}