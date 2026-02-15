package testdata

import (
	"github.com/fishedee/fishgo-boost/language"
	"time"
)

func queryColumnMap_12cb9c765f30b6d40fa9080910c80cea2cce66f7(data interface{}, column string) interface{} {
	dataIn := data.([]QueryInnerStruct2)
	result := make(map[int]QueryInnerStruct2, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		single := dataIn[i]
		result[single.QueryInnerStruct.MM] = single
	}
	return result
}

func queryColumnMap_163b58b970f87da0d1c228cced6790dd60ad2896(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make(map[string]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		single := dataIn[i]
		result[single.Name] = single
	}
	return result
}

func queryColumnMap_2d7e0ce966e7e5fb1bd62167c50fc81590fd3a1e(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make(map[int]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		single := dataIn[i]
		result[single.Age] = single
	}
	return result
}

func queryColumnMap_3923b792e276005e09637544ecb3aec8be870f41(data interface{}, column string) interface{} {
	dataIn := data.([]string)
	result := make(map[string]string, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		single := dataIn[i]
		result[single] = single
	}
	return result
}

func queryColumnMap_440473c7ae32877de566d281340e92c7a42fa3c2(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make(map[float64]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		single := dataIn[i]
		result[single.CardMoney] = single
	}
	return result
}

func queryColumnMap_621f96588f04633a85c32b0497fdd5b7cfaaebdf(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	result := make(map[string][]ContentType, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Name
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		result[bufferData[kbegin].Name] = bufferData[kbegin:k]

	}
	return result
}

func queryColumnMap_63871bbce37249e280e427591a85c168607bcfb1(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make(map[string]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		single := dataIn[i]
		result[single.Name] = single
	}
	return result
}

func queryColumnMap_704d73f9c17886c384e0a16be36b7b5959fea51f(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[bool]int, len(dataIn))
	result := make(map[bool][]ContentType, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Ok
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		result[bufferData[kbegin].Ok] = bufferData[kbegin:k]

	}
	return result
}

func queryColumnMap_725664359c0503f05a9971dca3064bc1d143d4fb(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make(map[bool]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		single := dataIn[i]
		result[single.Ok] = single
	}
	return result
}

func queryColumnMap_77cc920402a2e411fd707de6c0797eecc2947f6e(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[float32]int, len(dataIn))
	result := make(map[float32][]ContentType, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Money
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		result[bufferData[kbegin].Money] = bufferData[kbegin:k]

	}
	return result
}

func queryColumnMap_91dacd60e87431951940b4b4c51428e7c1e5c1f2(data interface{}, column string) interface{} {
	dataIn := data.([]int)
	result := make(map[int]int, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		single := dataIn[i]
		result[single] = single
	}
	return result
}

func queryColumnMap_bd864bf11acf9e4deb95377e4098c24bd32a1308(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[float64]int, len(dataIn))
	result := make(map[float64][]ContentType, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].CardMoney
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		result[bufferData[kbegin].CardMoney] = bufferData[kbegin:k]

	}
	return result
}

func queryColumnMap_ca2b7b23bd1c78dcf7d146a73399ab1744521b59(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make(map[float32]ContentType, len(dataIn))

	for i := len(dataIn) - 1; i >= 0; i-- {
		single := dataIn[i]
		result[single.Money] = single
	}
	return result
}

func queryColumnMap_da12d428eb4ff4f1c8585a24fead914b37d902b8(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	result := make(map[string][]ContentType, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Name
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		result[bufferData[kbegin].Name] = bufferData[kbegin:k]

	}
	return result
}

func queryColumnMap_dd430d90e19ca65443db28bfbf3eb2336f9a403f(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make(map[int][]ContentType, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Age
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		result[bufferData[kbegin].Age] = bufferData[kbegin:k]

	}
	return result
}

func queryColumnMap_f37bd7ac5974dfb3bba9c883aa742e29168e8dc7(data interface{}, column string) interface{} {
	dataIn := data.([]QueryInnerStruct2)
	bufferData := make([]QueryInnerStruct2, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	result := make(map[int][]QueryInnerStruct2, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].QueryInnerStruct.MM
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		result[bufferData[kbegin].QueryInnerStruct.MM] = bufferData[kbegin:k]

	}
	return result
}

func queryColumn_12cb9c765f30b6d40fa9080910c80cea2cce66f7(data interface{}, column string) interface{} {
	dataIn := data.([]QueryInnerStruct2)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.QueryInnerStruct.MM
	}
	return result
}

func queryColumn_163b58b970f87da0d1c228cced6790dd60ad2896(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
	}
	return result
}

func queryColumn_20e1f01d4859d8fc4c20210c90dffffcbbcf9fc5(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumn_2d7e0ce966e7e5fb1bd62167c50fc81590fd3a1e(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Age
	}
	return result
}

func queryColumn_35ec5ddff65a5a470c2b5f2e98feaa54169f430f(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Age
	}
	return result
}

