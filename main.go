package main

import (
	"gateioRobot/lib"
	"fmt"
	"errors"
	"github.com/fatih/color"
	"github.com/shopspring/decimal"
	"time"
	"os"
	"io/ioutil"
	"encoding/json"
)

const SHOW_ALL_DEPTH  = false
const DDD_COUNT  = "200"
const PAIR  = "ddd_eth"
var NewOrderNumber = make(chan int)
// asks 是委卖单, bids 是委买单。
func main() {
	cfg := LoadCfg()
	gateio := &lib.GateApi{cfg.Key,cfg.Secret}

	go Timer()
	Test()

	select{}
}


// 每两秒取一次
func Timer() {
	ticker := time.NewTicker(2 * time.Second)
	for _ = range ticker.C {
		// 当前订单状态
		fmt.Println("每两秒一次...")
	}

}


func OrderWathcer() {
	for {
		select {
		case orderNumber := <- NewOrderNumber:
			{
				for {
					// 已下单,监控订单状态
					st,err := lib.GetOrderStatus(string(orderNumber), PAIR)
					if err != nil {
						panic(err)
					}

					fmt.Println("订单状态 ==>" , st.Oder.Status)
				}

			}

		}
	}
}

func Test() {
	b,err := lib.GetBalances()
	if err != nil {
		panic(err)
	}


	/*c,err := lib.Buy(PAIR, "0.001","1")
	if err != nil {
		panic(err)
	}

	fmt.Println(c)*/
}

/*
	备注:买入手续费 = 购买花费ETH总值 * 0.0018(0.18%)
   	卖出手续费 = 出售DDD总数量 *  0.0018(0.18%)
*/

func Auto() {
	ask1, bid1,orderBook := GetBidAskInfo(PAIR)
	fmt.Println("卖1:", ask1, " 买1:", bid1)
	// 1. 挂比买1多一点价格
	SetOrder(orderBook,ask1,bid1)

	// 扫描器, 扫描挂单是否成交, 如果挂单时候我的挂单被超过
}

// 买了就跌，这种情况处理？？

func SetOrder(o *lib.OrderBookStruct, ask1,bid1 decimal.Decimal) {
	add := decimal.NewFromFloat(0.000001)
	buy1_str := bid1.Add(add).StringFixed(6)
	buyPerEth,err := decimal.NewFromString(buy1_str)
	if err != nil {
		panic(err)
	}


	fmt.Println("挂买1==>", buyPerEth)

	// 防止出错或者溢出什么的, 挂的买1价格不能高于卖3
	sell3 :=decimal.NewFromFloat(o.Asks[len(o.Asks) - 3][0])
	if buyPerEth.Cmp(sell3) <= 0 {
		str := fmt.Sprintf("准备挂单价格:%s", buyPerEth.String(), " 当前买一价格:", bid1.String())
		panic(errors.New(str))
	}


	b,err := lib.Buy(PAIR, buyPerEth.String(), DDD_COUNT)
	if err != nil {
		panic(err)
	}


	fmt.Println(b)

}

func GetBidAskInfo(pair string)(ask1,bid1 decimal.Decimal,orders *lib.OrderBookStruct) {
	s,err := lib.OrderBook(pair)
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

	return a1,b1,s
}



type Cfg struct {
	Key string
	Secret string
}

func LoadCfg() *Cfg {
	f,err := os.Open("cfg.json")
	if err != nil {
		panic(err)
	}

	b,err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	x := &Cfg{}
	err = json.Unmarshal(b,x)
	if err != nil {
		panic(err)
	}
	return x

}

