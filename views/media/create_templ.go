// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package media

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func CreateMedia() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"text-2xl font-bold text-center mb-8\">Enter Media</h1><section class=\"max-w-2xl w-4/5 h-96 mx-auto bg-slate-600 rounded-lg shadow-xl\"><form class=\"rounded-xl flex flex-col gap-4 w-11/12 p-4 mx-auto\" action=\"\" method=\"post\" hx-swap=\"transition:true\"><label class=\"flex flex-col justify-start gap-2\">Media Id: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"media_id\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Title: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"title\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Url: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"url\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Destination: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"destination\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Mimetype: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"mimetype\" required autofocus minlength=\"3\" maxlength=\"64\"></label><div class=\"form-control\"><label class=\"label cursor-pointer\"><span class=\"label-text\">Is Image</span> <input type=\"checkbox\" name=\"is_image\" class=\"checkbox\" required></label></div><label class=\"flex flex-col justify-start gap-2\">Uploaded By: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"uploaded_by\" required autofocus minlength=\"3\" maxlength=\"64\"></label><footer class=\"card-actions flex gap-4 justify-end\"><button class=\"btn btn-primary hover:scale-[1.1]\">Save</button> <a href=\"/media/list\" class=\"btn btn-error hover:scale-[1.1]\">Cancel</a></footer></form></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
