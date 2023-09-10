package handler

type REQ_TYPE_ENUM int32

const (
	TYPE_GENERATE_TOKEN REQ_TYPE_ENUM = 200
	TYPE_REQUESTS_OCR   REQ_TYPE_ENUM = 201
)

func (reqType REQ_TYPE_ENUM) String() *string {
	switch reqType {
	case TYPE_GENERATE_TOKEN:
		message := "generate_token"
		return &message
	case TYPE_REQUESTS_OCR:
		message := "request token ocr"
		return &message
	default:
		return nil
	}
}
