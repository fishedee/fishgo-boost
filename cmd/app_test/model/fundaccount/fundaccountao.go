package fundaccount

import (
	. "github.com/fishedee/fishgo-boost/cmd/app_test/model/unit"
)

type FundAccountAo struct {
}

func NewFundAccountAo() IFundAccountAo {
	return &FundAccountAo{}
}

func (this *FundAccountAo) Search(where FundAccount, where2 Unit) FundAccounts {
	return FundAccounts{}
}

func (this *FundAccountAo) Get(fundAccountId int) FundAccount {
	return FundAccount{}
}

func (this *FundAccountAo) Add(fundAccount FundAccount) {

}

func (this *FundAccountAo) Del(fundAccountId int) {

}

func (this *FundAccountAo) Mod(fundAccountId int, unit FundAccount) {

}
