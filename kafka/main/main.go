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
	client, err := sarama.NewSyncProducer([]string{"10.10.27.147:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()
	for {
		sendMsg(client, "mysql2es_t_spu")
	}
}

func sendMsg(client sarama.SyncProducer, topic_name string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic_name

	//id := getId()
	//phone订阅 +I
	//str := `{"data":{"app_id":"appwpor7xql8808","id":"i_607e728660b211dfe4ab1e32","name":"务必马到成功2256","img_url":"http://wechatapppro-1252524126.file.myqcloud.com/appwpor7xql8808/image/kr329l4907xw.jpg","img_url_compressed_larger":"http://wechatapppro-1252524126.file.myqcloud.com/appwpor7xql8808/image/coms/640480808kr329l4907xw.jpg","img_url_compressed":"http://wechatapppro-1252524126.file.myqcloud.com/appwpor7xql8808/image/compress/414320457kr329l4907xw.jpg","summary":"app优化专栏","sell_type":1,"price":100,"line_price":100,"is_stop_sell":0,"recycle_bin_state":0,"invite_poster":"","distribute_default":0,"created_at":"2021-09-03 17:56:54","updated_at":"2022-01-10 13:04:17","can_select":1,"is_chosen":0,"is_discount":0,"is_distribute":0,"is_show_resourcecount":1,"visible_on":1,"wx_share":null,"resource_statistic":null,"show_in_menu":1,"have_password":0,"resource_password":"","loop_resource":0,"sell_mode":0,"auth_type":0,"relation_type":0},"op":"+U"}`
	//union订阅 +U
	//str := `{"data":{"app_id":"app38itOR341547","id":"ex_5c04d37c62da0_gEd3f3Tw","name":"11","img_url":"http://wechatapppro-1252524126.cossh.myqcloud.com/app38itOR341547/image/9c609aee00be31fe4fb724df2acd5d98.png","img_url_compressed":"http://wechatapppro-1252524126.file.myqcloud.com/app38itOR341547/image/compress/2001129c609aee00be31fe4fb724df2acd5d98.png","introduction":"","detail":"<p>11</p>","describ":null,"resource_type":0,"resource_id":"","resource_name":"","total_question":1,"total_score":2,"pass_score_precent":0,"exam_time":0,"exam_chance":1,"reexam_interval":0,"state":2,"open_notice":0,"publish_time":null,"notify_time":null,"end_time":null,"answer_display_type":0,"is_push":0,"participate_count":0,"commit_count":0,"comment_count":0,"is_rel_paper":1,"exam_time_radio":0,"exam_start_time":null,"exam_end_time":null,"qes_score_display":1,"res_score_disp_type":2,"res_qes_disp_type":1,"error_ques_display":1,"creator":"","universal_union_id":"","create_source":0,"sell_mode":0,"created_at":"2018-12-03 14:55:56","updated_at":"2023-07-06 10:06:41","assign_type":0},"op":"+I"}`
	str := `{"data":{"id":1,"app_id":"app04ErzzLP3543","spu_id":"cz_5d9be93b96f8d_DfKyEqCqyO","spu_type":"OLC","resource_id":"cz_5d9be93b96f8d_DfKyEqCqyO","resource_type":29,"goods_sn":"","goods_category_id":"","wx_goods_category_id":"","goods_name":"测试","goods_img":"[\"http://wechatapppro-1252524126.file.myqcloud.com/app04ErzzLP3543/image/cmVzb3VyY2UtZmFjZVRvRmFjZS04MzExMTYyMA.jpg\"]","custom_cover":"","goods_brief_text":"","goods_detail_text":"","detail_cos_url":"","sell_type":1,"is_goods_package":0,"price_low":0,"price_high":0,"price_line":0,"visit_num":0,"goods_tag":"线下课","goods_tag_is_show":1,"sale_status":0,"is_timing_sale":0,"timing_sale":"","sale_at":"","is_timing_off":0,"timing_offtime":null,"timing_off":null,"has_distribute":0,"video_url":"","video_img_url":"","period":-1,"is_display":1,"is_stop_sell":0,"is_forbid":0,"is_ignore":0,"limit_purchase":0,"stock_deduct_mode":0,"appraise_num":0,"show_stock":1,"is_best":0,"is_hot":0,"is_new":0,"is_recom":0,"distribution_pattern":1,"freight":0,"is_uniform_freight":1,"freight_template_id":0,"img_url_compressed":"","is_password":0,"is_free":0,"can_sold_start":null,"can_sold_end":null,"sell_mode":1,"parent_spu_id":"","parent_app_id":"","is_public":1,"is_single":1,"spu_extend":"","is_deleted":0,"audit_user_id":null,"audit_reason":"","audit_time":null,"created_at":"2019-10-08 09:41:15","updated_at":"2019-10-08 09:42:54"},"op":"+I"}`
	//str := `{"EventType":"BAccountCancel","EventVersion":"1.0.0","EventTime":1691572055986,"BAccountCancelEvent":{"b_user_id":"b_u_cghrilsuq6lb5pnjb6rg","b_user_info":{"b_user_id":"b_u_cghrilsuq6lb5pnjb6rg","phone":"13694205591","nation_code":"86","cancel_type":"bind","wx_open_id":"omJa5v_kGQwGigHG7pebAGhvHIeM","wx_union_id":"ozStBt1MHPeldgJztdL4LIzJms0Q"}}}`
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
