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

func easyjson95a81519DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetSovereigntyStructures200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetSovereigntyStructures200OkList, 0, 1)
			} else {
				*out = GetSovereigntyStructures200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetSovereigntyStructures200Ok
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
func easyjson95a81519EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetSovereigntyStructures200OkList) {
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
func (v GetSovereigntyStructures200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson95a81519EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetSovereigntyStructures200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson95a81519EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetSovereigntyStructures200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson95a81519DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetSovereigntyStructures200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson95a81519DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson95a81519DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetSovereigntyStructures200Ok) {
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
		case "alliance_id":
			out.AllianceId = int32(in.Int32())
		case "solar_system_id":
			out.SolarSystemId = int32(in.Int32())
		case "structure_id":
			out.StructureId = int64(in.Int64())
		case "structure_type_id":
			out.StructureTypeId = int32(in.Int32())
		case "vulnerability_occupancy_level":
			out.VulnerabilityOccupancyLevel = float32(in.Float32())
		case "vulnerable_end_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.VulnerableEndTime).UnmarshalJSON(data))
			}
		case "vulnerable_start_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.VulnerableStartTime).UnmarshalJSON(data))
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
func easyjson95a81519EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetSovereigntyStructures200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AllianceId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"alliance_id\":")
		out.Int32(int32(in.AllianceId))
	}
	if in.SolarSystemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"solar_system_id\":")
		out.Int32(int32(in.SolarSystemId))
	}
	if in.StructureId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"structure_id\":")
		out.Int64(int64(in.StructureId))
	}
	if in.StructureTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"structure_type_id\":")
		out.Int32(int32(in.StructureTypeId))
	}
	if in.VulnerabilityOccupancyLevel != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"vulnerability_occupancy_level\":")
		out.Float32(float32(in.VulnerabilityOccupancyLevel))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"vulnerable_end_time\":")
		out.Raw((in.VulnerableEndTime).MarshalJSON())
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"vulnerable_start_time\":")
		out.Raw((in.VulnerableStartTime).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetSovereigntyStructures200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson95a81519EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetSovereigntyStructures200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson95a81519EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetSovereigntyStructures200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson95a81519DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetSovereigntyStructures200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson95a81519DecodeGithubComCurzonjGoesiEsi1(l, v)
}
