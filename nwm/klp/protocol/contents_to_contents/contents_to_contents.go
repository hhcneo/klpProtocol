// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package contents_to_contents

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
  MessageType_kForwardToObject MessageType = 1
  MessageType_kForwardToMessagingChannel MessageType = 2
)

func (p MessageType) String() string {
  switch p {
  case MessageType_kForwardToObject: return "kForwardToObject"
  case MessageType_kForwardToMessagingChannel: return "kForwardToMessagingChannel"
  }
  return "<UNSET>"
}

func MessageTypeFromString(s string) (MessageType, error) {
  switch s {
  case "kForwardToObject": return MessageType_kForwardToObject, nil 
  case "kForwardToMessagingChannel": return MessageType_kForwardToMessagingChannel, nil 
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
type ObjectType int64
const (
  ObjectType_kInvalid ObjectType = 0
  ObjectType_kUser ObjectType = 1
  ObjectType_kRoom ObjectType = 2
  ObjectType_kDiy ObjectType = 3
  ObjectType_kMax ObjectType = 4
)

func (p ObjectType) String() string {
  switch p {
  case ObjectType_kInvalid: return "kInvalid"
  case ObjectType_kUser: return "kUser"
  case ObjectType_kRoom: return "kRoom"
  case ObjectType_kDiy: return "kDiy"
  case ObjectType_kMax: return "kMax"
  }
  return "<UNSET>"
}

func ObjectTypeFromString(s string) (ObjectType, error) {
  switch s {
  case "kInvalid": return ObjectType_kInvalid, nil 
  case "kUser": return ObjectType_kUser, nil 
  case "kRoom": return ObjectType_kRoom, nil 
  case "kDiy": return ObjectType_kDiy, nil 
  case "kMax": return ObjectType_kMax, nil 
  }
  return ObjectType(0), fmt.Errorf("not a valid ObjectType string")
}


func ObjectTypePtr(v ObjectType) *ObjectType { return &v }

func (p ObjectType) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *ObjectType) UnmarshalText(text []byte) error {
q, err := ObjectTypeFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *ObjectType) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = ObjectType(v)
return nil
}

func (p * ObjectType) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
type MessagingChannelId string

func MessagingChannelIdPtr(v MessagingChannelId) *MessagingChannelId { return &v }

// Attributes:
//  - Type
//  - Serial
type ObjectId struct {
  Type ObjectType `thrift:"type,1,required" db:"type" json:"type"`
  Serial int64 `thrift:"serial,2,required" db:"serial" json:"serial"`
}

func NewObjectId() *ObjectId {
  return &ObjectId{}
}


func (p *ObjectId) GetType() ObjectType {
  return p.Type
}

func (p *ObjectId) GetSerial() int64 {
  return p.Serial
}
func (p *ObjectId) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetType bool = false;
  var issetSerial bool = false;

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
        issetType = true
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
        issetSerial = true
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
  if !issetType{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Type is not set"));
  }
  if !issetSerial{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Serial is not set"));
  }
  return nil
}

func (p *ObjectId)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := ObjectType(v)
  p.Type = temp
}
  return nil
}

func (p *ObjectId)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Serial = v
}
  return nil
}

func (p *ObjectId) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "ObjectId"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ObjectId) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "type", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:type: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Type)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.type (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:type: ", p), err) }
  return err
}

func (p *ObjectId) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "serial", thrift.I64, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:serial: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.Serial)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.serial (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:serial: ", p), err) }
  return err
}

func (p *ObjectId) Equals(other *ObjectId) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Type != other.Type { return false }
  if p.Serial != other.Serial { return false }
  return true
}

func (p *ObjectId) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ObjectId(%+v)", *p)
}

func (p *ObjectId) Validate() error {
  return nil
}
// Attributes:
//  - Ssn
//  - Source
//  - Destination
//  - Message
type ForwardToObject struct {
  Ssn protocol.Ssn `thrift:"ssn,1,required" db:"ssn" json:"ssn"`
  Source *ObjectId `thrift:"source,2,required" db:"source" json:"source"`
  Destination *ObjectId `thrift:"destination,3,required" db:"destination" json:"destination"`
  Message protocol.Buffer `thrift:"message,4,required" db:"message" json:"message"`
}

func NewForwardToObject() *ForwardToObject {
  return &ForwardToObject{}
}


func (p *ForwardToObject) GetSsn() protocol.Ssn {
  return p.Ssn
}
var ForwardToObject_Source_DEFAULT *ObjectId
func (p *ForwardToObject) GetSource() *ObjectId {
  if !p.IsSetSource() {
    return ForwardToObject_Source_DEFAULT
  }
return p.Source
}
var ForwardToObject_Destination_DEFAULT *ObjectId
func (p *ForwardToObject) GetDestination() *ObjectId {
  if !p.IsSetDestination() {
    return ForwardToObject_Destination_DEFAULT
  }
return p.Destination
}

func (p *ForwardToObject) GetMessage() protocol.Buffer {
  return p.Message
}
func (p *ForwardToObject) IsSetSource() bool {
  return p.Source != nil
}

func (p *ForwardToObject) IsSetDestination() bool {
  return p.Destination != nil
}

