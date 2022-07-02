package main

import (
	"fmt"
)

// 设想有一个戏剧演出团，演员们经常要去各种场合表演戏剧，通常客户（customer）会指定几出剧目，而剧团则根据观众（audience）人数及剧目类型来向客户收费。
// 该团目前出演两种戏剧：悲剧（tragedy）和喜剧（comedy）。给客户发出账单时，剧团还会根据到场观众的数量给出"观众量积分"（volume credit）优惠，
// 下次客户再请剧团表演时可以使用积分获得折扣——你可以把它看作一种提升客户忠诚度的方式

// Play ...
type Play struct {
	Name string
	Type string
}

// Performance ...
type Performance struct {
	PlayID   string
	Audience int
}

type Invoice struct {
	Customer     string
	Performances []Performance
}

// 剧目数据
var plays = map[string]Play{
	"hamlet": {
		Name: "Hamlet",
		Type: "tragedy",
	},
	"as-like": {
		Name: "As You Like It",
		Type: "comedy",
	},
	"othello": {
		Name: "Othello",
		Type: "tragedy",
	},
}

var performances = []Performance{
	{PlayID: "hamlet", Audience: 55}, {PlayID: "as-like", Audience: 35}, {PlayID: "othello", Audience: 40},
}

// 账单信息
var invoice = Invoice{
	Customer:     "BigGo",
	Performances: performances,
}

//func main() {
//	statement(invoice, plays)
//	//statementApplyAmountFor(invoice, plays)
//}

func numberFormat(number int) string {
	return fmt.Sprintf("%d.00", number)
}

func statement(invoice Invoice, plays map[string]Play) {
	var totalAmount = 0
	var volumeCredits = 0
	result := fmt.Sprintf("Statement for %s \n", invoice.Customer)
	var format = numberFormat
	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
		thisAmount := 0

		switch play.Type {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * (perf.Audience - 30)
			}
			break
		case "comedy":
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += 10000 + 500*(perf.Audience-20)
			}
			thisAmount += 300 * perf.Audience
			break
		default:
			fmt.Printf("unknown type : %s", play.Type)
			return
		}

		// add volume credits
		if perf.Audience-30 > 0 {
			volumeCredits += perf.Audience - 30
		}
		// add extra credit for every ten comedy attendees
		if play.Type == "comedy" {
			volumeCredits += perf.Audience / 5
		}

		// print line for this order
		result += fmt.Sprintf(" %s: %s (%d seats)\n", play.Name, format(thisAmount/100), perf.Audience)
		totalAmount += thisAmount
	}

	result += fmt.Sprintf("Amount owed is %s\n", format(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)
	fmt.Println(result)
	return
}

// 需求驱动重构，如果一段代码可以运行而且永远不需要修改，那么这段代码就不需要重构
// 假设有一个需求，打印HTML格式的账单，该如何实现？

// ----------------------------------------第一阶段：分解statement函数-----------------------------------------------------
// 1.提炼出amountFor函数(提炼函数)
func amountFor(perf Performance, play Play) int {
	thisAmount := 0
	switch play.Type {
	case "tragedy":
		thisAmount = 40000
		if perf.Audience > 30 {
			thisAmount += 1000 * (perf.Audience - 30)
		}
		break
	case "comedy":
		thisAmount = 30000
		if perf.Audience > 20 {
			thisAmount += 10000 + 500*(perf.Audience-20)
		}
		thisAmount += 300 * perf.Audience
		break
	default:
		return 0
	}
	return thisAmount
}

// 2.应用提炼出的函数
func statementApplyAmountFor(invoice Invoice, plays map[string]Play) {
	var totalAmount = 0
	var volumeCredits = 0
	result := fmt.Sprintf("Statement for %s \n", invoice.Customer)
	var format = numberFormat
	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
		thisAmount := amountFor(perf, play)

		// add volume credits
		if perf.Audience-30 > 0 {
			volumeCredits += perf.Audience - 30
		}
		// add extra credit for every ten comedy attendees
		if play.Type == "comedy" {
			volumeCredits += perf.Audience / 5
		}

		// print line for this order
		result += fmt.Sprintf(" %s: %s (%d seats)\n", play.Name, format(thisAmount/100), perf.Audience)
		totalAmount += thisAmount
	}

	result += fmt.Sprintf("Amount owed is %s\n", format(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)
	fmt.Println(result)
	return
}

// 编译，测试，提交
// 小步修改，以及它带来的频繁反馈，正是防止混乱的关键
// 重构技术就是以微小的步伐修改程序。如果你犯了错误，很容易便可发现它。

// 3.对提炼出的amountFor函数进行优化
// rename thisAmount -> result
// rename perf -> aPerformance
// 傻瓜都能写出计算机可以理解的代码。唯有写出人类容易理解的代码时，才是优秀的程序员  ----命名的重要性
func amountForRename(aPerformance Performance, play Play) int {
	result := 0
	switch play.Type {
	case "tragedy":
		result = 40000
		if aPerformance.Audience > 30 {
			result += 1000 * (aPerformance.Audience - 30)
		}
		break
	case "comedy":
		result = 30000
		if aPerformance.Audience > 20 {
			result += 10000 + 500*(aPerformance.Audience-20)
		}
		result += 300 * aPerformance.Audience
		break
	default:
		return 0
	}
	return result
}

