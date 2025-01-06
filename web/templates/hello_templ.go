// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Index() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Document</title><link rel=\"stylesheet\" href=\"/style.css\"><link href=\"https://cdn.jsdelivr.net/npm/quill@2.0.3/dist/quill.snow.css\" rel=\"stylesheet\"></head><body><nav class=\"border-b border-black bg-white\"><div class=\"mx-auto max-w-7xl px-4 sm:px-6 lg:px-8\"><div class=\"flex h-16 items-center justify-between\"><!-- Left Section: Text Input --><div class=\"flex items-center\"><input type=\"text\" placeholder=\"Title\" class=\"rounded-md border border-gray-300 px-4 py-2 text-sm focus:border-transparent focus:outline-none focus:ring-1 focus:ring-black\"></div><!-- Right Section: Save Button --><!-- Right Section: Save Button and Dropdown --><div class=\"flex items-center space-x-4\"><!-- Save Button --><button class=\"rounded-md bg-black px-4 py-2 text-white hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-500\">Save</button><!-- Dropdown --><div class=\"group relative\"><button class=\"rounded-md bg-black px-4 py-2 text-white hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-500\">Settings</button><!-- Dropdown Menu --><div class=\"invisible absolute right-0 mt-2 w-48 rounded-md border border-gray-200 bg-white opacity-0 shadow-lg transition-all duration-200 group-hover:visible group-hover:opacity-100\"><a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100\">Option 1</a> <a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100\">Option 2</a> <a href=\"#\" class=\"block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100\">Option 3</a></div></div></div></div></div></nav><div class=\"grid grid-cols-2\"><textarea id=\"raw\" class=\"h-screen border-e border-black p-1\"></textarea><div class=\"grid grid-rows-12 h-screen\"><!--<div id=\"toolbar\" class=\"border-b border-black\"></div>--><div id=\"editor\" class=\"row-span-11\"></div></div></div></body><!-- Include the Quill library --><script src=\"https://cdn.jsdelivr.net/npm/quill@2.0.3/dist/quill.js\"></script><!-- Initialize Quill editor --><script>\r\n  const quill = new Quill('#editor', {\r\n      theme: 'snow'\r\n    });\r\n\r\n    const raw = document.getElementById(\"raw\");\r\n\r\n    // Update the textarea when Quill content changes\r\n    quill.on('text-change', function() {\r\n      raw.value = quill.root.innerHTML;\r\n    });\r\n\r\n    // Update the Quill editor when the textarea content changes\r\n    raw.addEventListener('input', function() {\r\n      quill.root.innerHTML = raw.value;\r\n    });\r\n</script><script async src=\"https://platform.twitter.com/widgets.js\" charset=\"utf-8\"></script></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
