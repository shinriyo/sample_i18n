package routers

import (
	"sample_i18n/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

// setLangVer sets site language version.
func (this *baseRouter) setLangVer() bool {
    isNeedRedir := false
    hasCookie := false

    // 1. Check URL arguments.
    lang := this.Input().Get("lang")

    // 2. Get language information from cookies.
    if len(lang) == 0 {
        lang = this.Ctx.GetCookie("lang")
        hasCookie = true
    } else {
        isNeedRedir = true
    }

    // Check again in case someone modify by purpose.
    if !i18n.IsExist(lang) {
        lang = ""
        isNeedRedir = false
        hasCookie = false
    }

    // 3. Get language information from 'Accept-Language'.
    if len(lang) == 0 {
        al := this.Ctx.Request.Header.Get("Accept-Language")
        if len(al) > 4 {
            al = al[:5] // Only compare first 5 letters.
            if i18n.IsExist(al) {
                lang = al
            }
        }
    }

    // 4. Default language is English.
    if len(lang) == 0 {
        lang = "en-US"
        isNeedRedir = false
    }

    curLang := langType{
        Lang: lang,
    }

    // Save language information in cookies.
    if !hasCookie {
        this.Ctx.SetCookie("lang", curLang.Lang, 1<<31-1, "/")
    }

    restLangs := make([]*langType, 0, len(langTypes)-1)
    for _, v := range langTypes {
        if lang != v.Lang {
            restLangs = append(restLangs, v)
        } else {
            curLang.Name = v.Name
        }
    }

    // Set language properties.
    this.Lang = lang
    this.Data["Lang"] = curLang.Lang
    this.Data["CurLang"] = curLang.Name
    this.Data["RestLangs"] = restLangs

    return isNeedRedir
}
