package helps

import (
	"bytes"
	m "fix2man/models"
	"strconv"
	"strings"
)

//HTMLTemplate _
const HTMLTemplate = `<tr>
							<td>{name}</td>
							<td>
								<div class="btn-group">
									{action}
								</div>
							</td>                             
						</tr>`

//HTMLActionEnable _
const HTMLActionEnable = `<button type="button" class="btn btn-sm btn-primary" onclick='editNormal({id})'>แก้ไข</button>
						<button type="button" class="btn btn-sm btn-danger" onclick='deleteNormal({id})'>ลบ</button>`

//HTMLActionDisable _
const HTMLActionDisable = `<button type="button" class="btn btn-sm btn-primary disabled" >แก้ไข</button>
						 <button type="button" class="btn btn-sm btn-danger disabled" >ลบ</button>`

//HTMLNotFoundRows _
const HTMLNotFoundRows = `<tr><td></td><td>*** ไม่พบข้อมูล ***</td><td></td></tr>`

//HTMLPermissionDenie _
const HTMLPermissionDenie = `<tr><td></td><td>*** ไม่อนุญาติ ใน entity อื่น ***</td><td></td></tr>`

//HTMLError _
const HTMLError = `<tr><td></td><td>{err}</td><td></td></tr>`

//GetEntityTitle _
func GetEntityTitle(entity string) string {
	switch entity {
	case "roles":
		return "สิทธิ์"
	case "units":
		return "หน่วย"
	case "status":
		return "สถานะการซ่อม"
	case "branchs":
		return "สาขา/ไซต์"
	case "departs":
		return "แผนก"
	default:
		return ""
	}
}

//GetEntityParentField _
func GetEntityParentField(entity string) string {
	switch entity {
	case "roles":
		return ""
	case "units":
		return ""
	case "status":
		return ""
	case "branchs":
		return ""
	case "departs":
		return ""
	case "buildings":
		return "branch_id"
	case "class":
		return "building_id"
	case "rooms":
		return "class_id"
	default:
		return ""
	}
}

//GenEntityHTML _
func GenEntityHTML(lists []m.NormalEntity) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLTemplate, "{name}", val.Name, -1)
		if !val.Lock {
			temp = strings.Replace(temp, "{action}", strings.Replace(HTMLActionEnable, "{id}", strconv.Itoa(val.ID), -1), -1)
		} else {
			temp = strings.Replace(temp, "{action}", HTMLActionDisable, -1)
		}
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}
