package chat

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (a AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if len(url) == 0 {
		return "", ErrorNoAvatarURL
	}

	return url, nil
}
