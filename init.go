package dyntpl_i18n

import "github.com/koykov/dyntpl"

func init() {
	dyntpl.RegisterModFnNS("i18n", "translate", "t", modTranslate).
		WithParam("key string", "Machine-redable key in i18n database.").
		WithParam("default string", "Default value if key not found in database.").
		WithParam("placeholders object", "Associative array with placeholders.").
		WithDescription("Translate key to current locale.").
		WithExample(`{%= i18n::t("key", "default value", {"!placeholder0": "replacement", "!placeholder1": object.Label, ...}) %}`)
	dyntpl.RegisterModFnNS("i18n", "translatePlural", "tp", modTranslatePlural).
		WithParam("key string", "Machine-redable key in i18n database.").
		WithParam("default string", "Default value if key not found in database.").
		WithParam("count int", "Value to choose correct plural form.").
		WithParam("placeholders object", "Associative array with placeholders.").
		WithDescription("Translate key to current locale considering pluralization rules.").
		WithExample(`i18nDB.Set("key", "{0} There are none|[1,19] There are some|[20,*] There are many")
{%= i18n::tp("key", "default value", 15, {...}) %} // There are some`)
	dyntpl.RegisterModFnNS("i18n", "locale", "l", modSetLocale).
		WithParam("locale string", "Locale name known by i18n database.").
		WithDescription("Set up the current locale inside template.")

	// Legacy modifiers support.
	dyntpl.RegisterModFn("translate", "t", modTranslate).
		WithDescription("Alias of `i18n::translate`. Legacy support issue.")
	dyntpl.RegisterModFn("translatePlural", "tp", modTranslatePlural).
		WithDescription("Alias of `i18n::translatePlural`. Legacy support issue.")
	dyntpl.RegisterModFn("locale", "", modSetLocale).
		WithDescription("Alias of `i18n::locale`. Legacy support issue.")
}
