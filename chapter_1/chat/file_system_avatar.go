package chat

import (
	"io/ioutil"
	"path"
)

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	files, err := ioutil.ReadDir("avatars")
	if err != nil {
		return "", ErrorNoAvatarURL
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if match, _ := path.Match(u.UniqueID()+"*", file.Name()); match {
			return "/avatars/" + file.Name(), nil
		}
	}

	return "", ErrorNoAvatarURL
}
