// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package klpgateway_to_server

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
	klperror  "github.com/hhcneo/klpProtocol/nwm/klp/protocol/error"

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
var _ = klperror.GoUnusedProtection__
type MessageType int64
const (
  MessageType_kNtfKlpGatewayLogicTo MessageType = 1
)

func (p MessageType) String() string {
  switch p {
  case MessageType_kNtfKlpGatewayLogicTo: return "kNtfKlpGatewayLogicTo"
  }
  return "<UNSET>"
}

func MessageTypeFromString(s string) (MessageType, error) {
  switch s {
  case "kNtfKlpGatewayLogicTo": return MessageType_kNtfKlpGatewayLogicTo, nil 
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
//  - Result_
//  - Type
//  - Ssn
//  - Gsn
//  - RoutingKey
//  - RoomID
//  - Command
//  - Body
type NtfKlpGatewayLogicTo struct {
  Result_ *klperror.Error `thrift:"result,1,required" db:"result" json:"result"`
  Type protocol.Buffer `thrift:"type,2,required" db:"type" json:"type"`
  Ssn protocol.Ssn `thrift:"ssn,3,required" db:"ssn" json:"ssn"`
  Gsn protocol.Gsn `thrift:"gsn,4,required" db:"gsn" json:"gsn"`
  RoutingKey protocol.Buffer `thrift:"routing_key,5,required" db:"routing_key" json:"routing_key"`
  RoomID protocol.RoomId `thrift:"room_id,6,required" db:"room_id" json:"room_id"`
  Command protocol.Buffer `thrift:"command,7,required" db:"command" json:"command"`
  Body protocol.Buffer `thrift:"body,8,required" db:"body" json:"body"`
}

func NewNtfKlpGatewayLogicTo() *NtfKlpGatewayLogicTo {
  return &NtfKlpGatewayLogicTo{}
}

var NtfKlpGatewayLogicTo_Result__DEFAULT *klperror.Error
func (p *NtfKlpGatewayLogicTo) GetResult_() *klperror.Error {
  if !p.IsSetResult_() {
    return NtfKlpGatewayLogicTo_Result__DEFAULT
  }
return p.Result_
}

func (p *NtfKlpGatewayLogicTo) GetType() protocol.Buffer {
  return p.Type
}

func (p *NtfKlpGatewayLogicTo) GetSsn() protocol.Ssn {
  return p.Ssn
}

func (p *NtfKlpGatewayLogicTo) GetGsn() protocol.Gsn {
  return p.Gsn
}

func (p *NtfKlpGatewayLogicTo) GetRoutingKey() protocol.Buffer {
  return p.RoutingKey
}

func (p *NtfKlpGatewayLogicTo) GetRoomID() protocol.RoomId {
  return p.RoomID
}

func (p *NtfKlpGatewayLogicTo) GetCommand() protocol.Buffer {
  return p.Command
}

func (p *NtfKlpGatewayLogicTo) GetBody() protocol.Buffer {
  return p.Body
}
func (p *NtfKlpGatewayLogicTo) IsSetResult_() bool {
  return p.Result_ != nil
}

func (p *NtfKlpGatewayLogicTo) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetResult_ bool = false;
  var issetType bool = false;
  var issetSsn bool = false;
  var issetGsn bool = false;
  var issetRoutingKey bool = false;
  var issetRoomID bool = false;
  var issetCommand bool = false;
  var issetBody bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetResult_ = true
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
        issetType = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetSsn = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField4(ctx, iprot); err != nil {
          return err
        }
        issetGsn = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField5(ctx, iprot); err != nil {
          return err
        }
        issetRoutingKey = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 6:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField6(ctx, iprot); err != nil {
          return err
        }
        issetRoomID = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 7:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField7(ctx, iprot); err != nil {
          return err
        }
        issetCommand = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 8:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField8(ctx, iprot); err != nil {
          return err
        }
        issetBody = true
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
  if !issetResult_{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Result_ is not set"));
  }
  if !issetType{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Type is not set"));
  }
  if !issetSsn{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Ssn is not set"));
  }
  if !issetGsn{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Gsn is not set"));
  }
  if !issetRoutingKey{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field RoutingKey is not set"));
  }
  if !issetRoomID{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field RoomID is not set"));
  }
  if !issetCommand{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Command is not set"));
  }
  if !issetBody{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Body is not set"));
  }
  return nil
}

func (p *NtfKlpGatewayLogicTo)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  p.Result_ = &klperror.Error{}
  if err := p.Result_.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Result_), err)
  }
  return nil
}

