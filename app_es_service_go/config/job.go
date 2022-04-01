package config

import (
	"context"
	"app_es_service_go/job"
)

var JobConfigs = map[string]func(ctx context.Context){
	"shop_msg_consume": job.StartShopMsgConsume,
}
