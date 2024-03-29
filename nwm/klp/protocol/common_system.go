// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package protocol

import (
	"bytes"
	"context"
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

type PmbAddress string

func PmbAddressPtr(v PmbAddress) *PmbAddress { return &v }

type RequestKey int64

func RequestKeyPtr(v RequestKey) *RequestKey { return &v }

type UserState int32

func UserStatePtr(v UserState) *UserState { return &v }

// Attributes:
//  - IP
//  - Port
type InternetAddress struct {
  IP string `thrift:"ip,1,required" db:"ip" json:"ip"`
  Port int32 `thrift:"port,2,required" db:"port" json:"port"`
}

func NewInternetAddress() *InternetAddress {
  return &InternetAddress{}
}


func (p *InternetAddress) GetIP() string {
  return p.IP
}

func (p *InternetAddress) GetPort() int32 {
  return p.Port
}
func (p *InternetAddress) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetIP bool = false;
  var issetPort bool = false;

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
        issetIP = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetPort = true
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
  if !issetIP{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field IP is not set"));
  }
  if !issetPort{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Port is not set"));
  }
  return nil
}

func (p *InternetAddress)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.IP = v
}
  return nil
}

func (p *InternetAddress)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Port = v
}
  return nil
}

func (p *InternetAddress) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "InternetAddress"); err != nil {
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

func (p *InternetAddress) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "ip", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ip: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.IP)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.ip (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ip: ", p), err) }
  return err
}

func (p *InternetAddress) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "port", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:port: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Port)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.port (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:port: ", p), err) }
  return err
}

func (p *InternetAddress) Equals(other *InternetAddress) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.IP != other.IP { return false }
  if p.Port != other.Port { return false }
  return true
}

func (p *InternetAddress) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("InternetAddress(%+v)", *p)
}

func (p *InternetAddress) Validate() error {
  return nil
}
// Attributes:
//  - RoomInfo
//  - PmbAddress
//  - TCPAddress
type RoomInfoWithLocation struct {
  RoomInfo *RoomInfo `thrift:"room_info,1,required" db:"room_info" json:"room_info"`
  PmbAddress PmbAddress `thrift:"pmb_address,2,required" db:"pmb_address" json:"pmb_address"`
  TCPAddress *InternetAddress `thrift:"tcp_address,3,required" db:"tcp_address" json:"tcp_address"`
}

func NewRoomInfoWithLocation() *RoomInfoWithLocation {
  return &RoomInfoWithLocation{}
}

var RoomInfoWithLocation_RoomInfo_DEFAULT *RoomInfo
func (p *RoomInfoWithLocation) GetRoomInfo() *RoomInfo {
  if !p.IsSetRoomInfo() {
    return RoomInfoWithLocation_RoomInfo_DEFAULT
  }
return p.RoomInfo
}

func (p *RoomInfoWithLocation) GetPmbAddress() PmbAddress {
  return p.PmbAddress
}
var RoomInfoWithLocation_TCPAddress_DEFAULT *InternetAddress
func (p *RoomInfoWithLocation) GetTCPAddress() *InternetAddress {
  if !p.IsSetTCPAddress() {
    return RoomInfoWithLocation_TCPAddress_DEFAULT
  }
return p.TCPAddress
}
func (p *RoomInfoWithLocation) IsSetRoomInfo() bool {
  return p.RoomInfo != nil
}

func (p *RoomInfoWithLocation) IsSetTCPAddress() bool {
  return p.TCPAddress != nil
}

func (p *RoomInfoWithLocation) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetRoomInfo bool = false;
  var issetPmbAddress bool = false;
  var issetTCPAddress bool = false;

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
        issetRoomInfo = true
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
        issetPmbAddress = true
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
        issetTCPAddress = true
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
  if !issetRoomInfo{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field RoomInfo is not set"));
  }
  if !issetPmbAddress{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field PmbAddress is not set"));
  }
  if !issetTCPAddress{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field TCPAddress is not set"));
  }
  return nil
}

func (p *RoomInfoWithLocation)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  p.RoomInfo = &RoomInfo{}
  if err := p.RoomInfo.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RoomInfo), err)
  }
  return nil
}

func (p *RoomInfoWithLocation)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := PmbAddress(v)
  p.PmbAddress = temp
}
  return nil
}

func (p *RoomInfoWithLocation)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  p.TCPAddress = &InternetAddress{}
  if err := p.TCPAddress.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.TCPAddress), err)
  }
  return nil
}

func (p *RoomInfoWithLocation) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "RoomInfoWithLocation"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RoomInfoWithLocation) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "room_info", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:room_info: ", p), err) }
  if err := p.RoomInfo.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.RoomInfo), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:room_info: ", p), err) }
  return err
}

func (p *RoomInfoWithLocation) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "pmb_address", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:pmb_address: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.PmbAddress)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.pmb_address (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:pmb_address: ", p), err) }
  return err
}

