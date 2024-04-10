package algorithm

import (
	"container/list"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMax(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

// ====================================两数之和==================================
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

// ====================================三数之和==================================

// ThreeSum 三数之和，双指针解法
// 需要保证三个位置上都不出现重复的元素
func ThreeSum(nums []int) [][]int {
	var res [][]int
	length := len(nums)
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j] // 升序排序
	})
	for i := 0; i < length; i++ {
		// 如果当前数字和前一个数字重复的话，直接跳过
		// 这里必须和前面已经遍历过的数比较，否则会漏掉结果
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 如果nums[i]大于0，后面的组合不可能有0的情况了，直接break
		if nums[i] > 0 {
			break
		}
		left := i + 1
		right := length - 1
		for left < right {
			// 不能提前判断是否有重复数字，否则会漏掉结果
			// 必须在sum等于0的时候判断后面是否有重复数字，因为此时任何重复的left或者right都会导致结果重复
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				// 判断left是否和后面的数据重复，不能和前面的数比较，否则会导致left和i比较，跳过left和i相同的三元组
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 如果三数之和等于0，直接跳到最后一个重复的数字
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 记录结果后，左右指针分别向中间移动一位
				// 因为是排过序的，并且左右指针的元素不能重复，所以不能只移动左指针或者右指针，否则三数相加必然不等于零
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resList := &ListNode{}
	pHead := resList
	carry := 0
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

// 最长无重复子串，闭区间双指针+哈希表
func lengthOfLongestSubstringV3(s string) int {
	max, left := 0, 0
	indexMap := make(map[byte]int, len(s))
	for right := 0; right < len(s); right++ {
		cur := s[right]
		if v, ok := indexMap[cur]; ok {
			// 有重复元素，但是不在双指针范围内
			if v < left {
				max = getMax(max, right-left+1) // 双指针闭区间
			} else {
				// 有重复元素，但是在双指针范围内，左指针移动到重复元素的下一个位置
				left = v + 1
			}
			// 更新重复元素的下标
			indexMap[cur] = right
		} else {
			// 没有重复元素
			max = getMax(max, right-left+1) // 双指针闭区间
			// 新元素加入map，并记录下标
			indexMap[cur] = right
		}
	}
	return max
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
	// 如果已经存在，直接更新值并移动到头部，不用考虑容量是否超出
	if v, ok := l.nodeMap[key]; ok {
		// 更新值后，移动到头部
		v.val = value
		l.remove(v)
		l.addToFront(v)
		return
	}
	newNode := &MyListNode{
		key: key,
		val: value,
	}
	// 不存在，但是容量满了，先删除末尾元素，再添加
	if len(l.nodeMap) >= l.cap {
		lastNode := l.tail.pre
		l.removeLast()
		delete(l.nodeMap, lastNode.key)
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

//================================== 排序数组 ==================================

// HeapSort 堆排解法
func HeapSort(nums []int) []int {
	// 构建大顶堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, i, len(nums)-1)
	}
	// 循环取堆顶元素，放到数组末尾
	for j := len(nums) - 1; j > 0; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		heapify(nums, 0, j-1)
	}
	return nums
}

// 这个函数有两个作用：
// 一是递归零次，可以调节当前非叶子节点，使得左右子节点小于父节点；
// 二是递归多次，将当前节点的值放到合适的位置，使得大顶堆定义成立
// curNodeIdx: 当前父节点在数组中的下标
// length: 待排序数组的最后一个可用下标
func heapify(arr []int, root, length int) {
	// 递归终止条件：root已经到了待排序数组的最后
	if root >= length {
		return
	}
	left := root*2 + 1
	right := root*2 + 2
	maxIndex := root
	if left <= length && arr[left] > arr[maxIndex] {
		maxIndex = left
	}
	if right <= length && arr[right] > arr[maxIndex] {
		maxIndex = right
	}
	if maxIndex == root {
		return
	}
	arr[maxIndex], arr[root] = arr[root], arr[maxIndex]
	heapify(arr, maxIndex, length)
}

// QuickSort 常规快排(会超时)
func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, start, end int) {
	// 递归终止条件：分区内的元素小于等于1
	if start >= end {
		return
	}
	i, j := start, end
	pivot := arr[start] // 默认取第一个作为基准数
	for i < j {
		// 右指针向左移动，找到第一个比mid小的数
		for i < j && arr[j] > pivot {
			j--
		}
		// 左指针向右移动，找到第一个比mid大的数
		// 这里要用小于等于的原因是存在i==left的情况
		// 当i==left时需要跳过继续比较后面的数，防止arr[left]和arr[i]提前交换了位置
		for i < j && arr[i] <= pivot {
			i++
		}
		// 指针没有越界，并且左指针的值大于右指针的值，交换左右指针的值
		if i < j && arr[i] > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// 走到这里说明i==j，将mid与左指针交换
	if arr[start] > arr[i] {
		arr[start], arr[i] = arr[i], arr[start]
	}
	// 继续递归，处理后面的分区
	quickSort(arr, start, i-1)
	quickSort(arr, i+1, end)
}

// 优化版，不会超时
func quickSortV2(arr []int, start, end int) {
	// 递归终止条件，分区内元素小于等于1
	if start >= end {
		return
	}
	left, right := start, end

	midIdx := (rand.Int() % (right - start + 1)) + left // 随机选择基准数的下标
	arr[midIdx], arr[left] = arr[left], arr[midIdx]     // 将随机的基准数换到最左边，下面的逻辑保持不变

	mid := arr[left]
	for left < right {
		// 从右往左找到一个小于mid的数的下标
		for left < right && arr[right] > mid {
			right--
		}
		// 从左往右找到一个大于mid的数的下标
		for left < right && arr[left] <= mid {
			left++
		}
		// 走到这里说明找到了或者指针相遇了，交换左右指针的值
		arr[left], arr[right] = arr[right], arr[left]
	}
	// 走到这里说明指针相遇了，将mid和指针所在位置交换，使得左边的数都小于mid，右边的数都大于mid
	arr[left], arr[start] = arr[start], arr[left]

	// 优化下中轴范围
	for left > 0 && arr[left] == arr[left-1] {
		left--
	}
	for right < len(arr)-1 && arr[right] == arr[right+1] {
		right++
	}

	// 递归处理左分区和右分区
	quickSort(arr, start, left-1)
	quickSort(arr, right+1, end)
}

//================================== 二叉树的三种遍历 ==================================

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
		// 遍历根节点的左子树，将根节点和所有左节点都入栈
		for root != nil {
			res = append(res, root.Val) // 记录根节点或者左节点的值
			stack.PushFront(root)
			root = root.Left
		}
		// 弹出的栈顶元素有两种情况：有右子树的左节点，没有右子树的左节点
		// 这两种情况都不需要记录值，因为前面遍历的时候已经记录过了
		top := stack.Remove(stack.Front()).(*TreeNode)
		// 如果有右子树，转到右子树重复上面的步骤，
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
		// 弹出的栈顶元素有两种情况：有右子树的左节点和没有右子树的左节点
		top := stack.Remove(stack.Front()).(*TreeNode)
		// 无论左节点有没有右子树，处理的操作都是一样，都是记录元素的值
		res = append(res, top.Val)
		// 如果存在右子树，转到右子树重复上面的过程
		if top.Right != nil {
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
		// 这里弹出的栈顶元素有三种情况：有右子树的左节点(第一次遇到)，有右子树的左节点(第二次遇到)，左节点
		top := stack.Remove(stack.Front()).(*TreeNode)
		// 将根节点的左节点和第二次遇到的根节点这两种情况合并，将prev指向当前节点，并记录当前节点的值，然后回到上一层：弹出栈顶元素
		// top.Right == nil 没有右子树的左节点
		// top.Right == prev 第二次遇有右子树的左节点
		if top.Right == nil || top.Right == prev {
			res = append(res, top.Val)
			prev = top
		} else { // 第一次遇到有右子树的左节点，将节点入栈，转到节点的右子树，重复上面的步骤
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
	pathTargetDFS(root, target, path, &res)
	return res
}

func pathTargetDFS(root *TreeNode, target int, path []int, res *[][]int) {
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
	pathTargetDFS(root.Left, target, path, res)
	pathTargetDFS(root.Right, target, path, res)
	// 本来这里应该要删除前面在path中添加的元素，方便回溯时不会对前面的路径产生干扰
	// 但是因为go的切片特性不需要处理
}

//================================== 数组中第K个最大元素 ==================================

// FindKthLargest 快排解法
func FindKthLargest(arr []int, k int) int {
	qSort(arr, 0, len(arr)-1, len(arr)-k)
	return arr[len(arr)-k]
}

func qSort(arr []int, left, right, target int) {
	if left >= right {
		return
	}
	i, j := left, right
	mid := arr[left]
	for i < j {
		for i < j && arr[j] > mid {
			j--
		}
		for i < j && arr[i] <= mid {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[left] = arr[left], arr[i]
	if i == target {
		return
	}
	if i > target {
		qSort(arr, left, i-1, target)
	}
	if i < target {
		qSort(arr, i+1, right, target)
	}
}

//=================================== k个一组翻转链表 =======================================

func ReverseKGroup(head *ListNode, k int) *ListNode {
	// 整体思路：定义两个指针，prev和head
	// prev为辅助指针，指向已翻转链表的最后一个节点，
	// head为遍历指针依次遍历链表的每个节点，直到为nil
	// 翻转时需要记录三个指针变量：start, end, next
	// start和end分别为本次待翻转链表段的起始和结束节点
	// next为下一段未翻转链表的起始节点
	// 翻转时，先将待翻转链表段和原链表断开，翻转结束后再重新连接，并依次更新四个指针的值

	// 在链表头前增加一个辅助节点
	assistPtr := &ListNode{}
	assistPtr.Next = head

	prev := assistPtr
	for head != nil {
		for i := 0; i < k-1; i++ {
			head = head.Next
			// 如果head为nil，说明待翻转链表节点数少于k个，直接退出本次循环
			if head == nil {
				return assistPtr.Next
			}
		}
		start := prev.Next
		end := head
		next := head.Next
		// 解开链表，进行翻转
		end.Next = nil
		prev.Next = nil
		reverse(start)
		// 翻转后重新连接
		prev.Next = end
		start.Next = next
		// 更新辅助指针和遍历指针的值
		prev = start
		head = next
	}
	return assistPtr.Next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head // 这里需要注意，翻转链表时不要修改传入节点head的值
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

//=================================== 最大子数组和 =======================================

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 解法思路：动态规划
	// max[i]表示开始到当前位置的最大子序和，根据前一位的最大子序和来判断是继续累加还是另起一段区间
	// 状态转移方程：max[i] = MAX(max[i-1]+nums[i], nums[i])
	max := nums[0]
	preMax := nums[0]
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			curMax := getMax(preMax+nums[i], nums[i])
			preMax = curMax
			max = getMax(max, curMax)
		}
	}
	return max
}

// MaxSubArrayV2 解法二：线段树分治求解
func MaxSubArrayV2(nums []int) int {
	return get(nums, 0, len(nums)).mSum
}

type Status struct {
	lSum int // 左端点开始的最大子段和
	rSum int // 右端点开始的最大子段和
	mSum int // 区间内的最大子段和
	aSum int // 区间和
}

func get(nums []int, left, right int) Status {
	if left == right {
		return Status{lSum: nums[left], rSum: nums[left], mSum: nums[left], aSum: nums[left]}
	}
	mid := (left + right) / 2
	lStatus := get(nums, left, mid)
	rStatus := get(nums, mid+1, right)
	// push up操作：分别求lSum,rSum,mSum,aSum
	aSum := lStatus.aSum + rStatus.aSum                     // 左右子区间aSum相加
	lSum := getMax(lStatus.lSum, lStatus.aSum+rStatus.lSum) // max(左子区间的lSum,左子区间的aSum+右子区间的lSum)
	rSum := getMax(rStatus.rSum, rStatus.aSum+lStatus.rSum) // max(右子区间的rSum,右子区间的aSum+左子区间的rSum)
	mSum := getMax(getMax(lStatus.mSum, rStatus.mSum), lStatus.rSum+rStatus.lSum)
	return Status{lSum: lSum, rSum: rSum, mSum: mSum, aSum: aSum}
}

//=================================== 最长递增子序列 =======================================

// LengthOfLIS 动态规划解法
func LengthOfLIS(nums []int) int {
	// 定义dp[i]为以nums[i]结尾的最长递增子序列，nums[i]必须被选取
	// dp[i] = max(dp[j])+1，如果nums[i] > nums[j]，0 <= j <= i-1
	// 初始化条件，dp[0] = 1
	dp := make([]int, len(nums), len(nums))
	dp[0] = 1
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		tempMax := 0
		for j := i - 1; j >= 0; j-- {
			// 找到满足nums[i] > nums[j]条件的所有dp[j]中的最大值
			if nums[i] > nums[j] {
				tempMax = getMax(dp[j], tempMax)
			}
		}
		dp[i] = tempMax + 1
		// 这里取整个dp数组的最大值
		// 并不是序列越长，最长递增子序列的长度就越长，也就是dp数组不是递增的
		// 因为每加入一个新的元素，且新加入的元素必须被选取，会导致最长递增子序列发生变化，可能变长，也可能变短
		maxLen = getMax(maxLen, dp[i])
	}
	return maxLen
}

// LengthOfLISV2 解法二：贪心+二分
// 也就是每次让最长递增子序列上升得尽可能慢，最后遍历完所有元素留下的就是最长递增子序列
func LengthOfLISV2(nums []int) int {
	// 定义d[i]为长度为i的最长递增子序列的最后一个元素的最小值
	d := make([]int, len(nums)+1, len(nums)+1)
	length := 1
	d[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > d[length] {
			d[length+1] = nums[i]
			length++
		} else {
			// 使用二分查找，在d[0...length]中找到第一个比nums[i]小的数d[j]，并更新d[j+1] = nums[i]
			l, r, pos := 1, length, 0
			for l <= r {
				mid := (l + r) / 2
				if d[mid] < nums[i] { // 当l==r并且d[mid]<nums[i]时，找到了第一个比nums[i]小的d[mid]
					pos = mid
					l = mid + 1
				} else if d[mid] == nums[i] {
					r = mid - 1
				} else {
					r = mid - 1
				}
			}
			// 如果pos为零说明所有数都比nums[i]大，此时要将d[1]设置为nums[i]
			d[pos+1] = nums[i]
		}
	}
	return length
}

//===================================== 多线程按序打印123 =========================================

func PrintInOrder() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		ch3 <- 1
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-ch3
			fmt.Println(1)
			ch1 <- 1
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-ch1
			fmt.Println(2)
			ch2 <- 2
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-ch2
			fmt.Println(3)
			ch3 <- 3
		}
	}()
	time.Sleep(3 * time.Second)
}

