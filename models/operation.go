package models

// AddPakreq adds a pakreq to the whole array
func AddPakreq(pakreq Pakreq) bool {
	if PakreqExists(pakreq) == true {
		return false
	}
	CurConfig.CurPakreqs = append(CurConfig.CurPakreqs, pakreq)
	return true
}

// PakreqExists returns if the pakreq already exists in the database
func PakreqExists(pakreq Pakreq) bool {
	for _, pkg := range CurConfig.CurPakreqs {
		if pkg.PkgName == pakreq.PkgName {
			return true
		}
	}
	return false
}
