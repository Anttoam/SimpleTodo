// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package auth

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/Anttoam/golang-htmx-todos/views/base"

func Login() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"ja\" class=\"h-full bg-white\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = base.Header().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"h-full\"><div class=\"flex min-h-full flex-col justify-center px-6 py-12 lg:px-8\"><div class=\"sm:mx-auto sm:w-full sm:max-w-sm\"><h2 class=\"mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900\">ログイン</h2></div><div class=\"mt-10 sm:mx-auto sm:w-full sm:max-w-sm\"><form class=\"space-y-6\" action=\"/login\" method=\"post\"><div><label for=\"email\" class=\"block text-sm font-medium leading-6 text-gray-600\">メールアドレス</label><div class=\"mt-2\"><input id=\"email\" name=\"email\" type=\"email\" autocomplete=\"email\" required class=\"block w-full rounded-md border-0 py-1.5 pl-3 text-gray-600 shadow-sm ring-1 ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-gray-600 sm:text-sm sm:leading-6\"></div></div><div><div class=\"flex items-center justify-between\"><label for=\"password\" class=\"block text-sm font-medium leading-6 text-gray-600\">パスワード</label><div class=\"text-sm\"><a href=\"#\" class=\"font-semibold text-gray-600 hover:text-gray-300\">パスワードをお忘れですか？</a></div></div><div class=\"mt-2\"><input id=\"password\" name=\"password\" type=\"password\" autocomplete=\"current-password\" required class=\"block w-full rounded-md border-0 py-1.5 pl-3 text-gray-600 shadow-sm ring-1 ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-gray-600 sm:text-sm sm:leading-6\"></div></div><div><button type=\"submit\" class=\"flex w-full justify-center rounded-md bg-gray-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-gray-300 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gray-600\">サインアップ</button></div></form><p class=\"mt-10 text-center text-sm text-gray-500\">アカウントをお持ちですか？ <a href=\"/signup\" class=\"font-bold leading-6 text-gray-600 hover:text-gray-300\">アカウント作成</a></p></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
