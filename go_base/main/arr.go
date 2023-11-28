package main

import (
	"fmt"
	"strings"
	"time"
)

func GetMinutesAgo() string {
	currentTime := time.Now()
	t, _ := time.ParseDuration("-8h")
	result := currentTime.Add(t)
	return result.Format("2006-01-02 15:04:05.000000000")
}

func GetDateV4(goBack string) string {
	currentTime := time.Now()
	if goBack != "" {
		t, _ := time.ParseDuration(goBack)
		result := currentTime.Add(t)
		return result.Format("2006-01-02 00:00:00")
	}
	return currentTime.Format("2006-01-02 00:00:00")
}

func HttpBuildQuery(params map[string]string) (param_str string) {
	params_arr := make([]string, 0, len(params))
	for k, v := range params {
		params_arr = append(params_arr, fmt.Sprintf("%s=%s", k, v))
	}
	//fmt.Println(params_arr)
	param_str = strings.Join(params_arr, "&")
	return param_str
}

// Explode explode()
func Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}

type LiveSubscription struct {
	UnlockProduct    []UnlockProduct `json:"unlock_product"`
	ClassId          string          `json:"class_id"`
	LessonState      int             `json:"lesson_state"`
	IsLock           int             `json:"is_lock"`
	ResourceID       string          `json:"id"`
	AppID            string          `json:"app_id"`
	RoomID           string          `json:"room_id"`
	Title            string          `json:"title"`
	AliveType        int             `json:"type"`
	StartAt          time.Time       `json:"-"`
	StopAt           time.Time       `json:"-"`
	StartAtStr       string          `json:"start_at"`
	StopAtStr        string          `json:"stop_at"`
	ImgURL           string          `json:"img_url"`
	ImgURLCompressed string          `json:"img_url_compressed"`
	IsLookback       int             `json:"is_lookback"`
	PlayURL          string          `json:"play_url"`
	PushState        int             `json:"push_state"`
	CUserID          string          `json:"c_user_id"`
	IsShowTimeline   bool            `json:"is_show_timeline"`
	ResourceType     int             `json:"resource_type"`
	AliveState       int             `json:"alive_state"`
	IsTeacher        bool            `json:"is_teacher"`
	ContentAppId     string          `json:"content_app_id"`
	IsUser           bool            `json:"is_user"`
	ShopName         string          `json:"shop_name"`
	ShopLogo         string          `json:"shop_logo"`
	StartHour        string          `json:"start_hour"`
	StopHour         string          `json:"stop_hour"`
	Level            []int           `json:"level"`
	Parents          []Parent        `json:"parents"`
	H5Url            string          `json:"h5_url"`
}

type Parent struct {
	ProductID   string `json:"product_id"`
	ProductType int    ` json:"product_type"`
}

type UnlockProduct struct {
	AppID        string `json:"app_id"`
	ResourceID   string `json:"resource_id"`
	Title        string `json:"title"`
	ResourceType int    `json:"resource_type"`
	ImgURL       string `json:"img_url"`
	ShopName     string `json:"shop_name"`
	ShopLogo     string `json:"shop_logo"`
	H5URL        string `json:"h5_url"`
}

func FilterLiveSubscriptionByTime(subs []LiveSubscription, startTimeStr, endTimeStr string) ([]LiveSubscription, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, loc)
	if err != nil {
		return nil, err
	}
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, loc)
	if err != nil {
		return nil, err
	}
	var result []LiveSubscription
	for _, sub := range subs {

		fmt.Println("sub.StartAt", sub.StartAt, "startTime", startTime)
		//只判断直播开始时间
		if sub.StartAt.After(startTime) && sub.StartAt.Before(endTime) || sub.StartAt.Equal(startTime) || sub.StartAt.Equal(endTime) {
			result = append(result, sub)
		}
	}
	return result, nil
}

