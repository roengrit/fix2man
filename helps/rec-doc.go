package helps

import (
	"bytes"
	m "fix2man/models"
	"strconv"
	"strings"
)

//HTMLRecTemplate _
const HTMLRecTemplate = `<tr>
							<td>{branch}</td>
							<td>{docno}</td>
							<td>{reqname}</td> 
							<td>{reqdate}</td>                             
							<td>{eventdate}</td>        
							<td>{details}</td>                           
							<td>{status}</td> 
							<td>
								<div class="btn-group">
									{action}
								</div>
							</td>                             
						</tr>`

//HTMLRecActionEnable _
const HTMLRecActionEnable = `<a   class="btn bg-purple" title="รายละเอียด" target="_blank" href="/request/read/?id={id}&r=1"><i class="fa fa-file-text-o"></i></a>
							 <a   class="btn btn-danger " title="แก้ไข"  href="/create-request/?id={id}"><i class="fa fa-edit"></i></a>						
							 <button type="button" class="btn btn-danger dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
								<span class="caret"></span>
								<span class="sr-only">Toggle Dropdown</span>
							</button>
							<ul class="dropdown-menu" role="menu">
								<li><a href="#">ยกเลิก</a></li>								 							
							</ul>
								 `

//HTMLReqActionDisable _
//const HTMLReqActionDisable = `<button type="button" class="btn btn-sm btn-primary disabled" >แก้ไข</button>
//						      <button type="button" class="btn btn-sm btn-danger disabled" >ลบ</button>`

//HTMLRecNotFoundRows _
const HTMLRecNotFoundRows = `<tr> <td  colspan="8" style="text-align:center;">*** ไม่พบข้อมูล ***</td></tr>`

//HTMLReqPermissionDenie _
///const HTMLReqPermissionDenie = `<tr> <td  colspan="8" style="text-align:center;">*** ไม่อนุญาติ ใน entity อื่น ***</td></tr>`

//HTMLRecError _
const HTMLRecError = `<tr> <td  colspan="8" style="text-align:center;">{err}</td></tr>`

//GenRecHTML _
func GenRecHTML(lists []m.RequestList) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLReqTemplate, "{branch}", val.Branch, -1)
		temp = strings.Replace(temp, "{docno}", val.DocNo, -1)
		temp = strings.Replace(temp, "{reqname}", val.ReqName, -1)
		temp = strings.Replace(temp, "{reqdate}", val.DocDate.Format("02-01-2006"), -1)
		temp = strings.Replace(temp, "{eventdate}", val.EventDate.Format("02-01-2006"), -1)
		temp = strings.Replace(temp, "{details}", val.Details, -1)
		temp = strings.Replace(temp, "{status}", val.Status, -1)
		temp = strings.Replace(temp, "{action}", strings.Replace(HTMLReqActionEnable, "{id}", strconv.Itoa(val.ID), -1), -1)
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}
