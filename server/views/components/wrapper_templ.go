// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "axox/gaingauge/views/styles"

func sideBar() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{ navOpen: false }\"><div class=\"mobile-topbar\"><svg x-on:click=\"navOpen = !navOpen\" width=\"30\" height=\"30\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M4 18L20 18\" stroke=\"#000000\" stroke-width=\"2\" stroke-linecap=\"round\"></path> <path d=\"M4 12L20 12\" stroke=\"#000000\" stroke-width=\"2\" stroke-linecap=\"round\"></path> <path d=\"M4 6L20 6\" stroke=\"#000000\" stroke-width=\"2\" stroke-linecap=\"round\"></path></svg></div><div class=\"sidenav-container\" :class=\"navOpen ? &#39;open&#39; : &#39;&#39;\"><div class=\"sidenav-links\"><a href=\"/\"><span class=\"sidenav-icon\">💪 </span> <span class=\"sidenav-text\">Dashboard </span></a> <a href=\"/log\"><span class=\"sidenav-icon\">🗒</span> <span class=\"sidenav-text\">Log </span></a></div></div></div><style>\n    .sidenav-container {\n        position: fixed;\n        top: 0;\n        left: 0;\n        display: flex;\n        flex-direction: column;\n        box-shadow: 2px 0 4px rgba(0, 0, 0, .25);\n        background-color: var(--accent);\n        width: 64px;\n        height: 100%;\n        overflow: hidden;\n        transition: width 0.2s;\n    }\n\n    .mobile-topbar {\n        display: none;\n    }\n\n    .open {\n        width: 100% !important;\n    }\n\n    .open .sidenav-text {\n        opacity: 1 !important;\n    }\n\n    @media (max-width: 768px),\n    (hover: none) {\n        .sidenav-container {\n            width: 0px;\n            margin-top: 50px;\n        }\n\n        .content-wrapper {\n            margin-left: 0;\n        }\n\n        .mobile-topbar {\n            display: flex;\n            height: 50px;\n            flex-direction: row;\n            background-color: var(--accent);\n            align-items: center;\n            padding-left: 20px;\n        }\n    }\n\n    .sidenav-links {\n        display: flex;\n        flex-direction: column;\n        width: 200px;\n        row-gap: 20px;\n    }\n\n    .sidenav-icon {\n        width: 64px;\n        display: inline-block;\n        text-align: center;\n    }\n\n    .sidenav-text {\n        transition: opacity 0.2s;\n        font-size: 1.5rem;\n        margin-left: -12px;\n    }\n\n    .sidenav-links>a {\n        font-size: 2rem;\n        text-decoration: none;\n    }\n\n    .sidenav-container .sidenav-text {\n        opacity: 0;\n    }\n\n    .sidenav-container:hover {\n        width: 250px;\n    }\n\n    .sidenav-container:hover .sidenav-text {\n        opacity: 1;\n    }\n</style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func PageWrapper() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><script src=\"https://unpkg.com/htmx.org@2.0.1\" integrity=\"sha384-QWGpdj554B4ETpJJC9z+ZHJcA/i59TyjxEPXiiUgN2WmTyV5OEZWCD6gQhgkdpB/\" crossorigin=\"anonymous\"></script><script defer src=\"https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js\"></script><head><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><title>GainGauge</title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = styles.RootStyling().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n        .content-wrapper {\n            margin-left: 64px;\n        }\n\n        body {\n            width: 100%;\n            height: 100%;\n            margin: 0;\n        }\n    </style></head><body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = sideBar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"content-wrapper\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var2.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
