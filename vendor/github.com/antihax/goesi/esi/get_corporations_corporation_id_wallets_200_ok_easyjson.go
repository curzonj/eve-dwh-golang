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

func easyjson6c932674DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdWallets200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdWallets200OkList, 0, 8)
			} else {
				*out = GetCorporationsCorporationIdWallets200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdWallets200Ok
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
func easyjson6c932674EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdWallets200OkList) {
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
func (v GetCorporationsCorporationIdWallets200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6c932674EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdWallets200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6c932674EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdWallets200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6c932674DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdWallets200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6c932674DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson6c932674DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdWallets200Ok) {
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
		case "balance":
			out.Balance = float32(in.Float32())
		case "division":
			out.Division = int32(in.Int32())
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
func easyjson6c932674EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdWallets200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Balance != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"balance\":")
		out.Float32(float32(in.Balance))
	}
	if in.Division != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"division\":")
		out.Int32(int32(in.Division))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdWallets200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6c932674EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdWallets200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6c932674EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdWallets200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6c932674DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdWallets200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6c932674DecodeGithubComCurzonjGoesiEsi1(l, v)
}
