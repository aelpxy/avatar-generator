package lib

import (
	"fmt"
	"hash/fnv"

	"github.com/lucasb-eyer/go-colorful"
)

func stringToColor(str string) string {
	h := fnv.New32a()
	h.Write([]byte(str))
	hash := h.Sum32()
	color := colorful.Hsv(float64(hash%360), 0.5, 1)
	return color.Hex()
}

func generateColor(str string) (string, string) {
	st1 := str[:len(str)/2]
	st2 := str[len(str)/2:]
	colorOne := stringToColor(st1)
	colorTwo := stringToColor(st2)

	return colorOne, colorTwo
}

func GenerateImage(str string) string {
	size := 512
	colorOne, colorTwo := generateColor(str)
	image := fmt.Sprintf(`
    <svg width="%d" height="%d" viewBox="0 0 %d %d" fill="none" xmlns="http://www.w3.org/2000/svg">
      <circle cx="%d" cy="%d" r="%d" fill="url(#gradient)" />
      <defs>
        <linearGradient id="gradient" x1="0" y1="0" x2="%d" y2="%d" gradientUnits="userSpaceOnUse">
          <stop stop-color="%s" />
          <stop offset="1" stop-color="%s"  />
        </linearGradient>
      </defs>
    </svg>
    `, size, size, size, size, size/2, size/2, size/2, size, size, colorOne, colorTwo)

	return image

}
