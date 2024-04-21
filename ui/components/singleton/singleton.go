package singleton

type ValueServiceSingleton struct {
	valueService *ValueService
}

func (v *ValueServiceSingleton) GetService() *ValueService {
	if v.valueService == nil {
		v.valueService = NewValueService()
	}
	return v.valueService
}

type ValueService struct {
	value string
}

func NewValueService() *ValueService {
	return &ValueService{value: ""}
}

func (v *ValueService) GetValue() string {
	return v.value
}

func (v *ValueService) SetValue(input string) {
	v.value = input
}
