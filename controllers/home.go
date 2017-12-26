package controllers

//AppController _
type AppController struct {
	BaseController
}

//Get Home page
func (c *AppController) Get() {
	c.Layout = "layout.html"
	c.TplName = "main/index.html"
	c.Render()
}
