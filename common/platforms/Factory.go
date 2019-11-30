package platforms

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) NewPlatform(pt PlatformType, username, password, nickname string, picBlob []byte) *Platform {
	return &Platform{
		Type:           pt,
		Username:       username,
		Password:       password,
		Nickname:       nickname,
		ProfilePicture: picBlob,
	}
}
