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

func easyjsonB318302DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdStructuresCurrentVulList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdStructuresCurrentVulList, 0, 8)
			} else {
				*out = GetCorporationsCorporationIdStructuresCurrentVulList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdStructuresCurrentVul
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
func easyjsonB318302EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdStructuresCurrentVulList) {
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
func (v GetCorporationsCorporationIdStructuresCurrentVulList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB318302EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdStructuresCurrentVulList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB318302EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdStructuresCurrentVulList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB318302DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdStructuresCurrentVulList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB318302DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjsonB318302DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdStructuresCurrentVul) {
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
func easyjsonB318302EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdStructuresCurrentVul) {
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
func (v GetCorporationsCorporationIdStructuresCurrentVul) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB318302EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdStructuresCurrentVul) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB318302EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdStructuresCurrentVul) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB318302DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdStructuresCurrentVul) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB318302DecodeGithubComCurzonjGoesiEsi1(l, v)
}
