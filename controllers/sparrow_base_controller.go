package controllers

type SparrowBaseController struct {
	Data    map[interface{}]interface{}
	TplName string
}

func (c *SparrowBaseController) Render(tpl_name string) {

}

func (c *SparrowBaseController) RedirectTo(url string) {

}