func (p *ForwardToObject) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetSsn bool = false;
  var issetSource bool = false;
  var issetDestination bool = false;
  var issetMessage bool = false;

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
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetSource = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetDestination = true
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
        issetMessage = true
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
  if !issetSource{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Source is not set"));
  }
  if !issetDestination{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Destination is not set"));
  }
  if !issetMessage{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Message is not set"));
  }
  return nil
}

func (p *ForwardToObject)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := protocol.Ssn(v)
  p.Ssn = temp
}
  return nil
}

func (p *ForwardToObject)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  p.Source = &ObjectId{}
  if err := p.Source.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Source), err)
  }
  return nil
}

func (p *ForwardToObject)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  p.Destination = &ObjectId{}
  if err := p.Destination.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Destination), err)
  }
  return nil
}

func (p *ForwardToObject)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  temp := protocol.Buffer(v)
  p.Message = temp
}
  return nil
}

func (p *ForwardToObject) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "ForwardToObject"); err != nil {
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

func (p *ForwardToObject) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "ssn", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ssn: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Ssn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.ssn (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ssn: ", p), err) }
  return err
}

func (p *ForwardToObject) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "source", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:source: ", p), err) }
  if err := p.Source.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Source), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:source: ", p), err) }
  return err
}

func (p *ForwardToObject) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "destination", thrift.STRUCT, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:destination: ", p), err) }
  if err := p.Destination.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Destination), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:destination: ", p), err) }
  return err
}

func (p *ForwardToObject) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "message", thrift.STRING, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:message: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.Message); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.message (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:message: ", p), err) }
  return err
}

func (p *ForwardToObject) Equals(other *ForwardToObject) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Ssn != other.Ssn { return false }
  if !p.Source.Equals(other.Source) { return false }
  if !p.Destination.Equals(other.Destination) { return false }
  if bytes.Compare(p.Message, other.Message) != 0 { return false }
  return true
}

func (p *ForwardToObject) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ForwardToObject(%+v)", *p)
}

func (p *ForwardToObject) Validate() error {
  return nil
}
// Attributes:
//  - Ssn
//  - Source
//  - Destination
//  - Message
type ForwardToMessagingChannel struct {
  Ssn protocol.Ssn `thrift:"ssn,1,required" db:"ssn" json:"ssn"`
  Source *ObjectId `thrift:"source,2,required" db:"source" json:"source"`
  Destination MessagingChannelId `thrift:"destination,3,required" db:"destination" json:"destination"`
  Message protocol.Buffer `thrift:"message,4,required" db:"message" json:"message"`
}

func NewForwardToMessagingChannel() *ForwardToMessagingChannel {
  return &ForwardToMessagingChannel{}
}


func (p *ForwardToMessagingChannel) GetSsn() protocol.Ssn {
  return p.Ssn
}
var ForwardToMessagingChannel_Source_DEFAULT *ObjectId
func (p *ForwardToMessagingChannel) GetSource() *ObjectId {
  if !p.IsSetSource() {
    return ForwardToMessagingChannel_Source_DEFAULT
  }
return p.Source
}

func (p *ForwardToMessagingChannel) GetDestination() MessagingChannelId {
  return p.Destination
}

func (p *ForwardToMessagingChannel) GetMessage() protocol.Buffer {
  return p.Message
}
func (p *ForwardToMessagingChannel) IsSetSource() bool {
  return p.Source != nil
}

func (p *ForwardToMessagingChannel) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetSsn bool = false;
  var issetSource bool = false;
  var issetDestination bool = false;
  var issetMessage bool = false;

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
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetSource = true
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
        issetDestination = true
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
        issetMessage = true
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
  if !issetSource{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Source is not set"));
  }
  if !issetDestination{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Destination is not set"));
  }
  if !issetMessage{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Message is not set"));
  }
  return nil
}

func (p *ForwardToMessagingChannel)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := protocol.Ssn(v)
  p.Ssn = temp
}
  return nil
}

func (p *ForwardToMessagingChannel)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  p.Source = &ObjectId{}
  if err := p.Source.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Source), err)
  }
  return nil
}

func (p *ForwardToMessagingChannel)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  temp := MessagingChannelId(v)
  p.Destination = temp
}
  return nil
}

func (p *ForwardToMessagingChannel)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  temp := protocol.Buffer(v)
  p.Message = temp
}
  return nil
}

func (p *ForwardToMessagingChannel) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "ForwardToMessagingChannel"); err != nil {
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

func (p *ForwardToMessagingChannel) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "ssn", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ssn: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Ssn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.ssn (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ssn: ", p), err) }
  return err
}

func (p *ForwardToMessagingChannel) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "source", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:source: ", p), err) }
  if err := p.Source.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Source), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:source: ", p), err) }
  return err
}

func (p *ForwardToMessagingChannel) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "destination", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:destination: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Destination)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.destination (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:destination: ", p), err) }
  return err
}

func (p *ForwardToMessagingChannel) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "message", thrift.STRING, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:message: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.Message); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.message (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:message: ", p), err) }
  return err
}

func (p *ForwardToMessagingChannel) Equals(other *ForwardToMessagingChannel) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Ssn != other.Ssn { return false }
  if !p.Source.Equals(other.Source) { return false }
  if p.Destination != other.Destination { return false }
  if bytes.Compare(p.Message, other.Message) != 0 { return false }
  return true
}

func (p *ForwardToMessagingChannel) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ForwardToMessagingChannel(%+v)", *p)
}

func (p *ForwardToMessagingChannel) Validate() error {
  return nil
}
