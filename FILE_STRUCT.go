package main

// ^ This struct will be used to organise the data which will be stored in DB
type FILE_STRUCT struct {
	UserID      string `json:"userID"`
	Void        string `json:"void"`
	FileName    string `json:"fileName"`
	FileType    string `json:"fileType"`
	FileContent string `json:"fileContent"`
}

// ^ This function will update the sturct
func (F *FILE_STRUCT) updateFileStruct(updateUserID string, updateVoid string, updateFilename string, updateFileType string, updateFileContent string) {
	F.UserID = updateUserID
	F.Void = updateVoid
	F.FileName = updateFilename
	F.FileType = updateFileType
	F.FileContent = updateFileContent
}

func createFileStruct() FILE_STRUCT {
	F := FILE_STRUCT{
		UserID:      "",
		Void:        "",
		FileName:    "",
		FileType:    "",
		FileContent: "",
	}

	return F
}
