package i18n

import (
	"context"
	"github.com/102er/apiserver/pkg/i18n/en"
	"github.com/102er/apiserver/pkg/i18n/zh"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type languageKey struct{}

func NewI18nContext(ctx context.Context, locale string) context.Context {
	var lang = "en"
	switch locale {
	case "zh", "zh-CN":
		lang = "zh"
	}
	return context.WithValue(ctx, languageKey{}, lang)
}

func FromI18nContext(ctx context.Context) (string, bool) {
	lang, ok := ctx.Value(languageKey{}).(string)
	return lang, ok
}

type LocalizeOptions struct {
	MsgParams map[string]interface{} //自定义替换字段参数
	RawMsg    string                 //某找到翻译时 返回的信息 兼容php版本
}

type LocalizeOptionsFunc func(option *LocalizeOptions)

func Localize(ctx context.Context, ID string, opts ...LocalizeOptionsFunc) string {
	option := &LocalizeOptions{
		MsgParams: make(map[string]interface{}),
		RawMsg:    "",
	}
	for _, f := range opts {
		f(option)
	}
	var lang string
	var ok bool
	tag := language.English
	lang, ok = FromI18nContext(ctx)
	if ok && lang == "zh" {
		tag = language.Chinese
	}
	template, ok := templates(lang, ID)
	if !ok {
		if option.RawMsg != "" {
			return option.RawMsg
		}
		return ID
	}
	return i18n.NewLocalizer(i18n.NewBundle(tag), tag.String()).MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    ID,
			Other: template,
		},
		TemplateData: option.MsgParams,
	})
}

func templates(lang, key string) (string, bool) {
	data := en.TranslateToml
	if lang == "zh" {
		data = zh.TranslateToml
	}
	v, ok := data[key]
	return v, ok
}
