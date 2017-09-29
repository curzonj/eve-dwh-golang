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

func easyjsonA53f7cDecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetInsurancePrices200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetInsurancePrices200OkList, 0, 2)
			} else {
				*out = GetInsurancePrices200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetInsurancePrices200Ok
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
func easyjsonA53f7cEncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetInsurancePrices200OkList) {
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
func (v GetInsurancePrices200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA53f7cEncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetInsurancePrices200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA53f7cEncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetInsurancePrices200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA53f7cDecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetInsurancePrices200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA53f7cDecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjsonA53f7cDecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetInsurancePrices200Ok) {
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
		case "levels":
			if in.IsNull() {
				in.Skip()
				out.Levels = nil
			} else {
				in.Delim('[')
				if out.Levels == nil {
					if !in.IsDelim(']') {
						out.Levels = make([]GetInsurancePricesLevel, 0, 2)
					} else {
						out.Levels = []GetInsurancePricesLevel{}
					}
				} else {
					out.Levels = (out.Levels)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetInsurancePricesLevel
					easyjsonA53f7cDecodeGithubComCurzonjGoesiEsi2(in, &v4)
					out.Levels = append(out.Levels, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "type_id":
			out.TypeId = int32(in.Int32())
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
func easyjsonA53f7cEncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetInsurancePrices200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Levels) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"levels\":")
		if in.Levels == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Levels {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjsonA53f7cEncodeGithubComCurzonjGoesiEsi2(out, v6)
			}
			out.RawByte(']')
		}
	}
	if in.TypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"type_id\":")
		out.Int32(int32(in.TypeId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetInsurancePrices200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA53f7cEncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetInsurancePrices200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA53f7cEncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetInsurancePrices200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA53f7cDecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetInsurancePrices200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA53f7cDecodeGithubComCurzonjGoesiEsi1(l, v)
}
func easyjsonA53f7cDecodeGithubComCurzonjGoesiEsi2(in *jlexer.Lexer, out *GetInsurancePricesLevel) {
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
			out.Cost = float32(in.Float32())
		case "name":
			out.Name = string(in.String())
		case "payout":
			out.Payout = float32(in.Float32())
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
func easyjsonA53f7cEncodeGithubComCurzonjGoesiEsi2(out *jwriter.Writer, in GetInsurancePricesLevel) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Cost != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cost\":")
		out.Float32(float32(in.Cost))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.Payout != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"payout\":")
		out.Float32(float32(in.Payout))
	}
	out.RawByte('}')
}