func queryColumn_3923b792e276005e09637544ecb3aec8be870f41(data interface{}, column string) interface{} {
	dataIn := data.([]string)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumn_440473c7ae32877de566d281340e92c7a42fa3c2(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumn_63871bbce37249e280e427591a85c168607bcfb1(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Name
	}
	return result
}

func queryColumn_68e685a02a5d491113488beda9d1f86ea17b9f17(data interface{}, column string) interface{} {
	dataIn := data.([]QueryInnerStruct2)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.MM
	}
	return result
}

func queryColumn_725664359c0503f05a9971dca3064bc1d143d4fb(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]bool, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Ok
	}
	return result
}

func queryColumn_906ee8d8789edf186c177303f2b128f962913f12(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.CardMoney
	}
	return result
}

func queryColumn_91dacd60e87431951940b4b4c51428e7c1e5c1f2(data interface{}, column string) interface{} {
	dataIn := data.([]int)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single
	}
	return result
}

func queryColumn_ca2b7b23bd1c78dcf7d146a73399ab1744521b59(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Money
	}
	return result
}

func queryColumn_fba4271ecd4689e141e5e7912ebfa4df33dbf4df(data interface{}, column string) interface{} {
	dataIn := data.([]ContentType)
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = single.Money
	}
	return result
}

func queryCombine_201ce2f3d02a38a74c829435c5a37643ad1c65a0(leftData interface{}, rightData interface{}, combineFunctor interface{}) interface{} {
	leftDataIn := leftData.([]ContentType)
	rightDataIn := rightData.([]int)
	combineFunctorIn := combineFunctor.(func(ContentType, int) ContentType)
	newData := make([]ContentType, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryCombine_ea8b019a0a29189fa3ebd210e2511c1c02fb7e6b(leftData interface{}, rightData interface{}, combineFunctor interface{}) interface{} {
	leftDataIn := leftData.([]ContentType)
	rightDataIn := rightData.([]ContentType)
	combineFunctorIn := combineFunctor.(func(ContentType, ContentType) ContentType)
	newData := make([]ContentType, len(leftDataIn), len(leftDataIn))

	for i := 0; i != len(leftDataIn); i++ {
		newData[i] = combineFunctorIn(leftDataIn[i], rightDataIn[i])
	}
	return newData
}

func queryGroup_0c6547761bb0973a4383d22843b1d0709fba94ff(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[bool]int, len(dataIn))
	groupFunctorIn := groupFunctor.(func([]ContentType) []ContentType)
	result := make([]ContentType, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Ok
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)

	}
	return result
}

func queryGroup_35b60c12a41776ebe8b8613c6ba9f1f8601788c0(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[bool]int, len(dataIn))
	groupFunctorIn := groupFunctor.(func([]ContentType) []ContentType)
	result := make([]ContentType, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Ok
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)

	}
	return result
}

func queryGroup_61fdff2f19be1e88af76335f2f4e5a87e4f2f6e8(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[time.Time]int, len(dataIn))
	groupFunctorIn := groupFunctor.(func([]ContentType) []ContentType)
	result := make([]ContentType, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Register
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)

	}
	return result
}

func queryGroup_7670c81859dba936c3e0abf6ae6f86787d131520(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	groupFunctorIn := groupFunctor.(func([]ContentType) float32)
	result := make([]float32, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Name
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)

	}
	return result
}

func queryGroup_7959aac2ba701c92b02938af82c21599cbf58c3d(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]int)
	bufferData := make([]int, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor.(func([]int) int)
	result := make([]int, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i]
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)

	}
	return result
}

func queryGroup_80e3458868e4efbe8804936628265c1ebe9c299e(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor.(func([]ContentType) float64)
	result := make([]float64, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Age
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)

	}
	return result
}

func queryGroup_a06bb70800026a2782990b907c179096565c86b6(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor.(func([]ContentType) []ContentType)
	result := make([]ContentType, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Age
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)

	}
	return result
}

func queryGroup_b83692f0dc3c6c1880d677b53ecd5a1ab96c6900(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]string)
	bufferData := make([]string, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	groupFunctorIn := groupFunctor.(func([]string) ContentType)
	result := make([]ContentType, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i]
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)

	}
	return result
}

func queryGroup_cb3c9e707f62b6dbaef6a9a28e6c92f5c427bca4(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]QueryInnerStruct2)
	bufferData := make([]QueryInnerStruct2, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor.(func([]QueryInnerStruct2) []QueryInnerStruct2)
	result := make([]QueryInnerStruct2, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].QueryInnerStruct.MM
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)

	}
	return result
}

func queryGroup_d2ed8aa28b5e2c955476e38bf294a9130af6a8b0(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[int]int, len(dataIn))
	groupFunctorIn := groupFunctor.(func([]ContentType) []float64)
	result := make([]float64, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Age
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)

	}
	return result
}

func queryGroup_d30e6ca8b91693c4d7644c61f70b4aae7b6790de(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[time.Time]int, len(dataIn))
	groupFunctorIn := groupFunctor.(func([]ContentType) int)
	result := make([]int, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Register
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single)

	}
	return result
}

