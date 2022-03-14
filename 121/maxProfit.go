package main

import "math"

func main() {

}

// 股票最大收益
func maxProfit(prices []int) int {
	//记录【今天之前买入的最小值】
	//计算【今天之前最小值买入，今天卖出的获利】，也即【今天卖出的最大获利】
	//比较【每天的最大获利】，取最大值即可

	length := len(prices)
	if length <= 1 {
		return 0
	}

	min := prices[0]
	max := 0
	for i := 0; i < len(prices); i++ {
		max = int(math.Max(float64(max), float64(prices[i]-min)))
		min = int(math.Min(float64(prices[i]), float64(min)))
	}

	return max

}
