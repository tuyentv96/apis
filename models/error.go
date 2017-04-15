package model


type Error struct {
	Messeage string	`json:"message"`
}

type Erro struct {
	Messeage string	`json:"message"`
}

func Err(errIn int) Error {

	switch errIn {
	case 100:
		return Error{Messeage: "wrong format"}
		break
	case 104:
		return Error{Messeage: "user not found"}
		break

	case 200:
		return Error{Messeage: "success"}
		break

	case 401:
		return Error{Messeage: "token is expired"}
		break

	case 410:
		return Error{Messeage: "wrong password"}
		break


	default:
		return Error{Messeage: "invalid error"}
		break
	}

	return Error{Messeage: " invalid error"}


}
