package main

import "fmt"

type User struct {
	Name string
	ID string
	Age int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price int
}

func main() {
	user := User{ "송하나", "hana", 23 }
	vip := VIPUser{
		User { "화랑", "hwarang", 40 },
		3,
		2350, // 여러줄 초기화 할시 마지막 값 뒤에 쉼표 달기
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s skdl %d VIP 레벨: %d VIP 가격: %d 만 원\n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
		vip.Price,
	)
}
