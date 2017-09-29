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

func easyjson303db286DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdMedals200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdMedals200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdMedals200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdMedals200Ok
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
func easyjson303db286EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdMedals200OkList) {
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
func (v GetCharactersCharacterIdMedals200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson303db286EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdMedals200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson303db286EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdMedals200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson303db286DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdMedals200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson303db286DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson303db286DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdMedals200Ok) {
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
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		case "description":
			out.Description = string(in.String())
		case "graphics":
			if in.IsNull() {
				in.Skip()
				out.Graphics = nil
			} else {
				in.Delim('[')
				if out.Graphics == nil {
					if !in.IsDelim(']') {
						out.Graphics = make([]GetCharactersCharacterIdMedalsGraphic, 0, 2)
					} else {
						out.Graphics = []GetCharactersCharacterIdMedalsGraphic{}
					}
				} else {
					out.Graphics = (out.Graphics)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdMedalsGraphic
					easyjson303db286DecodeGithubComCurzonjGoesiEsi2(in, &v4)
					out.Graphics = append(out.Graphics, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "issuer_id":
			out.IssuerId = int32(in.Int32())
		case "medal_id":
			out.MedalId = int32(in.Int32())
		case "reason":
			out.Reason = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "title":
			out.Title = string(in.String())
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
func easyjson303db286EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdMedals200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CorporationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"corporation_id\":")
		out.Int32(int32(in.CorporationId))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"date\":")
		out.Raw((in.Date).MarshalJSON())
	}
	if in.Description != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"description\":")
		out.String(string(in.Description))
	}
	if len(in.Graphics) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"graphics\":")
		if in.Graphics == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Graphics {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson303db286EncodeGithubComCurzonjGoesiEsi2(out, v6)
			}
			out.RawByte(']')
		}
	}
	if in.IssuerId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"issuer_id\":")
		out.Int32(int32(in.IssuerId))
	}
	if in.MedalId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"medal_id\":")
		out.Int32(int32(in.MedalId))
	}
	if in.Reason != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"reason\":")
		out.String(string(in.Reason))
	}
	if in.Status != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"status\":")
		out.String(string(in.Status))
	}
	if in.Title != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"title\":")
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdMedals200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson303db286EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdMedals200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson303db286EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdMedals200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson303db286DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdMedals200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson303db286DecodeGithubComCurzonjGoesiEsi1(l, v)
}
func easyjson303db286DecodeGithubComCurzonjGoesiEsi2(in *jlexer.Lexer, out *GetCharactersCharacterIdMedalsGraphic) {
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
		case "color":
			out.Color = int32(in.Int32())
		case "graphic":
			out.Graphic = string(in.String())
		case "layer":
			out.Layer = int32(in.Int32())
		case "part":
			out.Part = int32(in.Int32())
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
func easyjson303db286EncodeGithubComCurzonjGoesiEsi2(out *jwriter.Writer, in GetCharactersCharacterIdMedalsGraphic) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Color != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"color\":")
		out.Int32(int32(in.Color))
	}
	if in.Graphic != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"graphic\":")
		out.String(string(in.Graphic))
	}
	if in.Layer != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"layer\":")
		out.Int32(int32(in.Layer))
	}
	if in.Part != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"part\":")
		out.Int32(int32(in.Part))
	}
	out.RawByte('}')
}
