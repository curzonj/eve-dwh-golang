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

func easyjson87807fe2DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetUniverseSchematicsSchematicIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseSchematicsSchematicIdOkList, 0, 2)
			} else {
				*out = GetUniverseSchematicsSchematicIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseSchematicsSchematicIdOk
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
func easyjson87807fe2EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetUniverseSchematicsSchematicIdOkList) {
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
func (v GetUniverseSchematicsSchematicIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87807fe2EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseSchematicsSchematicIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87807fe2EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseSchematicsSchematicIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87807fe2DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseSchematicsSchematicIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87807fe2DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson87807fe2DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetUniverseSchematicsSchematicIdOk) {
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
		case "cycle_time":
			out.CycleTime = int32(in.Int32())
		case "schematic_name":
			out.SchematicName = string(in.String())
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
func easyjson87807fe2EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetUniverseSchematicsSchematicIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CycleTime != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cycle_time\":")
		out.Int32(int32(in.CycleTime))
	}
	if in.SchematicName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"schematic_name\":")
		out.String(string(in.SchematicName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseSchematicsSchematicIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87807fe2EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseSchematicsSchematicIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87807fe2EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseSchematicsSchematicIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87807fe2DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseSchematicsSchematicIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87807fe2DecodeGithubComCurzonjGoesiEsi1(l, v)
}
