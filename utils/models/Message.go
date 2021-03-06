package models

//1 - сообщение
//2 - чувак набирает

type Message struct {
	ID            uint64 `json:"id"`
	MessageType   int    `json:"message_type"`
	Text          string `json:"text"`
	AuthorID      uint64 `json:"author_id"`
	MessageTime   string `json:"message_time"`
	ChatID        uint64 `json:"chat_id"`
	FileID        string `json:"file_id"`
	StickerID     string `json:"sticker_id"`
	FileType      string `json:"file_type"`
	HideForAuthor bool   `json:"-"`
	Likes         uint64 `json:"likes"`
}

type Messages struct {
	Messages []*Message
}
