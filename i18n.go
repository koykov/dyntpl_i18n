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
		db *i18n.DB
		pr *i18n.PlaceholderReplacer
		ok bool
	)

	raw := dyntpl.GetGlobal(DatabaseKey)
	if raw == nil {
		return nil
	}
	if db, ok = raw.(*i18n.DB); !ok {
		return nil
	}

	if raw = dyntpl.GetGlobal(PlaceholderReplacerKey); raw == nil {
		return nil
	}
	if pr, ok = raw.(*i18n.PlaceholderReplacer); !ok {
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
			if kv, ok := args[i].(*ctxKV); ok {
				ctx.BufAcc.StakeOut().WriteX(kv.v)
				pr.AddKV(fastconv.B2S(kv.k), ctx.BufAcc.StakedString())
			}
		}
	}

	// Compute the key with preceding locale.
	if len(key) == 0 {
		return nil
	}
	lkey := ctx.BufAcc.StakeOut().WriteStr(ctx.loc).WriteByte('.').WriteStr(key).StakedString()

	// Get translation from DB.
	if plural {
		t = db.GetPluralWR(lkey, def, count, &pr)
	} else {
		t = db.GetWR(lkey, def, &pr)
	}
	ctx.BufModStrOut(buf, t)

	return nil
}
