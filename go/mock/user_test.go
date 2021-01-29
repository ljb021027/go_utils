package mock

import (
	"errors"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
)

// 静态设置返回值
func TestReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserRepository(ctrl)
	// 期望FindOne(1)返回张三用户
	repo.EXPECT().FindOne(1).Return(&User{Name: "张三"}, nil)
	// 期望FindOne(2)返回李四用户
	repo.EXPECT().FindOne(2).Return(&User{Name: "李四"}, nil)
	// 期望给FindOne(3)返回找不到用户的错误
	repo.EXPECT().FindOne(3).Return(nil, errors.New("user not found"))
	// 验证一下结果
	log.Println(repo.FindOne(1)) // 这是张三
	log.Println(repo.FindOne(2)) // 这是李四
	log.Println(repo.FindOne(3)) // user not found
	log.Println(repo.FindOne(4)) //没有设置4的返回值，却执行了调用，测试不通过
}
