package fundaccount

import (
	. "github.com/fishedee/fishgo-boost/app/proxy"
	. "github.com/fishedee/fishgo-boost/cmd/app_test/model/unit"
)

type IFundAccountAo interface {
	Add(fundAccount FundAccount)
	Del(fundAccountId int)
	Get(fundAccountId int) FundAccount
	Mod(fundAccountId int, unit FundAccount)
	Search(where FundAccount, where2 Unit) FundAccounts
}

type FundAccountAoMock struct {
	AddHandler    func(fundAccount FundAccount)
	DelHandler    func(fundAccountId int)
	GetHandler    func(fundAccountId int) FundAccount
	ModHandler    func(fundAccountId int, unit FundAccount)
	SearchHandler func(where FundAccount, where2 Unit) FundAccounts
}

func (this *FundAccountAoMock) Add(fundAccount FundAccount) {
	this.AddHandler(fundAccount)
}

func (this *FundAccountAoMock) Del(fundAccountId int) {
	this.DelHandler(fundAccountId)
}

func (this *FundAccountAoMock) Get(fundAccountId int) FundAccount {
	return this.GetHandler(fundAccountId)
}

func (this *FundAccountAoMock) Mod(fundAccountId int, unit FundAccount) {
	this.ModHandler(fundAccountId, unit)
}

func (this *FundAccountAoMock) Search(where FundAccount, where2 Unit) FundAccounts {
	return this.SearchHandler(where, where2)
}

type IFundAccountDb interface {
	Add(fundAccount FundAccount)
	Del(fundAccountId int)
	Get(fundAccountId int) FundAccount
	Mod(fundAccountId int, unit FundAccount)
	Search(where FundAccount) FundAccounts
}

type FundAccountDbMock struct {
	AddHandler    func(fundAccount FundAccount)
	DelHandler    func(fundAccountId int)
	GetHandler    func(fundAccountId int) FundAccount
	ModHandler    func(fundAccountId int, unit FundAccount)
	SearchHandler func(where FundAccount) FundAccounts
}

func (this *FundAccountDbMock) Add(fundAccount FundAccount) {
	this.AddHandler(fundAccount)
}

func (this *FundAccountDbMock) Del(fundAccountId int) {
	this.DelHandler(fundAccountId)
}

func (this *FundAccountDbMock) Get(fundAccountId int) FundAccount {
	return this.GetHandler(fundAccountId)
}

func (this *FundAccountDbMock) Mod(fundAccountId int, unit FundAccount) {
	this.ModHandler(fundAccountId, unit)
}

func (this *FundAccountDbMock) Search(where FundAccount) FundAccounts {
	return this.SearchHandler(where)
}

func init() {
	RegisterProxyMock([]IFundAccountAo{&FundAccountAoMock{}})
	RegisterProxyMock([]IFundAccountDb{&FundAccountDbMock{}})
}