func (p *NtfKlpGatewayLogicTo)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := protocol.Buffer(v)
  p.Type = temp
}
  return nil
}

func (p *NtfKlpGatewayLogicTo)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  temp := protocol.Ssn(v)
  p.Ssn = temp
}
  return nil
}

func (p *NtfKlpGatewayLogicTo)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  temp := protocol.Gsn(v)
  p.Gsn = temp
}
  return nil
}

func (p *NtfKlpGatewayLogicTo)  ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  temp := protocol.Buffer(v)
  p.RoutingKey = temp
}
  return nil
}

func (p *NtfKlpGatewayLogicTo)  ReadField6(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 6: ", err)
} else {
  temp := protocol.RoomId(v)
  p.RoomID = temp
}
  return nil
}

func (p *NtfKlpGatewayLogicTo)  ReadField7(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 7: ", err)
} else {
  temp := protocol.Buffer(v)
  p.Command = temp
}
  return nil
}

func (p *NtfKlpGatewayLogicTo)  ReadField8(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 8: ", err)
} else {
  temp := protocol.Buffer(v)
  p.Body = temp
}
  return nil
}

func (p *NtfKlpGatewayLogicTo) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "NtfKlpGatewayLogicTo"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
    if err := p.writeField4(ctx, oprot); err != nil { return err }
    if err := p.writeField5(ctx, oprot); err != nil { return err }
    if err := p.writeField6(ctx, oprot); err != nil { return err }
    if err := p.writeField7(ctx, oprot); err != nil { return err }
    if err := p.writeField8(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *NtfKlpGatewayLogicTo) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "result", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:result: ", p), err) }
  if err := p.Result_.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Result_), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:result: ", p), err) }
  return err
}

func (p *NtfKlpGatewayLogicTo) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "type", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:type: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.Type); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.type (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:type: ", p), err) }
  return err
}

func (p *NtfKlpGatewayLogicTo) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "ssn", thrift.I32, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:ssn: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Ssn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.ssn (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:ssn: ", p), err) }
  return err
}

func (p *NtfKlpGatewayLogicTo) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "gsn", thrift.I64, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:gsn: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.Gsn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.gsn (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:gsn: ", p), err) }
  return err
}

func (p *NtfKlpGatewayLogicTo) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "routing_key", thrift.STRING, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:routing_key: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.RoutingKey); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.routing_key (5) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:routing_key: ", p), err) }
  return err
}

func (p *NtfKlpGatewayLogicTo) writeField6(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "room_id", thrift.I64, 6); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:room_id: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.RoomID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.room_id (6) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 6:room_id: ", p), err) }
  return err
}

func (p *NtfKlpGatewayLogicTo) writeField7(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "command", thrift.STRING, 7); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:command: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.Command); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.command (7) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 7:command: ", p), err) }
  return err
}

func (p *NtfKlpGatewayLogicTo) writeField8(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "body", thrift.STRING, 8); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:body: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.Body); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.body (8) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 8:body: ", p), err) }
  return err
}

func (p *NtfKlpGatewayLogicTo) Equals(other *NtfKlpGatewayLogicTo) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if !p.Result_.Equals(other.Result_) { return false }
  if bytes.Compare(p.Type, other.Type) != 0 { return false }
  if p.Ssn != other.Ssn { return false }
  if p.Gsn != other.Gsn { return false }
  if bytes.Compare(p.RoutingKey, other.RoutingKey) != 0 { return false }
  if p.RoomID != other.RoomID { return false }
  if bytes.Compare(p.Command, other.Command) != 0 { return false }
  if bytes.Compare(p.Body, other.Body) != 0 { return false }
  return true
}

func (p *NtfKlpGatewayLogicTo) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("NtfKlpGatewayLogicTo(%+v)", *p)
}

func (p *NtfKlpGatewayLogicTo) Validate() error {
  return nil
}
