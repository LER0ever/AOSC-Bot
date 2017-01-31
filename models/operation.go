package models

// AddPakreq adds a pakreq to the whole array
func AddPakreq(pakreq Pakreq) bool {
	if PakreqExists(pakreq) == true {
		return false
	}
	CurConfig.CurPakreqs = append(CurConfig.CurPakreqs, pakreq)
	ExportToFile()
	return true
}
