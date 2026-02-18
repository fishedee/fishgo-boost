package unit

import (
	. "github.com/fishedee/fishgo-boost/app/proxy"
	"time"
)

type IUnitAo interface {
	Add(unit Unit)
	Del(unitId int)
	Get(unitId int) Unit
	Mod(unitId int, unit Unit)
	Search(where Unit, t time.Time) Units
}

type UnitAoMock struct {
	AddHandler    func(unit Unit)
	DelHandler    func(unitId int)
	GetHandler    func(unitId int) Unit
	ModHandler    func(unitId int, unit Unit)
	SearchHandler func(where Unit, t time.Time) Units
}

func (this *UnitAoMock) Add(unit Unit) {
	this.AddHandler(unit)
}

func (this *UnitAoMock) Del(unitId int) {
	this.DelHandler(unitId)
}

func (this *UnitAoMock) Get(unitId int) Unit {
	return this.GetHandler(unitId)
}

func (this *UnitAoMock) Mod(unitId int, unit Unit) {
	this.ModHandler(unitId, unit)
}

func (this *UnitAoMock) Search(where Unit, t time.Time) Units {
	return this.SearchHandler(where, t)
}

type IUnitDb interface {
	Add(unit Unit)
	Del(unitId int)
	Get(unitId int) Unit
	Mod(unitId int, unit Unit)
	Search(where Unit) Units
}

type UnitDbMock struct {
	AddHandler    func(unit Unit)
	DelHandler    func(unitId int)
	GetHandler    func(unitId int) Unit
	ModHandler    func(unitId int, unit Unit)
	SearchHandler func(where Unit) Units
}

func (this *UnitDbMock) Add(unit Unit) {
	this.AddHandler(unit)
}

func (this *UnitDbMock) Del(unitId int) {
	this.DelHandler(unitId)
}

func (this *UnitDbMock) Get(unitId int) Unit {
	return this.GetHandler(unitId)
}

func (this *UnitDbMock) Mod(unitId int, unit Unit) {
	this.ModHandler(unitId, unit)
}

func (this *UnitDbMock) Search(where Unit) Units {
	return this.SearchHandler(where)
}

func init() {
	RegisterProxyMock([]IUnitAo{&UnitAoMock{}})
	RegisterProxyMock([]IUnitDb{&UnitDbMock{}})
}
