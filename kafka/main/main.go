package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
	"time"
)

func main() {

	//生产者
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	//client, err := sarama.NewSyncProducer([]string{"192.168.1.8:9092"}, config)
	client, err := sarama.NewSyncProducer([]string{"10.10.13.224:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()
	for {
		//sendMsg(client,"t_image_text")
		//sendMsg(client,"t_audio")
		sendMsg(client,"t_video_2")
		//sendMsg(client,"alive_data_change")
		//sendMsg(client,"t_ebook")
		//sendMsg(client,"t_pay_products")
		//sendMsg(client,"t_term")
		//sendMsg(client,"t_interaction")
		//sendMsg(client,"t_activity")
		//sendMsg(client,"t_clock_work")
		//sendMsg(client,"t_community")
		//sendMsg(client,"t_course_offline")
		//sendMsg(client,"t_examination")
		//sendMsg(client,"t_exercises")
		//sendMsg(client,"t_forms")
		//sendMsg(client,"t_practice")
		//sendMsg(client,"t_que_question")
	}
}


func sendMsg(client sarama.SyncProducer,topic_name string){
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic_name

	//id := getId()
	//phone订阅 +I
	str := `{"data":{"app_id":"appwpor7xql8808","id":"i_607e728660b211dfe4ab1e32","name":"务必马到成功2256","img_url":"http://wechatapppro-1252524126.file.myqcloud.com/appwpor7xql8808/image/kr329l4907xw.jpg","img_url_compressed_larger":"http://wechatapppro-1252524126.file.myqcloud.com/appwpor7xql8808/image/coms/640480808kr329l4907xw.jpg","img_url_compressed":"http://wechatapppro-1252524126.file.myqcloud.com/appwpor7xql8808/image/compress/414320457kr329l4907xw.jpg","summary":"app优化专栏","sell_type":1,"price":100,"line_price":100,"is_stop_sell":0,"recycle_bin_state":0,"invite_poster":"","distribute_default":0,"created_at":"2021-09-03 17:56:54","updated_at":"2022-01-10 13:04:17","can_select":1,"is_chosen":0,"is_discount":0,"is_distribute":0,"is_show_resourcecount":1,"visible_on":1,"wx_share":null,"resource_statistic":null,"show_in_menu":1,"have_password":0,"resource_password":"","loop_resource":0,"sell_mode":0,"auth_type":0,"relation_type":0},"op":"+U"}`
	//union订阅 +U
	//str := `{"data":{"id":500001,"app_id":"apppcHqlTPT3482","content_app_id":"","resource_id":"i_5f51a0a360b266ca4e9b9269","user_id":"u_5f55092de60b4_WYNlXe8eKe","universal_union_id":"ozStBt1BVbbkIJBfsjUV5sfIJ-1I","status":1,"resource_type":1,"created_at":"2020-09-07 18:35:43","updated_at":"2020-09-07 18:35:43"},"op":"+U"}`

	msg.Value = sarama.StringEncoder(str)
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	time.Sleep(5 * time.Second)
}




func getId() (id string){
	data := strconv.Itoa(int(time.Now().Unix()))
	id = data[len(data)-6:]
	return
}