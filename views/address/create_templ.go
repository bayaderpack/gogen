// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package address

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CreateAddress() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"text-2xl font-bold text-center mb-8\">Enter Address\r</h1><section class=\"max-w-2xl w-4/5 h-96 mx-auto bg-slate-600 rounded-lg shadow-xl\"><form class=\"rounded-xl flex flex-col gap-4 w-11/12 p-4 mx-auto\" action=\"\" method=\"post\" hx-swap=\"transition:true\"><label class=\"flex flex-col justify-start gap-2\">Address Id: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"address_id\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Customer Id: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"customer_id\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Firstname: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"firstname\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Lastname: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"lastname\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Country Code: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"country_code\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Mobile: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"mobile\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Line1: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"line1\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Location: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"location\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Country: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"country\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">City: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"city\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">State Code: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"state_code\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Post Code: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"post_code\" required autofocus minlength=\"3\" maxlength=\"64\"></label><div class=\"form-control\"><label class=\"label cursor-pointer\"><span class=\"label-text\">Is Primary</span> <input type=\"checkbox\" name=\"is_primary\" class=\"checkbox\" required></label></div><footer class=\"card-actions flex gap-4 justify-end\"><button class=\"btn btn-primary hover:scale-[1.1]\">Save\r</button> <a href=\"/address/list\" class=\"btn btn-error hover:scale-[1.1]\">Cancel\r</a></footer></form></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