func PrintInOrderV2() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-ch
			fmt.Println(1)
			ch <- 1
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-ch
			fmt.Println(2)
			ch <- 2
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-ch
			fmt.Println(3)
			ch <- 3
		}
	}()
	time.Sleep(3 * time.Second)
}

//========================================= 反转url =========================================

func ReverseUrl(url []string) []string {
	reverseStr(url, 0, len(url)-1)
	fmt.Println(url)
	pre := -1
	for i := 0; i < len(url); i++ {
		if pre == -1 && url[i] == "." {
			reverseStr(url, 0, i-1)
			pre = i - 1
		}
	}
	return url
}

func reverseStr(str []string, start, end int) []string {
	for start < end {
		str[start], str[end] = str[end], str[start]
		start++
		end--
	}
	return str
}

//========================================= 最长递增子序列二 =========================================

// LengthOfLISII dp解法会超时
func LengthOfLISII(nums []int, k int) int {
	// 定义dp[i]为以nums[i]结尾的最长递增子序列
	// dp[i] = max(dp[j]+1, 1) (0 <= j < i) if nums[i] - nums[j] <= k
	dp := make([]int, len(nums), len(nums))
	maxN := 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = 1
			continue
		}
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && nums[i]-nums[j] <= k {
				dp[i] = getMax(dp[i], dp[j]+1)
			}
		}
		maxN = getMax(maxN, dp[i])
	}
	return maxN
}

