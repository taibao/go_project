package main

import (
	"fmt"
	"go_project/github.com/Shopify/sarama"
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
	client, err := sarama.NewSyncProducer([]string{"10.10.42.114:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()
	for {
		sendMsg(client, "app_apm_server")
	}
}

func sendMsg(client sarama.SyncProducer, topic_name string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic_name

	//id := getId()
	//phone订阅 +I
	//str := `{"data":{"app_id":"appwpor7xql8808","id":"i_607e728660b211dfe4ab1e32","name":"务必马到成功2256","img_url":"http://wechatapppro-1252524126.file.myqcloud.com/appwpor7xql8808/image/kr329l4907xw.jpg","img_url_compressed_larger":"http://wechatapppro-1252524126.file.myqcloud.com/appwpor7xql8808/image/coms/640480808kr329l4907xw.jpg","img_url_compressed":"http://wechatapppro-1252524126.file.myqcloud.com/appwpor7xql8808/image/compress/414320457kr329l4907xw.jpg","summary":"app优化专栏","sell_type":1,"price":100,"line_price":100,"is_stop_sell":0,"recycle_bin_state":0,"invite_poster":"","distribute_default":0,"created_at":"2021-09-03 17:56:54","updated_at":"2022-01-10 13:04:17","can_select":1,"is_chosen":0,"is_discount":0,"is_distribute":0,"is_show_resourcecount":1,"visible_on":1,"wx_share":null,"resource_statistic":null,"show_in_menu":1,"have_password":0,"resource_password":"","loop_resource":0,"sell_mode":0,"auth_type":0,"relation_type":0},"op":"+U"}`
	//union订阅 +U
	//str := `{"data":{"id":500001,"app_id":"apppcHqlTPT3482","content_app_id":"","resource_id":"i_5f51a0a360b266ca4e9b9269","user_id":"u_5f55092de60b4_WYNlXe8eKe","universal_union_id":"ozStBt1BVbbkIJBfsjUV5sfIJ-1I","status":1,"resource_type":1,"created_at":"2020-09-07 18:35:43","updated_at":"2020-09-07 18:35:43"},"op":"+U"}`

	//str := `{"data":{"id":7952,"shop_id":"appjXbh1T3A5945","course_id":"course_28bBsT00G3ufh5RVdKJ7UuCGcc3","title":"新建训练营pro2022-05-02 15:13:05","summary":"新建训练营pro","img_url":"https://wechatapppro-1252524126.file.myqcloud.com/appjXbh1T3A5945/image/b_u_5bf76183e88e8_s62XrQSB/kmlh0nij090m.png","img_url_compressed":"https://wechatapppro-1252524126.file.myqcloud.com/appjXbh1T3A5945/image/b_u_5bf76183e88e8_s62XrQSB/kmlh0nij090m.png","img_url_compressed_larger":"https://wechatapppro-1252524126.file.myqcloud.com/appjXbh1T3A5945/image/b_u_5bf76183e88e8_s62XrQSB/kmlh0nij090m.png","created_at":"2022-05-02 15:13:31","updated_at":"2022-05-02 15:13:31"},"op":"+I"}`

	str := `{"token":"935595909a3f876cd4ecc5ed74db8a2c","time_stamp":111212122,"b_user_id":"b_u_5fd87d4f12eb9_og9m7CMS","app_id":"appgKvmP9gT7183","phone_name":"","event":"stuck","duration":3,"platform":"ios","play_session":{"stuck_count":3,"stuck_time":3,"error_code":0,"error_code_name":"stuck","app_cpu_avg":20.0,"sys_cpu_avg ":30.0,"extras":"{\"seesion_id\":\"werwerwerwe\"}"},"download_session":{"success":true,"speed_avg":20.2,"app_cpu_avg":20.0,"sys_cpu_avg":20.0,"extras":"{\"seesion_id\":\"werwerwerwe\"}"},"cast_screen_session":{"success":true,"type":0,"extras":"{\"seesion_id\":\"werwerwerwe\"}"},"client_info":{"batteryLevel":"1.0","deviceName":"\u5c0f\u4e25\u7684\u5b66\u4e60\u5c0f\u9a6c\u8fbe","phoneBrand":"iPadOS","phoneModel":"iPad (9th generation)","systemVersion":"15.3.1"}}`
	msg.Value = sarama.StringEncoder(str)
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	time.Sleep(5 * time.Second)
}

func getId() (id string) {
	data := strconv.Itoa(int(time.Now().Unix()))
	id = data[len(data)-6:]
	return
}
