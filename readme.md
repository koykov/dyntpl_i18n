# Dyntpl i18n bindings

Provide [i18n](https://github.com/koykov/i18n) support to use in [dyntpl](https://github.com/koykov/dyntpl) templates.

Internationalization support provides by [i18n](https://github.com/koykov/i18n) package.

I18n enables by setup context with variables like this:
```go
var db *i18n.DB
ctx := dyntpl.AcquireCtx()
defer dyntpl.ReleaseCtx(ctx)
ctx.SetStatic(dyntpl_i18n.DatabaseKey, db).
	SetString(dyntpl_i18n.LocaleKey, "en")
...
```
or by using shorthand
```go
var db *i18n.DB
ctx := dyntpl_i18n.AcquireCtx("en", "ltr", db)
defer dyntpl_i18n.Release(ctx)
...
```

For simple translate use function `i18n::template` or shorthand `i18n::t`:
```
{%= i18n::t("key", "default value", {"!placeholder0": "replacement", "!placeholder1": object.Label, ...}) %}
```
You may omit default value and replacements, only first argument is required.

For plural translation use function `i18n::translatePlural` or shorthand `i18n::tp`:
```
{%= i18n::tp("key", "default value", 15, {...}) %}
```
Third argument is a count for a plural formula. It's required as a `key` argument.
