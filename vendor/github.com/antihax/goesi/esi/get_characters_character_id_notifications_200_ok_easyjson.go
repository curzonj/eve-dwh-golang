// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

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

func easyjson6cda8914DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdNotifications200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdNotifications200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdNotifications200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdNotifications200Ok
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6cda8914EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdNotifications200OkList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdNotifications200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6cda8914EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdNotifications200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6cda8914EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdNotifications200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6cda8914DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdNotifications200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6cda8914DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson6cda8914DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdNotifications200Ok) {
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
		case "is_read":
			out.IsRead = bool(in.Bool())
		case "notification_id":
			out.NotificationId = int64(in.Int64())
		case "sender_id":
			out.SenderId = int32(in.Int32())
		case "sender_type":
			out.SenderType = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "timestamp":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "type":
			out.Type_ = string(in.String())
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
func easyjson6cda8914EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdNotifications200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.IsRead {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"is_read\":")
		out.Bool(bool(in.IsRead))
	}
	if in.NotificationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"notification_id\":")
		out.Int64(int64(in.NotificationId))
	}
	if in.SenderId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"sender_id\":")
		out.Int32(int32(in.SenderId))
	}
	if in.SenderType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"sender_type\":")
		out.String(string(in.SenderType))
	}
	if in.Text != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"text\":")
		out.String(string(in.Text))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"timestamp\":")
		out.Raw((in.Timestamp).MarshalJSON())
	}
	if in.Type_ != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"type\":")
		out.String(string(in.Type_))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdNotifications200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6cda8914EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdNotifications200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6cda8914EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdNotifications200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6cda8914DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdNotifications200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6cda8914DecodeGithubComAntihaxGoesiEsi1(l, v)
}
