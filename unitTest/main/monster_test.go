package main

import "testing"
//单元测试测试文件
// go test -v monster_test.go monster.go


func TestStore(t *testing.T){
	monster := Monster{
		Name:"红孩儿",
		Age:10,
		Skill:"吐火",
	}
	res := monster.Store()
	if !res{
		t.Fatalf("保存序列化失败")
	}
	t.Logf("测试成功")
}

func TestReStore(t *testing.T){
	//先创建一个Monster 实例，不需要指定字段的值
	var monster Monster
	res :=monster.Restore()
	if !res{
		t.Fatalf("获取序列化失败")
	}

	//进一步判断
	if monster.Name != "红孩儿"{
		t.Fatalf("保存序列化失败")
	}

	t.Logf("测试成功"+monster.Name)
}