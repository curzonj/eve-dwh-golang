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

func easyjson8920fc7eDecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdAttributesOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdAttributesOkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdAttributesOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdAttributesOk
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
func easyjson8920fc7eEncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdAttributesOkList) {
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
func (v GetCharactersCharacterIdAttributesOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8920fc7eEncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdAttributesOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8920fc7eEncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdAttributesOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8920fc7eDecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdAttributesOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8920fc7eDecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson8920fc7eDecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdAttributesOk) {
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
		case "accrued_remap_cooldown_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.AccruedRemapCooldownDate).UnmarshalJSON(data))
			}
		case "bonus_remaps":
			out.BonusRemaps = int32(in.Int32())
		case "charisma":
			out.Charisma = int32(in.Int32())
		case "intelligence":
			out.Intelligence = int32(in.Int32())
		case "last_remap_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastRemapDate).UnmarshalJSON(data))
			}
		case "memory":
			out.Memory = int32(in.Int32())
		case "perception":
			out.Perception = int32(in.Int32())
		case "willpower":
			out.Willpower = int32(in.Int32())
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
func easyjson8920fc7eEncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdAttributesOk) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accrued_remap_cooldown_date\":")
		out.Raw((in.AccruedRemapCooldownDate).MarshalJSON())
	}
	if in.BonusRemaps != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"bonus_remaps\":")
		out.Int32(int32(in.BonusRemaps))
	}
	if in.Charisma != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"charisma\":")
		out.Int32(int32(in.Charisma))
	}
	if in.Intelligence != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"intelligence\":")
		out.Int32(int32(in.Intelligence))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"last_remap_date\":")
		out.Raw((in.LastRemapDate).MarshalJSON())
	}
	if in.Memory != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"memory\":")
		out.Int32(int32(in.Memory))
	}
	if in.Perception != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"perception\":")
		out.Int32(int32(in.Perception))
	}
	if in.Willpower != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"willpower\":")
		out.Int32(int32(in.Willpower))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdAttributesOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8920fc7eEncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdAttributesOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8920fc7eEncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdAttributesOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8920fc7eDecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdAttributesOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8920fc7eDecodeGithubComCurzonjGoesiEsi1(l, v)
}
