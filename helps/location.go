package helps

import (
	"bytes"
	m "fix2man/models"
	"strconv"
	"strings"
)

//HTMLDepartTemplate _
const HTMLDepartTemplate = `<tr>
								<td>{branch_name}</td>
								<td>{name}</td>
								<td>
									<div class="btn-group">
										{action}
									</div>
								</td>
							</tr>`

//HTMLDepartActionEnable _
const HTMLDepartActionEnable = `<a   class="btn btn-sm btn-primary " title="แก้ไข"  href="/location/depart/?id={id}"><i class="fa fa-edit"></i></a>
                                <a   class="btn btn-sm btn-danger" title="ลบ" href="#"  onclick='confirmDeleteGlobal({id},"/location/depart/delete/")'><i class="fa fa-trash-o"></i></a>`

//HTMLDepartNotFoundRows _
const HTMLDepartNotFoundRows = `<tr><td colspan="3">*** ไม่พบข้อมูล ***</td> </tr>`

//HTMLDepartPermissionDenie _
const HTMLDepartPermissionDenie = `<tr><td colspan="3">*** ไม่อนุญาติ ใน entity อื่น ***</td></tr>`

//HTMLDepartError _
const HTMLDepartError = `<tr><td colspan="3"> {err}</td></tr>`

//GenDepartHTML _
func GenDepartHTML(lists []m.Departs) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLDepartTemplate, "{branch_name}", val.Branch.Name, -1)
		temp = strings.Replace(temp, "{name}", val.Name, -1)
		tempAction := strings.Replace(HTMLDepartActionEnable, "{id}", strconv.Itoa(val.ID), -1)
		temp = strings.Replace(temp, "{action}", tempAction, -1)
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}

/////////////////////////////////////////////  อาคาร /////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////

//HTMLBuildingTemplate _
const HTMLBuildingTemplate = `<tr>
								<td>{branch_name}</td>
								<td>{name}</td>
								<td>
									<div class="btn-group">
										{action}
									</div>
								</td>
							</tr>`

//HTMLBuildingActionEnable _
const HTMLBuildingActionEnable = `<a   class="btn btn-sm btn-primary " title="แก้ไข"  href="/location/building/?id={id}"><i class="fa fa-edit"></i></a>
                                <a   class="btn btn-sm btn-danger" title="ลบ" href="#"  onclick='confirmDeleteGlobal({id},"/location/building/delete/")'><i class="fa fa-trash-o"></i></a>`

//HTMLBuildingNotFoundRows _
const HTMLBuildingNotFoundRows = `<tr><td colspan="3">*** ไม่พบข้อมูล ***</td> </tr>`

//HTMLBuildingPermissionDenie _
const HTMLBuildingPermissionDenie = `<tr><td colspan="3">*** ไม่อนุญาติ ใน entity อื่น ***</td></tr>`

//HTMLBuildingError _
const HTMLBuildingError = `<tr><td colspan="3"> {err}</td></tr>`

//GenBuildingHTML _
func GenBuildingHTML(lists []m.Buildings) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLBuildingTemplate, "{branch_name}", val.Branch.Name, -1)
		temp = strings.Replace(temp, "{name}", val.Name, -1)
		tempAction := strings.Replace(HTMLBuildingActionEnable, "{id}", strconv.Itoa(val.ID), -1)
		temp = strings.Replace(temp, "{action}", tempAction, -1)
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}

/////////////////////////////////////////////  ชั้น /////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////

//HTMLClassTemplate _
const HTMLClassTemplate = `<tr>
								<td>{branch_name}</td>
								<td>{building_name}</td>
								<td>{name}</td>
								<td>
									<div class="btn-group">
										{action}
									</div>
								</td>
							</tr>`

//HTMLClassActionEnable _
const HTMLClassActionEnable = `<a   class="btn btn-sm btn-primary " title="แก้ไข"  href="/location/class/?id={id}"><i class="fa fa-edit"></i></a>
                                <a   class="btn btn-sm btn-danger" title="ลบ" href="#"  onclick='confirmDeleteGlobal({id},"/location/class/delete/")'><i class="fa fa-trash-o"></i></a>`

//HTMLClassNotFoundRows _
const HTMLClassNotFoundRows = `<tr><td colspan="4">*** ไม่พบข้อมูล ***</td> </tr>`

//HTMLClassPermissionDenie _
const HTMLClassPermissionDenie = `<tr><td colspan="4">*** ไม่อนุญาติ ใน entity อื่น ***</td></tr>`

//HTMLClassError _
const HTMLClassError = `<tr><td colspan="4"> {err}</td></tr>`

//GenClassHTML _
func GenClassHTML(lists []m.Class) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLClassTemplate, "{branch_name}", val.Building.Branch.Name, -1)
		temp = strings.Replace(temp, "{building_name}", val.Building.Name, -1)
		temp = strings.Replace(temp, "{name}", val.Name, -1)
		tempAction := strings.Replace(HTMLClassActionEnable, "{id}", strconv.Itoa(val.ID), -1)
		temp = strings.Replace(temp, "{action}", tempAction, -1)
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}
