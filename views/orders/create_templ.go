// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package orders

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func CreateOrders() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"text-2xl font-bold text-center mb-8\">Enter Orders</h1><section class=\"max-w-2xl w-4/5 h-96 mx-auto bg-slate-600 rounded-lg shadow-xl\"><form class=\"rounded-xl flex flex-col gap-4 w-11/12 p-4 mx-auto\" action=\"\" method=\"post\" hx-swap=\"transition:true\"><label class=\"flex flex-col justify-start gap-2\">Order Id: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"order_id\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Invoice No: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"invoice_no\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Customer Id: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"customer_id\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Name: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"name\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Country Code: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"country_code\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Mobile: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"mobile\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Country: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"country\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">City: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"city\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Line1: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"line1\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Location: <input class=\"input input-bordered input-primary\" type=\"number\" name=\"location\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Payment Method: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"payment_method\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Charge Id: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"charge_id\" required autofocus minlength=\"3\" maxlength=\"64\"></label><div class=\"form-control\"><label class=\"label cursor-pointer\"><span class=\"label-text\">Status Id</span> <input type=\"checkbox\" name=\"status_id\" class=\"checkbox\" required></label></div><label class=\"flex flex-col justify-start gap-2\">Tracking: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"tracking\" required autofocus minlength=\"3\" maxlength=\"64\"></label><label class=\"flex flex-col justify-start gap-2\">Tracking Company: <input class=\"input input-bordered input-primary\" type=\"text\" name=\"tracking_company\" required autofocus minlength=\"3\" maxlength=\"64\"></label><footer class=\"card-actions flex gap-4 justify-end\"><button class=\"btn btn-primary hover:scale-[1.1]\">Save</button> <a href=\"/orders/list\" class=\"btn btn-error hover:scale-[1.1]\">Cancel</a></footer></form></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
