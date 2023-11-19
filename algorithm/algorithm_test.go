package algorithm

import (
	"fmt"
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
