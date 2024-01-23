package algorithm

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"testing"
)

func Test_twoSum(t *testing.T) {

}

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				s: " ",
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				s: "abcabab",
			},
			want: 3,
		},
		{
			name: "case3",
			args: args{
				s: "abc",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRU(t *testing.T) {
	cache := Constructor(2)
	// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	// [null,null,null,1,null,-1,null,-1,3,4]
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}

func Test_ConcurrentPrintDataInOrder(t *testing.T) {
	ch := make(chan string, 3)
	count := 0
	printSlice := []string{"A", "B", "C"}
	wg := &sync.WaitGroup{}
	go func() {
		for value := range ch {
			fmt.Println(value)
			wg.Done()
			count++
		}
	}()
	for i := 0; i < 500; i++ {
		for _, v := range printSlice {
			wg.Add(1)
			ch <- v
		}
	}
	wg.Wait()
	close(ch)
	fmt.Println(count)
}

func TestOrderMap(t *testing.T) {
	orderMap := NewOrderMap(3)
	orderMap.Put(&BidirectionalList{
		key:   "1",
		value: 1,
	})
	orderMap.Put(&BidirectionalList{
		key:   "2",
		value: 2,
	})
	orderMap.Put(&BidirectionalList{
		key:   "3",
		value: 3,
	})
	fmt.Println(orderMap.Get("3").(*BidirectionalList).value)
	orderMap.Remove("3")
	fmt.Println(orderMap.Get("3"))

}

func Test_lengthOfLongestSubstringV2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringV2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valid case",
			args: args{
				arr: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.arr)
			// 打印一下arr
			for _, i := range tt.args.arr {
				fmt.Println(i)
			}
		})
	}
}

func TestPreOrderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	res := PreOrderTraversal(root) // want: 32451, got: 32451
	fmt.Println(res)
}

func TestPreOrderTraversalV2(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	res := PreOrderTraversalV2(root) // want: 32451, got: 32451
	fmt.Println(res)
}

func TestInOrderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	res := InOrderTraversal(root) // want: 42531
	fmt.Println(res)
}

func TestInOrderTraversalV2(t *testing.T) {
	root := &TreeNode{
		Val:  3,
		Left: nil,
		Right: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	res := InOrderTraversalV2(root) // want: 312
	fmt.Println(res)
}

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := LevelOrderTraversal(root)
	fmt.Println(res)
}

func TestPathTarget(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	res := PathTarget(root, 22)
	fmt.Println(res)
}

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case01",
			args: args{
				arr: []int{
					2, 5, 1, 23, 45, 1, 3, 7,
				},
			},
			want: []int{
				1, 1, 2, 3, 5, 7, 23, 45,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSortBoundary(t *testing.T) {
	var qSort func(arr []int, left, right int)
	qSort = func(arr []int, left, right int) {
		if left >= right {
			return
		}
		rIndex := left + rand.Int()%(right-left+1)
		arr[rIndex], arr[left] = arr[left], arr[rIndex]
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
		qSort(arr, left, i-1)
		qSort(arr, i+1, right)
	}
	nums := []int{3, 2, 1, 5, 6, 4}
	qSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func Test_lengthOfLongestSubstringV3(t *testing.T) {
	s := "tmmzuxt"
	fmt.Println(lengthOfLongestSubstringV3(s))
}

func TestReverseKGroup(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	newHead := ReverseKGroup(head, 2)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}

func TestThreeSum(t *testing.T) {
	nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	res := ThreeSum(nums)
	fmt.Println(res)

}

func TestPrintInOrder(t *testing.T) {
	PrintInOrder()
}

func TestPrintInOrderV2(t *testing.T) {
	PrintInOrderV2()
}

func TestReverseUrl(t *testing.T) {
	type args struct {
		url []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "valid case",
			args: args{
				url: []string{"w", "w", "w", ".", "1", "2", "3", "4", ".", "c", "n"},
			},
			want: []string{"c", "n", ".", "4", "3", "2", "1", ".", "w", "w", "w"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseUrl(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleFourOperation(t *testing.T) {
	type args struct {
		expression []string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "case 1",
			args: args{
				expression: []string{"1", "+", "2", "*", "3", "-", "4", "/", "5", "*", "6"},
			},
			want: 2.1999999999999993,
		},
		{
			name: "case 2",
			args: args{
				expression: []string{"2", "*", "3", "-", "1", "/", "3", "*", "6"},
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				expression: []string{"1", "*", "2", "*", "3"},
			},
			want: 6,
		},
		{
			name: "case 4",
			args: args{
				expression: []string{"1", "+", "2", "+", "3"},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleFourOperation(tt.args.expression); got != tt.want {
				t.Errorf("SimpleFourOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "valid case",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
