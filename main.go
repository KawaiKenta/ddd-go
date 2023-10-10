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

	for i := 0; i < 3; i++ {
		_, err := userApplicationService.Register(fmt.Sprintf("test%d", i), fmt.Sprintf("test%d@gmail.com", i))
		if err != nil {
			panic(err)
		}
	}
	repo.ShowAll()
	println("--------------------")
	// 重複登録テスト、emailが重複するとエラーになる
	// user, err := userApplicationService.Register("test", "aaa@exadmple.com")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("user: %+v\n", user)
	// sameUser, err := userApplicationService.Register("test", "aaa@exadmple.com")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("user2: %+v\n", sameUser)

	// updateテスト 成功例
	// _, err := userApplicationService.Update(1, "updated_test1", "updatetest@gmail.com")
	// if err != nil {
	// 	panic(err)
	// }
	// repo.ShowAll()

	// updateテスト 失敗例
	// すでに存在するemailを指定するとエラーになる
	// _, err = userApplicationService.Update(2, "updated_test2", "updatetest@gmail.com")
	// if err != nil {
	// 	panic(err)
	// }
	// repo.ShowAll()

	// deleteテスト 成功例
	if err := userApplicationService.Delete(1); err != nil {
		panic(err)
	}
	repo.ShowAll()
	// deleteテスト 失敗例
	// 存在しないidを指定するとエラーになる
	if err := userApplicationService.Delete(4); err != nil {
		panic(err)
	}
}
