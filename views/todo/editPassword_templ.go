// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package todo

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/Anttoam/golang-htmx-todos/views/base"
	"github.com/a-h/templ"
)

func EditPassword(id string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"ja\" class=\"h-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = base.Header().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"bg-gray-200\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = NavBar(id).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"page\" class=\"flex min-h-full flex-col justify-center px-6 py-12 lg:px-8\"><div class=\"sm:mx-auto sm:w-full sm:max-w-sm\"><h2 class=\"mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900\">パスワード編集</h2></div><div class=\"mt-10 sm:mx-auto sm:w-full sm:max-w-sm\"><form class=\"space-y-6\" hx-put=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/user/password/%s", id))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/todo/editPassword.templ`, Line: 19, Col: 74}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#edit-page\" hx-swap=\"outerHTML\"><div><label for=\"password\" class=\"block text-sm font-medium leading-6 text-gray-600\">現在のパスワード</label><div class=\"mt-2\"><input id=\"password\" name=\"password\" type=\"password\" autocomplete=\"password\" placeholder=\"現在のパスワードを入力してください\" required class=\"block w-full rounded-md border-0 py-1.5 pl-3 text-gray-600 shadow-sm ring-1 ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-gray-600 sm:text-sm sm:leading-6\"></div></div><div><label for=\"password\" class=\"block text-sm font-medium leading-6 text-gray-600\">新しいパスワード</label><div class=\"mt-2\"><input id=\"password\" name=\"new_password\" type=\"password\" autocomplete=\"new-password\" placeholder=\"新しいパスワードを入力してください\" required class=\"block w-full rounded-md border-0 py-1.5 pl-3 text-gray-600 shadow-sm ring-1 ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-gray-600 sm:text-sm sm:leading-6\"></div></div><button type=\"submit\" class=\"flex w-full justify-center rounded-md bg-gray-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-gray-300 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gray-600\">更新</button></form><div hx-boost=\"true\"><p class=\"mt-10 text-center text-sm text-gray-500\">パスワードの編集をキャンセルしますか？ <a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 templ.SafeURL = templ.SafeURL(fmt.Sprintf("/user/%s", id))
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var3)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"font-bold leading-6 text-gray-600 hover:text-gray-300\">はい</a></p></div></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}