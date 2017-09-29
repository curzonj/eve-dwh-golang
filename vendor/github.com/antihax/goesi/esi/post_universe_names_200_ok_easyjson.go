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

func easyjson84341653DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *PostUniverseNames200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostUniverseNames200OkList, 0, 1)
			} else {
				*out = PostUniverseNames200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostUniverseNames200Ok
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
func easyjson84341653EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in PostUniverseNames200OkList) {
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
func (v PostUniverseNames200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84341653EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUniverseNames200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84341653EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUniverseNames200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84341653DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUniverseNames200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84341653DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson84341653DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *PostUniverseNames200Ok) {
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
		case "category":
			out.Category = string(in.String())
		case "id":
			out.Id = int32(in.Int32())
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
func easyjson84341653EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in PostUniverseNames200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Category != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"category\":")
		out.String(string(in.Category))
	}
	if in.Id != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.Int32(int32(in.Id))
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
func (v PostUniverseNames200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson84341653EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUniverseNames200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson84341653EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUniverseNames200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson84341653DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUniverseNames200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson84341653DecodeGithubComCurzonjGoesiEsi1(l, v)
}
