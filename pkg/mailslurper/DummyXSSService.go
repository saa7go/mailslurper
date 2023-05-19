package mailslurper

type DummyXSSService struct {
}

func NewDummyXSSService() *DummyXSSService {
	return &DummyXSSService{}
}

func (service *DummyXSSService) SanitizeString(input string) string {
	return input
}
