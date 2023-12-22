package dyntpl_i18n

import "github.com/koykov/dyntpl"

// AcquireCtx takes ctx from default pool and wraps it with locale/direction vars.
func AcquireCtx(locale, direction string) *dyntpl.Ctx {
	ctx := dyntpl.AcquireCtx()
	if len(locale) > 0 {
		ctx.SetLocal(LocaleKey, locale)
	}
	if len(direction) > 0 {
		ctx.SetLocal(DirectionKey, direction)
	}
	return ctx
}

func ReleaseCtx(ctx *dyntpl.Ctx) {
	dyntpl.ReleaseCtx(ctx)
}

var _, _ = AcquireCtx, ReleaseCtx
