package main

import (
	"fmt"
	"github.com/syp25815/bpx-api-go/bpx"
	"strconv"
	"time"
)

func main() {
  // 替换key、secret
	c := bpx.NewClient("key", "secret")

  // 交易对
	Symbol := "WEN_USDC"

  // 间隔时间
	duration := time.NewTicker(time.Millisecond * 1888)
  // 数量
	quantity := "500000"
  // 几档成交价
	gear := 5

	for range duration.C {
		c.OrdersCancels(Symbol)
		fmt.Println(time.Hour)
		a := bpx.Depth(Symbol)
		if a == nil {
			continue
		}
		
		if len(a.Asks) < 5 || len(a.Bids) < 5 {
			continue
		}

		sellPrice, err := strconv.ParseFloat(a.Asks[gear][0], 64)
		if err != nil {
			fmt.Println("转换失败:", err)
			continue
		}

		buyPrice, err := strconv.ParseFloat(a.Bids[len(a.Bids)-gear-1][0], 64)
		if err != nil {
			fmt.Println("转换失败:", err)
			continue
		}

		fmt.Println("buyPrice", buyPrice)
		fmt.Println("sellPrice", sellPrice)

		go func() {
			aa := c.OrderExecute(Symbol, "Ask", "Limit", "", quantity, fmt.Sprintf("%.8f", sellPrice))
			fmt.Println("sell", aa)
		}()

		go func() {
			bb := c.OrderExecute(Symbol, "Bid", "Limit", "", quantity, fmt.Sprintf("%.8f", buyPrice))
			fmt.Println("buy", bb)
		}()

	}

}
