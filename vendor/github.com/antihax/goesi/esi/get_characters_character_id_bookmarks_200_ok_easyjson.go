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

func easyjsonF26fb59dDecodeGithubComCurzonjGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdBookmarks200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdBookmarks200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdBookmarks200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdBookmarks200Ok
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
func easyjsonF26fb59dEncodeGithubComCurzonjGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdBookmarks200OkList) {
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
func (v GetCharactersCharacterIdBookmarks200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF26fb59dEncodeGithubComCurzonjGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdBookmarks200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF26fb59dEncodeGithubComCurzonjGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdBookmarks200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF26fb59dDecodeGithubComCurzonjGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdBookmarks200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF26fb59dDecodeGithubComCurzonjGoesiEsi(l, v)
}
func easyjsonF26fb59dDecodeGithubComCurzonjGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdBookmarks200Ok) {
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
		case "bookmark_id":
			out.BookmarkId = int64(in.Int64())
		case "create_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreateDate).UnmarshalJSON(data))
			}
		case "creator_id":
			out.CreatorId = int32(in.Int32())
		case "folder_id":
			out.FolderId = int32(in.Int32())
		case "memo":
			out.Memo = string(in.String())
		case "note":
			out.Note = string(in.String())
		case "owner_id":
			out.OwnerId = int32(in.Int32())
		case "target":
			easyjsonF26fb59dDecodeGithubComCurzonjGoesiEsi2(in, &out.Target)
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
func easyjsonF26fb59dEncodeGithubComCurzonjGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdBookmarks200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.BookmarkId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"bookmark_id\":")
		out.Int64(int64(in.BookmarkId))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"create_date\":")
		out.Raw((in.CreateDate).MarshalJSON())
	}
	if in.CreatorId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"creator_id\":")
		out.Int32(int32(in.CreatorId))
	}
	if in.FolderId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"folder_id\":")
		out.Int32(int32(in.FolderId))
	}
	if in.Memo != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"memo\":")
		out.String(string(in.Memo))
	}
	if in.Note != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"note\":")
		out.String(string(in.Note))
	}
	if in.OwnerId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"owner_id\":")
		out.Int32(int32(in.OwnerId))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"target\":")
		easyjsonF26fb59dEncodeGithubComCurzonjGoesiEsi2(out, in.Target)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdBookmarks200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF26fb59dEncodeGithubComCurzonjGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdBookmarks200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF26fb59dEncodeGithubComCurzonjGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdBookmarks200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF26fb59dDecodeGithubComCurzonjGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdBookmarks200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF26fb59dDecodeGithubComCurzonjGoesiEsi1(l, v)
}
func easyjsonF26fb59dDecodeGithubComCurzonjGoesiEsi2(in *jlexer.Lexer, out *GetCharactersCharacterIdBookmarksTarget) {
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
		case "coordinates":
			easyjsonF26fb59dDecodeGithubComCurzonjGoesiEsi3(in, &out.Coordinates)
		case "item":
			easyjsonF26fb59dDecodeGithubComCurzonjGoesiEsi4(in, &out.Item)
		case "location_id":
			out.LocationId = int64(in.Int64())
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
func easyjsonF26fb59dEncodeGithubComCurzonjGoesiEsi2(out *jwriter.Writer, in GetCharactersCharacterIdBookmarksTarget) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"coordinates\":")
		easyjsonF26fb59dEncodeGithubComCurzonjGoesiEsi3(out, in.Coordinates)
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"item\":")
		easyjsonF26fb59dEncodeGithubComCurzonjGoesiEsi4(out, in.Item)
	}
	if in.LocationId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"location_id\":")
		out.Int64(int64(in.LocationId))
	}
	out.RawByte('}')
}
func easyjsonF26fb59dDecodeGithubComCurzonjGoesiEsi4(in *jlexer.Lexer, out *GetCharactersCharacterIdBookmarksItem) {
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
		case "item_id":
			out.ItemId = int64(in.Int64())
		case "type_id":
			out.TypeId = int32(in.Int32())
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
func easyjsonF26fb59dEncodeGithubComCurzonjGoesiEsi4(out *jwriter.Writer, in GetCharactersCharacterIdBookmarksItem) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ItemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"item_id\":")
		out.Int64(int64(in.ItemId))
	}
	if in.TypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"type_id\":")
		out.Int32(int32(in.TypeId))
	}
	out.RawByte('}')
}
func easyjsonF26fb59dDecodeGithubComCurzonjGoesiEsi3(in *jlexer.Lexer, out *GetCharactersCharacterIdBookmarksCoordinates) {
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
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		case "z":
			out.Z = float64(in.Float64())
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
func easyjsonF26fb59dEncodeGithubComCurzonjGoesiEsi3(out *jwriter.Writer, in GetCharactersCharacterIdBookmarksCoordinates) {
	out.RawByte('{')
	first := true
	_ = first
	if in.X != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"x\":")
		out.Float64(float64(in.X))
	}
	if in.Y != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"y\":")
		out.Float64(float64(in.Y))
	}
	if in.Z != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"z\":")
		out.Float64(float64(in.Z))
	}
	out.RawByte('}')
}
