// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package coupon

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CreateCoupon() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"text-2xl font-bold text-center mb-8\">Enter Coupon\r</h1><section class=\"max-w-2xl w-4/5 h-96 mx-auto bg-slate-600 rounded-lg shadow-xl\"><form class=\"rounded-xl flex flex-col gap-4 w-11/12 p-4 mx-auto\" action=\"\" method=\"post\" hx-swap=\"transition:true\"><label class=\"flex flex-col justify-start gap-2\">Coupon Id: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"coupon_id\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Title: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"title\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Code: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"code\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"form-control w-full max-w-xs\"><div class=\"label\"><span class=\"label-text\">Type</span></div><select class=\"select select-bordered\" name=\"type\" required><option disabled selected>Select one</option> <option>P</option> <option>F</option></select></label><label class=\"flex flex-col justify-start gap-2\">Date Start: <input class=\"input\" name=\"date_start\" type=\"datetime-local\" required></label><label class=\"flex flex-col justify-start gap-2\">Date End: <input class=\"input\" name=\"date_end\" type=\"datetime-local\" required></label><label class=\"flex flex-col justify-start gap-2\">Total Limit: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"total_limit\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Customer Limit: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"customer_limit\" required autofocus minlength=\"3\" maxlength=\"64\"></label><div class=\"form-control\"><label class=\"label cursor-pointer\"><span class=\"label-text\">Status</span> <input type=\"checkbox\" name=\"status\" class=\"checkbox\" required></label></div><footer class=\"card-actions flex gap-4 justify-end\"><button class=\"btn btn-primary hover:scale-[1.1]\">Save\r</button> <a href=\"/coupon/list\" class=\"btn btn-error hover:scale-[1.1]\">Cancel\r</a></footer></form></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
