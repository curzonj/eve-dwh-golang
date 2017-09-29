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

func easyjson135718d4DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetFwLeaderboardsCorporationsYesterday1List) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwLeaderboardsCorporationsYesterday1List, 0, 8)
			} else {
				*out = GetFwLeaderboardsCorporationsYesterday1List{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwLeaderboardsCorporationsYesterday1
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
func easyjson135718d4EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetFwLeaderboardsCorporationsYesterday1List) {
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
func (v GetFwLeaderboardsCorporationsYesterday1List) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson135718d4EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCorporationsYesterday1List) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson135718d4EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsYesterday1List) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson135718d4DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsYesterday1List) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson135718d4DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson135718d4DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetFwLeaderboardsCorporationsYesterday1) {
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
		case "amount":
			out.Amount = int32(in.Int32())
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
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
func easyjson135718d4EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetFwLeaderboardsCorporationsYesterday1) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"amount\":")
		out.Int32(int32(in.Amount))
	}
	if in.CorporationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"corporation_id\":")
		out.Int32(int32(in.CorporationId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsCorporationsYesterday1) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson135718d4EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCorporationsYesterday1) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson135718d4EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsYesterday1) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson135718d4DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsYesterday1) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson135718d4DecodeGithubComCurzonjGoesiEsi1(l, v)
}