func queryGroup_f98cba90270d119dd824b9d094279f1675a4628e(data interface{}, groupType string, groupFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	bufferData := make([]ContentType, len(dataIn), len(dataIn))
	mapData := make(map[string]int, len(dataIn))
	groupFunctorIn := groupFunctor.(func([]ContentType) []ContentType)
	result := make([]ContentType, 0, len(dataIn))

	length := len(dataIn)
	nextData := make([]int, length, length)
	for i := 0; i != length; i++ {
		single := dataIn[i].Name
		lastIndex, isExist := mapData[single]
		if isExist == true {
			nextData[lastIndex] = i
		}
		nextData[i] = -1
		mapData[single] = i
	}
	k := 0
	for i := 0; i != length; i++ {
		j := i
		if nextData[j] == 0 {
			continue
		}
		kbegin := k
		for nextData[j] != -1 {
			nextJ := nextData[j]
			bufferData[k] = dataIn[j]
			nextData[j] = 0
			j = nextJ
			k++
		}
		bufferData[k] = dataIn[j]
		k++
		nextData[j] = 0

		single := groupFunctorIn(bufferData[kbegin:k])
		result = append(result, single...)

	}
	return result
}

func queryJoin_30dc039066f6419d6957e80c70411d47f4b5ba38(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]ContentType2)
	joinFunctorIn := joinFunctor.(func(UserType, ContentType2) resultType)
	result := make([]resultType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := ContentType2{}
	joinPlace = "inner"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].UserName
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.Name
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_4241fc7851207aace786fc08431b2889505938b2(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]QueryInnerStruct2)
	rightDataIn := rightData.([]QueryInnerStruct2)
	joinFunctorIn := joinFunctor.(func(QueryInnerStruct2, QueryInnerStruct2) QueryInnerStruct2)
	result := make([]QueryInnerStruct2, 0, len(leftDataIn))

	emptyLeftData := QueryInnerStruct2{}
	emptyRightData := QueryInnerStruct2{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[int]int, len(rightDataIn))
	mapDataFirst := make(map[int]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].QueryInnerStruct.MM
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.QueryInnerStruct.MM
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_4c09fb28199d6fb68abff5b1f63b05dbefa478c8(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(UserType, UserType) UserType)
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := UserType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[time.Time]int, len(rightDataIn))
	mapDataFirst := make(map[time.Time]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Register
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.Register
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_4de3546ff37ccd0674dc0e293ba1e769036d2ced(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(UserType, UserType) UserType)
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := UserType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Name
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.Name
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_4ec9f8387590e45d2744b1c2070403a5f57a7a2a(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(UserType, UserType) UserType)
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := UserType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[float64]int, len(rightDataIn))
	mapDataFirst := make(map[float64]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Money
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.Money
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_69b119a315d160bab50488370dfbc5cbba1e0131(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]ExtendType)
	rightDataIn := rightData.([]ExtendType)
	joinFunctorIn := joinFunctor.(func(ExtendType, ExtendType) ExtendType)
	result := make([]ExtendType, 0, len(leftDataIn))

	emptyLeftData := ExtendType{}
	emptyRightData := ExtendType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[int]int, len(rightDataIn))
	mapDataFirst := make(map[int]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].ContentID
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.ContentID
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_6b2807f926ab60c7d2c5f9778cf2016ec6629f10(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]ContentType2)
	joinFunctorIn := joinFunctor.(func(UserType, ContentType2) resultType)
	result := make([]resultType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := ContentType2{}
	joinPlace = "outer"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].UserName
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.Name
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_796ffd9c222e1ce113eb608fab85d0aa8c3833ac(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]ContentType2)
	joinFunctorIn := joinFunctor.(func(UserType, ContentType2) resultType)
	result := make([]resultType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := ContentType2{}
	joinPlace = "right"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].UserName
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.Name
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_8b966d85170a23adae3fef0b0077ebb2be250354(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(UserType, UserType) UserType)
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := UserType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[float64]int, len(rightDataIn))
	mapDataFirst := make(map[float64]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Money
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.CardMoney
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_a688a4facd5a09920db8e94324979d4672c8a1a6(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]string)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(string, UserType) UserType)
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := ""
	emptyRightData := UserType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Name
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_b52dacd0e43e79057203119d86599f9c08c8eba5(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]int)
	rightDataIn := rightData.([]ExtendType)
	joinFunctorIn := joinFunctor.(func(int, ExtendType) ExtendType)
	result := make([]ExtendType, 0, len(leftDataIn))

	emptyLeftData := 0
	emptyRightData := ExtendType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[int]int, len(rightDataIn))
	mapDataFirst := make(map[int]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].ContentID
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_bef20457ee7b83717c635c495406fc2a81474a1b(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]ContentType2)
	joinFunctorIn := joinFunctor.(func(UserType, ContentType2) resultType)
	result := make([]resultType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := ContentType2{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].UserName
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.Name
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_c1841bda53132ab43596160c0c8a8ea30441db62(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]string)
	rightDataIn := rightData.([]ContentType2)
	joinFunctorIn := joinFunctor.(func(string, ContentType2) ContentType2)
	result := make([]ContentType2, 0, len(leftDataIn))

	emptyLeftData := ""
	emptyRightData := ContentType2{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[string]int, len(rightDataIn))
	mapDataFirst := make(map[string]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].UserName
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_c694c96d521e338da0cf10ff8a187f824ee16f7a(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(UserType, UserType) UserType)
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := UserType{}
	joinPlace = "left"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[bool]int, len(rightDataIn))
	mapDataFirst := make(map[bool]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Ok
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.Ok
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func queryJoin_caebed046b26f1196284f1fac2ff834a2dce91eb(leftData interface{}, rightData interface{}, joinPlace string, joinType string, joinFunctor interface{}) interface{} {
	leftDataIn := leftData.([]UserType)
	rightDataIn := rightData.([]UserType)
	joinFunctorIn := joinFunctor.(func(UserType, UserType) UserType)
	result := make([]UserType, 0, len(leftDataIn))

	emptyLeftData := UserType{}
	emptyRightData := UserType{}
	joinPlace = "right"

	nextData := make([]int, len(rightDataIn), len(rightDataIn))
	mapDataNext := make(map[int]int, len(rightDataIn))
	mapDataFirst := make(map[int]int, len(rightDataIn))

	for i := 0; i != len(rightDataIn); i++ {
		fieldValue := rightDataIn[i].Age
		lastIndex, isExist := mapDataNext[fieldValue]
		if isExist {
			nextData[lastIndex] = i
		} else {
			mapDataFirst[fieldValue] = i
		}
		nextData[i] = -1
		mapDataNext[fieldValue] = i
	}

	rightHaveJoin := make([]bool, len(rightDataIn), len(rightDataIn))
	for i := 0; i != len(leftDataIn); i++ {
		leftValue := leftDataIn[i]
		fieldValue := leftValue.Age
		rightIndex, isExist := mapDataFirst[fieldValue]
		if isExist {
			//找到右值
			j := rightIndex
			for nextData[j] != -1 {
				singleResult := joinFunctorIn(leftValue, rightDataIn[j])
				result = append(result, singleResult)
				rightHaveJoin[j] = true
				j = nextData[j]
			}
			singleResult := joinFunctorIn(leftValue, rightDataIn[j])
			result = append(result, singleResult)
			rightHaveJoin[j] = true
		} else {
			//找不到右值
			if joinPlace == "left" || joinPlace == "outer" {
				singleResult := joinFunctorIn(leftValue, emptyRightData)
				result = append(result, singleResult)
			}
		}
	}
	if joinPlace == "right" || joinPlace == "outer" {
		for j := 0; j != len(rightDataIn); j++ {
			if rightHaveJoin[j] {
				continue
			}
			singleResult := joinFunctorIn(emptyLeftData, rightDataIn[j])
			result = append(result, singleResult)
		}
	}
	return result
}

func querySelect_05e4b15cc3a69914fda8814f5872245157e0b985(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) int)
	result := make([]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelect_480a3efc4d2b87affc1a3a669ffac0acf8cf449e(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) bool)
	result := make([]bool, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelect_6085d6df5904966661594aeb23b1ce947424b16f(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) float32)
	result := make([]float32, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelect_7c6f88e03fb7c37cf22fe1f82aacf0e05a87c5dc(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) string)
	result := make([]string, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelect_a345e7ab88300ed2fa569ba92b5524034245a509(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) float64)
	result := make([]float64, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelect_a4b8c55e5808cc0fa09ae85d168d9ba1d67d59c6(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) time.Time)
	result := make([]time.Time, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelect_d533aaeeed37aa1681368203aaa3a4a0956b01e3(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) map[string]int)
	result := make([]map[string]int, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySelect_ee4857e23bc90bddbd99d92d75667bbbc9119553(data interface{}, selectFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	selectFunctorIn := selectFunctor.(func(ContentType) ContentType)
	result := make([]ContentType, len(dataIn), len(dataIn))

	for i, single := range dataIn {
		result[i] = selectFunctorIn(single)
	}
	return result
}

func querySort_12aab7b20c101ccdec30ed1a2e4ef9188296273b(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].Age < newData[j].Age {
			return 1
		} else if newData[i].Age > newData[j].Age {
			return -1
		}

		if newData[i].Ok == false && newData[j].Ok == true {
			return 1
		} else if newData[i].Ok == true && newData[j].Ok == false {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_27179b7563c3060f1b040933830d9412869d612c(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].CardMoney < newData[j].CardMoney {
			return -1
		} else if newData[i].CardMoney > newData[j].CardMoney {
			return 1
		}

		if newData[i].Register.Before(newData[j].Register) {
			return 1
		} else if newData[i].Register.After(newData[j].Register) {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_3946cca9dadc952f6146f008151c14dd9aed1dc7(data interface{}, sortType string) interface{} {
	dataIn := data.([]QueryInnerStruct2)
	newData := make([]QueryInnerStruct2, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].QueryInnerStruct.MM < newData[j].QueryInnerStruct.MM {
			return -1
		} else if newData[i].QueryInnerStruct.MM > newData[j].QueryInnerStruct.MM {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_3cbb995360e7a3448e99bd82b455f641085b28aa(data interface{}, sortType string) interface{} {
	dataIn := data.([]Student)
	newData := make([]Student, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		{
			tempDecimalCmp := newData[i].Score.Cmp(newData[j].Score)
			if tempDecimalCmp < 0 {
				return -1
			} else if tempDecimalCmp > 0 {
				return 1
			}
		}

		{
			tempDecimalCmp := newData[i].Score2.Cmp(newData[j].Score2)
			if tempDecimalCmp < 0 {
				return 1
			} else if tempDecimalCmp > 0 {
				return -1
			}
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_74654e8b45593005ef783b89255269f7c6ecc39b(data interface{}, sortType string) interface{} {
	dataIn := data.([]int)
	newData := make([]int, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i] < newData[j] {
			return -1
		} else if newData[i] > newData[j] {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_7820966153878ee7f83ad21258bd8245015746fa(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].Ok == false && newData[j].Ok == true {
			return 1
		} else if newData[i].Ok == true && newData[j].Ok == false {
			return -1
		}

		if newData[i].Name < newData[j].Name {
			return -1
		} else if newData[i].Name > newData[j].Name {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_8321af86ab4036933f4bb7188e99550d9b88b2f1(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].Name < newData[j].Name {
			return -1
		} else if newData[i].Name > newData[j].Name {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_a355e5961b46e18098a3d13389bfd6498a446cf5(data interface{}, sortType string) interface{} {
	dataIn := data.([]language.Decimal)
	newData := make([]language.Decimal, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		{
			tempDecimalCmp := newData[i].Cmp(newData[j])
			if tempDecimalCmp < 0 {
				return 1
			} else if tempDecimalCmp > 0 {
				return -1
			}
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_af891d058d5a2e0a3ac4b4b291ae9bb959364795(data interface{}, sortType string) interface{} {
	dataIn := data.([]int)
	newData := make([]int, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i] < newData[j] {
			return 1
		} else if newData[i] > newData[j] {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_b008424daf0bde59dfbe09b67fc521a1a5ee6c49(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].Money < newData[j].Money {
			return -1
		} else if newData[i].Money > newData[j].Money {
			return 1
		}

		if newData[i].Register.Before(newData[j].Register) {
			return 1
		} else if newData[i].Register.After(newData[j].Register) {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_beab8849a565ac45c9cda5e68f2a43207ad4c52c(data interface{}, sortType string) interface{} {
	dataIn := data.([]Student)
	newData := make([]Student, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].Name < newData[j].Name {
			return -1
		} else if newData[i].Name > newData[j].Name {
			return 1
		}

		{
			tempDecimalCmp := newData[i].Score.Cmp(newData[j].Score)
			if tempDecimalCmp < 0 {
				return 1
			} else if tempDecimalCmp > 0 {
				return -1
			}
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_c7f8509b5aef21d9f376b304e0b93e419397cec8(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].Money < newData[j].Money {
			return 1
		} else if newData[i].Money > newData[j].Money {
			return -1
		}

		if newData[i].Age < newData[j].Age {
			return -1
		} else if newData[i].Age > newData[j].Age {
			return 1
		}

		if newData[i].Name < newData[j].Name {
			return 1
		} else if newData[i].Name > newData[j].Name {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_d3e1ca8a84644854d1dddd0eca7a89bb6803cb6a(data interface{}, sortType string) interface{} {
	dataIn := data.([]QueryInnerStruct2)
	newData := make([]QueryInnerStruct2, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].MM < newData[j].MM {
			return 1
		} else if newData[i].MM > newData[j].MM {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_d4cc042b4a7e1344fe17bd968ce09ad536310de1(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].Name < newData[j].Name {
			return 1
		} else if newData[i].Name > newData[j].Name {
			return -1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func querySort_d684a18252070e8ca1b87d2481db44b426d34a54(data interface{}, sortType string) interface{} {
	dataIn := data.([]ContentType)
	newData := make([]ContentType, len(dataIn), len(dataIn))
	copy(newData, dataIn)

	language.QuerySortInternal(len(newData), func(i int, j int) int {
		if newData[i].Money < newData[j].Money {
			return 1
		} else if newData[i].Money > newData[j].Money {
			return -1
		}

		if newData[i].Age < newData[j].Age {
			return -1
		} else if newData[i].Age > newData[j].Age {
			return 1
		}

		return 0
	}, func(i int, j int) {
		newData[j], newData[i] = newData[i], newData[j]
	})
	return newData
}

func queryWhere_480a3efc4d2b87affc1a3a669ffac0acf8cf449e(data interface{}, whereFunctor interface{}) interface{} {
	dataIn := data.([]ContentType)
	whereFunctorIn := whereFunctor.(func(ContentType) bool)
	result := make([]ContentType, 0, len(dataIn))

	for _, single := range dataIn {
		shouldStay := whereFunctorIn(single)
		if shouldStay == true {
			result = append(result, single)
		}
	}
	return result
}

func init() {

	language.QueryColumnMapMacroRegister([]QueryInnerStruct2{}, "QueryInnerStruct.MM", queryColumnMap_12cb9c765f30b6d40fa9080910c80cea2cce66f7)

	language.QueryColumnMapMacroRegister([]ContentType{}, " Name ", queryColumnMap_163b58b970f87da0d1c228cced6790dd60ad2896)

	language.QueryColumnMapMacroRegister([]ContentType{}, "Age        ", queryColumnMap_2d7e0ce966e7e5fb1bd62167c50fc81590fd3a1e)

	language.QueryColumnMapMacroRegister([]string{}, " . ", queryColumnMap_3923b792e276005e09637544ecb3aec8be870f41)

	language.QueryColumnMapMacroRegister([]ContentType{}, "    CardMoney", queryColumnMap_440473c7ae32877de566d281340e92c7a42fa3c2)

	language.QueryColumnMapMacroRegister([]ContentType{}, " []Name ", queryColumnMap_621f96588f04633a85c32b0497fdd5b7cfaaebdf)

	language.QueryColumnMapMacroRegister([]ContentType{}, "     Name         ", queryColumnMap_63871bbce37249e280e427591a85c168607bcfb1)

	language.QueryColumnMapMacroRegister([]ContentType{}, "[]Ok        ", queryColumnMap_704d73f9c17886c384e0a16be36b7b5959fea51f)

	language.QueryColumnMapMacroRegister([]ContentType{}, "Ok        ", queryColumnMap_725664359c0503f05a9971dca3064bc1d143d4fb)

	language.QueryColumnMapMacroRegister([]ContentType{}, "    []Money  ", queryColumnMap_77cc920402a2e411fd707de6c0797eecc2947f6e)

	language.QueryColumnMapMacroRegister([]int{}, " . ", queryColumnMap_91dacd60e87431951940b4b4c51428e7c1e5c1f2)

	language.QueryColumnMapMacroRegister([]ContentType{}, "    []CardMoney", queryColumnMap_bd864bf11acf9e4deb95377e4098c24bd32a1308)

	language.QueryColumnMapMacroRegister([]ContentType{}, "    Money  ", queryColumnMap_ca2b7b23bd1c78dcf7d146a73399ab1744521b59)

	language.QueryColumnMapMacroRegister([]ContentType{}, "     [] Name         ", queryColumnMap_da12d428eb4ff4f1c8585a24fead914b37d902b8)

	language.QueryColumnMapMacroRegister([]ContentType{}, "[]Age        ", queryColumnMap_dd430d90e19ca65443db28bfbf3eb2336f9a403f)

	language.QueryColumnMapMacroRegister([]QueryInnerStruct2{}, "[]QueryInnerStruct.MM", queryColumnMap_f37bd7ac5974dfb3bba9c883aa742e29168e8dc7)

	language.QueryColumnMacroRegister([]QueryInnerStruct2{}, "QueryInnerStruct.MM", queryColumn_12cb9c765f30b6d40fa9080910c80cea2cce66f7)

	language.QueryColumnMacroRegister([]ContentType{}, " Name ", queryColumn_163b58b970f87da0d1c228cced6790dd60ad2896)

	language.QueryColumnMacroRegister([]ContentType{}, "  CardMoney  ", queryColumn_20e1f01d4859d8fc4c20210c90dffffcbbcf9fc5)

	language.QueryColumnMacroRegister([]ContentType{}, "Age        ", queryColumn_2d7e0ce966e7e5fb1bd62167c50fc81590fd3a1e)

	language.QueryColumnMacroRegister([]ContentType{}, "  Age  ", queryColumn_35ec5ddff65a5a470c2b5f2e98feaa54169f430f)

	language.QueryColumnMacroRegister([]string{}, " . ", queryColumn_3923b792e276005e09637544ecb3aec8be870f41)

	language.QueryColumnMacroRegister([]ContentType{}, "    CardMoney", queryColumn_440473c7ae32877de566d281340e92c7a42fa3c2)

	language.QueryColumnMacroRegister([]ContentType{}, "     Name         ", queryColumn_63871bbce37249e280e427591a85c168607bcfb1)

	language.QueryColumnMacroRegister([]QueryInnerStruct2{}, "  MM  ", queryColumn_68e685a02a5d491113488beda9d1f86ea17b9f17)

	language.QueryColumnMacroRegister([]ContentType{}, "Ok        ", queryColumn_725664359c0503f05a9971dca3064bc1d143d4fb)

	language.QueryColumnMacroRegister([]ContentType{}, "CardMoney  ", queryColumn_906ee8d8789edf186c177303f2b128f962913f12)

	language.QueryColumnMacroRegister([]int{}, " . ", queryColumn_91dacd60e87431951940b4b4c51428e7c1e5c1f2)

	language.QueryColumnMacroRegister([]ContentType{}, "    Money  ", queryColumn_ca2b7b23bd1c78dcf7d146a73399ab1744521b59)

	language.QueryColumnMacroRegister([]ContentType{}, "  Money  ", queryColumn_fba4271ecd4689e141e5e7912ebfa4df33dbf4df)

	language.QueryCombineMacroRegister([]ContentType{}, []int{}, (func(ContentType, int) ContentType)(nil), queryCombine_201ce2f3d02a38a74c829435c5a37643ad1c65a0)

	language.QueryCombineMacroRegister([]ContentType{}, []ContentType{}, (func(ContentType, ContentType) ContentType)(nil), queryCombine_ea8b019a0a29189fa3ebd210e2511c1c02fb7e6b)

	language.QueryGroupMacroRegister([]ContentType{}, " Ok ", (func([]ContentType) []ContentType)(nil), queryGroup_0c6547761bb0973a4383d22843b1d0709fba94ff)

	language.QueryGroupMacroRegister([]ContentType{}, "Ok", (func([]ContentType) []ContentType)(nil), queryGroup_35b60c12a41776ebe8b8613c6ba9f1f8601788c0)

	language.QueryGroupMacroRegister([]ContentType{}, "Register ", (func([]ContentType) []ContentType)(nil), queryGroup_61fdff2f19be1e88af76335f2f4e5a87e4f2f6e8)

	language.QueryGroupMacroRegister([]ContentType{}, "Name", (func([]ContentType) float32)(nil), queryGroup_7670c81859dba936c3e0abf6ae6f86787d131520)

	language.QueryGroupMacroRegister([]int{}, ".", (func([]int) int)(nil), queryGroup_7959aac2ba701c92b02938af82c21599cbf58c3d)

	language.QueryGroupMacroRegister([]ContentType{}, " Age ", (func([]ContentType) float64)(nil), queryGroup_80e3458868e4efbe8804936628265c1ebe9c299e)

	language.QueryGroupMacroRegister([]ContentType{}, " Age ", (func([]ContentType) []ContentType)(nil), queryGroup_a06bb70800026a2782990b907c179096565c86b6)

	language.QueryGroupMacroRegister([]string{}, ".", (func([]string) ContentType)(nil), queryGroup_b83692f0dc3c6c1880d677b53ecd5a1ab96c6900)

	language.QueryGroupMacroRegister([]QueryInnerStruct2{}, "QueryInnerStruct.MM", (func([]QueryInnerStruct2) []QueryInnerStruct2)(nil), queryGroup_cb3c9e707f62b6dbaef6a9a28e6c92f5c427bca4)

	language.QueryGroupMacroRegister([]ContentType{}, " Age ", (func([]ContentType) []float64)(nil), queryGroup_d2ed8aa28b5e2c955476e38bf294a9130af6a8b0)

	language.QueryGroupMacroRegister([]ContentType{}, "Register ", (func([]ContentType) int)(nil), queryGroup_d30e6ca8b91693c4d7644c61f70b4aae7b6790de)

	language.QueryGroupMacroRegister([]ContentType{}, "Name", (func([]ContentType) []ContentType)(nil), queryGroup_f98cba90270d119dd824b9d094279f1675a4628e)

	language.QueryJoinMacroRegister([]UserType{}, []ContentType2{}, "inner", "  Name  =  UserName ", (func(UserType, ContentType2) resultType)(nil), queryJoin_30dc039066f6419d6957e80c70411d47f4b5ba38)

	language.QueryJoinMacroRegister([]QueryInnerStruct2{}, []QueryInnerStruct2{}, "left", "QueryInnerStruct.MM = QueryInnerStruct.MM", (func(QueryInnerStruct2, QueryInnerStruct2) QueryInnerStruct2)(nil), queryJoin_4241fc7851207aace786fc08431b2889505938b2)

	language.QueryJoinMacroRegister([]UserType{}, []UserType{}, "left", " Register = Register ", (func(UserType, UserType) UserType)(nil), queryJoin_4c09fb28199d6fb68abff5b1f63b05dbefa478c8)

	language.QueryJoinMacroRegister([]UserType{}, []UserType{}, " left ", "  Name  =  Name ", (func(UserType, UserType) UserType)(nil), queryJoin_4de3546ff37ccd0674dc0e293ba1e769036d2ced)

	language.QueryJoinMacroRegister([]UserType{}, []UserType{}, "left", " Money=Money ", (func(UserType, UserType) UserType)(nil), queryJoin_4ec9f8387590e45d2744b1c2070403a5f57a7a2a)

	language.QueryJoinMacroRegister([]ExtendType{}, []ExtendType{}, " left ", "  ContentID  =  ContentID ", (func(ExtendType, ExtendType) ExtendType)(nil), queryJoin_69b119a315d160bab50488370dfbc5cbba1e0131)

	language.QueryJoinMacroRegister([]UserType{}, []ContentType2{}, "outer", "  Name  =  UserName ", (func(UserType, ContentType2) resultType)(nil), queryJoin_6b2807f926ab60c7d2c5f9778cf2016ec6629f10)

	language.QueryJoinMacroRegister([]UserType{}, []ContentType2{}, "right", "  Name  =  UserName ", (func(UserType, ContentType2) resultType)(nil), queryJoin_796ffd9c222e1ce113eb608fab85d0aa8c3833ac)

	language.QueryJoinMacroRegister([]UserType{}, []UserType{}, "left", " CardMoney = Money ", (func(UserType, UserType) UserType)(nil), queryJoin_8b966d85170a23adae3fef0b0077ebb2be250354)

	language.QueryJoinMacroRegister([]string{}, []UserType{}, "left", " . = Name", (func(string, UserType) UserType)(nil), queryJoin_a688a4facd5a09920db8e94324979d4672c8a1a6)

	language.QueryJoinMacroRegister([]int{}, []ExtendType{}, " left ", "  .  =  ContentID ", (func(int, ExtendType) ExtendType)(nil), queryJoin_b52dacd0e43e79057203119d86599f9c08c8eba5)

	language.QueryJoinMacroRegister([]UserType{}, []ContentType2{}, "left", "  Name  =  UserName ", (func(UserType, ContentType2) resultType)(nil), queryJoin_bef20457ee7b83717c635c495406fc2a81474a1b)

	language.QueryJoinMacroRegister([]string{}, []ContentType2{}, "left", "  .  =  UserName ", (func(string, ContentType2) ContentType2)(nil), queryJoin_c1841bda53132ab43596160c0c8a8ea30441db62)

	language.QueryJoinMacroRegister([]UserType{}, []UserType{}, "left", "Ok  =  Ok", (func(UserType, UserType) UserType)(nil), queryJoin_c694c96d521e338da0cf10ff8a187f824ee16f7a)

	language.QueryJoinMacroRegister([]UserType{}, []UserType{}, "right", "Age=Age", (func(UserType, UserType) UserType)(nil), queryJoin_caebed046b26f1196284f1fac2ff834a2dce91eb)

	language.QuerySelectMacroRegister([]ContentType{}, (func(ContentType) int)(nil), querySelect_05e4b15cc3a69914fda8814f5872245157e0b985)

	language.QuerySelectMacroRegister([]ContentType{}, (func(ContentType) bool)(nil), querySelect_480a3efc4d2b87affc1a3a669ffac0acf8cf449e)

	language.QuerySelectMacroRegister([]ContentType{}, (func(ContentType) float32)(nil), querySelect_6085d6df5904966661594aeb23b1ce947424b16f)

	language.QuerySelectMacroRegister([]ContentType{}, (func(ContentType) string)(nil), querySelect_7c6f88e03fb7c37cf22fe1f82aacf0e05a87c5dc)

	language.QuerySelectMacroRegister([]ContentType{}, (func(ContentType) float64)(nil), querySelect_a345e7ab88300ed2fa569ba92b5524034245a509)

	language.QuerySelectMacroRegister([]ContentType{}, (func(ContentType) time.Time)(nil), querySelect_a4b8c55e5808cc0fa09ae85d168d9ba1d67d59c6)

	language.QuerySelectMacroRegister([]ContentType{}, (func(ContentType) map[string]int)(nil), querySelect_d533aaeeed37aa1681368203aaa3a4a0956b01e3)

	language.QuerySelectMacroRegister([]ContentType{}, (func(ContentType) ContentType)(nil), querySelect_ee4857e23bc90bddbd99d92d75667bbbc9119553)

	language.QuerySortMacroRegister([]ContentType{}, "Age desc,Ok desc", querySort_12aab7b20c101ccdec30ed1a2e4ef9188296273b)

	language.QuerySortMacroRegister([]ContentType{}, "CardMoney,Register desc", querySort_27179b7563c3060f1b040933830d9412869d612c)

	language.QuerySortMacroRegister([]QueryInnerStruct2{}, "QueryInnerStruct.MM asc", querySort_3946cca9dadc952f6146f008151c14dd9aed1dc7)

	language.QuerySortMacroRegister([]Student{}, "Score asc,Score2 desc", querySort_3cbb995360e7a3448e99bd82b455f641085b28aa)

	language.QuerySortMacroRegister([]int{}, ". asc", querySort_74654e8b45593005ef783b89255269f7c6ecc39b)

	language.QuerySortMacroRegister([]ContentType{}, "Ok desc,Name", querySort_7820966153878ee7f83ad21258bd8245015746fa)

	language.QuerySortMacroRegister([]ContentType{}, "Name asc", querySort_8321af86ab4036933f4bb7188e99550d9b88b2f1)

	language.QuerySortMacroRegister([]language.Decimal{}, ". desc", querySort_a355e5961b46e18098a3d13389bfd6498a446cf5)

	language.QuerySortMacroRegister([]int{}, ". desc", querySort_af891d058d5a2e0a3ac4b4b291ae9bb959364795)

	language.QuerySortMacroRegister([]ContentType{}, "Money,Register desc", querySort_b008424daf0bde59dfbe09b67fc521a1a5ee6c49)

	language.QuerySortMacroRegister([]Student{}, "Name asc,Score desc", querySort_beab8849a565ac45c9cda5e68f2a43207ad4c52c)

	language.QuerySortMacroRegister([]ContentType{}, " Money desc,Age asc,Name desc", querySort_c7f8509b5aef21d9f376b304e0b93e419397cec8)

	language.QuerySortMacroRegister([]QueryInnerStruct2{}, "MM desc", querySort_d3e1ca8a84644854d1dddd0eca7a89bb6803cb6a)

	language.QuerySortMacroRegister([]ContentType{}, "Name desc", querySort_d4cc042b4a7e1344fe17bd968ce09ad536310de1)

	language.QuerySortMacroRegister([]ContentType{}, " Money desc,Age asc", querySort_d684a18252070e8ca1b87d2481db44b426d34a54)

	language.QueryWhereMacroRegister([]ContentType{}, (func(ContentType) bool)(nil), queryWhere_480a3efc4d2b87affc1a3a669ffac0acf8cf449e)

}
