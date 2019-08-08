package alice

const (
	// BigImageType одна большая картинка.
	BigImageType = "BigImage"

	// ItemsListType список с картинками.
	ItemsListType = "ItemsList"
)

// Image структура картинки.
type Image struct {
	ImageID     string       `json:"image_id"`
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	Button      *ImageButton `json:"button,omitempty"`
}

// ImageButton структура кнопки для добавления интерактивности картинке.
type ImageButton struct {
	Title   string      `json:"title,omitempty"`
	URL     string      `json:"url,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

// NewImageButton создает кнопку для добавления интерактивности картинке. Payload может быть проигнорирован. Если задано больше одного payload используется только первый.
func NewImageButton(title, url string, payload ...interface{}) ImageButton {
	var p interface{}
	if len(payload) > 0 {
		p = payload[0]
	}

	return ImageButton{
		Title:   title,
		URL:     url,
		Payload: p,
	}
}

// Card структура для добавления в ответ картинок.
type Card struct {
	ImageID     string       `json:"image_id,omitempty"`
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	Button      *ImageButton `json:"button,omitempty"`

	Type   string `json:"type"`
	Header *struct {
		Text string `json:"text,omitempty"`
	} `json:"header,omitempty"`
	Items  []Image `json:"items,omitempty"`
	Footer *struct {
		Text   string       `json:"text,omitempty"`
		Button *ImageButton `json:"button,omitempty"`
	} `json:"footer,omitempty"`
}

// BigImage добавляет к ответу одну большую картинку. Кнопка может быть проигнорирована. Если задано больше одной кнопки используется только первая. Требуемые размеры изображения: 1x — 388x172, 1.5x — 582x258, 2x — 776x344, 3x — 1164x516, 3.5x — 1358x602, 4x — 1552x688.
func (resp *Response) BigImage(id, title, desc string, button ...ImageButton) *Response {
	var b *ImageButton
	if len(button) > 0 {
		b = &button[0]
	}

	image := &Card{
		Type:        BigImageType,
		ImageID:     id,
		Title:       title,
		Description: desc,
		Button:      b,
	}

	resp.Response.Card = image
	return resp
}

// List хранилище картинок для создания списка.
type List struct {
	Images []Image
}

// Add добавляет картинку в заготовку для списка. Кнопка может быть проигнорирована. Если задано больше одной кнопки используется только первая.
func (l *List) Add(id, title, desc string, button ...ImageButton) *List {
	var b *ImageButton
	if len(button) > 0 {
		b = &button[0]
	}

	image := Image{
		ImageID:     id,
		Title:       title,
		Description: desc,
		Button:      b,
	}

	return l.AddImages(image)
}

// AddImages добавляет одну или несколько картинок в заготовку для списка.
func (l *List) AddImages(images ...Image) *List {
	for _, v := range images {
		l.Images = append(l.Images, v)
	}
	return l
}

// List добавляет к ответу список с картинками. Кнопка может быть проигнорирована. Если задано больше одной кнопки используется только первая.
func (resp *Response) List(header, footer string, l List, button ...ImageButton) *Response {
	var b *ImageButton
	if len(button) > 0 {
		b = &button[0]
	}

	if len(l.Images) > 5 {
		l.Images = l.Images[:5]
	}

	list := &Card{
		Type: ItemsListType,
		Header: &struct {
			Text string `json:"text,omitempty"`
		}{
			Text: header,
		},
		Items: l.Images,
		Footer: &struct {
			Text   string       `json:"text,omitempty"`
			Button *ImageButton `json:"button,omitempty"`
		}{
			Text:   footer,
			Button: b,
		},
	}

	resp.Response.Card = list
	return resp
}
