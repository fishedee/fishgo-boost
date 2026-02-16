package fundaccount

type FundAccountDb struct {
}

func NewFundAccountDb() IFundAccountDb {
	return &FundAccountDb{}
}

func (this *FundAccountDb) Search(where FundAccount) FundAccounts {
	return FundAccounts{}
}

func (this *FundAccountDb) Get(fundAccountId int) FundAccount {
	return FundAccount{}
}

func (this *FundAccountDb) Add(fundAccount FundAccount) {

}

func (this *FundAccountDb) Del(fundAccountId int) {

}

func (this *FundAccountDb) Mod(fundAccountId int, unit FundAccount) {

}
