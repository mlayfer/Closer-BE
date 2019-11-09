package platforms

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) NewPlatform(pt PlatformType, data string) *Platform {
	pd, err := LoadData(data)
	if err != nil {
		return nil
	}
	return &Platform{
		Type:          pt,
		PlatformsData: pd,
	}
}
