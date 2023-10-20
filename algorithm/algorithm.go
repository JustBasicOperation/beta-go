package algorithm

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		sub := target - n
		if v, ok := m[n]; ok {
			return []int{v, i}
		}
		m[sub] = i
	}
	return []int{}
}

func threeSum(nums []int) [][]int {
	var res [][]int
	// 特判
	if len(nums) < 3 {
		return res
	}
	// 排序
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if nums[0] > 0 {
		return res
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			if left < right {
				sum := nums[i] + nums[left] + nums[right]
				if sum == 0 {
					res = append(res, []int{nums[i], nums[left], nums[right]})
					left++
					right--
				} else if sum > 0 {
					right--
				} else {
					left++
				}
			} else {
				break
			}
		}
	}
	return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resList := &ListNode{}
	pHead := resList
	carry := 0
	// the characteristic of last list node is ptr != nil && ptr.Next = nil
	for l1 != nil && l2 != nil {
		v1 := l1.Val
		v2 := l2.Val
		var sum int
		if carry > 0 {
			sum = v1 + v2 + carry
			carry = 0
		} else {
			sum = v1 + v2
		}
		carry = sum / 10
		pHead.Next = &ListNode{
			Val: sum % 10,
		}
		pHead = pHead.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		var sum int
		pHead.Next = &ListNode{}
		if carry > 0 {
			sum = l1.Val + carry
			carry = 0
		} else {
			sum = l1.Val
		}
		carry = sum / 10
		pHead.Next.Val = sum % 10
		pHead = pHead.Next
		l1 = l1.Next
	}
	for l2 != nil {
		var sum int
		pHead.Next = &ListNode{}
		if carry > 0 {
			sum = l2.Val + carry
			carry = 0
		} else {
			sum = l2.Val
		}
		carry = sum / 10
		pHead.Next.Val = sum % 10
		pHead = pHead.Next
		l2 = l2.Next
	}
	if carry > 0 {
		pHead.Next = &ListNode{
			Val: carry,
		}
	}
	return resList.Next
}

func lengthOfLongestSubstring(s string) int {
	right, res := -1, 0
	tmpMap := make(map[byte]bool, len(s))
	for left := 0; left < len(s); left++ {
		if left != 0 {
			delete(tmpMap, s[left-1])
		}
		for right+1 < len(s) && !tmpMap[s[right+1]] {
			tmpMap[s[right+1]] = true
			right++
		}
		res = max(res, right+1-left)
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

type LRUCache struct {
	cap     int
	head    *MyListNode
	tail    *MyListNode
	nodeMap map[int]*MyListNode
}

type MyListNode struct {
	pre  *MyListNode
	next *MyListNode
	val  int
	key  int
}

func (this *LRUCache) remove(node *MyListNode) {
	pre := node.pre
	next := node.next

	pre.next = next
	next.pre = pre
	return
}

func (this *LRUCache) removeLast() {
	pre := this.tail.pre.pre
	pre.next = this.tail
	this.tail.pre = pre
	return
}

func (this *LRUCache) addToFront(node *MyListNode) {
	next := this.head.next
	this.head.next = node
	node.next = next

	next.pre = node
	node.pre = this.head
}

func Constructor(capacity int) LRUCache {
	h := &MyListNode{}
	t := &MyListNode{}

	h.next = t
	t.pre = h
	res := LRUCache{
		cap:     capacity,
		head:    h,
		tail:    t,
		nodeMap: make(map[int]*MyListNode, 0),
	}
	return res
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.nodeMap[key]; ok {
		this.remove(v)
		this.addToFront(v)
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.nodeMap[key]
	if ok {
		// 更新值后，移动到头部
		v.val = value
		this.remove(v)
		this.addToFront(v)
		return
	}
	newNode := &MyListNode{}
	newNode.val = value
	newNode.key = key
	// 不存在，但是容量满了
	if len(this.nodeMap) >= this.cap {
		lastNode := this.tail.pre
		this.removeLast()
		delete(this.nodeMap, lastNode.key)

		this.addToFront(newNode)
		this.nodeMap[key] = newNode
		return
	}
	// 不存在，容量没满，直接添加
	this.addToFront(newNode)
	this.nodeMap[key] = newNode
	return
}
