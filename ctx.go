package dyntpl_i18n

import (
	"github.com/koykov/dyntpl"
	"github.com/koykov/i18n"
)

// AcquireCtx takes ctx from default pool and wraps it with locale/direction vars.
func AcquireCtx(locale, direction string, db *i18n.DB) *dyntpl.Ctx {
	ctx := dyntpl.AcquireCtx()
	ctx.SetStatic(DatabaseKey, db)
	if len(locale) > 0 {
		ctx.SetString(LocaleKey, locale)
	}
	if len(direction) > 0 {
		ctx.SetString(DirectionKey, direction)
	}
	return ctx
}

func ReleaseCtx(ctx *dyntpl.Ctx) {
	dyntpl.ReleaseCtx(ctx)
}

var _, _ = AcquireCtx, ReleaseCtx
