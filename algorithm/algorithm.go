package algorithm

import (
	"container/list"
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

// ====================================最长无重复子串==================================
func getMax(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

// 最长无重复子串，基础版：双重循环+临时map暴力解法
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	dupMap := make(map[byte]bool)
	max := 1
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if _, ok := dupMap[s[j]]; ok {
				max = getMax(max, j-i)
				dupMap = make(map[byte]bool) // 重置map
				break
			} else {
				dupMap[s[j]] = true
				max = getMax(max, j-i+1)
			}
		}
	}
	return max
}

// 最长无重复子串，升级版：左滑动窗口法
func lengthOfLongestSubstringV2(s string) int {
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
		res = getMax(res, right+1-left)
	}

	return res
}

//====================================LRU缓存=====================================

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

func (l *LRUCache) remove(node *MyListNode) {
	pre := node.pre
	next := node.next

	pre.next = next
	next.pre = pre
	return
}

func (l *LRUCache) removeLast() {
	pre := l.tail.pre.pre
	pre.next = l.tail
	l.tail.pre = pre
	return
}

func (l *LRUCache) addToFront(node *MyListNode) {
	next := l.head.next
	l.head.next = node
	node.next = next

	next.pre = node
	node.pre = l.head
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

func (l *LRUCache) Get(key int) int {
	if v, ok := l.nodeMap[key]; ok {
		l.remove(v)
		l.addToFront(v)
		return v.val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	v, ok := l.nodeMap[key]
	if ok {
		// 更新值后，移动到头部
		v.val = value
		l.remove(v)
		l.addToFront(v)
		return
	}
	newNode := &MyListNode{}
	newNode.val = value
	newNode.key = key
	// 不存在，但是容量满了
	if len(l.nodeMap) >= l.cap {
		lastNode := l.tail.pre
		l.removeLast()
		delete(l.nodeMap, lastNode.key)

		l.addToFront(newNode)
		l.nodeMap[key] = newNode
		return
	}
	// 不存在，容量没满，直接添加
	l.addToFront(newNode)
	l.nodeMap[key] = newNode
	return
}

//=====================================有序map=====================================

type BidirectionalList struct {
	pre   *BidirectionalList
	next  *BidirectionalList
	key   string
	value interface{}
}

// OrderMap 有序map
type OrderMap struct {
	capacity int                           // 容量
	head     *BidirectionalList            // 双向链表的头部
	tail     *BidirectionalList            // 双向链表的尾部
	m        map[string]*BidirectionalList // map
}

func NewOrderMap(cap int) *OrderMap {
	head := &BidirectionalList{}
	tail := &BidirectionalList{}

	head.next = tail
	tail.pre = head

	return &OrderMap{
		capacity: cap,
		head:     head,
		tail:     tail,
		m:        make(map[string]*BidirectionalList, cap),
	}
}

// Put 添加一个key，存在则新增
func (o *OrderMap) Put(node *BidirectionalList) {
	v, ok := o.m[node.key]
	// 新增逻辑
	if !ok {
		// TODO 如果容量已满则需要先扩容再添加
		next := o.head.next
		o.head.next = node
		node.next = next

		node.pre = o.head
		next.pre = node

		o.m[node.key] = node
		return
	}
	// 更新逻辑
	v.value = node.value
	return
}

// Remove 删除一个key
func (o *OrderMap) Remove(key string) {
	v, ok := o.m[key]
	if !ok {
		return
	}
	pre := v.pre
	next := v.next

	pre.next = next
	next.pre = pre
	delete(o.m, key)
}

// Get 获取一个key
func (o *OrderMap) Get(key string) interface{} {
	if v, ok := o.m[key]; ok {
		return v
	}
	return nil
}

//================================== 堆排序 ==================================

func HeapSort(arr []int) {
	// 从下到上，从右往左，遍历所有非叶子节点，构建大顶堆
	// 最后一个非叶子节点的索引是：len(arr)/2 - 1
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}

	// 将堆顶元素与数组最后一个元素交换，然后重新构建大顶堆
	for j := len(arr) - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0]
		heapify(arr, 0, j)
	}
}

// 这个函数有两个作用：
// 一是递归零次，可以调节当前非叶子节点，使得左右子节点小于父节点；
// 二是递归多次，将当前节点的值放到合适的位置，使得大顶堆定义成立
// curNodeIdx: 当前父节点在数组中的下标
// length: 数组的长度
func heapify(arr []int, curNodeIdx, length int) {
	left := 2*curNodeIdx + 1
	// 递归终止条件，当前节点的左子节点不存在，说明后面已经没有要比较的元素了
	if left >= length {
		return
	}
	greater := left // 左右子节点较大值的索引暂定为left
	right := 2*curNodeIdx + 2
	if right < length && arr[greater] < arr[right] {
		greater = right
	}
	// 递归终止条件：此时说明当前节点的值已经大于左右子节点的值，到达了合适的位置，不用再递归下去了
	if arr[greater] < arr[curNodeIdx] {
		return
	}
	// 如果左右子节点中的较大值大于当前父节点的值，交换值
	if arr[greater] > arr[curNodeIdx] {
		arr[greater], arr[curNodeIdx] = arr[curNodeIdx], arr[greater]
	}
	// 走到这里说明上面的值交换逻辑一定执行到了，这个时候被交换下来的值(arr[greater])可能小于其左右子节点的值
	// 所以需要递归继续进行调整，直到到达了合适的位置
	heapify(arr, greater, length)
}

