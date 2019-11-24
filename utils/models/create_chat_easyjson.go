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

func easyjson40d320afDecodeGithubComGoParkMailRu20192CoolCodeMicroServicesUtilsModels(in *jlexer.Lexer, out *CreateChatModel) {
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
		case "user_id":
			out.UserID = uint64(in.Uint64())
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
func easyjson40d320afEncodeGithubComGoParkMailRu20192CoolCodeMicroServicesUtilsModels(out *jwriter.Writer, in CreateChatModel) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.UserID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateChatModel) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson40d320afEncodeGithubComGoParkMailRu20192CoolCodeMicroServicesUtilsModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateChatModel) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson40d320afEncodeGithubComGoParkMailRu20192CoolCodeMicroServicesUtilsModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateChatModel) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson40d320afDecodeGithubComGoParkMailRu20192CoolCodeMicroServicesUtilsModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateChatModel) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson40d320afDecodeGithubComGoParkMailRu20192CoolCodeMicroServicesUtilsModels(l, v)
}