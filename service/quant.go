package service

import (
	"gateioRobot/lib"
	"fmt"
	"os"
	"encoding/json"
)

const STATUS_OPEN = "open"
const STATUS_CANCELLED = "cancelled"
const STATUS_DONW  = "done"
const BUY_TYPE = "buy"
const SELL_TYPE = "sell"

var Watching *WatchingOrderInfo


// 当前监控订单信息
type WatchingOrderInfo struct {
	OrderNum string
	Type string // 买卖类型, sell卖出， buy买入
}


func SwitchWathchChecking() {

}

func OderStatusChecking(myOrders *lib.OpenOrders) {
	if myOrders.Result != "true" {
		// 获取订单状态失败，直接退出!?
		x,err := json.Marshal(myOrders)
		fmt.Println("订单内容===>", string(x), "  err==>", err)
		os.Exit(0)
	}

	if len(myOrders.Orders) < 1 {
		return
	}

	if Watching == nil {
		return
	}

	for _,v := range myOrders.Orders{

		// 如果当前订单是买单 --> 计算最佳挂卖单价格 --> 挂卖单
		if Watching.Type == BUY_TYPE {

		}

		// 如果当前订单是卖单 --> 计算最佳挂买单价格 --> 挂买单
		if Watching.Type == SELL_TYPE {

		}

		fmt.Println("组合:", v.CurrencyPair, " 状态", v.Status)
	}
}

func CalcBestPrice(order *lib.Oders) {

}

