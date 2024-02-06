// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package klp

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

type MessageType int64
const (
  MessageType_kReqPmangBusinessContents MessageType = 1
)

func (p MessageType) String() string {
  switch p {
  case MessageType_kReqPmangBusinessContents: return "kReqPmangBusinessContents"
  }
  return "<UNSET>"
}

func MessageTypeFromString(s string) (MessageType, error) {
  switch s {
  case "kReqPmangBusinessContents": return MessageType_kReqPmangBusinessContents, nil 
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
//  - URL
//  - BusinessContentsMessage
//  - Compress
type ReqPmangBusinessContents struct {
  RequestKey ClientRequestKey `thrift:"request_key,1,required" db:"request_key" json:"request_key"`
  URL string `thrift:"url,2,required" db:"url" json:"url"`
  BusinessContentsMessage Buffer `thrift:"business_contents_message,3,required" db:"business_contents_message" json:"business_contents_message"`
  Compress *bool `thrift:"compress,4" db:"compress" json:"compress,omitempty"`
}

func NewReqPmangBusinessContents() *ReqPmangBusinessContents {
  return &ReqPmangBusinessContents{}
}


func (p *ReqPmangBusinessContents) GetRequestKey() ClientRequestKey {
  return p.RequestKey
}

func (p *ReqPmangBusinessContents) GetURL() string {
  return p.URL
}

func (p *ReqPmangBusinessContents) GetBusinessContentsMessage() Buffer {
  return p.BusinessContentsMessage
}
var ReqPmangBusinessContents_Compress_DEFAULT bool
func (p *ReqPmangBusinessContents) GetCompress() bool {
  if !p.IsSetCompress() {
    return ReqPmangBusinessContents_Compress_DEFAULT
  }
return *p.Compress
}
func (p *ReqPmangBusinessContents) IsSetCompress() bool {
  return p.Compress != nil
}

func (p *ReqPmangBusinessContents) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetRequestKey bool = false;
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
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetURL = true
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
        issetBusinessContentsMessage = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.BOOL {
        if err := p.ReadField4(ctx, iprot); err != nil {
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
  if !issetURL{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field URL is not set"));
  }
  if !issetBusinessContentsMessage{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field BusinessContentsMessage is not set"));
  }
  return nil
}

func (p *ReqPmangBusinessContents)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := ClientRequestKey(v)
  p.RequestKey = temp
}
  return nil
}

func (p *ReqPmangBusinessContents)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.URL = v
}
  return nil
}

func (p *ReqPmangBusinessContents)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  temp := Buffer(v)
  p.BusinessContentsMessage = temp
}
  return nil
}

func (p *ReqPmangBusinessContents)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(ctx); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Compress = &v
}
  return nil
}

func (p *ReqPmangBusinessContents) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "ReqPmangBusinessContents"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
    if err := p.writeField4(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ReqPmangBusinessContents) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "request_key", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request_key: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.RequestKey)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.request_key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request_key: ", p), err) }
  return err
}

func (p *ReqPmangBusinessContents) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "url", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:url: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.URL)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.url (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:url: ", p), err) }
  return err
}

func (p *ReqPmangBusinessContents) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "business_contents_message", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:business_contents_message: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.BusinessContentsMessage); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.business_contents_message (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:business_contents_message: ", p), err) }
  return err
}

func (p *ReqPmangBusinessContents) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetCompress() {
    if err := oprot.WriteFieldBegin(ctx, "compress", thrift.BOOL, 4); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:compress: ", p), err) }
    if err := oprot.WriteBool(ctx, bool(*p.Compress)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.compress (4) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 4:compress: ", p), err) }
  }
  return err
}

func (p *ReqPmangBusinessContents) Equals(other *ReqPmangBusinessContents) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.RequestKey != other.RequestKey { return false }
  if p.URL != other.URL { return false }
  if bytes.Compare(p.BusinessContentsMessage, other.BusinessContentsMessage) != 0 { return false }
  if p.Compress != other.Compress {
    if p.Compress == nil || other.Compress == nil {
      return false
    }
    if (*p.Compress) != (*other.Compress) { return false }
  }
  return true
}

func (p *ReqPmangBusinessContents) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ReqPmangBusinessContents(%+v)", *p)
}

func (p *ReqPmangBusinessContents) Validate() error {
  return nil
}