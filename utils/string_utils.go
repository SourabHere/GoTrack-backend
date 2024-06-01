package utils

import "strings"

func ParseIssueFilesAttached(data []byte) []string {

	if data == nil {
		return []string{}
	}

	strData := string(data)

	strData = strData[1 : len(strData)-1]

	if strData == "" {
		return []string{}
	}

	files := strings.Split(strData, ",")

	return files

}
