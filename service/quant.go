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


	for _,v := range myOrders.Orders{
		fmt.Println("组合:", v.CurrencyPair, " 状态", v.Status)
	}

}

