package asset

import (
	"erros"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/faiface/pixel"
)

type Load struct {
	filesystem fs.FS
}

func NewLoad(filesystem fs.FS) *Load {
	return &Load{filesystem}
}

func (load *Load) Open(path string) (fs.File, error) {
	return load.filesystem.Open(path)
}

func (load *Load) Sprite(path string) (*pixel.Sprite, error) {
	file, err := load.filesystem.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	pic := pixel.PictureDataFromImage(img)

	return pixel.NewSprite(pic, pic.Bounds()), nil
}

type Spritesheet struct {
	picture pixel.Picture
	lookup  map[string]*pixel.Sprite
}

func NewSpritesheet(pic pixel.Picture, lookup map[string]*pixel.Sprite) *Spritesheet {
	return &Spritesheet{
		picture: pic,
		lookup:  lookup,
	}
}

func (s *Spritesheet) Get(name string) (*pixel.Sprite, error) {
	sprite, ok := s.lookup[name]
	if !ok {
		return erros.New("Invalid sprite name!")
	}
	return sprite
}
