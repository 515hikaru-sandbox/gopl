package uniq

func Uniq(strings []string) []string {
	var uniqStrings []string
	for i := range strings {
		if i == 0 {
			uniqStrings = append(uniqStrings, strings[i])
			continue
		}
		if strings[i-1] == strings[i] {
			continue
		} else {
			uniqStrings = append(uniqStrings, strings[i])
		}
	}
	return uniqStrings
}
