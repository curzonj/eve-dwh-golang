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

func easyjson41562567DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetDogmaEffectsEffectIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetDogmaEffectsEffectIdOkList, 0, 1)
			} else {
				*out = GetDogmaEffectsEffectIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetDogmaEffectsEffectIdOk
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
func easyjson41562567EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetDogmaEffectsEffectIdOkList) {
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
func (v GetDogmaEffectsEffectIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson41562567EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetDogmaEffectsEffectIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson41562567EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetDogmaEffectsEffectIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson41562567DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetDogmaEffectsEffectIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson41562567DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson41562567DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetDogmaEffectsEffectIdOk) {
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
		case "disallow_auto_repeat":
			out.DisallowAutoRepeat = bool(in.Bool())
		case "discharge_attribute_id":
			out.DischargeAttributeId = int32(in.Int32())
		case "display_name":
			out.DisplayName = string(in.String())
		case "duration_attribute_id":
			out.DurationAttributeId = int32(in.Int32())
		case "effect_category":
			out.EffectCategory = int32(in.Int32())
		case "effect_id":
			out.EffectId = int32(in.Int32())
		case "electronic_chance":
			out.ElectronicChance = bool(in.Bool())
		case "falloff_attribute_id":
			out.FalloffAttributeId = int32(in.Int32())
		case "icon_id":
			out.IconId = int32(in.Int32())
		case "is_assistance":
			out.IsAssistance = bool(in.Bool())
		case "is_offensive":
			out.IsOffensive = bool(in.Bool())
		case "is_warp_safe":
			out.IsWarpSafe = bool(in.Bool())
		case "modifiers":
			if in.IsNull() {
				in.Skip()
				out.Modifiers = nil
			} else {
				in.Delim('[')
				if out.Modifiers == nil {
					if !in.IsDelim(']') {
						out.Modifiers = make([]GetDogmaEffectsEffectIdModifier, 0, 1)
					} else {
						out.Modifiers = []GetDogmaEffectsEffectIdModifier{}
					}
				} else {
					out.Modifiers = (out.Modifiers)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetDogmaEffectsEffectIdModifier
					(v4).UnmarshalEasyJSON(in)
					out.Modifiers = append(out.Modifiers, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "name":
			out.Name = string(in.String())
		case "post_expression":
			out.PostExpression = int32(in.Int32())
		case "pre_expression":
			out.PreExpression = int32(in.Int32())
		case "published":
			out.Published = bool(in.Bool())
		case "range_attribute_id":
			out.RangeAttributeId = int32(in.Int32())
		case "range_chance":
			out.RangeChance = bool(in.Bool())
		case "tracking_speed_attribute_id":
			out.TrackingSpeedAttributeId = int32(in.Int32())
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
func easyjson41562567EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetDogmaEffectsEffectIdOk) {
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
	if in.DisallowAutoRepeat {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"disallow_auto_repeat\":")
		out.Bool(bool(in.DisallowAutoRepeat))
	}
	if in.DischargeAttributeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"discharge_attribute_id\":")
		out.Int32(int32(in.DischargeAttributeId))
	}
	if in.DisplayName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"display_name\":")
		out.String(string(in.DisplayName))
	}
	if in.DurationAttributeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"duration_attribute_id\":")
		out.Int32(int32(in.DurationAttributeId))
	}
	if in.EffectCategory != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"effect_category\":")
		out.Int32(int32(in.EffectCategory))
	}
	if in.EffectId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"effect_id\":")
		out.Int32(int32(in.EffectId))
	}
	if in.ElectronicChance {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"electronic_chance\":")
		out.Bool(bool(in.ElectronicChance))
	}
	if in.FalloffAttributeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"falloff_attribute_id\":")
		out.Int32(int32(in.FalloffAttributeId))
	}
	if in.IconId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"icon_id\":")
		out.Int32(int32(in.IconId))
	}
	if in.IsAssistance {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"is_assistance\":")
		out.Bool(bool(in.IsAssistance))
	}
	if in.IsOffensive {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"is_offensive\":")
		out.Bool(bool(in.IsOffensive))
	}
	if in.IsWarpSafe {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"is_warp_safe\":")
		out.Bool(bool(in.IsWarpSafe))
	}
	if len(in.Modifiers) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"modifiers\":")
		if in.Modifiers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Modifiers {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
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
	if in.PostExpression != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"post_expression\":")
		out.Int32(int32(in.PostExpression))
	}
	if in.PreExpression != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"pre_expression\":")
		out.Int32(int32(in.PreExpression))
	}
	if in.Published {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"published\":")
		out.Bool(bool(in.Published))
	}
	if in.RangeAttributeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"range_attribute_id\":")
		out.Int32(int32(in.RangeAttributeId))
	}
	if in.RangeChance {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"range_chance\":")
		out.Bool(bool(in.RangeChance))
	}
	if in.TrackingSpeedAttributeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"tracking_speed_attribute_id\":")
		out.Int32(int32(in.TrackingSpeedAttributeId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetDogmaEffectsEffectIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson41562567EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetDogmaEffectsEffectIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson41562567EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetDogmaEffectsEffectIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson41562567DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetDogmaEffectsEffectIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson41562567DecodeGithubComCurzonjGoesiEsi1(l, v)
}
