package helps

import (
	"bytes"
	m "fix2man/models"
	"strconv"
	"strings"
)

//HTMLPickUpTemplate _
const HTMLPickUpTemplate = `<tr>
							<td>{doc_date}</td>
							<td>{doc_no}</td>	
							<td>{remark}</td>					 
							<td>{active}</td>
							<td>{total_net_amount}</td> 
							<td>
								<div class="btn-group">
									{action}
								</div>
							</td>                             
						</tr>`

//HTMLPickUpActionEnable _
const HTMLPickUpActionEnable = `<a class="btn bg-purple" title="รายละเอียด" target="_blank" href="/pickup/read/?id={id}"><i class="fa fa-file-text-o"></i></a>
								 <a class="btn btn-primary " title="แก้ไข"  href="/pickup/?id={id}"><i class="fa fa-edit"></i></a>
								 <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
										<span class="caret"></span>
										<span class="sr-only">Toggle Dropdown</span>
									</button>
									<ul class="dropdown-menu" role="menu">
										<li><a href="#" onclick="cancelDoc({id})" title="ยกเลิก">ยกเลิก</a></li>
								</ul> `

//HTMLPickUpActionEditOnly _
const HTMLPickUpActionEditOnly = `<a class="btn bg-purple" title="รายละเอียด" target="_blank" href="/pickup/read/?id={id}"><i class="fa fa-file-text-o"></i></a>
								   <a class="btn btn-primary" title="แก้ไข"  target="_blank" href="/pickup/?id={id}"><i class="fa fa-edit"></i></a>
								   <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
										<span class="caret"></span>
										<span class="sr-only">Toggle Dropdown</span>
								   </button>
								  `

//HTMLPickUpNotFoundRows _
const HTMLPickUpNotFoundRows = `<tr> <td  colspan="6" style="text-align:center;">*** ไม่พบข้อมูล ***</td></tr>`

//HTMLPickUpError _
const HTMLPickUpError = `<tr> <td  colspan="6" style="text-align:center;">{err}</td></tr>`

//GenPickUpHTML _
func GenPickUpHTML(lists []m.PickUp) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLPickUpTemplate, "{doc_date}", val.DocDate.Format("02-01-2006"), -1)
		temp = strings.Replace(temp, "{doc_no}", val.DocNo, -1)
		temp = strings.Replace(temp, "{remark}", val.Remark, -1)
		if val.Active {
			temp = strings.Replace(temp, "{active}", "N", -1)
		} else {
			temp = strings.Replace(temp, "{active}", "C", -1)
		}
		temp = strings.Replace(temp, "{total_net_amount}", ThCommaSep(val.TotalNetAmount), -1)
		if val.Active {
			tempAction := strings.Replace(HTMLPickUpActionEnable, "{id}", strconv.Itoa(val.ID), -1)
			temp = strings.Replace(temp, "{action}", tempAction, -1)
		} else {
			tempAction := strings.Replace(HTMLPickUpActionEditOnly, "{id}", strconv.Itoa(val.ID), -1)
			temp = strings.Replace(temp, "{action}", tempAction, -1)
		}
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}