//========================================= 反转字符串中的单词 =========================================

func ReverseWords(s string) string {
	byteArr := trimSpaceAndToArr(s)
	// 先反转字符串
	reverseArr(byteArr, 0, len(byteArr)-1)
	// 反转单词
	start := 0
	for i := 0; i < len(byteArr); i++ {
		// 遇到空格传入开始和结束位置的下标，闭区间
		if byteArr[i] == ' ' {
			reverseArr(byteArr, start, i-1)
			// 重置开始指针下标
			start = i + 1
		}
	}
	// 反转最后一个单词
	reverseArr(byteArr, start, len(byteArr)-1)
	return string(byteArr)
}

// 字符串转数组并去除多余的空格
func trimSpaceAndToArr(s string) []byte {
	var res []byte
	for i := 0; i < len(s); i++ {
		// 处理第一个字符就是空格的情况
		if i == 0 && s[i] == ' ' {
			res = append(res, s[i])
			continue
		}
		// 去除重复的空格，只保留第一个
		if s[i] == ' ' && s[i-1] == ' ' {
			continue
		}
		res = append(res, s[i])
	}
	// 处理开头和结尾的空格
	left, right := 0, len(res)-1
	if res[0] == ' ' {
		left++
	}
	if res[right] == ' ' {
		right--
	}
	return res[left : right+1]
}

