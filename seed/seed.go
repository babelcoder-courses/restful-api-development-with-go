package seed

import (
	"course-go/config"
	"course-go/migrations"
	"course-go/models"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/bxcodec/faker/v3"
)

func Load() {
	db := config.GetDB()

	// Clean Database
	db.Migrator().DropTable("users", "articles", "categories", "migrations")
	migrations.Migrate()

	// Add Admin
	fmt.Println("Creating admin...")

	admin := models.User{
		Email:    "admin@babelcoder.com",
		Password: "passw0rd",
		Name:     "Admin",
		Role:     "Admin",
		Avatar:   "https://i.pravatar.cc/100",
	}

	admin.Password = admin.GenerateEncryptedPassword()
	db.Create(&admin)

	// Add normal users
	fmt.Println("Creating users...")

	numOfUsers := 50
	users := make([]models.User, 0, numOfUsers)
	userRoles := [2]string{"Editor", "Member"}

	for i := 1; i <= numOfUsers; i++ {
		user := models.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: "passw0rd",
			Avatar:   "https://i.pravatar.cc/100?" + strconv.Itoa(i),
			Role:     userRoles[rand.Intn(2)],
		}

		user.Password = user.GenerateEncryptedPassword()
		db.Create(&user)
		users = append(users, user)
	}

	// Add categories
	fmt.Println("Creating categories...")

	numOfCategories := 20
	categories := make([]models.Category, 0, numOfCategories)

	for i := 1; i <= numOfCategories; i++ {
		category := models.Category{
			Name: faker.Word(),
			Desc: faker.Paragraph(),
		}

		db.Create(&category)
		categories = append(categories, category)
	}

	// Add articles
	fmt.Println("Creating articles...")

	numOfArticles := 50
	articles := make([]models.Article, 0, numOfArticles)

	for i := 1; i <= numOfArticles; i++ {
		article := models.Article{
			Title:      faker.Sentence(),
			Excerpt:    faker.Sentence(),
			Body:       faker.Paragraph(),
			Image:      "https://source.unsplash.com/random/300x200?" + strconv.Itoa(i),
			CategoryID: uint(rand.Intn(numOfCategories) + 1),
			UserID:     uint(rand.Intn(numOfUsers) + 1),
		}

		db.Create(&article)
		articles = append(articles, article)
	}
}
