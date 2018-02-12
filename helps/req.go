package helps

import (
	"bytes"
	"errors"
	m "fix2man/models"
	"strconv"
	"strings"
	"time"
)

//HTMLReqTemplate _
const HTMLReqTemplate = `<tr>
							<td>
								<div class="btn-group">
									{action}
								</div>
							</td>
							<td class="is-col-toggle">{branch}</td>
							<td class="info-col" style="cursor: pointer;">{docno}</td>
							<td class="is-col-toggle">{reqname}</td>
							<td>{reqdate}</td>
							<td class="is-col-toggle">{eventdate}</td>
							<td class="is-col-toggle">{details}</td>
							<td >{status}</td>

						</tr>`

//HTMLReqActionEnable _
const HTMLReqActionEnable = `<a   class="btn bg-purple" title="รายละเอียด" target="_blank" href="/request/read/?id={id}&r=1"><i class="fa fa-file-text-o"></i></a>
							 <button type="button" class="btn bg-purple dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
								<span class="caret"></span>
								<span class="sr-only">Toggle Dropdown</span>
							</button>
							<ul class="dropdown-menu" role="menu">
							    <li><a href="/create-request/?id={id}" title="แก้ไข">แก้ไข</a></li>
								<li><a href="#" onclick="changeStatus({id})" title="เปลี่ยนสถานะ">เปลี่ยนสถานะ</a></li>
								<li><a href="/pickup/?doc_ref={docno}">เบิกอะไหล่</a></li>
							</ul>
								 `

//HTMLReqNotFoundRows _
const HTMLReqNotFoundRows = `<tr> <td  colspan="8" style="text-align:center;">*** ไม่พบข้อมูล ***</td></tr>`

//HTMLReqError _
const HTMLReqError = `<tr> <td  colspan="8" style="text-align:center;">{err}</td></tr>`

//GenReqHTML _
func GenReqHTML(lists []m.RequestList) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLReqTemplate, "{branch}", val.Branch, -1)
		temp = strings.Replace(temp, "{docno}", val.DocNo, -1)
		temp = strings.Replace(temp, "{reqname}", val.ReqName, -1)
		temp = strings.Replace(temp, "{reqdate}", val.ReqDate.Format("02-01-2006"), -1)
		temp = strings.Replace(temp, "{eventdate}", val.EventDate.Format("02-01-2006"), -1)
		temp = strings.Replace(temp, "{details}", val.Details, -1)
		temp = strings.Replace(temp, "{status}", val.Status, -1)
		tempAction := strings.Replace(HTMLReqActionEnable, "{id}", strconv.Itoa(val.ID), -1)
		tempAction = strings.Replace(tempAction, "{docno}", val.DocNo, -1)
		temp = strings.Replace(temp, "{action}", tempAction, -1)
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}

//ValidateReqData _
func ValidateReqData(reqDoc m.RequestDocument,
	reqDate, reqDateEvent, reqAppointmentDate,
	reqGoalDate, reqActionDate, reqCompleteDate string) (
	ret m.NormalModel,
	eventDate time.Time,
	requestDate time.Time,
	appointmentDate time.Time,
	goalDate time.Time,
	actionDate time.Time,
	completeDate time.Time,
) {

	ret.RetOK = true
	if reqDoc.ReqName == "" && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชื่อผู้แจ้ง"
	}
	if reqDoc.Tel == "" && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุเบอร์โทรศัพท์"
	}
	if reqDoc.Branch == nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุสาขา"
	}
	if reqDoc.Depart == nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุแผนก"
	}
	if reqDoc.Building == nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุตึก"
	}
	if reqDoc.Class == nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชั้น"
	}
	if reqDoc.Room == nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุห้อง"
	}
	errDate := errors.New("")
	sp := strings.Split(reqDate, "-")
	if len(sp) == 3 {
		requestDate, errDate = time.Parse("2006-02-01", sp[2]+"-"+sp[0]+"-"+sp[1])
	}
	if errDate != nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "วันที่แจ้งไม่ถูกต้อง (dd-mm-yyyy)"
	}

	sp = strings.Split(reqDateEvent, "-")
	if len(sp) == 3 {
		eventDate, errDate = time.Parse("2006-02-01", sp[2]+"-"+sp[0]+"-"+sp[1])
	}
	if errDate != nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "วันที่เสียไม่ถูกต้อง (dd-mm-yyyy)"
	}

	if reqDoc.Details == "" && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "รายละเอียดการชำรุด/ปัญหา/อาการเสีย"
	}

	sp = strings.Split(reqAppointmentDate, "-")
	if len(sp) == 3 && ret.RetOK {
		appointmentDate, errDate = time.Parse("2006-02-01", sp[2]+"-"+sp[0]+"-"+sp[1])
	}

	if errDate != nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "วันที่นัดดำเนินการ (dd-mm-yyyy)"
	}

	sp = strings.Split(reqGoalDate, "-")
	if len(sp) == 3 && ret.RetOK {
		goalDate, errDate = time.Parse("2006-02-01", sp[2]+"-"+sp[0]+"-"+sp[1])
	}

	if errDate != nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "วันที่คาดว่าจะเสร็จ (dd-mm-yyyy)"
	}

	sp = strings.Split(reqActionDate, "-")
	if len(sp) == 3 && ret.RetOK {
		actionDate, errDate = time.Parse("2006-02-01", sp[2]+"-"+sp[0]+"-"+sp[1])
	}

	if errDate != nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "วันที่ดำเนินการ (dd-mm-yyyy)"
	}

	sp = strings.Split(reqCompleteDate, "-")
	if len(sp) == 3 && ret.RetOK {
		completeDate, errDate = time.Parse("2006-02-01", sp[2]+"-"+sp[0]+"-"+sp[1])
	}

	if errDate != nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "วันที่ซ่อมเสร็จ (dd-mm-yyyy)"
	}
	// reqDate, reqDateEvent, reqAppointmentDate,
	// reqGoalDate, reqActionDate, reqCompleteDate string) (
	// ret m.NormalModel,
	// eventDate time.Time,
	// requestDate time.Time,
	// appointmentDate time.Time,
	// goalDate time.Time,
	// actionDate time.Time,
	// completeDate time.Time,
	return
}

//DateTimeParse _
func DateTimeParse(date string, timeStr string) (time.Time, error) {
	sp := strings.Split(date, "-")
	//retDate, errDate := time.Parse("2006-02-01 H:i:s", sp[2]+"-"+sp[0]+"-"+sp[1]+" "+timeStr)
	retDate, errDate := time.Parse(
		time.RFC3339, sp[2]+"-"+sp[0]+"-"+sp[1]+"T"+timeStr+":00+00:00")
	return retDate, errDate
}
