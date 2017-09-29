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

func easyjson2c2475fbDecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetLoyaltyStoresCorporationIdOffersRequiredItemList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetLoyaltyStoresCorporationIdOffersRequiredItemList, 0, 8)
			} else {
				*out = GetLoyaltyStoresCorporationIdOffersRequiredItemList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetLoyaltyStoresCorporationIdOffersRequiredItem
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
func easyjson2c2475fbEncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetLoyaltyStoresCorporationIdOffersRequiredItemList) {
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
func (v GetLoyaltyStoresCorporationIdOffersRequiredItemList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2c2475fbEncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetLoyaltyStoresCorporationIdOffersRequiredItemList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2c2475fbEncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetLoyaltyStoresCorporationIdOffersRequiredItemList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2c2475fbDecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetLoyaltyStoresCorporationIdOffersRequiredItemList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2c2475fbDecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson2c2475fbDecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetLoyaltyStoresCorporationIdOffersRequiredItem) {
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
		case "quantity":
			out.Quantity = int32(in.Int32())
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
func easyjson2c2475fbEncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetLoyaltyStoresCorporationIdOffersRequiredItem) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Quantity != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"quantity\":")
		out.Int32(int32(in.Quantity))
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
func (v GetLoyaltyStoresCorporationIdOffersRequiredItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2c2475fbEncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetLoyaltyStoresCorporationIdOffersRequiredItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2c2475fbEncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetLoyaltyStoresCorporationIdOffersRequiredItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2c2475fbDecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetLoyaltyStoresCorporationIdOffersRequiredItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2c2475fbDecodeGithubComCurzonjGoesiEsi1(l, v)
}
