// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package job_order

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CreateJobOrder() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"text-2xl font-bold text-center mb-8\">Enter Job Order\r</h1><section class=\"max-w-2xl w-4/5 h-96 mx-auto bg-slate-600 rounded-lg shadow-xl\"><form class=\"rounded-xl flex flex-col gap-4 w-11/12 p-4 mx-auto\" action=\"\" method=\"post\" hx-swap=\"transition:true\"><label class=\"flex flex-col justify-start gap-2\">Job Order Id: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"job_order_id\" required autofocus minlength=\"3\" maxlength=\"64\"></label><div class=\"form-control\"><label class=\"label cursor-pointer\"><span class=\"label-text\">Type</span> <input type=\"checkbox\" name=\"type\" class=\"checkbox\" required></label></div><label class=\"flex flex-col justify-start gap-2\">Deadline: <input class=\"input\" name=\"deadline\" type=\"datetime-local\" required></label><label class=\"flex flex-col justify-start gap-2\">Authorization Letter: <textarea class=\"textarea textarea-primary h-36 max-h-36\" authorization_letter name=\"\" maxlength=\"255\"></textarea></label><label class=\"flex flex-col justify-start gap-2\">Agency Letter: <textarea class=\"textarea textarea-primary h-36 max-h-36\" agency_letter name=\"\" maxlength=\"255\"></textarea></label><label class=\"flex flex-col justify-start gap-2\">Uuid: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"uuid\" autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Signature: <textarea class=\"textarea textarea-primary h-36 max-h-36\" signature name=\"\" maxlength=\"255\"></textarea></label><label class=\"flex flex-col justify-start gap-2\">Stamp: <textarea class=\"textarea textarea-primary h-36 max-h-36\" stamp name=\"\" maxlength=\"255\"></textarea></label><label class=\"flex flex-col justify-start gap-2\">Signed Date: <input class=\"input\" name=\"signed_date\" type=\"datetime-local\"></label><label class=\"flex flex-col justify-start gap-2\">Printed By: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"printed_by\" autofocus minlength=\"3\" maxlength=\"64\"></label><footer class=\"card-actions flex gap-4 justify-end\"><button class=\"btn btn-primary hover:scale-[1.1]\">Save\r</button> <a href=\"/job_order/list\" class=\"btn btn-error hover:scale-[1.1]\">Cancel\r</a></footer></form></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
