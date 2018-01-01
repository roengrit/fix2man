package helps

import (
	"bytes"
	m "fix2man/models"
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
const HTMLReqActionEnable = `<button type="button" class="btn btn-sm btn-primary" onclick='editReq({id})'>แก้ไข</button>
						<button type="button" class="btn btn-sm btn-danger" onclick='deleteNormal({id})'>ลบ</button>`

//HTMLReqActionDisable _
const HTMLReqActionDisable = `<button type="button" class="btn btn-sm btn-primary disabled" >แก้ไข</button>
						 <button type="button" class="btn btn-sm btn-danger disabled" >ลบ</button>`

//HTMLReqNotFoundRows _
const HTMLReqNotFoundRows = `<tr><td></td><td>*** ไม่พบข้อมูล ***</td><td></td></tr>`

//HTMLReqPermissionDenie _
const HTMLReqPermissionDenie = `<tr><td></td><td>*** ไม่อนุญาติ ใน entity อื่น ***</td><td></td></tr>`

//HTMLReqError _
const HTMLReqError = `<tr><td></td><td>{err}</td><td></td></tr>`

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
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}
