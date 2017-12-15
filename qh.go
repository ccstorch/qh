package qh

import (
	"fmt"
	"strings"
)

type ParamsMatching struct {
	Params       []string
	ParamsCount  int
	ParamsValues []interface{}
}

func (pm *ParamsMatching) Append(colName string, value interface{}) {
	pm.Params = append(pm.Params, fmt.Sprintf("%s = $%d", colName, pm.ParamsCount))
	pm.ParamsCount++
	pm.ParamsValues = append(pm.ParamsValues, value)
}

func (pm *ParamsMatching) AppendWithFunction(colName string, sqlFuncName string, value interface{}) {
	pm.Params = append(pm.Params, fmt.Sprintf("%s = %s($%d)", colName, sqlFuncName, pm.ParamsCount))
	pm.ParamsCount++
	pm.ParamsValues = append(pm.ParamsValues, value)
}

func (pm *ParamsMatching) JoinColNames() string {
	return strings.Join(pm.Params, ", ")
}
