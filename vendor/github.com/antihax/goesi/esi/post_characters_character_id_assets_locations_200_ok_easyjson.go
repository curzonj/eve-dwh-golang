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

func easyjson699ff4b8DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *PostCharactersCharacterIdAssetsLocations200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostCharactersCharacterIdAssetsLocations200OkList, 0, 2)
			} else {
				*out = PostCharactersCharacterIdAssetsLocations200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostCharactersCharacterIdAssetsLocations200Ok
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
func easyjson699ff4b8EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in PostCharactersCharacterIdAssetsLocations200OkList) {
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
func (v PostCharactersCharacterIdAssetsLocations200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson699ff4b8EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdAssetsLocations200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson699ff4b8EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdAssetsLocations200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson699ff4b8DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdAssetsLocations200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson699ff4b8DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson699ff4b8DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *PostCharactersCharacterIdAssetsLocations200Ok) {
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
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		case "z":
			out.Z = float64(in.Float64())
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
func easyjson699ff4b8EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in PostCharactersCharacterIdAssetsLocations200Ok) {
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
	if in.X != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"x\":")
		out.Float64(float64(in.X))
	}
	if in.Y != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"y\":")
		out.Float64(float64(in.Y))
	}
	if in.Z != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"z\":")
		out.Float64(float64(in.Z))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostCharactersCharacterIdAssetsLocations200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson699ff4b8EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdAssetsLocations200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson699ff4b8EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdAssetsLocations200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson699ff4b8DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdAssetsLocations200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson699ff4b8DecodeGithubComCurzonjGoesiEsi1(l, v)
}
