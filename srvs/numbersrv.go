package srvs

import (
	"trafilea/model"
	"trafilea/repos"
)

type NumberSrv struct {
	numberRepo repos.NumberRepo
}

func New(numberRepo repos.NumberRepo) NumberSrv {
	return NumberSrv{numberRepo: numberRepo}
}

func (srv *NumberSrv) GetNumber(number int) (model.Number, error) {
	return srv.numberRepo.GetNumber(number)
}

func (srv *NumberSrv) GetNumbers() []model.Number {
	return srv.numberRepo.GetNumbers()
}

func (srv *NumberSrv) SaveNumber(number model.Number) {
	numType := 1
	// Core logic
	for i := 2; i <= 5; i++ {
		if number.Value%i == 0 && (i == 3 || i == 5) {
			numType = numType * i
		}
	}

	// Type mapping
	switch numType {
	case 3:
		number.NumberType = "Type1"
	case 5:
		number.NumberType = "Type2"
	case 15:
		number.NumberType = "Type3"
	}

	srv.numberRepo.SaveNumber(number)
}
