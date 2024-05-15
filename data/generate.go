package data

func Generate(dataType string) string {
	switch dataType {
	case TYPE_NAME:
		return generateName()
	case TYPE_DATE:
		return generateDate()
	case TYPE_ADDRESS:
		return generateAddress()
	case TYPE_PHONE:
		return generatePhone()
	}

	return ""
}

func generateName() string {
	return ""
}

func generateDate() string {
	return ""
}

func generateAddress() string {
	return ""
}

func generatePhone() string {
	return ""
}
