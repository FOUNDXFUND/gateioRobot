package lib

type OrderPlace struct {
	Result string
	OrderNumber string//订单号码
	Rate string // 下单价格
	LeftAmount string //剩余数量
	FilledAmound string // 成交数量
	FilledRate string // 成交价格
}




type BalanceStruct struct {
	Result string
	Available map[string]interface{}
	Locked map[string]interface{}
}


type OrderBookStruct struct {
	Result string
	Asks [][]float64
	Bids [][]float64
}


type CancelOrderStruct struct {
	Result string
	Message string
}



type OrderStatus struct {
	Result string
	Oder *Order
	Message string
}

type Order struct {
	Id string
	Status string // 订单状态, cancelled已取消, done已完成
	CurrencyPair string
	Type string // 买卖类型, sell卖出， buy买入
	Rate float64 // 价格
	Amount string //买卖数量
	InitialRate float64 // 下单价格
	InitialAmount float64 // 下单数量
}