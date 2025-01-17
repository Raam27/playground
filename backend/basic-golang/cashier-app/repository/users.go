package repository

import (
	"fmt"
	"strconv"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
<<<<<<< HEAD
	//beginanswer
	records, err := u.db.Load("users")
	if err != nil {
		records = [][]string{
			{"username", "password", "loggedin"},
		}
		if err := u.db.Save("users", records); err != nil {
			return nil, err
		}
	}

	result := make([]User, 0)
	for i := 1; i < len(records); i++ {
		loggedIn, err := strconv.ParseBool(records[i][2])
		if err != nil {
			return nil, err
		}

		user := User{
			Username: records[i][0],
			Password: records[i][1],
			Loggedin: loggedIn,
		}
		result = append(result, user)
	}

	return result, nil
	//endanswer return []User{}, nil
}

func (u *UserRepository) SelectAll() ([]User, error) {
	//beginanswer
	users, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}
	return users, nil
	//endanswer return []User{}, nil
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	//beginanswer

	if err := u.LogoutAll(); err != nil {
		return nil, err
	}

	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Username == username && user.Password == password {
			u.changeStatus(username, true)
			return &username, nil
		}
	}

	return nil, fmt.Errorf("Login Failed")
	//endanswer return nil, nil
}

func (u *UserRepository) FindLoggedinUser() (*string, error) {
	//beginanswer
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Loggedin {
			return &user.Username, nil
		}
	}

	return nil, fmt.Errorf("no user is logged in")
	//endanswer return nil, nil
}

func (u *UserRepository) Logout(username string) error {
	//beginanswer
	userLogin, err := u.FindLoggedinUser()
	if err != nil {
		return err
	}

	return u.changeStatus(*userLogin, false)
	//endanswer return nil
}

func (u *UserRepository) Save(users []User) error {
	//beginanswer
	records := [][]string{
		{"username", "password", "loggedin"},
	}
	for i := 0; i < len(users); i++ {
		records = append(records, []string{
			users[i].Username,
			users[i].Password,
			strconv.FormatBool(users[i].Loggedin),
		})
	}
	return u.db.Save("users", records)
	//endanswer return nil
=======
	return []User{}, nil // TODO: replace this
}

func (u *UserRepository) SelectAll() ([]User, error) {
	return []User{}, nil // TODO: replace this
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	return nil, nil // TODO: replace this
}

func (u *UserRepository) FindLoggedinUser() (*string, error) {
	return nil, nil // TODO: replace this
}

func (u *UserRepository) Logout(username string) error {
	return nil // TODO: replace this
}

func (u *UserRepository) Save(users []User) error {
	return nil // TODO: replace this
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

func (u *UserRepository) changeStatus(username string, status bool) error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

<<<<<<< HEAD
	//beginanswer
	//update
	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			users[i].Loggedin = status
			return u.Save(users)
		}
	}
	return fmt.Errorf("user not found")
	//endanswer return nil
}

func (u *UserRepository) LogoutAll() error {
	//beginanswer
	users, err := u.SelectAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		u.Logout(user.Username)
	}
	return nil
	//endanswer return nil
=======
	return nil // TODO: replace this
}

func (u *UserRepository) LogoutAll() error {
	return nil // TODO: replace this
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}