// 4.移除play变量(以查询取代临时变量，内联变量)
func playFor(aPerformance Performance, plays map[string]Play) Play {
	return plays[aPerformance.PlayID]
}

func statementRemovePlay(invoice Invoice, plays map[string]Play) {
	var totalAmount = 0
	var volumeCredits = 0
	result := fmt.Sprintf("Statement for %s \n", invoice.Customer)
	var format = numberFormat
	for _, perf := range invoice.Performances {
		// thisAmount := amountFor(perf, playFor(perf)) // remove play and apply playFor

		// add volume credits
		if perf.Audience-30 > 0 {
			volumeCredits += perf.Audience - 30
		}
		// add extra credit for every ten comedy attendees
		if playFor(perf, plays).Type == "comedy" { // inline variable
			volumeCredits += perf.Audience / 5
		}

		// print line for this order
		result += fmt.Sprintf(" %s: %s (%d seats)\n", playFor(perf, plays).Name, // inline variable
			format(amountFor(perf, playFor(perf, plays))/100), perf.Audience) // inline variable
		totalAmount += amountFor(perf, playFor(perf, plays)) // inline variable
	}

	result += fmt.Sprintf("Amount owed is %s\n", format(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits", volumeCredits)
	fmt.Println(result)
	return
}

// 编译，测试，提交

// 5.提炼计算观众量积分的逻辑
func volumeCreditsFor(aPerformance Performance, play Play) int {
	result := 0
	if aPerformance.Audience-30 > 0 {
		result += aPerformance.Audience - 30
	}
	if play.Type == "comedy" {
		result += aPerformance.Audience / 5
	}
	return result
}

// 应用提炼出的 volumeCreditsFor 函数
func statementVolumeCreditsFor(invoice Invoice, plays map[string]Play) {
	var totalAmount = 0
	var volumeCredits = 0
	result := fmt.Sprintf("Statement for %s \n", invoice.Customer)
	var format = numberFormat
	for _, perf := range invoice.Performances {
		volumeCredits += volumeCreditsFor(perf, playFor(perf, plays))
		// print line for this order
		result += fmt.Sprintf(" %s: %s (%d seats)\n", playFor(perf, plays).Name,
			format(amountFor(perf, playFor(perf, plays))/100), perf.Audience)
		totalAmount += amountFor(perf, playFor(perf, plays))
	}
	result += fmt.Sprintf("Amount owed is %s\n", format(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits", volumeCredits)
	fmt.Println(result)
}

// 编译，测试，提交

// 6. 移除format变量
// func format -> func usd
func usd(aNumber int) string { // usd = United States Dollar
	return fmt.Sprintf("%d.00", aNumber)
}

func statementRemoveFormat(invoice Invoice, plays map[string]Play) {
	var totalAmount = 0
	var volumeCredits = 0
	result := fmt.Sprintf("Statement for %s \n", invoice.Customer)
	//var format = numberFormat  removed
	for _, perf := range invoice.Performances {
		volumeCredits += volumeCreditsFor(perf, playFor(perf, plays))
		// print line for this order
		result += fmt.Sprintf(" %s: %s (%d seats)\n", playFor(perf, plays).Name,
			usd(amountFor(perf, playFor(perf, plays))/100), perf.Audience)
		totalAmount += amountFor(perf, playFor(perf, plays))
	}
	result += fmt.Sprintf("Amount owed is %s\n", usd(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits", volumeCredits)
	fmt.Println(result)
}

// 7.移除观众量积分总和 volumeCredits(拆分循环)
func statementRemoveVolumeCredits(invoice Invoice, plays map[string]Play) {
	var totalAmount = 0
	//var volumeCredits = 0
	result := fmt.Sprintf("Statement for %s \n", invoice.Customer)
	for _, perf := range invoice.Performances {
		// print line for this order
		result += fmt.Sprintf(" %s: %s (%d seats)\n", playFor(perf, plays).Name,
			usd(amountFor(perf, playFor(perf, plays))/100), perf.Audience)
		totalAmount += amountFor(perf, playFor(perf, plays))
	}
	var volumeCredits = 0 // (移动语句)
	for _, perf := range invoice.Performances {
		volumeCredits += volumeCreditsFor(perf, playFor(perf, plays))
	}
	result += fmt.Sprintf("Amount owed is %s\n", usd(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits", volumeCredits)
	fmt.Println(result)
}

// 提炼出totalVolumeCredits函数
func totalVolumeCredits(invoice Invoice, plays map[string]Play) int {
	var result = 0 // (移动语句)
	for _, perf := range invoice.Performances {
		result += volumeCreditsFor(perf, playFor(perf, plays))
	}
	return result
}

// 提炼出totalAmount函数
func totalAmount(invoice Invoice, plays map[string]Play) int {
	var result = 0
	for _, perf := range invoice.Performances {
		// print line for this order
		result += amountFor(perf, playFor(perf, plays)) // inline variable
	}
	return result
}

// discuss: 关于拆分循环，一次循环变成三次循环可能带来的性能问题
// martin fowler的观点：大多数情况下可以忽略它。如果重构引入了性能损耗，先完成重构，再做性能优化
// 我的观点：基本赞同，但是要结合具体的业务场景进行分析。
// 这里其实折射出的一个问题是，有些情况下，可能需要在代码重构和代码性能之间进行取舍，但是代码性能的优先级应该大于重构的优先级
