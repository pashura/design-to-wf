package xtl_structs

func (group Element) IsRepeatableGroup() bool {
	if group.Name == "GROUPDEF" && group.Atts.Max != "1" {
		return true
	}
	return false
}
