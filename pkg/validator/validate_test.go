package validator

import "testing"

func TestValidateData(t *testing.T) {
	type testStruct struct {
		Name      string `json:"name" validate:"required" label:"函数名"`
		StartTime int    `json:"start_time" validate:"required,gt=0" label:"开始时间"`
		EndTime   int    `json:"end_time"`
		Interval  string `json:"interval" validate:"required" label:"时间间隔"`
	}
	d := testStruct{}
	err := ValidateData("en", d)
	if err != nil {
		return
	}
}
