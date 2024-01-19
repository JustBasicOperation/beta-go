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
		// 如果当前数字和前一个数字重复的话，直接跳过，这里要和前一位数比较
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
			// 必须在sum等于0的时候判断后面是否有重复数字并且直接跳到最后一个重复的数字
			// 不能提前判断是否有重复数字，否则会跳过像000这样的结果
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				// 如果三数之和等于0，直接跳到最后一个重复的数字，这里要和后一位数比较
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 如果三数之和等于0，直接跳到最后一个重复的数字，这里要和后一位数比较
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
	assistPtr := &ListNode{} // 定义一个辅助指针，指向head
	assistPtr.Next = head

	// 初始化双指针，前驱指针和后继指针都指向辅助指针
	// 前驱和后继指针的目的是为了将翻转后的链表重新插入正确的位置
	pre, end := assistPtr, assistPtr
	count := 0
	for end != nil {
		// end指针指向了待翻转链表的最后一个元素
		if count == k {
			next := end.Next  // 先记录下一组待翻转链表的第一个元素位置
			start := pre.Next // 确定本组待翻转链表的起始位置

			end.Next = nil // 断开连接后再去翻转链表
			pre.Next = nil // 断开连接后再去翻转链表

			// 翻转链表，翻转后，start虽然还是指向本组链表的第一个元素，但是需要放到原本end的位置
			newStart := reverse(start)

			pre.Next = newStart // 左边接上
			start.Next = next   // 右边接上

			// 重置pre，end指针和count计数器
			pre = start
			end = start
			count = 0
		} else {
			end = end.Next
			count++
		}
	}

	return assistPtr.Next
}

// 注意这个方法不能修改传入的head值
func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cursor := head
	for cursor != nil {
		next := cursor.Next
		cursor.Next = pre

		pre = cursor  // pre后移一个元素
		cursor = next // cursor后移一个元素
	}
	return pre // 遍历结束后pre指向了链表的最后一个元素
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
	// 定义dp[i]表示以nums[i]结尾的最长递增子序列
	// 状态转移方程：dp[i] = max(1, max(dp[j])+1)，其中0 <= j < i且nums[i]>nums[j]
	// 也就是往dp[0...i-1]中最长的上升子序列后面再加一个nums[i](需要保证nums[i]>nums[j])，dp[i]=dp[0...i-1]中的最大值+1
	dp := make([]int, len(nums), len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1 // 默认为自己，长度为1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = getMax(dp[i], dp[j]+1) // 这行代码可能会被执行多次，需要取其中最大的那次执行结果
			}
		}
	}
	var maxN int
	for i := 0; i < len(dp); i++ {
		if maxN < dp[i] {
			maxN = dp[i]
		}
	}
	return maxN
}

// LengthOfLISV2 解法二：贪心+二分
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
