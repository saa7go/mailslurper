package mailslurper

type DummyXSSService struct {
	data string
}

func NewDummyXSSService() *DummyXSSService {
	return &DummyXSSService{
		data: "",
	}
}

func (service *DummyXSSService) SanitizeString(input string) string {
	service.data = input
	return service.data
}