func (p *RoomInfoWithLocation) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "tcp_address", thrift.STRUCT, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:tcp_address: ", p), err) }
  if err := p.TCPAddress.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.TCPAddress), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:tcp_address: ", p), err) }
  return err
}

func (p *RoomInfoWithLocation) Equals(other *RoomInfoWithLocation) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if !p.RoomInfo.Equals(other.RoomInfo) { return false }
  if p.PmbAddress != other.PmbAddress { return false }
  if !p.TCPAddress.Equals(other.TCPAddress) { return false }
  return true
}

func (p *RoomInfoWithLocation) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("RoomInfoWithLocation(%+v)", *p)
}

func (p *RoomInfoWithLocation) Validate() error {
  return nil
}
// Attributes:
//  - Ssn
//  - CategoryID
//  - RoomID
//  - PmbAddress
//  - TCPAddress
type RoomLocation struct {
  Ssn Ssn `thrift:"ssn,1,required" db:"ssn" json:"ssn"`
  CategoryID CategoryId `thrift:"category_id,2,required" db:"category_id" json:"category_id"`
  RoomID RoomId `thrift:"room_id,3,required" db:"room_id" json:"room_id"`
  PmbAddress PmbAddress `thrift:"pmb_address,4,required" db:"pmb_address" json:"pmb_address"`
  TCPAddress *InternetAddress `thrift:"tcp_address,5,required" db:"tcp_address" json:"tcp_address"`
}

func NewRoomLocation() *RoomLocation {
  return &RoomLocation{}
}


func (p *RoomLocation) GetSsn() Ssn {
  return p.Ssn
}

func (p *RoomLocation) GetCategoryID() CategoryId {
  return p.CategoryID
}

func (p *RoomLocation) GetRoomID() RoomId {
  return p.RoomID
}

func (p *RoomLocation) GetPmbAddress() PmbAddress {
  return p.PmbAddress
}
var RoomLocation_TCPAddress_DEFAULT *InternetAddress
func (p *RoomLocation) GetTCPAddress() *InternetAddress {
  if !p.IsSetTCPAddress() {
    return RoomLocation_TCPAddress_DEFAULT
  }
return p.TCPAddress
}
func (p *RoomLocation) IsSetTCPAddress() bool {
  return p.TCPAddress != nil
}

func (p *RoomLocation) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetSsn bool = false;
  var issetCategoryID bool = false;
  var issetRoomID bool = false;
  var issetPmbAddress bool = false;
  var issetTCPAddress bool = false;

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
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetCategoryID = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetRoomID = true
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
        issetPmbAddress = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField5(ctx, iprot); err != nil {
          return err
        }
        issetTCPAddress = true
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
  if !issetCategoryID{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field CategoryID is not set"));
  }
  if !issetRoomID{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field RoomID is not set"));
  }
  if !issetPmbAddress{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field PmbAddress is not set"));
  }
  if !issetTCPAddress{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field TCPAddress is not set"));
  }
  return nil
}

func (p *RoomLocation)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := Ssn(v)
  p.Ssn = temp
}
  return nil
}

func (p *RoomLocation)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := CategoryId(v)
  p.CategoryID = temp
}
  return nil
}

func (p *RoomLocation)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  temp := RoomId(v)
  p.RoomID = temp
}
  return nil
}

func (p *RoomLocation)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  temp := PmbAddress(v)
  p.PmbAddress = temp
}
  return nil
}

func (p *RoomLocation)  ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
  p.TCPAddress = &InternetAddress{}
  if err := p.TCPAddress.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.TCPAddress), err)
  }
  return nil
}

func (p *RoomLocation) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "RoomLocation"); err != nil {
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

func (p *RoomLocation) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "ssn", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ssn: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Ssn)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.ssn (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ssn: ", p), err) }
  return err
}

func (p *RoomLocation) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "category_id", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:category_id: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.CategoryID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.category_id (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:category_id: ", p), err) }
  return err
}

func (p *RoomLocation) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "room_id", thrift.I64, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:room_id: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.RoomID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.room_id (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:room_id: ", p), err) }
  return err
}

func (p *RoomLocation) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "pmb_address", thrift.STRING, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:pmb_address: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.PmbAddress)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.pmb_address (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:pmb_address: ", p), err) }
  return err
}

func (p *RoomLocation) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "tcp_address", thrift.STRUCT, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:tcp_address: ", p), err) }
  if err := p.TCPAddress.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.TCPAddress), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:tcp_address: ", p), err) }
  return err
}

func (p *RoomLocation) Equals(other *RoomLocation) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Ssn != other.Ssn { return false }
  if p.CategoryID != other.CategoryID { return false }
  if p.RoomID != other.RoomID { return false }
  if p.PmbAddress != other.PmbAddress { return false }
  if !p.TCPAddress.Equals(other.TCPAddress) { return false }
  return true
}

func (p *RoomLocation) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("RoomLocation(%+v)", *p)
}

func (p *RoomLocation) Validate() error {
  return nil
}
