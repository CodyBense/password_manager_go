package app

type Login struct {
    website     string
    username    string
    password    string
}

func NewLogin(website, username, password string) Login {
    return Login{website: website, username: username, password: password}
}
