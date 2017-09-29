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

func easyjsonDce3bba4DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *PostCharactersCharacterIdAssetsNames200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostCharactersCharacterIdAssetsNames200OkList, 0, 2)
			} else {
				*out = PostCharactersCharacterIdAssetsNames200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostCharactersCharacterIdAssetsNames200Ok
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
func easyjsonDce3bba4EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in PostCharactersCharacterIdAssetsNames200OkList) {
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
func (v PostCharactersCharacterIdAssetsNames200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDce3bba4EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdAssetsNames200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDce3bba4EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdAssetsNames200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDce3bba4DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdAssetsNames200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDce3bba4DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjsonDce3bba4DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *PostCharactersCharacterIdAssetsNames200Ok) {
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
		case "item_id":
			out.ItemId = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
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
func easyjsonDce3bba4EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in PostCharactersCharacterIdAssetsNames200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ItemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"item_id\":")
		out.Int64(int64(in.ItemId))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostCharactersCharacterIdAssetsNames200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDce3bba4EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdAssetsNames200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDce3bba4EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdAssetsNames200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDce3bba4DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdAssetsNames200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDce3bba4DecodeGithubComCurzonjGoesiEsi1(l, v)
}
