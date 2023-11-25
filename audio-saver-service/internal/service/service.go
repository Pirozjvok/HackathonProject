package service

type Audio interface {
	GetAudio(id string) ([]byte, error)
}

type Services struct {
	Audio
}

type Dependencies struct {
	URL string
}

func NewServices(deps *Dependencies) *Services {
	return &Services{
		Audio: newAudioService(deps.URL),
	}
}
