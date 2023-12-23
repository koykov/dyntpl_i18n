package dyntpl_i18n

import "github.com/koykov/dyntpl"

func init() {
	dyntpl.RegisterModFnNS("i18n", "translate", "t", modTranslate)
	dyntpl.RegisterModFnNS("i18n", "translatePlural", "tp", modTranslatePlural)
	dyntpl.RegisterModFnNS("i18n", "locale", "l", modSetLocale)

	// Legacy modifiers support.
	dyntpl.RegisterModFn("translate", "t", modTranslate)
	dyntpl.RegisterModFn("translatePlural", "tp", modTranslatePlural)
	dyntpl.RegisterModFn("locale", "", modSetLocale)
}
