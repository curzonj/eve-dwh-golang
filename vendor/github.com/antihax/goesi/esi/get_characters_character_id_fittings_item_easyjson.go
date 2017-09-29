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

func easyjsonF56e5f30DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdFittingsItemList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdFittingsItemList, 0, 5)
			} else {
				*out = GetCharactersCharacterIdFittingsItemList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdFittingsItem
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
func easyjsonF56e5f30EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdFittingsItemList) {
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
func (v GetCharactersCharacterIdFittingsItemList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF56e5f30EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdFittingsItemList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF56e5f30EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdFittingsItemList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF56e5f30DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdFittingsItemList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF56e5f30DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjsonF56e5f30DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdFittingsItem) {
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
		case "flag":
			out.Flag = int32(in.Int32())
		case "quantity":
			out.Quantity = int32(in.Int32())
		case "type_id":
			out.TypeId = int32(in.Int32())
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
func easyjsonF56e5f30EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdFittingsItem) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Flag != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"flag\":")
		out.Int32(int32(in.Flag))
	}
	if in.Quantity != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"quantity\":")
		out.Int32(int32(in.Quantity))
	}
	if in.TypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"type_id\":")
		out.Int32(int32(in.TypeId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdFittingsItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF56e5f30EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdFittingsItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF56e5f30EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdFittingsItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF56e5f30DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdFittingsItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF56e5f30DecodeGithubComCurzonjGoesiEsi1(l, v)
}
