package mediator

import (
	"fmt"
	"time"
)

//ChatRoom 聊天室类，当做中介者
type ChatRoom struct{}

var chatRoom = NewChatRoom()

//User 用户类
type User struct {
	Name string
}

//NewUser 实例化用户类
func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

//SendMessage 用户类使用中介者发送消息
func (u *User) SendMessage(msg string) {
	chatRoom.ShowMessage(u, msg)
}

//GetName 获取用于昵称
func (u *User) GetName() string {
	return u.Name
}

//NewChatRoom 实例化中介者
func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}

//ShowMessage 中介者发送消息
func (cr *ChatRoom) ShowMessage(user *User, msg string) {
	fmt.Printf("%s: [ %s ]: %s \n",
		time.Now().Format("2006-01-02 15:04:05"),
		user.GetName(),
		msg)
}
