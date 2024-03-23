// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package option_value

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CreateOptionValue() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"text-2xl font-bold text-center mb-8\">Enter Option Value\r</h1><section class=\"max-w-2xl w-4/5 h-96 mx-auto bg-slate-600 rounded-lg shadow-xl\"><form class=\"rounded-xl flex flex-col gap-4 w-11/12 p-4 mx-auto\" action=\"\" method=\"post\" hx-swap=\"transition:true\"><label class=\"flex flex-col justify-start gap-2\">Option Value Id: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"option_value_id\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Option Id: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"option_id\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Media Id: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"media_id\" autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Sort Order: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"sort_order\" required autofocus minlength=\"3\" maxlength=\"64\"></label><footer class=\"card-actions flex gap-4 justify-end\"><button class=\"btn btn-primary hover:scale-[1.1]\">Save\r</button> <a href=\"/option_value/list\" class=\"btn btn-error hover:scale-[1.1]\">Cancel\r</a></footer></form></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
