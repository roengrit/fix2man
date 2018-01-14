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

/////////////////////////////////////////////  ห้อง /////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////

//HTMLRoomTemplate _
const HTMLRoomTemplate = `<tr>
								<td>{branch_name}</td>
								<td>{building_name}</td>
								<td>{class_name}</td>
								<td>{name}</td>
								<td>
									<div class="btn-group">
										{action}
									</div>
								</td>
							</tr>`

//HTMLRoomActionEnable _
const HTMLRoomActionEnable = `<a   class="btn btn-sm btn-primary " title="แก้ไข"  href="/location/room/?id={id}"><i class="fa fa-edit"></i></a>
                              <a   class="btn btn-sm btn-danger" title="ลบ" href="#"  onclick='confirmDeleteGlobal({id},"/location/room/delete/")'><i class="fa fa-trash-o"></i></a>`

//HTMLRoomNotFoundRows _
const HTMLRoomNotFoundRows = `<tr><td colspan="5">*** ไม่พบข้อมูล ***</td> </tr>`

//HTMLRoomPermissionDenie _
const HTMLRoomPermissionDenie = `<tr><td colspan="5">*** ไม่อนุญาติ ใน entity อื่น ***</td></tr>`

//HTMLRoomError _
const HTMLRoomError = `<tr><td colspan="5"> {err}</td></tr>`

//GenRoomHTML _
func GenRoomHTML(lists []m.Rooms) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		val.Class.Building.Branch, _ = m.GetBranchByID(val.Class.Building.Branch.ID)
		temp := strings.Replace(HTMLRoomTemplate, "{branch_name}", val.Class.Building.Branch.Name, -1)
		temp = strings.Replace(temp, "{building_name}", val.Class.Building.Name, -1)
		temp = strings.Replace(temp, "{class_name}", val.Class.Name, -1)
		temp = strings.Replace(temp, "{name}", val.Name, -1)
		tempAction := strings.Replace(HTMLRoomActionEnable, "{id}", strconv.Itoa(val.ID), -1)
		temp = strings.Replace(temp, "{action}", tempAction, -1)
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}
