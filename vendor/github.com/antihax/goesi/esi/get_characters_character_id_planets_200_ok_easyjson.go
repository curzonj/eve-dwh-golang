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

func easyjson5f788c63DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanets200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdPlanets200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdPlanets200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdPlanets200Ok
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
func easyjson5f788c63EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdPlanets200OkList) {
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
func (v GetCharactersCharacterIdPlanets200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5f788c63EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanets200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5f788c63EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanets200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5f788c63DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanets200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5f788c63DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson5f788c63DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanets200Ok) {
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
		case "last_update":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastUpdate).UnmarshalJSON(data))
			}
		case "num_pins":
			out.NumPins = int32(in.Int32())
		case "owner_id":
			out.OwnerId = int32(in.Int32())
		case "planet_id":
			out.PlanetId = int32(in.Int32())
		case "planet_type":
			out.PlanetType = string(in.String())
		case "solar_system_id":
			out.SolarSystemId = int32(in.Int32())
		case "upgrade_level":
			out.UpgradeLevel = int32(in.Int32())
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
func easyjson5f788c63EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdPlanets200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"last_update\":")
		out.Raw((in.LastUpdate).MarshalJSON())
	}
	if in.NumPins != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"num_pins\":")
		out.Int32(int32(in.NumPins))
	}
	if in.OwnerId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"owner_id\":")
		out.Int32(int32(in.OwnerId))
	}
	if in.PlanetId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"planet_id\":")
		out.Int32(int32(in.PlanetId))
	}
	if in.PlanetType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"planet_type\":")
		out.String(string(in.PlanetType))
	}
	if in.SolarSystemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"solar_system_id\":")
		out.Int32(int32(in.SolarSystemId))
	}
	if in.UpgradeLevel != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"upgrade_level\":")
		out.Int32(int32(in.UpgradeLevel))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdPlanets200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5f788c63EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanets200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5f788c63EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanets200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5f788c63DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanets200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5f788c63DecodeGithubComAntihaxGoesiEsi1(l, v)
}
