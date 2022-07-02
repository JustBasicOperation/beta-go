package stage03

// createStatementData 填充中转数据
func createStatementData(invoice Invoice, plays map[string]Play) statementData {
	newPerformances := make([]NewPerformance, 0, len(invoice.Performances))
	for _, perf := range invoice.Performances {
		newPerformances = append(newPerformances, EnrichPerformance(perf, plays))
	}
	return statementData{
		Customer:      invoice.Customer,
		Performances:  newPerformances,
		TotalAmount:   TotalAmount(invoice, plays),
		VolumeCredits: TotalVolumeCredits(invoice, plays),
	}
}

// EnrichPerformance 填充Performance
func EnrichPerformance(aPerformance Performance, plays map[string]Play) NewPerformance {
	return NewPerformance{
		Play:     PlayFor(aPerformance, plays),
		PlayID:   aPerformance.PlayID,
		Audience: aPerformance.Audience,
	}
}

// TotalVolumeCredits 提炼出totalVolumeCredits函数
func TotalVolumeCredits(invoice Invoice, plays map[string]Play) int {
	var result = 0 // (移动语句)
	for _, perf := range invoice.Performances {
		result += VolumeCreditsFor(perf, plays)
	}
	return result
}

// TotalAmount 提炼出totalAmount函数
func TotalAmount(invoice Invoice, plays map[string]Play) int {
	var result = 0
	for _, perf := range invoice.Performances {
		// print line for this order
		result += AmountFor(perf, PlayFor(perf, plays)) // inline variable
	}
	return result
}

func AmountFor(aPerformance Performance, play Play) int {
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

func PlayFor(aPerformance Performance, plays map[string]Play) Play {
	return plays[aPerformance.PlayID]
}

// VolumeCreditsFor ...
func VolumeCreditsFor(aPerformance Performance, plays map[string]Play) int {
	result := 0
	if aPerformance.Audience-30 > 0 {
		result += aPerformance.Audience - 30
	}
	if PlayFor(aPerformance, plays).Type == "comedy" {
		result += aPerformance.Audience / 5
	}
	return result
}

// -----------------------------------第三阶段 按类型重组计算过程(以多态取代条件表达式)----------------------------------------

// 第二个需求，除了喜剧和悲剧之外，还增加了其他剧种，且拥有不同的计算规则

// Calculator ...
type Calculator interface {
	amount() int
}

// PerformanceCalculator 构造演出计算器
type PerformanceCalculator struct {
	aPerformance Performance
	aPlay        Play
}

func (p *PerformanceCalculator) amount() int {
	switch p.aPlay.Type { // inline variable
	case "tragedy":
		t := &TragedyCalculator{p: p}
		return t.amount()
	case "comedy":
		c := &ComedyCalculator{p: p}
		return c.amount()
	default:
		return 0
	}
}

// TragedyCalculator ...
type TragedyCalculator struct {
	p *PerformanceCalculator
}

func (t *TragedyCalculator) amount() int {
	result := 40000
	if t.p.aPerformance.Audience > 30 {
		result += 1000 * (t.p.aPerformance.Audience - 30)
	}
	return result
}

// ComedyCalculator ...
type ComedyCalculator struct {
	p *PerformanceCalculator
}

func (c *ComedyCalculator) amount() int {
	result := 30000
	if c.p.aPerformance.Audience > 20 {
		result += 10000 + 500*(c.p.aPerformance.Audience-20)
	}
	result += 300 * c.p.aPerformance.Audience
	return result
}
