// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson51ee17fDecodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels(in *jlexer.Lexer, out *Messages) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Messages":
			if in.IsNull() {
				in.Skip()
				out.Messages = nil
			} else {
				in.Delim('[')
				if out.Messages == nil {
					if !in.IsDelim(']') {
						out.Messages = make([]*Message, 0, 8)
					} else {
						out.Messages = []*Message{}
					}
				} else {
					out.Messages = (out.Messages)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Message
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Message)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Messages = append(out.Messages, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson51ee17fEncodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels(out *jwriter.Writer, in Messages) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Messages\":"
		out.RawString(prefix[1:])
		if in.Messages == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Messages {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Messages) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson51ee17fEncodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Messages) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson51ee17fEncodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Messages) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson51ee17fDecodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Messages) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson51ee17fDecodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels(l, v)
}
func easyjson51ee17fDecodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels1(in *jlexer.Lexer, out *Message) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "message_type":
			out.MessageType = int(in.Int())
		case "text":
			out.Text = string(in.String())
		case "author_id":
			out.AuthorID = uint64(in.Uint64())
		case "message_time":
			out.MessageTime = string(in.String())
		case "chat_id":
			out.ChatID = uint64(in.Uint64())
		case "file_id":
			out.FileID = string(in.String())
		case "sticker_id":
			out.StickerID = uint64(in.Uint64())
		case "file_type":
			out.FileType = string(in.String())
		case "likes":
			out.Likes = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson51ee17fEncodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels1(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"message_type\":"
		out.RawString(prefix)
		out.Int(int(in.MessageType))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"author_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.AuthorID))
	}
	{
		const prefix string = ",\"message_time\":"
		out.RawString(prefix)
		out.String(string(in.MessageTime))
	}
	{
		const prefix string = ",\"chat_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ChatID))
	}
	{
		const prefix string = ",\"file_id\":"
		out.RawString(prefix)
		out.String(string(in.FileID))
	}
	{
		const prefix string = ",\"sticker_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.StickerID))
	}
	{
		const prefix string = ",\"file_type\":"
		out.RawString(prefix)
		out.String(string(in.FileType))
	}
	{
		const prefix string = ",\"likes\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Likes))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson51ee17fEncodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson51ee17fEncodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson51ee17fDecodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson51ee17fDecodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels1(l, v)
}
