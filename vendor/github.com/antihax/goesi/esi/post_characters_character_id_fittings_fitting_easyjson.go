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

func easyjson746f8c8DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *PostCharactersCharacterIdFittingsFittingList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostCharactersCharacterIdFittingsFittingList, 0, 1)
			} else {
				*out = PostCharactersCharacterIdFittingsFittingList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostCharactersCharacterIdFittingsFitting
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
func easyjson746f8c8EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in PostCharactersCharacterIdFittingsFittingList) {
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
func (v PostCharactersCharacterIdFittingsFittingList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson746f8c8EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdFittingsFittingList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson746f8c8EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdFittingsFittingList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson746f8c8DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdFittingsFittingList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson746f8c8DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson746f8c8DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *PostCharactersCharacterIdFittingsFitting) {
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
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]PostCharactersCharacterIdFittingsItem, 0, 5)
					} else {
						out.Items = []PostCharactersCharacterIdFittingsItem{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v4 PostCharactersCharacterIdFittingsItem
					easyjson746f8c8DecodeGithubComCurzonjGoesiEsi2(in, &v4)
					out.Items = append(out.Items, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "name":
			out.Name = string(in.String())
		case "ship_type_id":
			out.ShipTypeId = int32(in.Int32())
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
func easyjson746f8c8EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in PostCharactersCharacterIdFittingsFitting) {
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
	if len(in.Items) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"items\":")
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Items {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson746f8c8EncodeGithubComCurzonjGoesiEsi2(out, v6)
			}
			out.RawByte(']')
		}
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.ShipTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ship_type_id\":")
		out.Int32(int32(in.ShipTypeId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostCharactersCharacterIdFittingsFitting) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson746f8c8EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdFittingsFitting) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson746f8c8EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdFittingsFitting) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson746f8c8DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdFittingsFitting) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson746f8c8DecodeGithubComCurzonjGoesiEsi1(l, v)
}
func easyjson746f8c8DecodeGithubComCurzonjGoesiEsi2(in *jlexer.Lexer, out *PostCharactersCharacterIdFittingsItem) {
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
		case "flag":
			out.Flag = int32(in.Int32())
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
func easyjson746f8c8EncodeGithubComCurzonjGoesiEsi2(out *jwriter.Writer, in PostCharactersCharacterIdFittingsItem) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Flag != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"flag\":")
		out.Int32(int32(in.Flag))
	}
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
