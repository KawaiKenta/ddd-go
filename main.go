package main

import (
	"fmt"

	applicationservice "github.com/KawaiKenta/ddd-go/application_service"
	domainservice "github.com/KawaiKenta/ddd-go/domain_service"
	"github.com/KawaiKenta/ddd-go/repository"
)

func main() {
	repo := repository.NewInMemoryUserRepository()
	userService := domainservice.NewUserService(repo)
	userApplicationService := applicationservice.NewUserApplicationService(repo, userService)

	for i := 0; i < 10; i++ {
		_, err := userApplicationService.Register(fmt.Sprintf("test%d", i), fmt.Sprintf("test%d@gmail.com", i))
		if err != nil {
			panic(err)
		}
	}
	repo.ShowAll()

	// 重複登録テスト、emailが重複するとエラーになる
	user, err := userApplicationService.Register("test", "aaa@exadmple.com")
	if err != nil {
		panic(err)
	}
	fmt.Printf("user: %+v\n", user)
	sameUser, err := userApplicationService.Register("test", "aaa@exadmple.com")
	if err != nil {
		panic(err)
	}
	fmt.Printf("user2: %+v\n", sameUser)
}
