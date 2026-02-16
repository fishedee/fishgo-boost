package unit

type UnitDb struct {
}

func NewUnitDb() IUnitDb {
	return &UnitDb{}
}

func (this *UnitDb) Search(where Unit) Units {
	return Units{}
}

func (this *UnitDb) Get(unitId int) Unit {
	return Unit{}
}

func (this *UnitDb) Add(unit Unit) {

}

func (this *UnitDb) Del(unitId int) {

}

func (this *UnitDb) Mod(unitId int, unit Unit) {

}
