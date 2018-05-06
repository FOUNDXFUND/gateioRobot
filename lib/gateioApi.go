package lib

import (
	"crypto/hmac"
	"crypto/sha512"
	// "encoding/hex"
	// "encoding/json"
	"net/http"
	// "net/url"
	// "sort"
	"io/ioutil"
	"strings"
	"fmt"
	"encoding/json"
	"Sknife/utils"
)



type GateApi struct {
	Key string
	Secret string
}


func (s *GateApi) Buy(currencyPair string, rate string, amount string) (b *OrderPlace,err error) {
	method := "POST"
	url := "https://api.gateio.io/api2/1/private/buy"
	param := "currencyPair=" + currencyPair + "&rate=" + rate + "&amount=" + amount
	ret,err := s.httpGo(method,url,param)
	if err != nil {
		return
	}
	x := &OrderPlace{}
	str := fmt.Sprintf("[buy] %s", string(ret))
	utils.WriteLog("gate.response.log", str)
	err = json.Unmarshal(ret,x)
	return x,err
}


// 挂卖单
func (s *GateApi) Sell(currencyPair string, rate string, amount string) (sb *OrderPlace, err error) {
	method := "POST"
	url := "https://api.gateio.io/api2/1/private/sell"
	param := "currencyPair=" + currencyPair + "&rate=" + rate + "&amount=" + amount
	ret,err := s.httpGo(method,url,param)
	if err != nil {
		return
	}
	x := &OrderPlace{}
	err = json.Unmarshal(ret,x)
	return x,err
}


// 取消订单
func (s *GateApi) CancelOrder(orderNumber string, currencyPair string ) (c *CancelOrderStruct,err error) {
	method := "POST"
	url := "https://api.gateio.io/api2/1/private/cancelOrder"
	param := "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	ret,err := s.httpGo(method,url,param)
	x := &CancelOrderStruct{}
	err = json.Unmarshal(ret, x)
	return x,err
}



// 查询挂单状态
func (s *GateApi) GetOrderStatus (orderNumber string, currencyPair string) (sb *OrderStatus, err error) {
	method := "POST"
	url := "https://api.gateio.io/api2/1/private/getOrder"
	param := "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	ret,err := s.httpGo(method,url,param)
	if err != nil {
		return
	}
	x := &OrderStatus{}
	err = json.Unmarshal(ret,x)
	return x,err
}



// 查询余额
func (s *GateApi) GetBalances() (b *BalanceStruct,err error) {
	method := "POST"
	url := "https://api.gateio.io/api2/1/private/balances"
	param := ""
	ret,err := s.httpGo(method,url,param)
	x := &BalanceStruct{}
	fmt.Println(string(ret))
	str := fmt.Sprintf("[getBalance] %s", string(ret))
	utils.WriteLog("gate.response.log", str)
	err = json.Unmarshal(ret,x)
	return x,err
}

// 获取当前市场深度(所有买单，卖单列表)
func (s *GateApi) OrderBook(params string) (o *OrderBookStruct,err error) {
	method := "GET"
	url := "http://data.gateio.io/api2/1/orderBook/" + params
	param := ""
	ret,err := s.httpGo(method,url,param)
	if err != nil {
		return
	}

	x := &OrderBookStruct{}
	err = json.Unmarshal(ret,x)
	return x,err
}



// 读取用户当前所有挂单列表
func (s *GateApi) GetOpenOrders() (o *OpenOrders,err error) {
	method := "POST"
	url := "https://api.gateio.io/api2/1/private/openOrders"
	param := ""
	ret,err  := s.httpGo(method,url,param)
	//fmt.Println(string(ret))
	x := &OpenOrders{}
	err = json.Unmarshal(ret,x)
	return x,err
}




// all support pairs
func getPairs() string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/pairs"
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}

// Market Info
func marketinfo() string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/marketinfo"
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}


// Market Details
func Marketlist() string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/marketlist"
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}


// tickers
func tickers() string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/tickers"
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}


// ticker
func ticker(ticker string) string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/ticker" + "/" + ticker
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}


// Depth
func orderBooks() string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/orderBooks"
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}





// Trade History
func tradeHistory(params string) string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/tradeHistory/" + params
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}





// get deposit address
func depositAddress(currency string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/depositAddress"
	var param string = "currency=" + currency
	var ret string = httpDo(method,url,param)
	return ret
}


// get deposit withdrawal history
func depositsWithdrawals(start string, end string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/depositsWithdrawals"
	var param string = "start=" + start + "&end=" + end
	var ret string = httpDo(method,url,param)
	return ret
}







// Cancel all orders
func cancelAllOrders( types string, currencyPair string ) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/cancelAllOrders"
	var param string = "type=" + types + "&currencyPair=" + currencyPair
	var ret string = httpDo(method,url,param)
	return ret
}




// 获取我的24小时内成交记录
func myTradeHistory( currencyPair string, orderNumber string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/tradeHistory"
	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	var ret string = httpDo(method,url,param)
	return ret
}


// Get my last 24h trades
func withdraw( currency string, amount string, address string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/withdraw"
	var param string = "currency=" + currency + "&amount=" + amount + "address=" + address
	var ret string = httpDo(method,url,param)
	return ret
}


func (s *GateApi)getSign(params string) string {
	key := []byte(s.Secret)
	mac := hmac.New(sha512.New, key)
	mac.Write([]byte(params))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

/**
*  gate.io 官方垃圾代码,待删除
*/
func httpDo(method string,url string, param string) string {
	return ""
}

func (s *GateApi) httpGo(method string,url string, param string) (ret []byte,err error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(param))
	if err != nil {
		return
	}

	sign := s.getSign(param)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("key", s.Key)
	req.Header.Set("sign", sign)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return body,nil;
}