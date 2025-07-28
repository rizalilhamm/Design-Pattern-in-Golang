package main

type User struct {
    ID    int
    Name  string
    Email string
}

type Validator interface {
    ValidateEmail(email string) error
}

type Repository interface {
    SaveUser(user *User) error
}

type Notifier interface {
    SendWelcomeEmail(user *User) error
}

func RegisterUser(user *User, v Validator, r Repository, n Notifier) error {
    if err := v.ValidateEmail(user.Email); err != nil {
        return err
    }

    if err := r.SaveUser(user); err != nil {
        return err
    }

    return n.SendWelcomeEmail(user)
}
