// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package web_gateway_to_client

import (
	"bytes"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"strings"
	"regexp"
	"nwm/klp/protocol"
	"nwm/klp/protocol/error"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal
// (needed by validator.)
var _ = strings.Contains
var _ = regexp.MatchString

var _ = protocol.GoUnusedProtection__
var _ = error.GoUnusedProtection__
type MessageType int64
const (
  MessageType_kAnsPmangBusinessContents MessageType = 1
  MessageType_kNtfPmangBusinessContents MessageType = 2
)

func (p MessageType) String() string {
  switch p {
  case MessageType_kAnsPmangBusinessContents: return "kAnsPmangBusinessContents"
  case MessageType_kNtfPmangBusinessContents: return "kNtfPmangBusinessContents"
  }
  return "<UNSET>"
}

func MessageTypeFromString(s string) (MessageType, error) {
  switch s {
  case "kAnsPmangBusinessContents": return MessageType_kAnsPmangBusinessContents, nil 
  case "kNtfPmangBusinessContents": return MessageType_kNtfPmangBusinessContents, nil 
  }
  return MessageType(0), fmt.Errorf("not a valid MessageType string")
}


func MessageTypePtr(v MessageType) *MessageType { return &v }

func (p MessageType) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *MessageType) UnmarshalText(text []byte) error {
q, err := MessageTypeFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *MessageType) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = MessageType(v)
return nil
}

func (p * MessageType) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
// Attributes:
//  - RequestKey
//  - Result_
//  - URL
//  - BusinessContentsMessage
//  - Compressed
type AnsPmangBusinessContents struct {
  RequestKey protocol.ClientRequestKey `thrift:"request_key,1,required" db:"request_key" json:"request_key"`
  Result_ *error.Error `thrift:"result,2,required" db:"result" json:"result"`
  URL string `thrift:"url,3,required" db:"url" json:"url"`
  BusinessContentsMessage protocol.Buffer `thrift:"business_contents_message,4,required" db:"business_contents_message" json:"business_contents_message"`
  Compressed *bool `thrift:"compressed,5" db:"compressed" json:"compressed,omitempty"`
}

func NewAnsPmangBusinessContents() *AnsPmangBusinessContents {
  return &AnsPmangBusinessContents{}
}


func (p *AnsPmangBusinessContents) GetRequestKey() protocol.ClientRequestKey {
  return p.RequestKey
}
var AnsPmangBusinessContents_Result__DEFAULT *error.Error
func (p *AnsPmangBusinessContents) GetResult_() *error.Error {
  if !p.IsSetResult_() {
    return AnsPmangBusinessContents_Result__DEFAULT
  }
return p.Result_
}

func (p *AnsPmangBusinessContents) GetURL() string {
  return p.URL
}

func (p *AnsPmangBusinessContents) GetBusinessContentsMessage() protocol.Buffer {
  return p.BusinessContentsMessage
}
var AnsPmangBusinessContents_Compressed_DEFAULT bool
func (p *AnsPmangBusinessContents) GetCompressed() bool {
  if !p.IsSetCompressed() {
    return AnsPmangBusinessContents_Compressed_DEFAULT
  }
return *p.Compressed
}
func (p *AnsPmangBusinessContents) IsSetResult_() bool {
  return p.Result_ != nil
}

func (p *AnsPmangBusinessContents) IsSetCompressed() bool {
  return p.Compressed != nil
}

func (p *AnsPmangBusinessContents) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetRequestKey bool = false;
  var issetResult_ bool = false;
  var issetURL bool = false;
  var issetBusinessContentsMessage bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetRequestKey = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetResult_ = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetURL = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField4(ctx, iprot); err != nil {
          return err
        }
        issetBusinessContentsMessage = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.BOOL {
        if err := p.ReadField5(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetRequestKey{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field RequestKey is not set"));
  }
  if !issetResult_{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Result_ is not set"));
  }
  if !issetURL{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field URL is not set"));
  }
  if !issetBusinessContentsMessage{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field BusinessContentsMessage is not set"));
  }
  return nil
}

func (p *AnsPmangBusinessContents)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := protocol.ClientRequestKey(v)
  p.RequestKey = temp
}
  return nil
}

