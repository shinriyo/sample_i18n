// Initialized language type list.
langs := strings.Split(models.Cfg.MustValue("lang", "types"), "|")
names := strings.Split(models.Cfg.MustValue("lang", "names"), "|")
langTypes = make([]*langType, 0, len(langs))
for i, v := range langs {
    langTypes = append(langTypes, &langType{
        Lang: v,
        Name: names[i],
    })
}

for _, lang := range langs {
    beego.Trace("Loading language: " + lang)
    if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
        beego.Error("Fail to set message file: " + err.Error())
        return
    }
}
