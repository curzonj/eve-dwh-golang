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

func easyjsonF6d56f96DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdStructures200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdStructures200OkList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdStructures200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdStructures200Ok
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
func easyjsonF6d56f96EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdStructures200OkList) {
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
func (v GetCorporationsCorporationIdStructures200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF6d56f96EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdStructures200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF6d56f96EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdStructures200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF6d56f96DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdStructures200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF6d56f96DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjsonF6d56f96DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdStructures200Ok) {
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
		case "current_vul":
			if in.IsNull() {
				in.Skip()
				out.CurrentVul = nil
			} else {
				in.Delim('[')
				if out.CurrentVul == nil {
					if !in.IsDelim(']') {
						out.CurrentVul = make([]GetCorporationsCorporationIdStructuresCurrentVul, 0, 8)
					} else {
						out.CurrentVul = []GetCorporationsCorporationIdStructuresCurrentVul{}
					}
				} else {
					out.CurrentVul = (out.CurrentVul)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCorporationsCorporationIdStructuresCurrentVul
					easyjsonF6d56f96DecodeGithubComCurzonjGoesiEsi2(in, &v4)
					out.CurrentVul = append(out.CurrentVul, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "fuel_expires":
			out.FuelExpires = string(in.String())
		case "next_vul":
			if in.IsNull() {
				in.Skip()
				out.NextVul = nil
			} else {
				in.Delim('[')
				if out.NextVul == nil {
					if !in.IsDelim(']') {
						out.NextVul = make([]GetCorporationsCorporationIdStructuresNextVul, 0, 8)
					} else {
						out.NextVul = []GetCorporationsCorporationIdStructuresNextVul{}
					}
				} else {
					out.NextVul = (out.NextVul)[:0]
				}
				for !in.IsDelim(']') {
					var v5 GetCorporationsCorporationIdStructuresNextVul
					easyjsonF6d56f96DecodeGithubComCurzonjGoesiEsi3(in, &v5)
					out.NextVul = append(out.NextVul, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "profile_id":
			out.ProfileId = int32(in.Int32())
		case "services":
			if in.IsNull() {
				in.Skip()
				out.Services = nil
			} else {
				in.Delim('[')
				if out.Services == nil {
					if !in.IsDelim(']') {
						out.Services = make([]GetCorporationsCorporationIdStructuresService, 0, 2)
					} else {
						out.Services = []GetCorporationsCorporationIdStructuresService{}
					}
				} else {
					out.Services = (out.Services)[:0]
				}
				for !in.IsDelim(']') {
					var v6 GetCorporationsCorporationIdStructuresService
					easyjsonF6d56f96DecodeGithubComCurzonjGoesiEsi4(in, &v6)
					out.Services = append(out.Services, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "state_timer_end":
			out.StateTimerEnd = string(in.String())
		case "state_timer_start":
			out.StateTimerStart = string(in.String())
		case "structure_id":
			out.StructureId = int64(in.Int64())
		case "system_id":
			out.SystemId = int32(in.Int32())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "unanchors_at":
			out.UnanchorsAt = string(in.String())
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
func easyjsonF6d56f96EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdStructures200Ok) {
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
	if len(in.CurrentVul) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"current_vul\":")
		if in.CurrentVul == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.CurrentVul {
				if v7 > 0 {
					out.RawByte(',')
				}
				easyjsonF6d56f96EncodeGithubComCurzonjGoesiEsi2(out, v8)
			}
			out.RawByte(']')
		}
	}
	if in.FuelExpires != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fuel_expires\":")
		out.String(string(in.FuelExpires))
	}
	if len(in.NextVul) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"next_vul\":")
		if in.NextVul == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.NextVul {
				if v9 > 0 {
					out.RawByte(',')
				}
				easyjsonF6d56f96EncodeGithubComCurzonjGoesiEsi3(out, v10)
			}
			out.RawByte(']')
		}
	}
	if in.ProfileId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"profile_id\":")
		out.Int32(int32(in.ProfileId))
	}
	if len(in.Services) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"services\":")
		if in.Services == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Services {
				if v11 > 0 {
					out.RawByte(',')
				}
				easyjsonF6d56f96EncodeGithubComCurzonjGoesiEsi4(out, v12)
			}
			out.RawByte(']')
		}
	}
	if in.StateTimerEnd != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"state_timer_end\":")
		out.String(string(in.StateTimerEnd))
	}
	if in.StateTimerStart != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"state_timer_start\":")
		out.String(string(in.StateTimerStart))
	}
	if in.StructureId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"structure_id\":")
		out.Int64(int64(in.StructureId))
	}
	if in.SystemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"system_id\":")
		out.Int32(int32(in.SystemId))
	}
	if in.TypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"type_id\":")
		out.Int32(int32(in.TypeId))
	}
	if in.UnanchorsAt != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"unanchors_at\":")
		out.String(string(in.UnanchorsAt))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdStructures200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF6d56f96EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdStructures200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF6d56f96EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdStructures200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF6d56f96DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdStructures200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF6d56f96DecodeGithubComCurzonjGoesiEsi1(l, v)
}
func easyjsonF6d56f96DecodeGithubComCurzonjGoesiEsi4(in *jlexer.Lexer, out *GetCorporationsCorporationIdStructuresService) {
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
		case "name":
			out.Name = string(in.String())
		case "state":
			out.State = string(in.String())
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
func easyjsonF6d56f96EncodeGithubComCurzonjGoesiEsi4(out *jwriter.Writer, in GetCorporationsCorporationIdStructuresService) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.State != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"state\":")
		out.String(string(in.State))
	}
	out.RawByte('}')
}
func easyjsonF6d56f96DecodeGithubComCurzonjGoesiEsi3(in *jlexer.Lexer, out *GetCorporationsCorporationIdStructuresNextVul) {
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
		case "day":
			out.Day = int32(in.Int32())
		case "hour":
			out.Hour = int32(in.Int32())
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
func easyjsonF6d56f96EncodeGithubComCurzonjGoesiEsi3(out *jwriter.Writer, in GetCorporationsCorporationIdStructuresNextVul) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Day != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"day\":")
		out.Int32(int32(in.Day))
	}
	if in.Hour != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"hour\":")
		out.Int32(int32(in.Hour))
	}
	out.RawByte('}')
}
func easyjsonF6d56f96DecodeGithubComCurzonjGoesiEsi2(in *jlexer.Lexer, out *GetCorporationsCorporationIdStructuresCurrentVul) {
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
		case "day":
			out.Day = int32(in.Int32())
		case "hour":
			out.Hour = int32(in.Int32())
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
func easyjsonF6d56f96EncodeGithubComCurzonjGoesiEsi2(out *jwriter.Writer, in GetCorporationsCorporationIdStructuresCurrentVul) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Day != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"day\":")
		out.Int32(int32(in.Day))
	}
	if in.Hour != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"hour\":")
		out.Int32(int32(in.Hour))
	}
	out.RawByte('}')
}
