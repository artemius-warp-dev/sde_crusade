// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package event

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

func easyjsonF642ad3eDecodeCalendarServiceDomainValueObjectsEvent(in *jlexer.Lexer, out *Event) {
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
		case "attendees":
			if in.IsNull() {
				in.Skip()
				out.Attendees = nil
			} else {
				in.Delim('[')
				if out.Attendees == nil {
					if !in.IsDelim(']') {
						out.Attendees = make([]string, 0, 4)
					} else {
						out.Attendees = []string{}
					}
				} else {
					out.Attendees = (out.Attendees)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Attendees = append(out.Attendees, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "description":
			out.Description = string(in.String())
		case "end_time":
			out.EndTime = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "location":
			out.Location = string(in.String())
		case "start_time":
			out.StartTime = string(in.String())
		case "title":
			out.Title = string(in.String())
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
func easyjsonF642ad3eEncodeCalendarServiceDomainValueObjectsEvent(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"attendees\":"
		out.RawString(prefix[1:])
		if in.Attendees == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Attendees {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"end_time\":"
		out.RawString(prefix)
		out.String(string(in.EndTime))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"location\":"
		out.RawString(prefix)
		out.String(string(in.Location))
	}
	{
		const prefix string = ",\"start_time\":"
		out.RawString(prefix)
		out.String(string(in.StartTime))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeCalendarServiceDomainValueObjectsEvent(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeCalendarServiceDomainValueObjectsEvent(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeCalendarServiceDomainValueObjectsEvent(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeCalendarServiceDomainValueObjectsEvent(l, v)
}