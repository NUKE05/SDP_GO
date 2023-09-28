package main

import "fmt"

const (
	// Константы для аутендификации в систему
	correctUsernamePassword = "user-password"
	correctSMSPassword      = "correct_sms"
	correctEDCToken         = "123456" // EDC - это Electronic Digital Signature, либо просто аналог ЭЦП ключа
)

type AuthenticationStrategy interface {
	Authenticate(credentials string) bool
}

type UsernamePasswordAuthentication struct{}

func (upa *UsernamePasswordAuthentication) Authenticate(credentials string) bool {
	return credentials == correctUsernamePassword
}

type SocialLoginAuthentication struct{}

func (sla *SocialLoginAuthentication) Authenticate(credentials string) bool {
	return credentials == correctSMSPassword
}

type MultiFactorAuthentication struct{}

func (mfa *MultiFactorAuthentication) Authenticate(credentials string) bool {
	return credentials == correctEDCToken
}

type AuthenticationService struct {
	strategy AuthenticationStrategy
}

func (as *AuthenticationService) Login(credentials string) bool {
	return as.strategy.Authenticate(credentials)
}

func main() {
	var credentials string // Переменная хранящая данные вводимые пользователем для аутендификации
	var aut_choice int     // Переменная хранящая выбор метода аутендификации пользователя

	fmt.Print("Выберите метод аутендификации (1 для Пользователь-Пароль, 2 для СМС логина, 3 для использования ЭЦП Ключа): ")
	fmt.Scanln(&aut_choice)

	var authStrategy AuthenticationStrategy

	if aut_choice == 1 {
		authStrategy = &UsernamePasswordAuthentication{}
	} else if aut_choice == 2 {
		authStrategy = &SocialLoginAuthentication{}
	} else if aut_choice == 3 {
		authStrategy = &MultiFactorAuthentication{}
	} else {
		fmt.Println("Вы выбрали неправильный метод аутендификации.")
		return
	}

	fmt.Print("Введите ваши данные: ")
	fmt.Scanln(&credentials)

	authService := AuthenticationService{strategy: authStrategy}
	isAuthenticated := authService.Login(credentials)

	if isAuthenticated {
		fmt.Println("Успешная аутендификация!")
	} else {
		fmt.Println("Аутентификация не удалась.")
	}
}
