package usecase

type Chat interface {
}

func NewChat() Chat {
	return &chat{}
}

type chat struct {
}
