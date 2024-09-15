// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package navbartempl

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func (n NavbarTempl) Navbar(navbarStyl Navbarstyl, id string) templ.Component {
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
		var templ_7745c5c3_Var2 = []any{" mobile:px-10 lg:px-20 py-4 flex justify-between " + navbarStyl.Bg}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(id)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `business/domain/uinavbarbus/templs/navbartempl/navbar.templ`, Line: 4, Col: 13}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `business/domain/uinavbarbus/templs/navbartempl/navbar.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"flex items-center gap-10\"><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(navbarStyl.Logo)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `business/domain/uinavbarbus/templs/navbartempl/navbar.templ`, Line: 6, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-[6.188rem] h-7\" alt=\"logo\"><div class=\"flex items-center gap-8 mobile:hidden\"><a href=\"#\">services</a> <a href=\"#\">Support</a> <a href=\"#\">About us</a> <a href=\"#\">Connect us</a></div></div><div class=\"flex items-center gap-3 mobile:hidden\"><div id=\"languages_parent\" class=\"relative cursor-pointer\"><div id=\"inner_lang\" class=\"flex items-center gap-3\"><img src=\"/assets/public/icons/UpdateMain/english.svg\" alt=\"\"> <span>English</span> <img src=\"/assets/public/icons/UpdateMain/arrow.svg\" alt=\"\"></div><div id=\"languages\" class=\"flex hidden text-[#1D1B20] flex-col gap-5 py-5 z-50 top-10 px-5 justify-center items-center absolute bg-white rounded-lg shadow\"><div class=\"flex cursor-pointer items-center pr-10 gap-2\"><img src=\"/assets/public/icons/UpdateMain/farsi.svg\" alt=\"...\"> <span>Farsi</span></div><div class=\"flex cursor-pointer items-center gap-2\"><img src=\"/assets/public/icons/UpdateMain/arabic.svg\" alt=\"...\"> <span>Arabic</span></div><div class=\"flex cursor-pointer items-center gap-2\"><img src=\"/assets/public/icons/UpdateMain/english.svg\" alt=\"...\"> <span>English</span></div></div></div><div id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(n.Profile.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `business/domain/uinavbarbus/templs/navbartempl/navbar.templ`, Line: 41, Col: 25}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = n.Profile.Componenet.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div id=\"sidebar_menu_icon\" class=\"hidden mobile:block\"><img class=\"h-8 cursor-pointer\" src=\"/assets/public/icons/main/menu.svg\" alt=\"menu\"><div id=\"sidebar_menu\" class=\"fixed mobile:hidden left-0 top-0 w-[100vw] h-[100vh] bg-[#0000004D] backdrop-blur-[0.125rem] z-[100]\"><div class=\"w-[21.25rem] h-full bg-white rtl:rounded-l-[2rem] ltr:rounded-r-[2rem] flex flex-col justify-between\"><div><div class=\"mt-8 flex gap-4 items-center text-lg px-6 justify-between text-[2.15rem] font-bold text-[#1C0E07]\"><img class=\"h-8 cursor-pointer\" src=\"/assets/public/images/UpdateMain/logo.svg\" alt=\"search\"> <img class=\"h-8 cursor-pointer\" src=\"/assets/public/icons/main/close.svg\" alt=\"close\"></div><ul class=\"mt-10 text-xl text-[#5B4A42] px-6 space-y-3\"><li class=\"text-crs-first bg-crs-active h-14 rounded-[0.75rem] flex gap-4 px-4 items-center cursor-pointer\">Home</li><li class=\"h-14 rounded-[0.75rem] flex gap-4 px-4 items-center cursor-pointer\">Contact Us</li><li class=\"h-14 rounded-[0.75rem] flex gap-4 px-4 items-center cursor-pointer\">About Us</li></ul></div><div class=\"px-6\"><div dir=\"rtl\" class=\"flex p-1 mb-16 h-[3rem] bg-[#F7F7F9] rounded-[0.5rem] min-w-[13.25rem] text-[#514037]\"><div class=\"w-1/3 grow cursor-pointer flex items-center justify-center rounded-[0.5rem]\">فارسی</div><div class=\"w-1/3 grow cursor-pointer flex items-center justify-center rounded-[0.5rem]\">عربی</div><div class=\"w-1/3 grow cursor-pointer flex items-center justify-center rounded-[0.5rem] text-crs-first bg-white\">English</div></div></div></div></div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