//================================== 二叉树的三种遍历 ==================================

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrderTraversal 二叉树的前序遍历(递归版)
func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		inOrder(node.Left)
		inOrder(node.Right)
	}
	inOrder(root)
	return res
}

// PreOrderTraversalV2 二叉树的前序遍历(非递归版)
func PreOrderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := list.New()
	// 循环退出条件：遍历指针为nil并且栈为空，栈为空说明已经回到了最上层
	for root != nil || stack.Len() > 0 {
		// 遍历根节点的左子树，记录根节点或者左节点的值后再将它们全部入栈
		for root != nil {
			res = append(res, root.Val) // 记录根节点或者左节点的值
			stack.PushFront(root)
			root = root.Left
		}
		// 弹出的栈顶元素有两种情况：根节点或者根节点的左节点
		// 这两种情况都不需要记录值，因为前面遍历的时候已经记录过了
		top := stack.Remove(stack.Front()).(*TreeNode)
		// 如果是根节点，并且有右节点，遍历指针指向右节点，转到右子树重复上面的步骤，
		if top.Right != nil {
			root = top.Right
		}
	}
	return res
}

// InOrderTraversal 二叉树的中序遍历(递归版)
func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		res = append(res, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	return res
}

// InOrderTraversalV2 二叉树的中序遍历(非递归版)
func InOrderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := list.New()
	// 循环退出条件：遍历指针为nil并且栈为空，栈为空说明已经回到了最上层
	for root != nil || stack.Len() > 0 {
		// 遍历根节点的左子树，将根节点和所有左节点都入栈
		for root != nil {
			stack.PushFront(root)
			root = root.Left
		}
		// 弹出的栈顶元素有两种情况：根节点或者根节点的左节点
		top := stack.Remove(stack.Front()).(*TreeNode)
		// 无论是根节点还是根节点的左节点，处理的操作是一样，都是记录元素的值
		res = append(res, top.Val)
		if top.Right != nil { // 如果存在右节点，说明是右节点，遍历指针指向右节点，转到右子树重复上面的过程
			root = top.Right
		}
	}
	return res
}

// PostOrderTraversal 二叉树后序遍历(递归版)
func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var postOrder func(root *TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		res = append(res, node.Val)
	}
	postOrder(root)
	return res
}

// PostOrderTraversalV2 二叉树后序遍历(迭代版)
func PostOrderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var prev *TreeNode
	var res []int
	stack := list.New()
	// 循环终止条件：遍历指针为nil或者栈为空，栈为空说明已经回到了最上层
	for root != nil || stack.Len() > 0 {
		// 遍历根节点的左子树，将根节点和所有左节点都入栈
		for root != nil {
			stack.PushFront(root)
			root = root.Left
		}
		// 走到头了，回到上一层：弹出栈顶元素
		// 这里弹出的栈顶元素有三种情况：第一次遇到的根节点，第二次遇到的根节点，根节点的左节点
		top := stack.Remove(stack.Front()).(*TreeNode)
		// 将根节点的左节点和第二次遇到的根节点这两种情况合并，将prev指向当前节点，并记录当前节点的值，然后回到上一层：弹出栈顶元素
		// top.Right == nil 说明是根节点的左节点
		// top.Right == prev 说明是第二次遇到的根节点
		if top.Right == nil || top.Right == prev {
			res = append(res, top.Val)
			prev = top
			root = nil // 遍历指针置为nil，继续弹出栈顶元素
		} else { // 第一次遇到的根节点，将根节点入栈，转到根节点的右子树，重复上面的步骤
			stack.PushFront(top)
			root = top.Right // 遍历指针指向右节点，转到根节点的右子树
		}
	}
	return res
}

// LevelOrderTraversal 层序遍历
func LevelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := list.New()   // 队列
	queue.PushFront(root) // 根节点先入栈，从第一层开始遍历
	for queue.Len() > 0 {
		// 因为每层遍历都是复用同一个队列，这里需要先记录每次循环的队列长度，此时的队列长度也就是上一层的节点个数
		n := queue.Len()
		var temp []int           // 保存上一层的值
		for j := 0; j < n; j++ { // 循环n次，刚好将上一层的节点全部取出来
			back := queue.Remove(queue.Back()).(*TreeNode)
			temp = append(temp, back.Val)
			// 对于出队列的每一个节点，依次将节点的左节点和右节点放入队列
			if back.Left != nil {
				queue.PushFront(back.Left)
			}
			if back.Right != nil {
				queue.PushFront(back.Right)
			}
		}
		// 保存每一层的结果
		res = append(res, temp)
	}
	return res
}

//================================== 二叉树中和为目标值的路径 ==================================

// PathTarget ...
func PathTarget(root *TreeNode, target int) [][]int {
	var path []int // 保存路径中每个节点的值
	var res [][]int
	dfs(root, target, path, &res)
	return res
}

func dfs(root *TreeNode, target int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	tempPath := make([]int, 0)
	for _, n := range path {
		tempPath = append(tempPath, n)
	}
	// 计算路径和，如果满足条件，就记录下来
	var sum int
	for _, num := range tempPath {
		sum += num
	}
	// 递归到叶子节点并且和等于target才记录结果
	if sum == target && root.Left == nil && root.Right == nil {
		*res = append(*res, tempPath)
	}
	dfs(root.Left, target, path, res)
	dfs(root.Right, target, path, res)
	// 本来这里应该要删除前面在path中添加的元素，方便回溯时不会对前面的路径产生干扰
	// 但是因为go的切片特性不需要处理
}
