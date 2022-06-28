package data

import "server2/internal/biz"


type orderRepo struct {

}

var _ biz.OrderRepo = (*orderRepo)(nil)


func NewOrderRepo() biz.OrderRepo{
	return new(orderRepo)
}

