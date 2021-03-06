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

func easyjson191b3d72DecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetOpportunitiesTasksTaskIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetOpportunitiesTasksTaskIdOkList, 0, 1)
			} else {
				*out = GetOpportunitiesTasksTaskIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetOpportunitiesTasksTaskIdOk
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
func easyjson191b3d72EncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetOpportunitiesTasksTaskIdOkList) {
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
func (v GetOpportunitiesTasksTaskIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson191b3d72EncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetOpportunitiesTasksTaskIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson191b3d72EncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetOpportunitiesTasksTaskIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson191b3d72DecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetOpportunitiesTasksTaskIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson191b3d72DecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson191b3d72DecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetOpportunitiesTasksTaskIdOk) {
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
		case "name":
			out.Name = string(in.String())
		case "notification":
			out.Notification = string(in.String())
		case "task_id":
			out.TaskId = int32(in.Int32())
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
func easyjson191b3d72EncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetOpportunitiesTasksTaskIdOk) {
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
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.Notification != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"notification\":")
		out.String(string(in.Notification))
	}
	if in.TaskId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"task_id\":")
		out.Int32(int32(in.TaskId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetOpportunitiesTasksTaskIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson191b3d72EncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetOpportunitiesTasksTaskIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson191b3d72EncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetOpportunitiesTasksTaskIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson191b3d72DecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetOpportunitiesTasksTaskIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson191b3d72DecodeGithubComCurzonjGoesiEsi1(l, v)
}