// 反转数组
func reverseArr(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return
}

//========================================= 最长公共子序列 =========================================

func LongestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j]表示text[1...i]和text2[1...j]的最长公共子序列
	// dp[i][j] = max(dp[i-1][j],dp[i][j-1]) text1[i] != text2[j]
	// dp[i][j] = dp[i-1][j-1]+1 text[i] == text2[j]，这里不需要考虑dp[i-1][j]和dp[i][j-1]转移过去的情况了
	// dp[i][0] = 0, dp[0][j] = 0
	dp := make([][]int, len(text1)+1, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1, len(text2)+1)
	}
	for i := 0; i <= len(text1); i++ {
		for j := 0; j <= len(text2); j++ {
			if i == 0 {
				dp[i][j] = 0
				continue
			}
			if j == 0 {
				dp[i][j] = 0
				continue
			}
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = getMax(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

//========================================= 编辑距离 =========================================

func MinDistance(word1 string, word2 string) int {
	// 定义dp[i][j]表示word1的前i个字符和word2的前j个字符的最短编辑距离
	// 此时可以认为word1[1...i]和word2[1...j已经完全相同
	// 后续只需要考虑三种情况即可：word1加一个字符，word2加一个字符，word1和word2都加一个字符
	// if word2[i] == word2[j] dp[i][j] = min(dp[i-1][j]+1,dp[i][j-1]+1),dp[i-1][j-1]
	// if word1[i] != word2[j] dp[i][j] = min(dp[i-1][j]+1,dp[i][j-1]+1),dp[i-1][j-1]+1)
	// 初始状态：dp[0][j] = j,dp[i][0] = i
	len1 := len(word1)
	len2 := len(word2)
	dp := make([][]int, len1+1, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1, len2+1)
	}

	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j == 0 {
				dp[i][j] = i
				continue
			}
			if word1[i-1] == word2[j-1] {
				dp[i][j] = getMin(getMin(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1])
			} else {
				dp[i][j] = getMin(getMin(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func getMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//========================================= 合并两个有序链表 =========================================

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	cursor := res
	for list1 != nil && list2 != nil {
		node1 := list1.Val
		node2 := list2.Val
		// 每次从两个链表中取较小的节点
		if node1 < node2 {
			cursor.Next = list1
			list1 = list1.Next
		} else {
			cursor.Next = list2
			list2 = list2.Next
		}
		cursor = cursor.Next
	}
	for list1 != nil {
		cursor.Next = list1
		list1 = list1.Next
		cursor = cursor.Next
	}
	for list2 != nil {
		cursor.Next = list2
		list2 = list2.Next
		cursor = cursor.Next
	}
	return res.Next
}

//========================================= 简单四则运算 =========================================

type SymbolType int

var (
	unknownSymbol SymbolType = 0
	addSymbol     SymbolType = 1
	subSymbol     SymbolType = 2
	mulSymbol     SymbolType = 3
	divSymbol     SymbolType = 4
)

// SimpleFourOperation 简单四则运算，
// 解题的关键在于，每次遇到符号都判断一下前一个符号的优先级，高的话先计算前一个符号，再将当前符号入栈
func SimpleFourOperation(expression []string) float64 {
	symbolStack := list.New()
	numStack := list.New()
	for i := 0; i < len(expression); i++ {
		operator := expression[i]
		num, err := strconv.ParseFloat(operator, 64)
		if err != nil {
			// 遇到了符号
			st := parseSymbol(operator)
			if symbolStack.Len() > 0 {
				// 如果此时栈顶的符号是*或/，从数字栈中弹出两个数字进行计算
				// 即每次遇到符号，都判断一下此时栈顶是否有高优先级的符号，保证最终留在栈中的符号都是+或-
				topType := symbolStack.Remove(symbolStack.Front())
				if topType == mulSymbol {
					nextNum := numStack.Remove(numStack.Front()).(float64)
					preNum := numStack.Remove(numStack.Front()).(float64)
					res := preNum * nextNum
					numStack.PushFront(res)
				} else if topType == divSymbol {
					nextNum := numStack.Remove(numStack.Front()).(float64)
					preNum := numStack.Remove(numStack.Front()).(float64)
					res := preNum / nextNum
					numStack.PushFront(res)
				} else {
					// 栈顶符号不是*或/，将栈顶符号重新放回
					symbolStack.PushFront(topType)
				}
			}
			// 符号入栈
			symbolStack.PushFront(st)
		} else {
			// 遇到了数字
			numStack.PushFront(num)
		}
	}
	// 处理剩下的符号，注意剩下符号的第一个可能为*或者/，后面的都是+或者-
	for symbolStack.Len() > 0 {
		topType := symbolStack.Remove(symbolStack.Front())
		next := numStack.Remove(numStack.Front()).(float64)
		pre := numStack.Remove(numStack.Front()).(float64)
		if topType == addSymbol {
			res := pre + next
			numStack.PushFront(res)
		} else if topType == subSymbol {
			res := pre - next
			numStack.PushFront(res)
		} else if topType == mulSymbol {
			res := pre * next
			numStack.PushFront(res)
		} else {
			res := pre / next
			numStack.PushFront(res)
		}
	}
	return numStack.Remove(numStack.Front()).(float64)
}

func parseSymbol(s string) SymbolType {
	if s == "+" {
		return addSymbol
	}
	if s == "-" {
		return subSymbol
	}
	if s == "*" {
		return mulSymbol
	}
	if s == "/" {
		return divSymbol
	}
	return unknownSymbol
}

//========================================= 二分查找：等值查询，左边界和右边界 =========================================

// BSLeftBoundary 二分查找左边界(包含相等)：1 3 5 <7 9 11>，target = 7, x = 7
// 在递增数组nums中，找到一个数x >= target，下标为pos，使得nums[left,pos) < target，nums[pos,right] >= target
// 即x左边的数都小于target，x及其右边的数都大于等于target，所以叫做左边界查询
// 返回x的下标pos，如果没找到，返回-1
// 有两种极端情况需要考虑：nums中所有数都小于target，此时返回-1，nums中所有数都大于target，此时返回0，即最初的左边界
func BSLeftBoundary(nums []int, left, right, target int) int {
	rightBoundary := right
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 走到这里说明left > right，检查下left有没有越界，或者没找到的情况
	if left > rightBoundary || !(nums[left] >= target) {
		return -1
	}
	return left
}

// BSRightBoundary 二分查找右边界(包含相等)：<1 3 5> 7 9 11，target = 5, x = 5
// 在递增数组nums中，找到最后一个数x <= target，下标为pos，使得nums[left,pos] <= target，nums(pos,right) > target
// 即x及其左边的数都小于target，x右边的数都大于等于target，所以叫做右边界查询
// 返回x的下标pos，如果没找到，返回-1
// 这里也有两种极端情况需要考虑：数组nums中所有数都小于target，此时返回最初的右边界right，数组中所有数都大于target，此时返回-1
func BSRightBoundary(nums []int, left, right, target int) int {
	leftBoundary := left
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 走到这里说明left > right，检查下right有没有越界，或者没找到的情况
	if right < leftBoundary || !(nums[right] <= target) {
		return -1
	}
	return right
}

//============================================= 最长回文子串 ==================================================

// LongestPalindrome 最长回文子串，扩散法
// 定义左右指针，左右指针之间的子串是回文串
func LongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	// 定义nums[i]为以s[i]为中心的最长回文子串
	maxStr := ""
	for i := 1; i < len(s); i++ {
		tempStr := ""
		// 先向左扩散：识别aaaa或aaa这种情况，a为奇数个和偶数个都是回文串
		l := i - 1
		for {
			if l >= 0 && s[l] == s[i] {
				l--
			} else {
				break
			}
		}
		// 再向右扩散：识别aaaa或aaa这种情况，b为奇数个和偶数个都是回文串
		r := i + 1
		for {
			if r < len(s) && s[r] == s[i] {
				r++
			} else {
				break
			}
		}
		// 走到这里说明l和r之间的子串是回文串，有三种可能的情况，分别是一个a, 奇数个aa和偶数个aaa
		// 向两边扩散
		for l >= 0 && r < len(s) {
			if s[l] == s[r] {
				l--
				r++
			} else {
				break
			}
		}
		tempStr = s[l+1 : r]
		if len(tempStr) > len(maxStr) {
			maxStr = tempStr
		}
	}
	return maxStr
}

//============================================= 全排列 ==================================================

func Permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums), len(nums))
	permuteDFS(nums, []int{}, used, &res)
	return res
}

func permuteDFS(nums, path []int, used []bool, res *[][]int) {
	if len(path) == len(nums) {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		permuteDFS(nums, path, used, res)
		// 撤销操作
		path = path[:len(path)-1]
		used[i] = false
	}
}

//============================================= 全排列II ==================================================

func PermuteUnique(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums), len(nums))
	permute2DFS(nums, []int{}, used, &res)
	return res
}

