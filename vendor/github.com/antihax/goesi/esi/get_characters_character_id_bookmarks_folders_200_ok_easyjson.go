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

func easyjson74168e7fDecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdBookmarksFolders200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdBookmarksFolders200OkList, 0, 2)
			} else {
				*out = GetCharactersCharacterIdBookmarksFolders200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdBookmarksFolders200Ok
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
func easyjson74168e7fEncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdBookmarksFolders200OkList) {
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
func (v GetCharactersCharacterIdBookmarksFolders200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74168e7fEncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdBookmarksFolders200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74168e7fEncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdBookmarksFolders200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74168e7fDecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdBookmarksFolders200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74168e7fDecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjson74168e7fDecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdBookmarksFolders200Ok) {
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
		case "folder_id":
			out.FolderId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "owner_id":
			out.OwnerId = int32(in.Int32())
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
func easyjson74168e7fEncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdBookmarksFolders200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FolderId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"folder_id\":")
		out.Int32(int32(in.FolderId))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.OwnerId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"owner_id\":")
		out.Int32(int32(in.OwnerId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdBookmarksFolders200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74168e7fEncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdBookmarksFolders200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74168e7fEncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdBookmarksFolders200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74168e7fDecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdBookmarksFolders200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74168e7fDecodeGithubComCurzonjGoesiEsi1(l, v)
}
