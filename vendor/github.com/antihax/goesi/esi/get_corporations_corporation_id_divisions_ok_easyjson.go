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

func easyjson83718d43DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdDivisionsOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdDivisionsOkList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdDivisionsOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdDivisionsOk
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
func easyjson83718d43EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdDivisionsOkList) {
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
func (v GetCorporationsCorporationIdDivisionsOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83718d43EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdDivisionsOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83718d43EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdDivisionsOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83718d43DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdDivisionsOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83718d43DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson83718d43DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdDivisionsOk) {
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
		case "hangar":
			if in.IsNull() {
				in.Skip()
				out.Hangar = nil
			} else {
				in.Delim('[')
				if out.Hangar == nil {
					if !in.IsDelim(']') {
						out.Hangar = make([]GetCorporationsCorporationIdDivisionsHangar, 0, 2)
					} else {
						out.Hangar = []GetCorporationsCorporationIdDivisionsHangar{}
					}
				} else {
					out.Hangar = (out.Hangar)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCorporationsCorporationIdDivisionsHangar
					(v4).UnmarshalEasyJSON(in)
					out.Hangar = append(out.Hangar, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "wallet":
			if in.IsNull() {
				in.Skip()
				out.Wallet = nil
			} else {
				in.Delim('[')
				if out.Wallet == nil {
					if !in.IsDelim(']') {
						out.Wallet = make([]GetCorporationsCorporationIdDivisionsWallet, 0, 2)
					} else {
						out.Wallet = []GetCorporationsCorporationIdDivisionsWallet{}
					}
				} else {
					out.Wallet = (out.Wallet)[:0]
				}
				for !in.IsDelim(']') {
					var v5 GetCorporationsCorporationIdDivisionsWallet
					easyjson83718d43DecodeGithubComCurzonjGoesiEsi2(in, &v5)
					out.Wallet = append(out.Wallet, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson83718d43EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdDivisionsOk) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Hangar) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"hangar\":")
		if in.Hangar == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Hangar {
				if v6 > 0 {
					out.RawByte(',')
				}
				(v7).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if len(in.Wallet) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"wallet\":")
		if in.Wallet == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Wallet {
				if v8 > 0 {
					out.RawByte(',')
				}
				easyjson83718d43EncodeGithubComCurzonjGoesiEsi2(out, v9)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdDivisionsOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83718d43EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdDivisionsOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83718d43EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdDivisionsOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83718d43DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdDivisionsOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83718d43DecodeGithubComCurzonjGoesiEsi1(l, v)
}
func easyjson83718d43DecodeGithubComCurzonjGoesiEsi2(in *jlexer.Lexer, out *GetCorporationsCorporationIdDivisionsWallet) {
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
		case "division":
			out.Division = int32(in.Int32())
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
func easyjson83718d43EncodeGithubComCurzonjGoesiEsi2(out *jwriter.Writer, in GetCorporationsCorporationIdDivisionsWallet) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Division != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"division\":")
		out.Int32(int32(in.Division))
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
