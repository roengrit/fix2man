package controllers

import h "fix2man/helps"

type GetNameController struct {
	BaseController
}

func (c *GetNameController) Get() {

	c.Data["json"] = h.GetUser(c.Ctx.Request)
	c.ServeJSON()
}
