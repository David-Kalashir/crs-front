// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package logintempl

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func (l LoginTemplate) Login() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"max-w-md mx-auto p-8 bg-white rounded-md shadow-md\"><h2 class=\"text-2xl font-semibold mb-6\">Sign In</h2><form action=\"#\" method=\"POST\"><div class=\"mb-4\"><label for=\"name\" class=\"block text-gray-700 text-sm font-bold mb-2\">username</label> <input type=\"text\" id=\"name\" name=\"name\" placeholder=\"John Doe\" required class=\"w-full px-3 py-2 border rounded-md focus:outline-none focus:border-blue-500\"></div><div class=\"mb-4\"><label for=\"password\" class=\"block text-gray-700 text-sm font-bold mb-2\">password</label> <input type=\"password\" id=\"password\" name=\"password\" placeholder=\"********\" required class=\"w-full px-3 py-2 border rounded-md focus:outline-none focus:border-blue-500\"></div><button type=\"submit\" class=\"bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline-blue\">Login</button></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
