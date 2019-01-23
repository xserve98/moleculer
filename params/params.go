package params

import (
	. "github.com/moleculer-go/moleculer/common"
)

type ParamsImpl struct {
	source *interface{}
	values *map[string]interface{}
}

func (params ParamsImpl) Int(name string) int {
	if value, ok := (*params.values)[name]; ok {
		return value.(int)
	}
	return 0

}

func (params ParamsImpl) Int64(name string) int64 {
	if value, ok := (*params.values)[name]; ok {
		return value.(int64)
	}
	return 0
}

func (params ParamsImpl) Float(name string) float32 {
	if value, ok := (*params.values)[name]; ok {
		return value.(float32)
	}
	return 0
}

func (params ParamsImpl) Float64(name string) float64 {
	if value, ok := (*params.values)[name]; ok {
		return value.(float64)
	}
	return 0
}

func (params ParamsImpl) String(name string) string {
	if value, ok := (*params.values)[name]; ok {
		return value.(string)
	}
	return ""
}

func (params ParamsImpl) Get(name string) string {
	return params.String(name)
}

func CreateParams(source *interface{}) Params {
	//TODO parse source and feed into values

	values := make(map[string]interface{})
	return ParamsImpl{source, &values}
}
