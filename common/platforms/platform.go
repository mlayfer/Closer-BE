package platforms

import "encoding/json"

type PlatformType int

const (
	_                     = iota // ignore 0
	Facebook PlatformType = iota
	Instagram
	Linkedin
	Snapchat
	Tiktok
	Youtube
	Pinterest
	Google
)

type PlatformData struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Nickname   string `json:"nickname" binding:"required"`
	ProfilePic []byte `json:"profilepic" binding:"required"`
}

func (pd *PlatformData) String() (string, error) {
	b, err := json.Marshal(pd)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func LoadData(data string) (*PlatformData, error) {
	pd := &PlatformData{}
	err := json.Unmarshal([]byte(data), pd)
	if err != nil {
		return nil, err
	}
	return pd, nil
}

type Platform struct {
	Type          PlatformType `json:"type" binding:"required"`
	PlatformData *PlatformData `json:"platformdata" binding:"required"`
}
