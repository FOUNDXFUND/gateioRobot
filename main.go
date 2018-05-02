package main

import (
	"gateioRobot/lib"
	"fmt"
	"github.com/fatih/color"
	"github.com/shopspring/decimal"
)

const SHOW_ALL_DEPTH  = false

// asks 是委卖单, bids 是委买单。
func main() {
	ask1, bid1 := GetBidAskInfo()
	/*s,err := lib.OrderBook("ddd_eth")
	if err != nil {
		panic(err)
	}


	askLen := len(s.Asks)

	if SHOW_ALL_DEPTH {
		for k,v := range s.Asks {
			color.Green("卖%d|%v|%v",askLen  - k, v[0], v[1])
		}

		for k,v := range s.Bids {
			color.Red("买%d|%v|%v",k+1, v[0], v[1])
		}
	}


	// 求卖1买1价差
	ask1 := s.Asks[askLen-1][0]
	bid1 := s.Bids[0][0]
	a1 := decimal.NewFromFloat(ask1)
	b1 := decimal.NewFromFloat(bid1)



	diff1 := a1.Sub(b1).Truncate(6)
	fmt.Println(ask1," ", bid1, "")
	fmt.Println(diff1)
	fmt.Sprintf("卖1:%v, 买1:%v, 价差:%v", ask1, bid1, diff1)*/

	fmt.Println("卖1:", ask1, " 买1:", bid1)
	add := decimal.NewFromFloat(0.000001)
	// 1. 挂比买1多一点价格
	fmt.Println("挂买1==>", bid1.Add(add).StringFixed(6))
	// 2. 等待成交


	// 2. 继续读取挂单挂卖	1便宜一点价格
}


func GetBidAskInfo()(ask1,bid1 decimal.Decimal) {
	s,err := lib.OrderBook("ddd_eth")
	if err != nil {
		panic(err)
	}


	askLen := len(s.Asks)
	//bidLen := len(s.Bids)

	if SHOW_ALL_DEPTH {
		for k,v := range s.Asks {
			color.Green("卖%d|%v|%v",askLen  - k, v[0], v[1])
		}

		for k,v := range s.Bids {
			color.Red("买%d|%v|%v",k+1, v[0], v[1])
		}
	}


	// 求卖1买1价差
	a1_str := decimal.NewFromFloat(s.Asks[askLen-1][0]).StringFixed(6)
	b1_str := decimal.NewFromFloat(s.Bids[0][0]).StringFixed(6)

	a1,err := decimal.NewFromString(a1_str)
	if err != nil {
		panic(err)
	}
	b1,err := decimal.NewFromString(b1_str)
	if err != nil {
		panic(err)
	}
	//diff1 := a1.Sub(b1).Truncate(6)

	return a1,b1
}

