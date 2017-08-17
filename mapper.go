package mapper

import (
	"github.com/gedex/inflector"
	"github.com/go-xorm/core"
)

type InflectorMapper struct {
	Mapper core.IMapper
}

func NewInflectorMapper(mapper core.IMapper) InflectorMapper {
	return InflectorMapper{mapper}
}

func (mapper InflectorMapper) Obj2Table(name string) string {
	return mapper.Mapper.Obj2Table(inflector.Pluralize(name))
}

func (mapper InflectorMapper) Table2Obj(name string) string {
	return mapper.Mapper.Table2Obj(inflector.Singularize(name))
}
