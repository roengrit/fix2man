package helps

import (
	"bytes"
	m "fix2man/models"
	"strconv"
	"strings"
)

//HTMLMemTemplate _
const HTMLMemTemplate = `<tr>
							<td>{name}</td>
							<td>{contact}</td>
							<td>{tel}</td> 
							<td>{type}</td> 
							<td>
								<div class="btn-group">
									{action}
								</div>
							</td>                             
						</tr>`

//HTMLMemActionEnable _
const HTMLMemActionEnable = `<a class="btn bg-purple" title="รายละเอียด" target="_blank" href="/member/read/?id={id}"><i class="fa fa-file-text-o"></i></a>
							 <a class="btn btn-primary " title="แก้ไข"  href="/member/?id={id}"><i class="fa fa-edit"></i></a>
							 <a class="btn btn-danger" title="ลบ" href="#"  onclick='confirmDeleteGlobal({id},"/member")'><i class="fa fa-trash-o"></i></a>
								 `

//HTMLMemActionDisable _
const HTMLMemActionDisable = `<a class="btn bg-purple disabled"   title="รายละเอียด" target="_blank" href="/member/read/?id={id}"><i class="fa fa-file-text-o"></i></a>
							 <a class="btn btn-primary disabled"   title="แก้ไข"  href="/member/?id={id}"><i class="fa fa-edit"></i></a>
							 <a class="btn btn-danger disabled"   title="ลบ" href="#"  onclick='confirmDeleteGlobal({id},"/member")'><i class="fa fa-trash-o"></i></a>
								 `

//HTMLMemNotFoundRows _
const HTMLMemNotFoundRows = `<tr> <td  colspan="5" style="text-align:center;">*** ไม่พบข้อมูล ***</td></tr>`

//HTMLMemError _
const HTMLMemError = `<tr> <td  colspan="5" style="text-align:center;">{err}</td></tr>`

//GenMemHTML _
func GenMemHTML(lists []m.Member) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLMemTemplate, "{name}", val.Name, -1)
		temp = strings.Replace(temp, "{contact}", val.Contact, -1)
		temp = strings.Replace(temp, "{tel}", val.Tel, -1)
		var tempText string
		switch val.MemberType {
		case 0:
			tempText = "ลูกค้า"
		case 1:
			tempText = "ร้านค้า/Supplier"
		case 2:
			tempText = "ภายใน"

		}
		temp = strings.Replace(temp, "{type}", tempText, -1)
		if val.Lock {
			tempAction := strings.Replace(HTMLMemActionDisable, "{id}", strconv.Itoa(val.ID), -1)
			temp = strings.Replace(temp, "{action}", tempAction, -1)
		} else {
			tempAction := strings.Replace(HTMLMemActionEnable, "{id}", strconv.Itoa(val.ID), -1)
			temp = strings.Replace(temp, "{action}", tempAction, -1)
		}
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}
