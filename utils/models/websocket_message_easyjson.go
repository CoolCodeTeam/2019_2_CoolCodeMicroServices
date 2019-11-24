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

func easyjson717ebd13DecodeGithubComGoParkMailRu20192CoolCodeMicroServicesUtilsModels(in *jlexer.Lexer, out *WebsocketMessage) {
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
		case "event_type":
			out.WebsocketEventType = int(in.Int())
		case "body":
			(out.Body).UnmarshalEasyJSON(in)
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
func easyjson717ebd13EncodeGithubComGoParkMailRu20192CoolCodeMicroServicesUtilsModels(out *jwriter.Writer, in WebsocketMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"event_type\":"
		out.RawString(prefix[1:])
		out.Int(int(in.WebsocketEventType))
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		(in.Body).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WebsocketMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson717ebd13EncodeGithubComGoParkMailRu20192CoolCodeMicroServicesUtilsModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WebsocketMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson717ebd13EncodeGithubComGoParkMailRu20192CoolCodeMicroServicesUtilsModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WebsocketMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson717ebd13DecodeGithubComGoParkMailRu20192CoolCodeMicroServicesUtilsModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WebsocketMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson717ebd13DecodeGithubComGoParkMailRu20192CoolCodeMicroServicesUtilsModels(l, v)
}
