package model

type Registr struct{
	ID int
	Login string
	Password string
	ConfirmPassword string
}


type SignIn struct{
	ID int
	Login string
	Password string
}

type RegLogError struct{
	Login string
	Error string
	ErrorPass string
}


type SignInError struct{
	Login string
	ErrorLog string
	ErrorPass string
}
