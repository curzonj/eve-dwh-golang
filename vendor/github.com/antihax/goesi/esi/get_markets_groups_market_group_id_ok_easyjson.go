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

func easyjson288d22fbDecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetMarketsGroupsMarketGroupIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetMarketsGroupsMarketGroupIdOkList, 0, 1)
			} else {
				*out = GetMarketsGroupsMarketGroupIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetMarketsGroupsMarketGroupIdOk
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
func easyjson288d22fbEncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetMarketsGroupsMarketGroupIdOkList) {
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
func (v GetMarketsGroupsMarketGroupIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson288d22fbEncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetMarketsGroupsMarketGroupIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson288d22fbEncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetMarketsGroupsMarketGroupIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson288d22fbDecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetMarketsGroupsMarketGroupIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson288d22fbDecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson288d22fbDecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetMarketsGroupsMarketGroupIdOk) {
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
		case "description":
			out.Description = string(in.String())
		case "market_group_id":
			out.MarketGroupId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "parent_group_id":
			out.ParentGroupId = int32(in.Int32())
		case "types":
			if in.IsNull() {
				in.Skip()
				out.Types = nil
			} else {
				in.Delim('[')
				if out.Types == nil {
					if !in.IsDelim(']') {
						out.Types = make([]int32, 0, 16)
					} else {
						out.Types = []int32{}
					}
				} else {
					out.Types = (out.Types)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Types = append(out.Types, v4)
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
func easyjson288d22fbEncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetMarketsGroupsMarketGroupIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Description != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"description\":")
		out.String(string(in.Description))
	}
	if in.MarketGroupId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"market_group_id\":")
		out.Int32(int32(in.MarketGroupId))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.ParentGroupId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"parent_group_id\":")
		out.Int32(int32(in.ParentGroupId))
	}
	if len(in.Types) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"types\":")
		if in.Types == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Types {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetMarketsGroupsMarketGroupIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson288d22fbEncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetMarketsGroupsMarketGroupIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson288d22fbEncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetMarketsGroupsMarketGroupIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson288d22fbDecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetMarketsGroupsMarketGroupIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson288d22fbDecodeGithubComCurzonjGoesiEsi1(l, v)
}
