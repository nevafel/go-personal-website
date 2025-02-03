// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package styling

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

/*
In this file I write main CSS for all webpages.
Also I store there some variables that I use a lot.
Currently all elements are not sorted.
Be aware that class names in final html page will be random.
*/

var mainColor = "#ffb37f"
var bgColor = "#272727"
var bgVariation = "#1f1f1f"   // Color simmilar to bg
var contrastColor = "#92ffd1" // Contrast to mainColor
var textColor = "#FFFFFF"

func BlogFormInput() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`width:70%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`height:150px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:12px 20px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-sizing:border-box;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:2px solid { bgVariation };`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:4px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:#f8f8f8;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:13px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`resize:none;`)
	templ_7745c5c3_CSSID := templ.CSSID(`BlogFormInput`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func FileHeader() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`color`, mainColor)))
	templ_7745c5c3_CSSBuilder.WriteString(`text-align:center;`)
	templ_7745c5c3_CSSID := templ.CSSID(`FileHeader`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Header() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`color`, mainColor)))
	templ_7745c5c3_CSSBuilder.WriteString(`text-align:left;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Header`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func HighlightText() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`color`, contrastColor)))
	templ_7745c5c3_CSSBuilder.WriteString(`font-weight:bold;`)
	templ_7745c5c3_CSSID := templ.CSSID(`HighlightText`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Textcontainer() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`block-size:fit-content;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:5px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`text-wrap:wrap;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:10px;`)
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`color`, textColor)))
	templ_7745c5c3_CSSID := templ.CSSID(`Textcontainer`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func BlogContainer() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`border:2px dashed gray;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:20px;`)
	templ_7745c5c3_CSSID := templ.CSSID(`BlogContainer`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func HelloContainer() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`height:fit-content;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:50px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`color`, textColor)))
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`background`, bgVariation)))
	templ_7745c5c3_CSSID := templ.CSSID(`HelloContainer`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func BlockContainer() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`height:fit-content;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`background`, bgVariation)))
	templ_7745c5c3_CSSID := templ.CSSID(`BlockContainer`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func PostScriptum() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`text-align:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:12px;`)
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`color`, contrastColor)))
	templ_7745c5c3_CSSID := templ.CSSID(`PostScriptum`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func DefaultTable() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`border-collapse:collapse;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid;`)
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`color`, textColor)))
	templ_7745c5c3_CSSID := templ.CSSID(`DefaultTable`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func CenterPageContainer() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`padding:15px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`position:absolute;`)
	templ_7745c5c3_CSSBuilder.WriteString(`top:50%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`left:50%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`-ms-transform:translateX(-50%) translateY(-50%);`)
	templ_7745c5c3_CSSBuilder.WriteString(`-webkit-transform:translate(-50%,-50%);`)
	templ_7745c5c3_CSSBuilder.WriteString(`transform:translate(-50%,-50%);`)
	templ_7745c5c3_CSSID := templ.CSSID(`CenterPageContainer`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func CenterContainer() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSID := templ.CSSID(`CenterContainer`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

var _ = templruntime.GeneratedTemplate
