package unit

import "time"

type UnitAo struct {
	unitDb IUnitDb
}

func NewUnitAo(unitDb IUnitDb) IUnitAo {
	return &UnitAo{
		unitDb: unitDb,
	}
}

func (this *UnitAo) Search(where Unit, t time.Time) Units {
	return Units{}
}

func (this *UnitAo) Get(unitId int) Unit {
	return Unit{}
}

func (this *UnitAo) Add(unit Unit) {

}

func (this *UnitAo) Del(unitId int) {

}

func (this *UnitAo) Mod(unitId int, unit Unit) {

}
