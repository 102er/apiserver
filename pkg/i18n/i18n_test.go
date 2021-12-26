package i18n

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalize(t *testing.T) {
	ctx := context.TODO()
	ctx = NewI18nContext(ctx, "zh")
	assertions := assert.New(t)
	var tests = []struct {
		Name     string
		ID       string
		f        []LocalizeOptionsFunc
		expected string
	}{
		{
			Name:     "hello world no set name",
			ID:       "hello_world",
			f:        []LocalizeOptionsFunc{},
			expected: "你好 <no value>",
		},
		{
			Name: "hello set name",
			ID:   "hello_world",
			f: []LocalizeOptionsFunc{func(option *LocalizeOptions) {
				option.MsgParams = map[string]interface{}{
					"name": "apiserver",
				}
			}},
			expected: "你好 apiserver",
		},
		{
			Name: "set msg param",
			ID:   "time_duration",
			f: []LocalizeOptionsFunc{func(option *LocalizeOptions) {
				option.MsgParams = map[string]interface{}{
					"d": 11,
					"h": 11,
					"m": 11,
					"s": 11,
				}
			}},
			expected: "11天11时11分11秒",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			res := Localize(ctx, tt.ID, tt.f...)
			assertions.Equal(tt.expected, res)
			t.Log(res)
		})
	}
}
