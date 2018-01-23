package helps

import (
	"bytes"
	m "fix2man/models"
	"fmt"
	"strconv"
	"strings"
)

//HTMLTemplate _
const HTMLTemplate = `<tr> 
							<td>
							<div class="btn-group">
								{action}
							</div>
							</td> 
							<td>{name}</td>
							                            
						</tr>`

//HTMLActionEnable _
const HTMLActionEnable = `<button type="button" class="btn btn-sm btn-primary" onclick='editNormal({id})'><i class="fa fa-edit"></i></button>
						<button type="button" class="btn btn-sm btn-danger" onclick='deleteNormal({id})'><i class="fa fa-trash-o"></i></button>`

//HTMLActionDisable _
const HTMLActionDisable = `<button type="button" class="btn btn-sm btn-primary disabled" ><i class="fa fa-edit"></i></button>
						 <button type="button" class="btn btn-sm btn-danger disabled" ><i class="fa fa-trash-o"></i></button>`

//HTMLNotFoundRows _
const HTMLNotFoundRows = `<tr><td colspan="2">*** ไม่พบข้อมูล ***</td> </tr>`

//HTMLPermissionDenie _
const HTMLPermissionDenie = `<tr><td colspan="2">*** ไม่อนุญาติ ใน entity อื่น ***</td></tr>`

//HTMLError _
const HTMLError = `<tr><td colspan="2"> {err}</td></tr>`

//GetEntityTitle _
func GetEntityTitle(entity string) string {
	switch entity {
	case "roles":
		return "สิทธิ์"
	case "units":
		return "หน่วย"
	case "status":
		return "สถานะการซ่อม"
	case "branchs":
		return "สาขา/ไซต์"
	case "categorys":
		return "หมวดหมู่อุปกรณ์"
	default:
		return ""
	}
}

//GetEntityParentField _
func GetEntityParentField(entity string) string {
	switch entity {
	case "roles":
		return ""
	case "units":
		return ""
	case "status":
		return ""
	case "branchs":
		return ""
	case "departs":
		return "branch_id"
	case "buildings":
		return "branch_id"
	case "class":
		return "building_id"
	case "rooms":
		return "class_id"
	case "categorys":
		return ""
	default:
		return ""
	}
}

//GenEntityHTML _
func GenEntityHTML(lists []m.NormalEntity) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLTemplate, "{name}", val.Name, -1)
		if !val.Lock {
			temp = strings.Replace(temp, "{action}", strings.Replace(HTMLActionEnable, "{id}", strconv.Itoa(val.ID), -1), -1)
		} else {
			temp = strings.Replace(temp, "{action}", HTMLActionDisable, -1)
		}
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}

//IsMobile _
func IsMobile(useragent string) bool {
	// the list below is taken from
	mobiles := []string{
		"Mobile Explorer",
		"Palm",
		"Motorola",
		"Nokia",
		"Palm",
		"Apple iPhone",
		"iPad",
		"iPhone",
		"Apple iPod Touch",
		"Sony Ericsson",
		"Sony Ericsson",
		"BlackBerry",
		"O2 Cocoon",
		"Treo",
		"LG",
		"Amoi",
		"XDA",
		"MDA",
		"Vario",
		"HTC",
		"Samsung",
		"Sharp",
		"Siemens",
		"Alcatel",
		"BenQ",
		"HP iPaq",
		"Motorola",
		"PlayStation Portable",
		"PlayStation 3",
		"PlayStation Vita",
		"Danger Hiptop",
		"NEC",
		"Panasonic",
		"Philips",
		"Sagem",
		"Sanyo",
		"SPV",
		"ZTE",
		"Sendo",
		"Nintendo DSi",
		"Nintendo DS",
		"Nintendo 3DS",
		"Nintendo Wii",
		"Open Web",
		"OpenWeb",
		"Android",
		"Symbian",
		"SymbianOS",
		"Palm",
		"Symbian S60",
		"Windows CE",
		"Obigo",
		"Netfront Browser",
		"Openwave Browser",
		"Mobile Explorer",
		"Opera Mini",
		"Opera Mobile",
		"Firefox Mobile",
		"Digital Paths",
		"AvantGo",
		"Xiino",
		"Novarra Transcoder",
		"Vodafone",
		"NTT DoCoMo",
		"O2",
		"mobile",
		"wireless",
		"j2me",
		"midp",
		"cldc",
		"up.link",
		"up.browser",
		"smartphone",
		"cellphone",
		"Generic Mobile"}
	fmt.Println(useragent)
	for _, device := range mobiles {
		if strings.Index(useragent, device) > -1 {
			fmt.Println(true)
			return true
		}
	}
	return false
}
