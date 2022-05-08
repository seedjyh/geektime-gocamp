package user

import (
	"fmt"
	"regexp"
)

type BindId string

func (b BindId) String() string {
	return string(b)
}

type Number string

func (n Number) String() string {
	return string(n)
}

const msisdnRegex = "^((13[0-9])|(14[6-8])|(15([0-3]|[5-9]))|162|165|166|167|(17[0-8])|(18[0-9])|(19[0-3,5-9]))\\d{8}$"

func (n Number) IsMSISDN() bool {
	if len(n) != 11 {
		return false
	}
	if result, err := regexp.MatchString(msisdnRegex, string(n)); err != nil {
		return false
	} else {
		return result
	}
}

func (n Number) AssertValid() error {
	if !n.IsMSISDN() {
		return fmt.Errorf("[%+v] is not valid MSISDN", n)
	}
	return nil
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

func (p *BindParameter) String() string {
	return fmt.Sprintf("telA=[%+v], telX=[%+v], telB=[%+v]", p.TelA, p.TelX, p.TelB)
}

func (p *BindParameter) AssertValid() error {
	if err := p.TelA.AssertValid(); err != nil {
		return err
	}
	if err := p.TelX.AssertValid(); err != nil {
		return err
	}
	if err := p.TelB.AssertValid(); err != nil {
		return err
	}
	return nil
}

type BindDetail struct {
	TelA   Number
	TelX   Number
	TelB   Number
	BindId BindId
}
