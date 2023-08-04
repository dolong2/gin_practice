package repository

import (
	"gin_practice/src/practice/gin/domain/user/entity"
	"log"
)

var repo = make(map[string]entity.User)

func Save(user entity.User) {
	log.Println(user.Email)
	repo[user.Email] = user
}

func DeleteByEmail(email string) {
	delete(repo, email)
}

func GetByEmail(email string) (entity.User, bool) {
	user, success := repo[email]
	return user, success
}

func ExistsByEmail(email string) bool {
	_, exist := repo[email]
	return exist
}
