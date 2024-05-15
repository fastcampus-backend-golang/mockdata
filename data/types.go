package data

const (
	TYPE_NAME    = "name"
	TYPE_DATE    = "date"
	TYPE_ADDRESS = "address"
	TYPE_PHONE   = "phone"
)

var SupportedTypes = map[string]bool{
	TYPE_NAME:    true,
	TYPE_DATE:    true,
	TYPE_ADDRESS: true,
	TYPE_PHONE:   true,
}
