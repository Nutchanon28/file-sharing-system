package constants

type FilePolicy struct {
	FileType string
	MaxSize  int32
}

var AppFilePolicy = FilePolicy{
	FileType: "application/pdf|image/png|image/jpg|image/jpeg",
	MaxSize:  10 * 1024 * 1024, //10MB
}

var QuestionFilePolicy = FilePolicy{
	FileType: "image/png|image/jpg|image/jpeg",
	MaxSize:  10 * 1024 * 1024, //10MB
}

const (
	TempFolder string = "temporary/"
)
