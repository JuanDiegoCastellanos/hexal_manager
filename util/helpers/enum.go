package helpers

// ValueEnum is a struct used to simulate the enum string
type ValueEnum struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// EnumEntity is the interface that all the entities where their need the EnumValue
// they have to implement this interface
// EnumEntity has some single methods to add, delete, update, getAll and getOneSpecific
type EnumEntity interface {
	GetValues() []ValueEnum
	GetValue(enum ValueEnum) *ValueEnum
	AddValue(enum ValueEnum) error
	UpdateValue(enum ValueEnum) error
	DeleteValue(enum ValueEnum) error
}

// EnumAdapter is the adapter in order to each entity implements its specific enum
type EnumAdapter struct {
	Values []ValueEnum
}

func (ea *EnumAdapter) GetValues() *[]ValueEnum{
	return ea.
}
func (ea *EnumAdapter) GetValue(searchParam interface{}) interface{} {
	for idx, value := range ea.Values {
		if value == searchParam {
			return value
		}
	}
}

var Sate = []string{"free", "busy", "remodeling"}
var ParkingType = []string{"large", "medium", "small"}
var SocialSpaces = []string{"pool_are", "fitness_"}
