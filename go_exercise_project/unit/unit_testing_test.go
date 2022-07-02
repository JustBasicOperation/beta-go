package unit

import (
	"errors"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/golang/mock/gomock"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func Test_checkUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid Username",
			args: args{
				username: "faustzhao",
			},
			want: true,
		},
		{
			name: "invalid Username",
			args: args{
				username: "张三",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkUsername(tt.args.username)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_checkEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid email",
			args: args{
				email: "faustzhao@tencent.com",
			},
			want: true,
		},
		{
			name: "invalid email",
			args: args{
				email: "tencent.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkEmail(tt.args.email)
			assert.Equal(t, tt.want, got)
		})
	}
}

// 使用go mock来mock接口：mockgen -destination ... -package ... -source ...
func Test_getPersonDetailRedis(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *PersonDetail
		wantErr bool
	}{
		{
			name:    "redis.Do err",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Unmarshal err",
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			want: &PersonDetail{
				Username: "steven",
				Email:    "123456@qq.com",
			},
			wantErr: false,
		},
	}
	// 生成控制器
	controller := gomock.NewController(t)
	defer controller.Finish()
	// 1.生成符合redis.Conn接口的mockConn
	mockConn := NewMockConn(controller)
	// 2.给接口打桩序列
	gomock.InOrder(
		mockConn.EXPECT().Do("GET", gomock.Any()).AnyTimes().Return("", errors.New("redis.Do err")),
		mockConn.EXPECT().Close().AnyTimes().Return(nil),
		mockConn.EXPECT().Do("GET", gomock.Any()).AnyTimes().Return("123", nil),
		mockConn.EXPECT().Close().AnyTimes().Return(nil),
		mockConn.EXPECT().Do("GET", gomock.Any()).AnyTimes().Return([]byte(`{"username":"steven","email":"123456@qq.com"}`), nil),
		mockConn.EXPECT().Close().AnyTimes().Return(nil),
	)
	// 3.给redis.Dail函数打桩
	cells := []gomonkey.OutputCell{
		{
			Values: gomonkey.Params{mockConn, nil},
			Times:  3,
		},
	}
	patches := gomonkey.ApplyFuncSeq(redis.Dial, cells)
	defer patches.Reset()
	// 4.断言
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPersonDetailRedis(tt.name)
			// assert.Equal()能够对结构体进行deep diff
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

// 为外部函数打一个桩序列
func TestGetPersonDetail(t *testing.T) {
	type args struct {
		username string
	}
	// 测试用例：主要针对函数的输入输出做设计
	tests := []struct {
		name    string
		args    args
		want    *PersonDetail
		wantErr bool
	}{
		{
			name: "invalid Username",
			args: args{
				username: "steven jam",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid email",
			args: args{
				username: "invalid_email",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "throw err",
			args: args{
				username: "throw_err",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid return",
			args: args{
				username: "steven",
			},
			want: &PersonDetail{
				Username: "steven",
				Email:    "123456@qq.com",
			},
			wantErr: false,
		},
	}
	// 为外部函数打桩序列，也就是这里的getPersonDetailRedis方法
	cells := []gomonkey.OutputCell{
		{
			Values: gomonkey.Params{
				&PersonDetail{
					Username: "invalid email",
					Email:    "test.com",
				},
				nil,
			},
		},
		{
			Values: gomonkey.Params{
				nil,
				errors.New("redis err"),
			},
		},
		{
			Values: gomonkey.Params{
				&PersonDetail{
					Username: "steven",
					Email:    "123456@qq.com",
				},
				nil,
			},
		},
	}
	//applyFunc := gomonkey.ApplyFunc(getPersonDetailRedis, func(username string) (*PersonDetail, error) {
	//	return &PersonDetail{
	//		Username: "steven",
	//		Email:    "123456@qq.com",
	//	}, nil
	//})
	patches := gomonkey.ApplyFuncSeq(getPersonDetailRedis, cells)
	// 执行完毕后释放桩序列
	//defer applyFunc.Reset()
	defer patches.Reset()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPersonDetail(tt.args.username)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

// 为一个全局变量打桩
func Test_max(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "case1",
			args: args{
				a: 1,
			},
			wantRes: 101,
		},
		{
			name: "case2",
			args: args{
				a: 10,
			},
			wantRes: 110,
		},
	}
	patches := gomonkey.ApplyGlobalVar(&b, 100)
	defer patches.Reset()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := max(tt.args.a); gotRes != tt.wantRes {
				t.Errorf("max() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

// 为一个函数变量打桩
func Test_square(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name       string
		args       args
		wantSquare int
	}{
		{
			name: "valid input",
			args: args{
				a: 10,
			},
			wantSquare: 100,
		},
	}
	patches := gomonkey.ApplyFuncVar(&Fun, func(a int) (square int) {
		return 100
	})
	defer patches.Reset()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSquare := square(tt.args.a); gotSquare != tt.wantSquare {
				t.Errorf("square() = %v, want %v", gotSquare, tt.wantSquare)
			}
		})
	}
}
