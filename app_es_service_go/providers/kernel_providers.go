package providers

import (
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var Env *library.Env

var CoreDB *library.DB
var SecondDB *library.DB

var CoreGorm *library.GormDB
var SecondGorm *library.GormDB

var DefaultLogger *library.Log
var RequestLogger *library.Log
var CallApiLogger *library.Log

var DefaultRedis *library.RedisClient
var ConfhubRedis *library.RedisClient

var HttpClient *library.HttpClient

var KafkaShopMsgSyncProducer *library.KafkaSyncProducer
var KafkaShopMsgGroupConsumer *library.KafkaGroupConsumer
