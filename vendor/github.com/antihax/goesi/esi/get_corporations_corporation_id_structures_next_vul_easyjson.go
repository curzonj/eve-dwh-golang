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

func easyjsonD96c3a4eDecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdStructuresNextVulList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdStructuresNextVulList, 0, 8)
			} else {
				*out = GetCorporationsCorporationIdStructuresNextVulList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdStructuresNextVul
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
func easyjsonD96c3a4eEncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdStructuresNextVulList) {
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
func (v GetCorporationsCorporationIdStructuresNextVulList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD96c3a4eEncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdStructuresNextVulList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD96c3a4eEncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdStructuresNextVulList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD96c3a4eDecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdStructuresNextVulList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD96c3a4eDecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjsonD96c3a4eDecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdStructuresNextVul) {
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
		case "day":
			out.Day = int32(in.Int32())
		case "hour":
			out.Hour = int32(in.Int32())
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
func easyjsonD96c3a4eEncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdStructuresNextVul) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Day != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"day\":")
		out.Int32(int32(in.Day))
	}
	if in.Hour != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"hour\":")
		out.Int32(int32(in.Hour))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdStructuresNextVul) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD96c3a4eEncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdStructuresNextVul) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD96c3a4eEncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdStructuresNextVul) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD96c3a4eDecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdStructuresNextVul) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD96c3a4eDecodeGithubComCurzonjGoesiEsi1(l, v)
}
