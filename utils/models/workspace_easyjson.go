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

func easyjson66c9e915DecodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels(in *jlexer.Lexer, out *Workspace) {
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
		case "name":
			out.Name = string(in.String())
		case "channels":
			if in.IsNull() {
				in.Skip()
				out.Channels = nil
			} else {
				in.Delim('[')
				if out.Channels == nil {
					if !in.IsDelim(']') {
						out.Channels = make([]*Channel, 0, 8)
					} else {
						out.Channels = []*Channel{}
					}
				} else {
					out.Channels = (out.Channels)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Channel
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Channel)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Channels = append(out.Channels, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "members":
			if in.IsNull() {
				in.Skip()
				out.Members = nil
			} else {
				in.Delim('[')
				if out.Members == nil {
					if !in.IsDelim(']') {
						out.Members = make([]uint64, 0, 8)
					} else {
						out.Members = []uint64{}
					}
				} else {
					out.Members = (out.Members)[:0]
				}
				for !in.IsDelim(']') {
					var v2 uint64
					v2 = uint64(in.Uint64())
					out.Members = append(out.Members, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "admins":
			if in.IsNull() {
				in.Skip()
				out.Admins = nil
			} else {
				in.Delim('[')
				if out.Admins == nil {
					if !in.IsDelim(']') {
						out.Admins = make([]uint64, 0, 8)
					} else {
						out.Admins = []uint64{}
					}
				} else {
					out.Admins = (out.Admins)[:0]
				}
				for !in.IsDelim(']') {
					var v3 uint64
					v3 = uint64(in.Uint64())
					out.Admins = append(out.Admins, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "creator_id":
			out.CreatorID = uint64(in.Uint64())
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
func easyjson66c9e915EncodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels(out *jwriter.Writer, in Workspace) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"channels\":"
		out.RawString(prefix)
		if in.Channels == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Channels {
				if v4 > 0 {
					out.RawByte(',')
				}
				if v5 == nil {
					out.RawString("null")
				} else {
					(*v5).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"members\":"
		out.RawString(prefix)
		if in.Members == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Members {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.Uint64(uint64(v7))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"admins\":"
		out.RawString(prefix)
		if in.Admins == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Admins {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.Uint64(uint64(v9))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"creator_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.CreatorID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Workspace) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson66c9e915EncodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Workspace) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson66c9e915EncodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Workspace) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson66c9e915DecodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Workspace) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson66c9e915DecodeGithubComCoolCodeTeam20192CoolCodeMicroServicesUtilsModels(l, v)
}