func permute2DFS(nums, path []int, used []bool, res *[][]int) {
	if len(path) == len(nums) {
		*res = append(*res, append([]int{}, path...))
		return
	}
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		// 剪枝，之前已经遍历过的元素直接跳过，防止重复
		if _, ok := m[nums[i]]; ok {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		m[nums[i]] = struct{}{}
		permute2DFS(nums, path, used, res)
		// 撤销操作
		path = path[:len(path)-1]
		used[i] = false
	}
}

//============================================= 跳跃游戏 ==================================================

// CanJump 跳跃游戏
func CanJump(nums []int) bool {
	var max int
	for i := 0; i < len(nums); i++ {
		// 当max大于当前位置时才可以执行下面的逻辑
		if max >= i {
			temp := i + nums[i]
			if temp > max {
				max = temp
			}
			if max >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}

//============================================= 跳跃游戏II ==================================================

// JumpII 贪心反向查找解法（时间复杂度高）
func JumpII(nums []int) int {
	pos := len(nums) - 1
	var step int
	for pos > 0 {
		// 从左往右找前一个能到达pos位置的起跳点，直到pos位置等于0
		// 为什么要从左往右：前一个起跳点越接近下标0，才能使得跳跃次数越少，贪心思想
		// 找到了就更新pos的位置，并且增加跳跃次数
		for i := 0; i < pos; i++ {
			if i+nums[i] >= pos {
				pos = i
				step++
			}
		}
	}
	return step
}

// JumpIIV2 贪心正向查找解法
func JumpIIV2(nums []int) int {
	var start int
	var end int
	var step int
	// 注意这里的循环条件为end<len(nums-1)，只需要到倒数第二个位置即可
	// 因为每次循环结束，下标已经跳到了下一个区间，如果到最后一个区间才结束循环会多跳一次
	for end < len(nums)-1 {
		var maxPos int
		// 遍历区间[start,end]，找到最大的起跳点并更新能够跳到的最远距离
		for i := start; i <= end; i++ {
			// 更新能够跳到的最远距离
			maxPos = getMax(maxPos, i+nums[i])
		}
		start = end + 1 // 下一个区间的开始下标
		end = maxPos    // 下一个区间的结束下标
		step++          // 增加一次跳跃次数，因为从上一个区间跳到了下一个区间
	}
	return step
}

//============================================= 买卖股票的最佳时机 ==================================================

func MaxProfit(prices []int) int {
	// 本解法本质上是动态规划解法的空间优化版
	// 定义dp[i]为前i天的最低价格，dp[i] = min(dp[i-1], prices[i])
	// 第i天的最大利润就是第i天的价格减去前i天的最低价格
	// profit[i] = max(profit[i-1], prices[i]-dp[i])
	var minPrices, maxProfit int
	minPrices = prices[0]
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			continue
		}
		if prices[i] < minPrices {
			minPrices = prices[i]
		} else {
			maxProfit = getMax(maxProfit, prices[i]-minPrices)
		}
	}
	return maxProfit
}

