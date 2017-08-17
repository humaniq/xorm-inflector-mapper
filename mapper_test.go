package mapper_test

import (
	"github.com/go-xorm/core"
	"github.com/humaniq/xorm-inflector-mapper"
	"testing"
)

var (
	convertionTests = []struct {
		objName, tableName string
	}{
		{"User", "users"},
		{"Person", "people"},
		{"Customer", "customers"},
		{"Gallery", "galleries"},
		{"UserRole", "user_roles"},
	}
)

func TestInflectorMapper_Obj2Table(t *testing.T) {
	testMapper := mapper.NewInflectorMapper(core.SnakeMapper{})

	for _, tt := range convertionTests {
		actual := testMapper.Obj2Table(tt.objName)
		if actual != tt.tableName {
			t.Errorf("Obj2Table(%s) expected %s, actual: %s", tt.objName, actual, tt.tableName)
		}
	}
}

func TestInflectorMapper_Table2Obj(t *testing.T) {
	testMapper := mapper.NewInflectorMapper(core.SnakeMapper{})

	for _, tt := range convertionTests {
		actual := testMapper.Table2Obj(tt.tableName)
		if actual != tt.objName {
			t.Errorf("Table2Obj(%s) expected %s, actual: %s", tt.tableName, actual, tt.objName)
		}
	}
}
