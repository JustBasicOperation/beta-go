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