//============================================= 买卖股票的最佳时机II ==================================================

func maxProfitII(prices []int) int {
	var max int
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			continue
		}
		sub := prices[i] - prices[i-1]
		if sub > 0 {
			max += sub
		}
	}
	return max
}

//============================================= H指数 ==================================================

// HIndex H指数
func HIndex(citations []int) int {
	sort.Ints(citations) // 升序排序
	var h int
	n := len(citations)
	// 从后往前遍历
	for i := n - 1; i >= 0; i-- {
		// 如果citations[i] >= n-i，则说明citations[i]及其后面的共n-i个数都大于等于n-i，此时n-i就是新的h指数
		// 更新h的值，不断向前遍历，找到最大的h。
		if citations[i] >= n-i {
			h = n - i
		}
	}
	return h
}

//============================================= O(1) 时间插入、删除和获取随机元素 ==============================================

type RandomizedSet struct {
	m    map[int]int
	nums []int
}

func NewRandomizedSet() RandomizedSet {
	return RandomizedSet{
		m:    map[int]int{},
		nums: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.m[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; ok {
		// 将数组中最后一个元素放到被删除元素的位置
		// 同时更新最后一个元素在map中的下标
		idx := this.m[val]
		lastEle := this.nums[len(this.nums)-1]
		this.nums[idx] = lastEle
		this.m[lastEle] = idx

		this.nums = this.nums[:len(this.nums)-1] // 移除数组最后一个元素
		delete(this.m, val)                      // 删除map中的被删除元素
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	r := rand.Intn(len(this.nums))
	return this.nums[r]
}

//============================================= 除自身以外数组的乘积 ==============================================

// ProductExceptSelf
func ProductExceptSelf(nums []int) []int {
	// left[i]表示nums[i]左侧所有元素的乘积，不包括nums[i]
	left := make([]int, len(nums), len(nums))
	// right[j]表示nums[j]右侧所有元素的乘积，包括nums[j]
	right := make([]int, len(nums), len(nums))
	left[0] = 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}
		left[i] = left[i-1] * nums[i-1]
	}
	right[len(nums)-1] = 1
	for j := len(nums) - 1; j >= 0; j-- {
		if j == len(nums)-1 {
			continue
		}
		right[j] = right[j+1] * nums[j+1]
	}
	res := make([]int, len(nums), len(nums))
	for k := 0; k < len(nums); k++ {
		res[k] = left[k] * right[k]
	}
	return res
}
