package repos

import (
	"errors"

	"trafilea/model"
)

type NumberRepo struct {
	numberMap map[int]model.Number
}

func New() NumberRepo {
	numberMap := make(map[int]model.Number)
	return NumberRepo{numberMap: numberMap}
}

func (repo *NumberRepo) GetNumber(number int) (model.Number, error) {
	num, ok := repo.numberMap[number]
	if !ok {
		return model.Number{}, errors.New("number not found")
	}
	return num, nil
}

func (repo *NumberRepo) GetNumbers() []model.Number {
	var nums []model.Number
	for _, v := range repo.numberMap {
		nums = append(nums, v)
	}
	return nums
}

func (repo *NumberRepo) SaveNumber(num model.Number) {
	repo.numberMap[num.Value] = num
}
