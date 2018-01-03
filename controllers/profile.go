package controllers

import h "fix2man/helps"

//GetNameController _
type GetNameController struct {
	BaseController
}

//GetName -
func (c *GetNameController) GetName() {

	c.Data["json"] = h.GetUser(c.Ctx.Request)
	c.ServeJSON()
}

//ChangeBranch -
func (c *GetNameController) ChangeBranch() {

	c.Data["json"] = h.GetUser(c.Ctx.Request)
	c.ServeJSON()
}
