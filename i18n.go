package dyntpl_i18n

import (
	"bytes"

	"github.com/koykov/dyntpl"
	"github.com/koykov/fastconv"
	"github.com/koykov/i18n"
)

const (
	DatabaseKey            = "__i18n_db__"
	PlaceholderReplacerKey = "__i18n_pr__"
	LocaleKey              = "i18n.locale"
)

var (
	defEmpty  = []byte(`""`)
	defaultPR i18n.PlaceholderReplacer
)

// Translate label.
func modTranslate(ctx *dyntpl.Ctx, buf *any, _ any, args []any) error {
	return trans(ctx, buf, args, false)
}

// Translate label with plural formula.
func modTranslatePlural(ctx *dyntpl.Ctx, buf *any, _ any, args []any) error {
	return trans(ctx, buf, args, true)
}

// Generic translate function.
func trans(ctx *dyntpl.Ctx, buf *any, args []any, plural bool) error {
	var (
		raw any
		db  *i18n.DB
		pr  *i18n.PlaceholderReplacer
		loc string
		ok  bool
	)

	if raw, ok = getVar(ctx, DatabaseKey); !ok {
		return nil
	}
	if db, ok = raw.(*i18n.DB); !ok || db == nil {
		return nil
	}

	if raw, ok = getVar(ctx, PlaceholderReplacerKey); !ok {
		return nil
	}
	if pr, ok = raw.(*i18n.PlaceholderReplacer); !ok {
		return nil
	}
	if pr == nil {
		pr = &defaultPR
	}

	if len(args) == 0 {
		return dyntpl.ErrModNoArgs
	}

	var (
		key, def, t string
		count       = 1
	)
	// Try to get the key.
	if raw, ok := args[0].(*[]byte); ok && len(*raw) > 0 {
		key = fastconv.B2S(*raw)
	} else if err := ctx.BufAcc.StakeOut().WriteX(args[0]).Error(); err == nil {
		key = ctx.BufAcc.StakedString()
	}
	args = args[1:]
	// Try to get the default value.
	if len(args) > 0 {
		if raw, ok := args[0].(*[]byte); ok && len(*raw) > 0 && !bytes.Equal(*raw, defEmpty) {
			def = fastconv.B2S(*raw)
		}
		args = args[1:]
	}
	// Try to get count to use in plural formula.
	if plural {
		if len(args) > 0 {
			if raw, ok := args[0].(int); ok {
				count = raw
			}
			args = args[1:]
		}
	}

	// Collect placeholder replacements.
	pr.Reset()
	if len(args) > 0 {
		_ = args[len(args)-1]
		for i := 0; i < len(args); i++ {
			if kv, ok := args[i].(*dyntpl.KV); ok {
				ctx.BufAcc.StakeOut().WriteX(kv.V)
				pr.AddKV(fastconv.B2S(kv.K), ctx.BufAcc.StakedString())
			}
		}
	}

	// Compute the key with preceding locale.
	if len(key) == 0 {
		return nil
	}
	if raw, ok = getVar(ctx, LocaleKey); !ok {
		return nil
	}
	if loc, ok = raw.(string); !ok || len(loc) == 0 {
		return nil
	}
	lkey := ctx.BufAcc.StakeOut().WriteStr(loc).WriteByte('.').WriteStr(key).StakedString()

	// Get translation from DB.
	if plural {
		t = db.GetPluralWR(lkey, def, count, pr)
	} else {
		t = db.GetWR(lkey, def, pr)
	}
	ctx.BufModStrOut(buf, t)

	return nil
}

func getVar(ctx *dyntpl.Ctx, name string) (any, bool) {
	v := ctx.GetLocal(name)
	if v != nil {
		return v, true
	}
	if v = dyntpl.GetGlobal(name); v != nil {
		return v, true
	}
	return nil, false
}
