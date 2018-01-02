package helps

import (
	"bytes"
	m "fix2man/models"
	"strconv"
	"strings"
)

//HTMLReqTemplate _
const HTMLReqTemplate = `<tr>
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

//HTMLReqActionEnable _
const HTMLReqActionEnable = `<a   class="btn bg-purple" title="รายละเอียด" target="_blank" href="/request/read/?id={id}&r=1"><i class="fa fa-file-text-o"></i></a>
							 <a   class="btn btn-primary" title="เปลี่ยนสถานะ"><i class="fa fa-list-ol"></i></a>
							 <a   class="btn btn-danger " title="แก้ไข" target="_blank" href="/create-request/?id={id}"><i class="fa fa-edit"></i></a>
							 <button type="button" class="btn btn-danger dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
								<span class="caret"></span>
								<span class="sr-only">Toggle Dropdown</span>
							</button>
							<ul class="dropdown-menu" role="menu">
								<li><a href="#">รับ/แจกจ่ายงาน</a></li>
								<li><a href="#">ใบแจ้งงานต่อเนื่อง</a></li>
								<li><a href="#">ประเมินวันที่แล้วเสร็จ</a></li>
								<li><a href="#">เบิกอะไหล่</a></li>
								<li><a href="#">เบิกเครื่องสำรอง</a></li>
								<li><a href="#">ประเมินผลการซ่อม</a></li>								
							</ul>
							</div>
								 `

//HTMLReqActionDisable _
//const HTMLReqActionDisable = `<button type="button" class="btn btn-sm btn-primary disabled" >แก้ไข</button>
//						      <button type="button" class="btn btn-sm btn-danger disabled" >ลบ</button>`

//HTMLReqNotFoundRows _
const HTMLReqNotFoundRows = `<tr> <td  colspan="8" style="text-align:center;">*** ไม่พบข้อมูล ***</td></tr>`

//HTMLReqPermissionDenie _
///const HTMLReqPermissionDenie = `<tr> <td  colspan="8" style="text-align:center;">*** ไม่อนุญาติ ใน entity อื่น ***</td></tr>`

//HTMLReqError _
const HTMLReqError = `<tr> <td  colspan="8" style="text-align:center;">{err}</td></tr>`

//GenReqHTML _
func GenReqHTML(lists []m.RequestList) string {
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