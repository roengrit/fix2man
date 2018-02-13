package helps

import (
	"bytes"
	m "fix2man/models"
	"strconv"
	"strings"
)

//HTMLRecTemplate _
const HTMLRecTemplate = `<tr>
							<td class="is-col-toggle">{branch}</td>
							<td>{docno}</td>
							<td class="is-col-toggle">{reqname}</td> 
							<td >{reqdate}</td>                             
							<td class="is-col-toggle">{eventdate}</td>        
							<td class="is-col-toggle">{details}</td>                           
							<td class="is-col-toggle">{status}</td> 
							<td>
								<div class="btn-group">
									{action}
								</div>
							</td>                             
						</tr>`

//HTMLRecActionEnable _
const HTMLRecActionEnable = ` 
							 <a   class="btn btn-danger " title="แก้ไข"  href="/create-request/?id={id}"><i class="fa fa-edit"></i></a>						
							 <button type="button" class="btn btn-danger dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
								<span class="caret"></span>
								<span class="sr-only">Toggle Dropdown</span>
							</button>
							<ul class="dropdown-menu" role="menu">
								<li><a href="#">ยกเลิก</a></li>								 							
							</ul>
								 `

//HTMLRecNotFoundRows _
const HTMLRecNotFoundRows = `<tr> <td  colspan="8" style="text-align:center;">*** ไม่พบข้อมูล ***</td></tr>`

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
