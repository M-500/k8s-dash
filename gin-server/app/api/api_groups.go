//@Author: wulinlin
//@Description:
//@File:  api_groups
//@Version: 1.0.0
//@Date: 2024/01/02 21:26

package api

type ApiGroup struct {
}

func NewApiGroups() *ApiGroup {
	return &ApiGroup{}
}

var MainAPPGroups = NewApiGroups()
