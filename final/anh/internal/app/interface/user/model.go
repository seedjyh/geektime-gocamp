package user

type BindId string

func (b BindId) String() string {
	return string(b)
}

type Number string

func (n Number) String() string {
	return string(n)
}

type UserId string

func (id UserId) String() string {
	return string(id)
}

type BindParameter struct {
	TelA Number
	TelX Number
	TelB Number
}

type BindDetail struct {
	TelA   Number
	TelX   Number
	TelB   Number
	BindId BindId
}
