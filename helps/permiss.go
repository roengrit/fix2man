package helps

import (
	"fix2man/models"
	"net/http"

	c "github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/gorilla/securecookie"
)

var s = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

//KeepLogin login
func KeepLogin(w *c.Response, username string, role string) (ok bool, err string) {
	value := map[string]string{
		"username": username,
		"role":     role,
	}
	if encoded, errs := s.Encode("fixman", value); errs != nil {
		ok = false
		err = errs.Error()
	} else {
		cookie := http.Cookie{
			Name:     "fixman",
			Value:    encoded,
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(w.ResponseWriter, &cookie)
		ok = true
		err = ""
	}
	return ok, err
}

//GetUser get user from cookie
func GetUser(r *http.Request) string {
	if cookie, err := r.Cookie("fixman"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("fixman", cookie.Value, &value); err == nil {
			return value["username"]
		}
	}
	return ""
}

//GetRole get role from cookie
func GetRole(r *http.Request) string {

	if cookie, err := r.Cookie("fixman"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("fixman", cookie.Value, &value); err == nil {
			return value["role"]
		}
	}
	return ""
}

//CheckPermiss check permission
func CheckPermiss(roleID, menuID int) bool {
	var ret bool
	o := orm.NewOrm()
	Permiss := models.Permiss{RoleID: roleID, MenuID: menuID}
	err := o.Read(&Permiss, "RoleID", "MenuID")
	if err == orm.ErrNoRows {
		ret = false
	} else {
		ret = Permiss.Active
	}
	return ret
}
