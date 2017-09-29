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

func easyjson429021f6DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdChatChannelsBlockedList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdChatChannelsBlockedList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdChatChannelsBlockedList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdChatChannelsBlocked
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
func easyjson429021f6EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdChatChannelsBlockedList) {
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
func (v GetCharactersCharacterIdChatChannelsBlockedList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson429021f6EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdChatChannelsBlockedList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson429021f6EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdChatChannelsBlockedList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson429021f6DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdChatChannelsBlockedList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson429021f6DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson429021f6DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdChatChannelsBlocked) {
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
		case "accessor_id":
			out.AccessorId = int32(in.Int32())
		case "accessor_type":
			out.AccessorType = string(in.String())
		case "end_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EndAt).UnmarshalJSON(data))
			}
		case "reason":
			out.Reason = string(in.String())
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
func easyjson429021f6EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdChatChannelsBlocked) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AccessorId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accessor_id\":")
		out.Int32(int32(in.AccessorId))
	}
	if in.AccessorType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accessor_type\":")
		out.String(string(in.AccessorType))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"end_at\":")
		out.Raw((in.EndAt).MarshalJSON())
	}
	if in.Reason != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"reason\":")
		out.String(string(in.Reason))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdChatChannelsBlocked) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson429021f6EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdChatChannelsBlocked) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson429021f6EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdChatChannelsBlocked) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson429021f6DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdChatChannelsBlocked) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson429021f6DecodeGithubComCurzonjGoesiEsi1(l, v)
}
