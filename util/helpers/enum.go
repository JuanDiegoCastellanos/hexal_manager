package helpers

import (
	"errors"
)

// ValueEnum is a struct used to simulate the enum string
type ValueEnum struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func NewValueEnum(id int32, name string) *ValueEnum {
	return &ValueEnum{
		ID:   id,
		Name: name,
	}
}

// EnumEntity is the interface that all the entities where their need the EnumValue
// they have to implement this interface
// EnumEntity has some single methods to add, delete, update, getAll and getOneSpecific
type EnumEntity interface {
	GetValues() []ValueEnum
	GetValue(enum *ValueEnum) (*ValueEnum, error)
	AddValue(enum *ValueEnum) error
	UpdateValue(enum *ValueEnum) error
	DeleteValue(enum *ValueEnum) error
}

// EnumAdapter is the adapter in order to each entity implements its specific enum
type EnumAdapter struct {
	Values []ValueEnum
}

func NewEnumAdapter(valuesEnum []ValueEnum) *EnumAdapter {
	return &EnumAdapter{
		Values: valuesEnum,
	}
}

func (ea *EnumAdapter) GetValues() []ValueEnum {
	return ea.Values
}
func (ea *EnumAdapter) GetValue(valueParam ValueEnum) (*ValueEnum, error) {
	var valEnm ValueEnum
	var err error
	if valueParam.ID != 0 && valueParam.Name != "" {
		for _, valueEnum := range ea.Values {
			if valueEnum.Name == valueParam.Name || valueEnum.ID == valueParam.ID {
				valEnm = valueEnum
				err = nil
			}
			err = errors.New("valueEnum it doesn't exist")
		}
	}
	return &valEnm, err
}
func (ea *EnumAdapter) AddValue(paramValue ValueEnum) error {
	if paramValue.Name != "" && paramValue.ID != 0 {
		ea.Values = append(ea.Values, paramValue)
		return nil
	}
	return errors.New("the parameter newValue may have some errors")
}
func (ea *EnumAdapter) UpdateValue(paramValue ValueEnum) error {
	var err error
	if paramValue.Name != "" && paramValue.ID != 0 {
		for _, enValue := range ea.Values {
			if enValue.ID == paramValue.ID {
				enValue = paramValue
				err = nil
			}
		}
		err = errors.New("the ValueEnum with the ID parameter it doesn't exist")
	}
	err = errors.New("the ValueEnum fields are empty")
	return err
}
func (ea *EnumAdapter) DeleteValue(paramValue ValueEnum) error {
	var err error
	if paramValue.Name != "" && paramValue.ID != 0 {
		for idx, value := range ea.Values {
			if value.Name == paramValue.Name {
				tempLeft := ea.Values[idx-1]
				temRight := ea.Values[idx+1]
				ea.Values = []ValueEnum{}
				ea.Values = append(ea.Values, tempLeft)
				ea.Values = append(ea.Values, temRight)
				err = nil
			}
		}
		err = errors.New("the ValueEnum doesn't exist or the ValueEnum fields are wrong")
	}
	return err
}
