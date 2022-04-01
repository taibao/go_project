package request

type GetShopInfoRequest struct {
	AppId string `json:"app_id" binding:"required"`
}