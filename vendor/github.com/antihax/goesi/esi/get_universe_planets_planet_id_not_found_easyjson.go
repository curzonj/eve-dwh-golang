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

func easyjsonE0452e36DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetUniversePlanetsPlanetIdNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniversePlanetsPlanetIdNotFoundList, 0, 4)
			} else {
				*out = GetUniversePlanetsPlanetIdNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniversePlanetsPlanetIdNotFound
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
func easyjsonE0452e36EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetUniversePlanetsPlanetIdNotFoundList) {
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
func (v GetUniversePlanetsPlanetIdNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE0452e36EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniversePlanetsPlanetIdNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE0452e36EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniversePlanetsPlanetIdNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE0452e36DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniversePlanetsPlanetIdNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE0452e36DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjsonE0452e36DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetUniversePlanetsPlanetIdNotFound) {
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
		case "error":
			out.Error_ = string(in.String())
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
func easyjsonE0452e36EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetUniversePlanetsPlanetIdNotFound) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Error_ != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"error\":")
		out.String(string(in.Error_))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniversePlanetsPlanetIdNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE0452e36EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniversePlanetsPlanetIdNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE0452e36EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniversePlanetsPlanetIdNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE0452e36DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniversePlanetsPlanetIdNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE0452e36DecodeGithubComCurzonjGoesiEsi1(l, v)
}
