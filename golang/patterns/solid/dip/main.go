package main

import (
	"fmt"
)

//Bad code

// type MySQL struct {
// }

// func (db MySQL) QueryMySQL() []string {
// 	return []string{"alex", "john", "mike"}
// }

// type PostgreSQL struct {
// }

// func (db PostgreSQL) QueryPostgreSQL() map[string]string {
// 	return map[string]string{
// 		"a3f69c2b-d153-48fd-b10c-5b641657477b": "alex",
// 		"a4f69c2b-d153-48fd-b10c-5b641657477a": "john",
// 		"a5f69c2b-d153-48fd-b10c-5b641657477c": "mike",
// 	}
// }

// type UsersRepository struct {
// 	db MySQL
// 	//db PostgreSQL
// }

// func (r UsersRepository) GetUsers() []string {
// 	res := r.db.QueryMySQL()
// 	//res := r.db.QueryPostgreSQL()
// 	//var users []string
// 	//for _, u := range res {
// 	//	users = append(users, u)
// 	//}
// 	//return users
// 	return res
// }
// func main() {
// 	mysqlDB := MySQL{}
// 	//postgreSQLDB := PostgreSQL{}
// 	repo := UsersRepository{db: mysqlDB}
// 	//repo := UsersRepository{db: postgreSQLDB}
// 	fmt.Println(repo.GetUsers())
// }

//Good code

type MySQL struct {
}

func (db MySQL) Query() interface{} {
	return []string{"alex", "john", "mike"}
}

type PostgreSQL struct {
}

func (db PostgreSQL) Query() interface{} {
	return map[string]string{
		"a3f69c2b-d153-48fd-b10c-5b641657477b": "alex",
		"a4f69c2b-d153-48fd-b10c-5b641657477a": "john",
		"a5f69c2b-d153-48fd-b10c-5b641657477c": "mike",
	}
}

type DBConn interface {
	Query() interface{}
}

type UsersRepository struct {
	db DBConn
}

func (r UsersRepository) GetUsers() []string {
	var users []string
	res := r.db.Query()

	switch res.(type) {
	case map[string]string:
		for _, u := range res.(map[string]string) {
			users = append(users, u)
		}
		return users
	case []string:
		return res.([]string)
	}

	return []string{}
}

func main() {
	mysqlDB := MySQL{}
	postgreSQLDB := PostgreSQL{}
	repo1 := UsersRepository{db: mysqlDB}
	repo2 := UsersRepository{db: postgreSQLDB}
	fmt.Println(repo1.GetUsers())
	fmt.Println(repo2.GetUsers())
}
