package fundaccount

import (
	"time"
)

type FundAccount struct {
	FundAccountId int `sqlf:"autoincr"`
	Name          string
	Remark        string
	CreateTime    time.Time `sqlf:"created"`
	ModifyTime    time.Time `sqlf:"updated"`
}

type FundAccounts struct {
	Count int
	Data  []FundAccount
}