func (p *AnsPmangBusinessContents)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  p.Result_ = &error.Error{}
  if err := p.Result_.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Result_), err)
  }
  return nil
}

func (p *AnsPmangBusinessContents)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.URL = v
}
  return nil
}

func (p *AnsPmangBusinessContents)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  temp := protocol.Buffer(v)
  p.BusinessContentsMessage = temp
}
  return nil
}

func (p *AnsPmangBusinessContents)  ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(ctx); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.Compressed = &v
}
  return nil
}

func (p *AnsPmangBusinessContents) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "AnsPmangBusinessContents"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
    if err := p.writeField4(ctx, oprot); err != nil { return err }
    if err := p.writeField5(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *AnsPmangBusinessContents) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "request_key", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request_key: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.RequestKey)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.request_key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request_key: ", p), err) }
  return err
}

func (p *AnsPmangBusinessContents) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "result", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:result: ", p), err) }
  if err := p.Result_.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Result_), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:result: ", p), err) }
  return err
}

func (p *AnsPmangBusinessContents) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "url", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:url: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.URL)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.url (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:url: ", p), err) }
  return err
}

func (p *AnsPmangBusinessContents) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "business_contents_message", thrift.STRING, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:business_contents_message: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.BusinessContentsMessage); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.business_contents_message (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:business_contents_message: ", p), err) }
  return err
}

func (p *AnsPmangBusinessContents) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetCompressed() {
    if err := oprot.WriteFieldBegin(ctx, "compressed", thrift.BOOL, 5); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:compressed: ", p), err) }
    if err := oprot.WriteBool(ctx, bool(*p.Compressed)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.compressed (5) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 5:compressed: ", p), err) }
  }
  return err
}

func (p *AnsPmangBusinessContents) Equals(other *AnsPmangBusinessContents) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.RequestKey != other.RequestKey { return false }
  if !p.Result_.Equals(other.Result_) { return false }
  if p.URL != other.URL { return false }
  if bytes.Compare(p.BusinessContentsMessage, other.BusinessContentsMessage) != 0 { return false }
  if p.Compressed != other.Compressed {
    if p.Compressed == nil || other.Compressed == nil {
      return false
    }
    if (*p.Compressed) != (*other.Compressed) { return false }
  }
  return true
}

func (p *AnsPmangBusinessContents) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AnsPmangBusinessContents(%+v)", *p)
}

func (p *AnsPmangBusinessContents) Validate() error {
  return nil
}
// Attributes:
//  - BusinessContentsMessage
type NtfPmangBusinessContents struct {
  BusinessContentsMessage protocol.Buffer `thrift:"business_contents_message,1,required" db:"business_contents_message" json:"business_contents_message"`
}

func NewNtfPmangBusinessContents() *NtfPmangBusinessContents {
  return &NtfPmangBusinessContents{}
}


func (p *NtfPmangBusinessContents) GetBusinessContentsMessage() protocol.Buffer {
  return p.BusinessContentsMessage
}
func (p *NtfPmangBusinessContents) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetBusinessContentsMessage bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetBusinessContentsMessage = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetBusinessContentsMessage{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field BusinessContentsMessage is not set"));
  }
  return nil
}

func (p *NtfPmangBusinessContents)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := protocol.Buffer(v)
  p.BusinessContentsMessage = temp
}
  return nil
}

func (p *NtfPmangBusinessContents) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "NtfPmangBusinessContents"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *NtfPmangBusinessContents) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "business_contents_message", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:business_contents_message: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.BusinessContentsMessage); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.business_contents_message (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:business_contents_message: ", p), err) }
  return err
}

func (p *NtfPmangBusinessContents) Equals(other *NtfPmangBusinessContents) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if bytes.Compare(p.BusinessContentsMessage, other.BusinessContentsMessage) != 0 { return false }
  return true
}

func (p *NtfPmangBusinessContents) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("NtfPmangBusinessContents(%+v)", *p)
}

func (p *NtfPmangBusinessContents) Validate() error {
  return nil
}
