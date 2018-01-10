package helps

import (
	"bytes"
	m "fix2man/models"
	"strconv"
	"strings"
)

//HTMLSupTemplate _
const HTMLSupTemplate = `<tr>
							<td>{name}</td>
							<td>{contact}</td>
							<td>{tel}</td> 
							<td>
								<div class="btn-group">
									{action}
								</div>
							</td>                             
						</tr>`

//HTMLSupActionEnable _
const HTMLSupActionEnable = `<a   class="btn bg-purple" title="รายละเอียด" target="_blank" href="/supplier/read/?id={id}&r=1"><i class="fa fa-file-text-o"></i></a>
							 <a   class="btn btn-primary " title="แก้ไข"  href="/create-supplier/?id={id}"><i class="fa fa-edit"></i></a>
							 <a   class="btn btn-danger" title="ลบ" href="#"  onclick='confirmDeleteGlobal({id},"/supplier/delete")'><i class="fa fa-trash-o"></i></a>
								 `

//HTMLSupNotFoundRows _
const HTMLSupNotFoundRows = `<tr> <td  colspan="4" style="text-align:center;">*** ไม่พบข้อมูล ***</td></tr>`

//HTMLSupError _
const HTMLSupError = `<tr> <td  colspan="4" style="text-align:center;">{err}</td></tr>`

//GenSupHTML _
func GenSupHTML(lists []m.Suppliers) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLSupTemplate, "{name}", val.Name, -1)
		temp = strings.Replace(temp, "{contact}", val.Contact, -1)
		temp = strings.Replace(temp, "{tel}", val.Tel, -1)
		tempAction := strings.Replace(HTMLSupActionEnable, "{id}", strconv.Itoa(val.ID), -1)
		temp = strings.Replace(temp, "{action}", tempAction, -1)
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}