func main() {
	// AppIdUsers := Explode(":","appviesn6py9620:u_646aec135da19_SYqkgw52Wz")
	// fmt.Println(AppIdUsers[1])

	//str := `[{"unlock_product":null,"class_id":"","lesson_state":0,"is_lock":0,"id":"l_64644572e4b0f2aa7deb6a4d","app_id":"appviesn6py9620","room_id":"XET#6ee4858f6000ea0","title":"24川农考研高分备考公开课","type":2,"start_at":"2023-05-22 19:30:00","stop_at":"2023-05-22 20:30:00","img_url":"https://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/b_u_618cfab08001f_2dkGe7PX/lhtzlojq0j7g.png","img_url_compressed":"http://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/compress/400204288lhtzlojq0j7g.png","is_lookback":1,"play_url":"[\"rtmp:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_6ee4858ecc000bct6\",\"http:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_6ee4858ecc000bct6.flv\",\"http:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_6ee4858ecc000bct6.m3u8\",\"http:\\/\\/wechatapppro-1252524126.file.myqcloud.com\\/appviesn6py9620\\/alive_auto_playurl\\/appviesn6py9620_aliveauto_5060_6ee4858ecc000bct6_5060_6ee4858ecc000bct6.m3u8\"]","push_state":0,"c_user_id":"u_646aec135da19_SYqkgw52Wz","is_show_timeline":true,"resource_type":4,"alive_state":3,"is_teacher":false,"content_app_id":"","is_user":true,"shop_name":"新航线书店","shop_logo":"https://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/b_u_618cfab08001f_2dkGe7PX/kvw9gtyv04n2.jpg","start_hour":"19:00","stop_hour":"20:00","level":[1],"parents":null,"h5_url":"https://appviesn6py9620.h5.xiaoeknow.com/v1/course/alive/l_64644572e4b0f2aa7deb6a4d?type=2"},{"unlock_product":null,"class_id":"","lesson_state":0,"is_lock":0,"id":"l_64999585e4b0cf39e6dce87a","app_id":"appviesn6py9620","room_id":"XET#722588b3500016e","title":"生物化学核心难点-糖酵解","type":2,"start_at":"2023-06-27 19:00:00","stop_at":"2023-06-27 20:00:00","img_url":"https://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/b_u_614dfaf1ce9ff_EaYvpkBq/ljcwp9gu01ng.jpg","img_url_compressed":"http://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/compress/400225597ljcwp9gu01ng.jpg","is_lookback":1,"play_url":"[\"rtmp:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_722588b274008d0ms\",\"http:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_722588b274008d0ms.flv\",\"http:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_722588b274008d0ms.m3u8\",\"http:\\/\\/wechatapppro-1252524126.file.myqcloud.com\\/appviesn6py9620\\/alive_auto_playurl\\/appviesn6py9620_aliveauto_5060_722588b274008d0ms_5060_722588b274008d0ms.m3u8\"]","push_state":0,"c_user_id":"u_646aec135da19_SYqkgw52Wz","is_show_timeline":true,"resource_type":4,"alive_state":3,"is_teacher":false,"content_app_id":"","is_user":true,"shop_name":"新航线书店","shop_logo":"https://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/b_u_618cfab08001f_2dkGe7PX/kvw9gtyv04n2.jpg","start_hour":"19:00","stop_hour":"20:00","level":[1],"parents":null,"h5_url":"https://appviesn6py9620.h5.xiaoeknow.com/v1/course/alive/l_64999585e4b0cf39e6dce87a?type=2"},{"unlock_product":null,"class_id":"","lesson_state":0,"is_lock":0,"id":"l_64c500d4e4b0d1e42e88204a","app_id":"appviesn6py9620","room_id":"XET#74cbf56e6800bad","title":"高分学霸带你考川农---854系列","type":2,"start_at":"2023-07-30 19:30:00","stop_at":"2023-07-30 20:30:00","img_url":"https://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/b_u_618cfab08001f_2dkGe7PX/sxxkdulko2g2vv.png","img_url_compressed":"http://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/compress/400202850sxxkdulko2g2vv.png","is_lookback":1,"play_url":"[\"rtmp:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_74cbf56da000dfeqo\",\"http:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_74cbf56da000dfeqo.flv\",\"http:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_74cbf56da000dfeqo.m3u8\",\"http:\\/\\/wechatapppro-1252524126.file.myqcloud.com\\/appviesn6py9620\\/alive_auto_playurl\\/appviesn6py9620_aliveauto_5060_74cbf56da000dfeqo_5060_74cbf56da000dfeqo.m3u8\"]","push_state":0,"c_user_id":"u_646aec135da19_SYqkgw52Wz","is_show_timeline":true,"resource_type":4,"alive_state":3,"is_teacher":false,"content_app_id":"","is_user":true,"shop_name":"新航线书店","shop_logo":"https://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/b_u_618cfab08001f_2dkGe7PX/kvw9gtyv04n2.jpg","start_hour":"19:00","stop_hour":"20:00","level":[1],"parents":null,"h5_url":"https://appviesn6py9620.h5.xiaoeknow.com/v1/course/alive/l_64c500d4e4b0d1e42e88204a?type=2"},{"unlock_product":null,"class_id":"","lesson_state":0,"is_lock":0,"id":"l_64d61894e4b0d1e42e8ccb96","app_id":"appviesn6py9620","room_id":"XET#75d7088b4c00d6f","title":"24备考规划+高分攻略","type":1,"start_at":"2023-08-09 20:00:00","stop_at":"2023-08-09 20:38:09","img_url":"https://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/b_u_618cfab08001f_2dkGe7PX/madm5mll6hpkdy.jpg","img_url_compressed":"http://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/compress/400224855madm5mll6hpkdy.jpg","is_lookback":1,"play_url":"http://c-vod-hw.xiaoeknow.com/asset/02435a4241d4e0cad356b7ae51174e58/451fe7fb9e0c8949e56e74e36ffc2d18.m3u8","push_state":2,"c_user_id":"u_646aec135da19_SYqkgw52Wz","is_show_timeline":true,"resource_type":4,"alive_state":3,"is_teacher":false,"content_app_id":"","is_user":true,"shop_name":"新航线书店","shop_logo":"https://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/b_u_618cfab08001f_2dkGe7PX/kvw9gtyv04n2.jpg","start_hour":"20:00","stop_hour":"20:00","level":[1],"parents":null,"h5_url":"https://appviesn6py9620.h5.xiaoeknow.com/v1/course/alive/l_64d61894e4b0d1e42e8ccb96?type=2"},{"unlock_product":null,"class_id":"","lesson_state":0,"is_lock":0,"id":"l_64d8830ce4b09d7237a7015d","app_id":"appviesn6py9620","room_id":"XET#75fcc810f0001bc","title":"公益直播第一场-蛋白质化学","type":2,"start_at":"2023-08-13 19:30:00","stop_at":"2023-08-13 20:30:00","img_url":"https://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/b_u_614dfaf1ce9ff_EaYvpkBq/ll93xxko0sac.jpg","img_url_compressed":"http://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/compress/301225425ll93xxko0sac.jpg","is_lookback":1,"play_url":"[\"rtmp:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_75fcc8102000ab5rE\",\"http:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_75fcc8102000ab5rE.flv\",\"http:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_75fcc8102000ab5rE.m3u8\",\"http:\\/\\/wechatapppro-1252524126.file.myqcloud.com\\/appviesn6py9620\\/alive_auto_playurl\\/appviesn6py9620_aliveauto_5060_75fcc8102000ab5rE_5060_75fcc8102000ab5rE.m3u8\"]","push_state":0,"c_user_id":"u_646aec135da19_SYqkgw52Wz","is_show_timeline":true,"resource_type":4,"alive_state":3,"is_teacher":false,"content_app_id":"","is_user":true,"shop_name":"新航线书店","shop_logo":"https://wechatapppro-1252524126.file.myqcloud.com/appviesn6py9620/image/b_u_618cfab08001f_2dkGe7PX/kvw9gtyv04n2.jpg","start_hour":"19:00","stop_hour":"20:00","level":[1],"parents":null,"h5_url":"https://appviesn6py9620.h5.xiaoeknow.com/v1/course/alive/l_64d8830ce4b09d7237a7015d?type=2"},{"unlock_product":null,"class_id":"","lesson_state":0,"is_lock":0,"id":"l_623d8acbe4b01a4851e906e3","app_id":"appmHqP1QzP9963","room_id":"XET#K3f681af6e","title":"2022年昆明市“公共就业服务进校园”活动系列之海归英才聚春城 创业分享新思想直播引才招聘","type":2,"start_at":"2022-03-31 19:00:00","stop_at":"2022-03-31 22:00:00","img_url":"https://wechatapppro-1252524126.file.myqcloud.com/appmHqP1QzP9963/image/b_u_5e5323695021b_alitLCoO/l1eex08d0vwm.png","img_url_compressed":"https://wechatapppro-1252524126.file.myqcloud.com/appmHqP1QzP9963/image/b_u_5e5323695021b_alitLCoO/l1eex08d0vwm.png","is_lookback":1,"play_url":"[\"rtmp:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_EEtNezW2Nm2shwKn\",\"http:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_EEtNezW2Nm2shwKn.flv\",\"http:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_EEtNezW2Nm2shwKn.m3u8\",\"http:\\/\\/wechatapppro-1252524126.file.myqcloud.com\\/appmHqP1QzP9963\\/alive_auto_playurl\\/appmHqP1QzP9963_aliveauto_5060_EEtNezW2Nm2shwKn_5060_EEtNezW2Nm2shwKn.m3u8\"]","push_state":0,"c_user_id":"u_6245930a4dae5_Fqok4hsJAH","is_show_timeline":true,"resource_type":4,"alive_state":3,"is_teacher":false,"content_app_id":"","is_user":true,"shop_name":"云上职播","shop_logo":"http://wechatapppro-1252524126.file.myqcloud.com/appmHqP1QzP9963/image/cmVzb3VyY2Utc2hvcFNldHRpbmctNjA3Mjk2NDU.png","start_hour":"19:00","stop_hour":"22:00","level":[1],"parents":null,"h5_url":"https://appmHqP1QzP9963.h5.xiaoeknow.com/v1/course/alive/l_623d8acbe4b01a4851e906e3?type=2"},{"unlock_product":null,"class_id":"","lesson_state":0,"is_lock":0,"id":"l_647e8b83e4b0b2d1c420765a","app_id":"appmHqP1QzP9963","room_id":"XET#707f0c624000aad","title":"三支一扶/西部志愿者必学-云南省省情专题","type":2,"start_at":"2023-06-07 19:00:00","stop_at":"2023-06-07 21:00:00","img_url":"https://wechatapppro-1252524126.file.myqcloud.com/appmHqP1QzP9963/image/b_u_5e5323695021b_alitLCoO/lijlks8i00df.jpeg","img_url_compressed":"http://wechatapppro-1252524126.file.myqcloud.com/appmHqP1QzP9963/image/compress/400170403lijlks8i00df.jpeg","is_lookback":1,"play_url":"[\"rtmp:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_707f0c61840010cNZ\",\"http:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_707f0c61840010cNZ.flv\",\"http:\\/\\/liveplay.xiaoeknow.com\\/live\\/5060_707f0c61840010cNZ.m3u8\",\"http:\\/\\/wechatapppro-1252524126.file.myqcloud.com\\/appmHqP1QzP9963\\/alive_auto_playurl\\/appmHqP1QzP9963_aliveauto_5060_707f0c61840010cNZ_5060_707f0c61840010cNZ.m3u8\"]","push_state":0,"c_user_id":"u_6245930a4dae5_Fqok4hsJAH","is_show_timeline":true,"resource_type":4,"alive_state":3,"is_teacher":false,"content_app_id":"","is_user":true,"shop_name":"云上职播","shop_logo":"http://wechatapppro-1252524126.file.myqcloud.com/appmHqP1QzP9963/image/cmVzb3VyY2Utc2hvcFNldHRpbmctNjA3Mjk2NDU.png","start_hour":"19:00","stop_hour":"21:00","level":[1],"parents":null,"h5_url":"https://appmHqP1QzP9963.h5.xiaoeknow.com/v1/course/alive/l_647e8b83e4b0b2d1c420765a?type=2"},{"unlock_product":null,"class_id":"bclass_64c8f23a43551_hUvIAd","lesson_state":2,"is_lock":0,"id":"lesson_nZLsCY9L2XY79JX0","app_id":"appviesn6py9620","room_id":"983599896","title":"第五章生物膜的结构与功能","type":1,"start_at":"2023-08-14 17:00:00","stop_at":"2023-08-14 19:00:00","img_url":"https://wechatapppro-1252524126.cdn.xiaoeknow.com/appviesn6py9620/image/b_u_614dfaf1ce9ff_EaYvpkBq/lks8q4970o63.jpg?imageMogr2/thumbnail/304x/quality/100%7CimageMogr2/ignore-error/1","img_url_compressed":"https://wechatapppro-1252524126.cdn.xiaoeknow.com/appviesn6py9620/image/b_u_614dfaf1ce9ff_EaYvpkBq/lks8q4970o63.jpg?imageMogr2/thumbnail/304x/quality/100%7CimageMogr2/ignore-error/1","is_lookback":0,"play_url":"https://encrypt-k-vod.xet.tech/522ff1e0vodcq1252524126/1606f4165576678019787236217/playlist_eof.m3u8?sign=7624a67a011476d6c6d4ea906b348378\u0026t=64e48e7e\u0026us=JrGlDzOZIS","push_state":0,"c_user_id":"u_646aec135da19_SYqkgw52Wz","is_show_timeline":true,"resource_type":35,"alive_state":0,"is_teacher":true,"content_app_id":"","is_user":false,"shop_name":"新航线书店","shop_logo":"https://wechatapppro-1252524126.cdn.xiaoeknow.com/appviesn6py9620/image/b_u_618cfab08001f_2dkGe7PX/kvw9gtyv04n2.jpg?imageMogr2/thumbnail/304x/quality/100%7CimageMogr2/ignore-error/1","start_hour":"17:00","stop_hour":"19:00","level":null,"parents":null,"h5_url":"https://appviesn6py9620.h5.xiaoeknow.com/v1/course/alive/lesson_nZLsCY9L2XY79JX0?type=2"},{"unlock_product":null,"class_id":"bclass_64886371ecaa7_xyPzjr","lesson_state":2,"is_lock":0,"id":"lesson_Bm7RyxCpcrMZBBBk","app_id":"appviesn6py9620","room_id":"396436011","title":"病理-细菌性疾病病理","type":1,"start_at":"2023-08-14 18:30:00","stop_at":"2023-08-14 20:00:00","img_url":"https://wechatapppro-1252524126.cdn.xiaoeknow.com/appviesn6py9620/image/b_u_614dfaf1ce9ff_EaYvpkBq/liu9rjxp0csx.jpg?imageMogr2/thumbnail/304x/quality/100%7CimageMogr2/ignore-error/1","img_url_compressed":"https://wechatapppro-1252524126.cdn.xiaoeknow.com/appviesn6py9620/image/b_u_614dfaf1ce9ff_EaYvpkBq/liu9rjxp0csx.jpg?imageMogr2/thumbnail/304x/quality/100%7CimageMogr2/ignore-error/1","is_lookback":0,"play_url":"https://encrypt-k-vod.xet.tech/522ff1e0vodcq1252524126/e5888a355576678019790196226/playlist_eof.m3u8?sign=8bff3263eb22ee5d1aa76a2c956940db\u0026t=64e48e7e\u0026us=RmlrGdRNLC","push_state":0,"c_user_id":"u_646aec135da19_SYqkgw52Wz","is_show_timeline":true,"resource_type":35,"alive_state":0,"is_teacher":true,"content_app_id":"","is_user":false,"shop_name":"新航线书店","shop_logo":"https://wechatapppro-1252524126.cdn.xiaoeknow.com/appviesn6py9620/image/b_u_618cfab08001f_2dkGe7PX/kvw9gtyv04n2.jpg?imageMogr2/thumbnail/304x/quality/100%7CimageMogr2/ignore-error/1","start_hour":"18:00","stop_hour":"20:00","level":null,"parents":null,"h5_url":"https://appviesn6py9620.h5.xiaoeknow.com/v1/course/alive/lesson_Bm7RyxCpcrMZBBBk?type=2"}]`
	//
	//
	//var data []LiveSubscription
	//err := json.Unmarshal([]byte(str), &data)
	//if err != nil {
	//	fmt.Println("err",err)
	//}
	//
	//
	//liveSubscriptionByTime, err := FilterLiveSubscriptionByTime(data, "2023-08-14 00:00:00", "2023-08-14 23:59:59")
	//
	//fmt.Println("liveSubscriptionByTim",liveSubscriptionByTime)

	StartAt := "2023-08-11 19:30:00"
	StopAt := "2023-08-13 19:31:00"

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", StartAt, loc)
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", StopAt, loc)
	str := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 24, 0, 0, 0, endTime.Location())

	fmt.Println("输入", str)
	//fmt.Println(GetMinutesAgo())

	//str := map[string][]string{"first": {"value"}, "multi": {"foo bar", "baz"}}
	//decodeStr = HttpBuildQuery(str)

	//str :="eyJhcHBfaWQiOiJhcHBvdnlqdm5leTI2MDEiLCJkYXRhIjp7Ind4X3VuaW9uX2lkIjoib1RIVzV2NmFnZzFiOGh4M1Z5YkM5V21VM3dhOCIsInd4X25pY2tuYW1lIjoi5Y2T5ZiJ5a6+Iiwid3hfYXZhdGFyIjoiaHR0cHM6Ly90aGlyZHd4LnFsb2dvLmNuL21tb3Blbi92aV8zMi9RMGo0VHdHVGZUSlJ2RFZUekRMNlZBaWEyRUtFd09Xd1o1NDJ5STM2QTdjb2hGbVBrYXJlRXY2Z01IOHZnOTl5aWFGdHNRM2liemZqMnd5OVpEUUxMRWJBQS8xMzIifX0="
	//fmt.Println(string([]byte(str)))
	//arr := map[string]interface{}{
	//	"app_id":"",
	//
	//}
	//
	//var uri url.URL
	//q := uri.Query()
	//q.Add("name", "张三")
	//q.Add("age", "20")
	//q.Add("sex", "1")
	//q.Add("wew","1232")
	//queryStr := q.Encode()
	//fmt.Println(queryStr)

	//fmt.Println(91/10)

	//测试break
	//for i := 0; i < 2; i++ {
	//	j := 0
	//	for {
	//		if j >= 5 {
	//			break
	//		}
	//		j++
	//		fmt.Println("第", i, "列")
	//	}
	//}

	//jsons := `["1：修复直播黑屏，app闪退","2：优化ui"]`
	//var msg  []string
	//err := json.Unmarshal([]byte(jsons), &msg)
	//
	//fmt.Println(msg , err)
	//
	//
	//
	//os.Exit(1)
	//使用数组求平均
	//
	//var hens [6]float64 //定义数组
	//
	//hens[0] = 3.0
	//hens[1] = 5.0
	//hens[2] = 1.0
	//hens[3] = 3.4
	//hens[4] = 2.0
	//hens[5] = 50.0
	//
	//totalWeight2 := 0.0
	//for i:=0;i<len(hens);i++{
	//	totalWeight2 += hens[i]
	//}
	//
	////平均体重
	//avgWeight2 := fmt.Sprintf("%.2f",totalWeight2/float64(len(hens)))
	//fmt.Printf("totalWeight2 %v ,avgWeight2 = %v,",totalWeight2,avgWeight2)

	//var intArr [3]int
	////当我们定义完数组后，其实数组的各个元素有默认值0
	//fmt.Println(intArr)
	//fmt.Printf("intArr的地址=%p intArr[0] 地址%p intArr[1] 地址%p intArr[2] 地址%p",&intArr,&intArr[0],&intArr[1],&intArr[2])
	//

	//1:数组的地址可以通过数组名来获取&intArr
	//2.

	//var score [5]float64
	//
	//for i :=0;i<len(score);i++{
	//	fmt.Printf("请输入第%d个元素的值\n",i+1)
	//	fmt.Scanln(&score[i])
	//}
	//
	////变量数组打印
	//for i:=0;i<len(score);i++{
	//	fmt.Printf("score[%d]=%v",i,score[i])
	//}

	////四种初始化数组的方式
	//var numArr01 [3]int = [3]int{1,2,3}
	//fmt.Println("numArr01=",numArr01)
	//
	//var numArr02 = [3]int{5,6,7}
	//fmt.Println("numArr02=",numArr02)
	//
	//var numArr03 = [...]int{8,9,10} //三个点是固定写法，不能增减改变
	//fmt.Println("numArr03=",numArr03)
	//
	//var numArr04 = [...]int{1:800,0:900,2:299} //指定下标赋值
	//fmt.Println("numArr04=",numArr04)
	//
	////也可以类型推导
	//numArr05 :=  [...]int{1:800,0:900,2:299}
	//fmt.Println("numArr05=",numArr05)

	//heroes := [...]string{"宋江","吴用","卢俊义"}
	//
	//heroes2 := heroes
	//heroes2[0] = "卓嘉宾"
	//for i,v := range heroes2{
	//	fmt.Printf("heroes2[%d]=%v\n",i,v)
	//}
	//
	//for i,v := range heroes{
	//	fmt.Printf("heroes[%d]=%v\n",i,v)
	//}

	//var b [26]byte
	//index := 0
	//for i :='A';i<='Z';i++{
	//	b[index] = byte(i)
	//	index++
	//}
	//
	//fmt.Printf("输出值%c",b)

	//数组平均值
	//	var intArr = [...]int{1,-1,9,90,11}
	//	sum := 0
	//	for _,v := range intArr{
	//		//累计求和
	//		sum += v
	//	}
	//
	//	avg := float64(sum) /float64(len(intArr))
	//	fmt.Printf("sum=%v abg=%v",sum,avg)

	//var intArr3 [5]int
	////为了每次生成的随机数不一样，我们需要给一个seed值
	//rand.Seed(time.Now().UnixNano())
	//len := len(intArr3)
	//for i:=0;i<len;i++{
	//	intArr3[i] = rand.Intn(100)
	//}
	//fmt.Println(intArr3)
	////反转打印
	//temp :=0
	//for i:=0;i<len/2;i++{
	//	temp = intArr3[i]
	//	intArr3[i] = intArr3[len - i -1]
	//	intArr3[len - i -1] = temp
	//}
	//fmt.Println(intArr3)

	//二维数组

	//定义二位数组
	//var arr [4][6]int
	////fmt.Println(arr)
	//
	//for i:=0;i<len(arr);i++{
	//	for j:=0;j<len(arr[i]);j++{
	//		fmt.Print(" "+strconv.Itoa(arr[i][j]))
	//	}
	//	fmt.Println()
	//}

	//初始化数组
	//var arr = [2][3]int{{1,2,3},{4,5,6}}
	//fmt.Println("arr=",arr)
	//
	//data := []string{"apppcHqlTPT3482","13066867190"}
	//fmt.Println(data)
	//
	////二维数组遍历
	//for i,v := range arr{
	//	for j,v2 := range v{
	//		fmt.Printf("arr[%v][%v]=%v ",i,j,v2)
	//	}
	//	fmt.Println()
	//}
	//
	//	params := genSQLRangeStrByIntArr([]string{"123","13066867190"})
	//	fmt.Println(params)
	//	s :=  fmt.Sprintf("SELECT * FROM t_phone_subscribe_record_0 WHERE (app_id=? AND phone_name=?)" ,params)
	//	fmt.Println(s)
	//}

	//
	//
	//func genSQLRangeStrByIntArr(arr []string) (res string) {
	//	var tempStrArr = make([]string, len(arr))
	//	for k, v := range arr {
	//		tempStrArr[k] = fmt.Sprintf("%s", v)
	//	}
	//	res = "(" + strings.Join(tempStrArr, ",") + ")"
	//	return
	//}
	//

}
