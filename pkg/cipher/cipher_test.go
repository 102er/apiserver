package cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRsaRsaDecryptBase64(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		name     string
		value    string
		expected string
	}{
		{
			name:     "2",
			value:    `NsEJkb2REyGpYnQ8GJrbKOiKi96F73cDL6P0M2cnAw8ETmR0JKK551TPBbLtdPb8U5R23CqxzC89q+9ldOkRHgPKMaAW3/xOhjxdrOkdyEprfa0o7nvOufGl0PvGif3Mm5e0Irvk1DCapuhTSd8cFGKB49+2ttq+gxwzb5jMBVY=`,
			expected: `xxxxx`,
		},
		{
			name:     "3",
			value:    `dfXexieuGETDjTqZsokTKacP8eB/of+fDPVkDO2JzMbkU2PvCRGuPAYKUF03RGMLZPpxzHP2Rn7Fu/n1+Ys9bqHYFg6uBEJXBMLIhM3hYsHB+M/RP6Wlo+rivqHjw76jqBDO+kZw6P4RDnvCQZm0qiMVQjGjQ2t2VHwN8Zex39k=`,
			expected: "xxxxx",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := RsaDecryptBase64(tt.value)
			if err != nil {
				t.Error(err)
				return
			}
			assertions.Equal(tt.expected, res)
			t.Log(res)
		})
	}
}

func TestRsaRsaEecryptBase64(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		name     string
		value    string
		expected string
	}{
		{
			name:     "2",
			expected: `Uy97CdsvStVEfBd4UXoerjaCYgfeszShCHEYoI3+mHlhe3ssoP7SAxPOWnrfl1vkYXTwHK1YBPe2Z2P4l/OsQ64Ir1TUmNi35JANxuTHqhfkNRjNSIm/wzKzMju/wgd3C31uHLPmcjl+iauLDL4g0KG4yIhtKoQZbaoM7EDAGSw=`,
			value:    `xxxxx`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := RsaDecodeBase64(tt.value)
			if err != nil {
				t.Error(err)
				return
			}
			d, err := RsaDecryptBase64(res)
			if err != nil {
				t.Error(err)
				return
			}
			assertions.Equal(d, tt.value)
			t.Log(res)
		})
	}
}
