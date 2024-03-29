// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package user_gateway_to_friends

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
	"github.com/hhcneo/klpProtocol/nwm/klp/protocol"

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
type MessageType int64
const (
  MessageType_kForwardMessage MessageType = 1
)

func (p MessageType) String() string {
  switch p {
  case MessageType_kForwardMessage: return "kForwardMessage"
  }
  return "<UNSET>"
}

func MessageTypeFromString(s string) (MessageType, error) {
  switch s {
  case "kForwardMessage": return MessageType_kForwardMessage, nil 
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
//  - Ssn
//  - Gsn
//  - EnvelopedMessage
//  - Exclusive
type ForwardMessage struct {
  Ssn protocol.Ssn `thrift:"ssn,1,required" db:"ssn" json:"ssn"`
  Gsn protocol.Gsn `thrift:"gsn,2,required" db:"gsn" json:"gsn"`
  EnvelopedMessage protocol.KlpProtocolBuffer `thrift:"enveloped_message,3,required" db:"enveloped_message" json:"enveloped_message"`
  Exclusive int32 `thrift:"exclusive,4,required" db:"exclusive" json:"exclusive"`
}

func NewForwardMessage() *ForwardMessage {
  return &ForwardMessage{}
}


func (p *ForwardMessage) GetSsn() protocol.Ssn {
  return p.Ssn
}

func (p *ForwardMessage) GetGsn() protocol.Gsn {
  return p.Gsn
}

func (p *ForwardMessage) GetEnvelopedMessage() protocol.KlpProtocolBuffer {
  return p.EnvelopedMessage
}

func (p *ForwardMessage) GetExclusive() int32 {
  return p.Exclusive
}
func (p *ForwardMessage) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetSsn bool = false;
  var issetGsn bool = false;
  var issetEnvelopedMessage bool = false;
  var issetExclusive bool = false;

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
        issetSsn = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetGsn = true
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
        issetEnvelopedMessage = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField4(ctx, iprot); err != nil {
          return err
        }
        issetExclusive = true
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
  if !issetSsn{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Ssn is not set"));
  }
  if !issetGsn{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Gsn is not set"));
  }
  if !issetEnvelopedMessage{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field EnvelopedMessage is not set"));
  }
  if !issetExclusive{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Exclusive is not set"));
  }
  return nil
}

func (p *ForwardMessage)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := protocol.Ssn(v)
  p.Ssn = temp
}
  return nil
}

func (p *ForwardMessage)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := protocol.Gsn(v)
  p.Gsn = temp
}
  return nil
}

func (p *ForwardMessage)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  temp := protocol.KlpProtocolBuffer(v)
  p.EnvelopedMessage = temp
}
  return nil
}

func (p *ForwardMessage)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Exclusive = v
}
  return nil
}

func (p *ForwardMessage) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "ForwardMessage"); err != nil {
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

func (p *ForwardMessage) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "ssn", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ssn: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Ssn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.ssn (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ssn: ", p), err) }
  return err
}

func (p *ForwardMessage) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "gsn", thrift.I64, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:gsn: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.Gsn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.gsn (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:gsn: ", p), err) }
  return err
}

func (p *ForwardMessage) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "enveloped_message", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:enveloped_message: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.EnvelopedMessage); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.enveloped_message (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:enveloped_message: ", p), err) }
  return err
}

func (p *ForwardMessage) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "exclusive", thrift.I32, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:exclusive: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Exclusive)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.exclusive (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:exclusive: ", p), err) }
  return err
}

func (p *ForwardMessage) Equals(other *ForwardMessage) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Ssn != other.Ssn { return false }
  if p.Gsn != other.Gsn { return false }
  if bytes.Compare(p.EnvelopedMessage, other.EnvelopedMessage) != 0 { return false }
  if p.Exclusive != other.Exclusive { return false }
  return true
}

func (p *ForwardMessage) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ForwardMessage(%+v)", *p)
}

func (p *ForwardMessage) Validate() error {
  return nil
}
