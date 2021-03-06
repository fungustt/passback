// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjsonC80ae7adDecodePassbackCmdPassHttpModel(in *jlexer.Lexer, out *GeneratePasswordResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "password":
			out.Password = string(in.String())
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
func easyjsonC80ae7adEncodePassbackCmdPassHttpModel(out *jwriter.Writer, in GeneratePasswordResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix[1:])
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GeneratePasswordResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodePassbackCmdPassHttpModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GeneratePasswordResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodePassbackCmdPassHttpModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GeneratePasswordResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodePassbackCmdPassHttpModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GeneratePasswordResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodePassbackCmdPassHttpModel(l, v)
}
func easyjsonC80ae7adDecodePassbackCmdPassHttpModel1(in *jlexer.Lexer, out *GeneratePasswordRequest) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "length":
			out.Length = int(in.Int())
		case "uppercase":
			out.SymbolsUppercase = bool(in.Bool())
		case "numbers":
			out.Numbers = bool(in.Bool())
		case "special_symbols":
			out.SpecialSymbols = bool(in.Bool())
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
func easyjsonC80ae7adEncodePassbackCmdPassHttpModel1(out *jwriter.Writer, in GeneratePasswordRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"length\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Length))
	}
	{
		const prefix string = ",\"uppercase\":"
		out.RawString(prefix)
		out.Bool(bool(in.SymbolsUppercase))
	}
	{
		const prefix string = ",\"numbers\":"
		out.RawString(prefix)
		out.Bool(bool(in.Numbers))
	}
	{
		const prefix string = ",\"special_symbols\":"
		out.RawString(prefix)
		out.Bool(bool(in.SpecialSymbols))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GeneratePasswordRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodePassbackCmdPassHttpModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GeneratePasswordRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodePassbackCmdPassHttpModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GeneratePasswordRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodePassbackCmdPassHttpModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GeneratePasswordRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodePassbackCmdPassHttpModel1(l, v)
}
