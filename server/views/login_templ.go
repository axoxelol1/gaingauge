// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Login() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><script src=\"https://unpkg.com/htmx.org@2.0.1\" integrity=\"sha384-QWGpdj554B4ETpJJC9z+ZHJcA/i59TyjxEPXiiUgN2WmTyV5OEZWCD6gQhgkdpB/\" crossorigin=\"anonymous\"></script><head><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><title>Login</title></head><body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = LoginForm().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body><style>\n    :root {\n        --primary: #FFFFFF;\n        --accent: #00b8ff;\n        --dark-text: #000000;\n        --light-text: #FFFFFF;\n        font-family: sans-serif;\n        font-optical-sizing: auto;\n        font-style: normal;\n    }\n\n    body {\n        display: flex;\n        justify-content: center;\n        align-content: center;\n    }\n\n    button {\n        background-color: var(--accent);\n        padding: 1rem;\n        border-radius: 10px;\n        transition-duration: 0.2s;\n    }\n\n    input {\n        width: 20em\n    }\n\n    button:hover {\n        box-shadow: 0 12px 16px 0 rgba(0, 0, 0, 0.24), 0 17px 50px 0 rgba(0, 0, 0, 0.19);\n    }\n\n    form {\n        background-color: --white;\n        display: flex;\n        flex-direction: column;\n        justify-content: center;\n        margin-top: 2rem;\n        gap: 5px;\n        padding: 1rem;\n        border: 1px solid rgba(0, 0, 0, 0.19);\n        width: 500px;\n    }\n\n    #button-wrapper {\n        display: flex;\n        flex-direction: row;\n    }\n\n    #inputs {\n        display: flex;\n        flex-direction: column;\n        align-items: start;\n    }\n</style></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func LoginForm() templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form id=\"form\"><div id=\"inputs\"><div class=\"labeltext\"><label for=\"username\">Username</label> <input id=\"username\" name=\"username\" type=\"text\"></div><div class=\"labeltext\"><label for=\"password\">Password</label> <input id=\"password\" name=\"password\" type=\"password\"></div></div><div id=\"button-wrapper\"><button type=\"submit\" hx-post=\"/login\" hx-target=\"#loginerror\">Log in</button> <button type=\"button\" hx-get=\"/register\" hx-target=\"#form\" hx-swap=\"outerHTML\">Sign up instead</button></div><div id=\"loginerror\"></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func RegisterForm() templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form id=\"form\"><div id=\"inputs\"><div class=\"labeltext\"><label for=\"username\">Username</label> <input id=\"username\" name=\"username\" type=\"text\"></div><div class=\"labeltext\"><label for=\"password\">Password</label> <input id=\"password\" name=\"password\" type=\"password\"></div><div class=\"labeltext\"><label for=\"firstname\">First name</label> <input id=\"firstname\" name=\"firstname\" type=\"text\"></div><div class=\"labeltext\"><label for=\"lastname\">Last name</label> <input id=\"lastname\" name=\"lastname\" type=\"text\"></div></div><div id=\"button-wrapper\"><button type=\"submit\" hx-post=\"/register\" hx-target=\"#registererror\">Register</button> <button type=\"button\" hx-get=\"/login\" hx-target=\"#form\" hx-swap=\"outerHTML\">Already have an account?</button></div><div id=\"registererror\"></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
