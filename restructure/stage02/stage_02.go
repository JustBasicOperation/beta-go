package main

import (
	"fmt"
)

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

var invoice = Invoice{
	Customer:     "BigGo",
	Performances: performances,
}

func main() {
	statementStage01(invoice, plays)
	statementStage02(invoice, plays)
}

// StatementStage01 第一阶段重构后的代码全貌
func statementStage01(invoice Invoice, plays map[string]Play) {
	result := fmt.Sprintf("Statement for %s \n", invoice.Customer)
	for _, perf := range invoice.Performances {
		// print line for this order
		result += fmt.Sprintf(" %s: %s (%d seats)\n", playFor(perf, plays).Name,
			usd(amountFor(perf, playFor(perf, plays))/100), perf.Audience)
	}
	result += fmt.Sprintf("Amount owed is %s\n", usd(totalAmount(invoice, plays)/100))
	result += fmt.Sprintf("You earned %d credits\n", totalVolumeCredits(invoice, plays))
	fmt.Println(result)
}

func playFor(aPerformance Performance, plays map[string]Play) Play {
	return plays[aPerformance.PlayID]
}

func usd(aNumber int) string { // usd = United States Dollar
	return fmt.Sprintf("%d.00", aNumber)
}

func amountFor(aPerformance Performance, play Play) int {
	result := 0
	switch play.Type { // inline variable
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

func totalAmount(invoice Invoice, plays map[string]Play) int {
	var result = 0
	for _, perf := range invoice.Performances {
		// print line for this order
		result += amountFor(perf, playFor(perf, plays)) // inline variable
	}
	return result
}

func totalVolumeCredits(invoice Invoice, plays map[string]Play) int {
	var result = 0 // (移动语句)
	for _, perf := range invoice.Performances {
		result += volumeCreditsFor(perf, plays)
	}
	return result
}

func volumeCreditsFor(aPerformance Performance, plays map[string]Play) int {
	result := 0
	if aPerformance.Audience-30 > 0 {
		result += aPerformance.Audience - 30
	}
	if playFor(aPerformance, plays).Type == "comedy" {
		result += aPerformance.Audience / 5
	}
	return result
}

// --------------------------------------------第二阶段 拆分计算阶段和格式化阶段---------------------------------------------
// 创建中转数据结构
// 格式化只需要知道五个参数
// customer, playName, audience, totalAmount, totalVolumeCredits
type statementData struct {
	Customer      string
	Performances  []newPerformance
	TotalAmount   int
	VolumeCredits int
}

type newPerformance struct {
	Play     Play
	PlayID   string
	Audience int
}

// 提炼出renderPlainText函数
func statementStage02(invoice Invoice, plays map[string]Play) {
	renderPlainText(createStatementData(invoice, plays))
}

// 对传入的数据进行格式化
func renderPlainText(data statementData) {
	result := fmt.Sprintf("Statement for %s \n", data.Customer)
	for _, perf := range data.Performances {
		result += fmt.Sprintf(" %s: %s (%d seats)\n",
			perf.Play.Name,
			usd(data.TotalAmount/100),
			perf.Audience)
	}
	result += fmt.Sprintf("Amount owed is %s\n", usd(data.TotalAmount/100))
	result += fmt.Sprintf("You earned %d credits", data.VolumeCredits)
	fmt.Println(result)
}

// 填充中转数据
func createStatementData(invoice Invoice, plays map[string]Play) statementData {
	newPerformances := make([]newPerformance, 0, len(invoice.Performances))
	for _, perf := range invoice.Performances {
		newPerformances = append(newPerformances, enrichPerformance(perf, plays))
	}
	return statementData{
		Customer:      invoice.Customer,
		Performances:  newPerformances,
		TotalAmount:   totalAmount(invoice, plays),
		VolumeCredits: totalVolumeCredits(invoice, plays),
	}
}

// 填充Performance
func enrichPerformance(aPerformance Performance, plays map[string]Play) newPerformance {
	return newPerformance{
		Play:     playFor(aPerformance, plays),
		PlayID:   aPerformance.PlayID,
		Audience: aPerformance.Audience,
	}
}

// 打印HTML格式的账单
func htmlStatement(invoice Invoice, plays map[string]Play) {
	renderHtmlText(createStatementData(invoice, plays))
}

func renderHtmlText(data statementData) {
	// TODO ...
}

// 代码分离，将格式化和计算的代码分别移动到两个文件中
// 营地法则：保证你离开时的代码库一定比来时更健康
// 好代码的检验标准就是人们是否能轻而易举地修改它
