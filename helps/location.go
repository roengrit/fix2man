package helps

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

//HTMLDepartNotFoundRows _
const HTMLDepartNotFoundRows = `<tr><td colspan="3">*** ไม่พบข้อมูล ***</td> </tr>`

//HTMLDepartPermissionDenie _
const HTMLDepartPermissionDenie = `<tr><td colspan="3">*** ไม่อนุญาติ ใน entity อื่น ***</td></tr>`

//HTMLDepartError _
const HTMLDepartError = `<tr><td colspan="3"> {err}</td></tr>`
