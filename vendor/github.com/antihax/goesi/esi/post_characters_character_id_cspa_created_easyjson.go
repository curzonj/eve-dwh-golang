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

func easyjson891887feDecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *PostCharactersCharacterIdCspaCreatedList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostCharactersCharacterIdCspaCreatedList, 0, 8)
			} else {
				*out = PostCharactersCharacterIdCspaCreatedList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostCharactersCharacterIdCspaCreated
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
func easyjson891887feEncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in PostCharactersCharacterIdCspaCreatedList) {
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
func (v PostCharactersCharacterIdCspaCreatedList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson891887feEncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdCspaCreatedList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson891887feEncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdCspaCreatedList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson891887feDecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdCspaCreatedList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson891887feDecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson891887feDecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *PostCharactersCharacterIdCspaCreated) {
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
		case "cost":
			out.Cost = int64(in.Int64())
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
func easyjson891887feEncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in PostCharactersCharacterIdCspaCreated) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Cost != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cost\":")
		out.Int64(int64(in.Cost))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostCharactersCharacterIdCspaCreated) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson891887feEncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdCspaCreated) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson891887feEncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdCspaCreated) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson891887feDecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdCspaCreated) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson891887feDecodeGithubComCurzonjGoesiEsi1(l, v)
}
