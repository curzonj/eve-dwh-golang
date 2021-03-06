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

func easyjson67f82948DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetAlliancesAllianceIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetAlliancesAllianceIdOkList, 0, 1)
			} else {
				*out = GetAlliancesAllianceIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetAlliancesAllianceIdOk
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
func easyjson67f82948EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetAlliancesAllianceIdOkList) {
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
func (v GetAlliancesAllianceIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson67f82948EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAlliancesAllianceIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson67f82948EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAlliancesAllianceIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson67f82948DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAlliancesAllianceIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson67f82948DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson67f82948DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetAlliancesAllianceIdOk) {
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
		case "alliance_name":
			out.AllianceName = string(in.String())
		case "date_founded":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateFounded).UnmarshalJSON(data))
			}
		case "executor_corp":
			out.ExecutorCorp = int32(in.Int32())
		case "ticker":
			out.Ticker = string(in.String())
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
func easyjson67f82948EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetAlliancesAllianceIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AllianceName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"alliance_name\":")
		out.String(string(in.AllianceName))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"date_founded\":")
		out.Raw((in.DateFounded).MarshalJSON())
	}
	if in.ExecutorCorp != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"executor_corp\":")
		out.Int32(int32(in.ExecutorCorp))
	}
	if in.Ticker != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ticker\":")
		out.String(string(in.Ticker))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetAlliancesAllianceIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson67f82948EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAlliancesAllianceIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson67f82948EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAlliancesAllianceIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson67f82948DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAlliancesAllianceIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson67f82948DecodeGithubComCurzonjGoesiEsi1(l, v)
}
