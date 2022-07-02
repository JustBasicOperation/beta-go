package stage03

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

type statementData struct {
	Customer      string
	Performances  []NewPerformance
	TotalAmount   int
	VolumeCredits int
}

type NewPerformance struct {
	Play     Play
	PlayID   string
	Audience int
}

func statement(invoice Invoice, plays map[string]Play) {
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
	result += fmt.Sprintf("You earned %d credits\n", data.VolumeCredits)
	fmt.Println(result)
}

func usd(aNumber int) string { // usd = United States Dollar
	return fmt.Sprintf("%d.00", aNumber)
}

// 完成需求，打印HTML格式的账单
func htmlStatement(invoice Invoice, plays map[string]Play) {
	renderHtmlText(createStatementData(invoice, plays))
}

func renderHtmlText(data statementData) {
	// TODO ...
}
