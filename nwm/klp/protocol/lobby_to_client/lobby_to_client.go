// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package lobby_to_client

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
  MessageType_kAnsCategoryList MessageType = 1
  MessageType_kAnsSearchRoom MessageType = 2
  MessageType_kAnsRoomInfoList MessageType = 3
  MessageType_kForwardFromLobbyContents MessageType = 4
)

func (p MessageType) String() string {
  switch p {
  case MessageType_kAnsCategoryList: return "kAnsCategoryList"
  case MessageType_kAnsSearchRoom: return "kAnsSearchRoom"
  case MessageType_kAnsRoomInfoList: return "kAnsRoomInfoList"
  case MessageType_kForwardFromLobbyContents: return "kForwardFromLobbyContents"
  }
  return "<UNSET>"
}

func MessageTypeFromString(s string) (MessageType, error) {
  switch s {
  case "kAnsCategoryList": return MessageType_kAnsCategoryList, nil 
  case "kAnsSearchRoom": return MessageType_kAnsSearchRoom, nil 
  case "kAnsRoomInfoList": return MessageType_kAnsRoomInfoList, nil 
  case "kForwardFromLobbyContents": return MessageType_kForwardFromLobbyContents, nil 
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
//  - CategoryID
//  - Title
//  - MaxUserCount
//  - CurrentUserCount
type CategoryInfo struct {
  CategoryID protocol.CategoryId `thrift:"category_id,1,required" db:"category_id" json:"category_id"`
  Title string `thrift:"title,2,required" db:"title" json:"title"`
  MaxUserCount int32 `thrift:"max_user_count,3,required" db:"max_user_count" json:"max_user_count"`
  CurrentUserCount int32 `thrift:"current_user_count,4,required" db:"current_user_count" json:"current_user_count"`
}

func NewCategoryInfo() *CategoryInfo {
  return &CategoryInfo{}
}


func (p *CategoryInfo) GetCategoryID() protocol.CategoryId {
  return p.CategoryID
}

func (p *CategoryInfo) GetTitle() string {
  return p.Title
}

func (p *CategoryInfo) GetMaxUserCount() int32 {
  return p.MaxUserCount
}

func (p *CategoryInfo) GetCurrentUserCount() int32 {
  return p.CurrentUserCount
}
func (p *CategoryInfo) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetCategoryID bool = false;
  var issetTitle bool = false;
  var issetMaxUserCount bool = false;
  var issetCurrentUserCount bool = false;

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
        issetCategoryID = true
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
        issetTitle = true
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
        issetMaxUserCount = true
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
        issetCurrentUserCount = true
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
  if !issetCategoryID{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field CategoryID is not set"));
  }
  if !issetTitle{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Title is not set"));
  }
  if !issetMaxUserCount{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field MaxUserCount is not set"));
  }
  if !issetCurrentUserCount{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field CurrentUserCount is not set"));
  }
  return nil
}

func (p *CategoryInfo)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := protocol.CategoryId(v)
  p.CategoryID = temp
}
  return nil
}

func (p *CategoryInfo)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Title = v
}
  return nil
}

func (p *CategoryInfo)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.MaxUserCount = v
}
  return nil
}

func (p *CategoryInfo)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.CurrentUserCount = v
}
  return nil
}

func (p *CategoryInfo) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "CategoryInfo"); err != nil {
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

func (p *CategoryInfo) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "category_id", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:category_id: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.CategoryID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.category_id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:category_id: ", p), err) }
  return err
}

func (p *CategoryInfo) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "title", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:title: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Title)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.title (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:title: ", p), err) }
  return err
}

func (p *CategoryInfo) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "max_user_count", thrift.I32, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:max_user_count: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.MaxUserCount)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.max_user_count (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:max_user_count: ", p), err) }
  return err
}

func (p *CategoryInfo) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "current_user_count", thrift.I32, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:current_user_count: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.CurrentUserCount)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.current_user_count (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:current_user_count: ", p), err) }
  return err
}

func (p *CategoryInfo) Equals(other *CategoryInfo) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.CategoryID != other.CategoryID { return false }
  if p.Title != other.Title { return false }
  if p.MaxUserCount != other.MaxUserCount { return false }
  if p.CurrentUserCount != other.CurrentUserCount { return false }
  return true
}

func (p *CategoryInfo) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("CategoryInfo(%+v)", *p)
}

func (p *CategoryInfo) Validate() error {
  return nil
}
// Attributes:
//  - RequestKey
//  - Result_
//  - Categories
type AnsCategoryList struct {
  RequestKey protocol.ClientRequestKey `thrift:"request_key,1,required" db:"request_key" json:"request_key"`
  Result_ *klperror.Error `thrift:"result,2,required" db:"result" json:"result"`
  Categories []*CategoryInfo `thrift:"categories,3,required" db:"categories" json:"categories"`
}

func NewAnsCategoryList() *AnsCategoryList {
  return &AnsCategoryList{}
}


func (p *AnsCategoryList) GetRequestKey() protocol.ClientRequestKey {
  return p.RequestKey
}
var AnsCategoryList_Result__DEFAULT *klperror.Error
func (p *AnsCategoryList) GetResult_() *klperror.Error {
  if !p.IsSetResult_() {
    return AnsCategoryList_Result__DEFAULT
  }
return p.Result_
}

func (p *AnsCategoryList) GetCategories() []*CategoryInfo {
  return p.Categories
}
func (p *AnsCategoryList) IsSetResult_() bool {
  return p.Result_ != nil
}

func (p *AnsCategoryList) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetRequestKey bool = false;
  var issetResult_ bool = false;
  var issetCategories bool = false;

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
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetCategories = true
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
  if !issetCategories{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Categories is not set"));
  }
  return nil
}

func (p *AnsCategoryList)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := protocol.ClientRequestKey(v)
  p.RequestKey = temp
}
  return nil
}

func (p *AnsCategoryList)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  p.Result_ = &klperror.Error{}
  if err := p.Result_.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Result_), err)
  }
  return nil
}

func (p *AnsCategoryList)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin(ctx)
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*CategoryInfo, 0, size)
  p.Categories =  tSlice
  for i := 0; i < size; i ++ {
    _elem0 := &CategoryInfo{}
    if err := _elem0.Read(ctx, iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
    }
    p.Categories = append(p.Categories, _elem0)
  }
  if err := iprot.ReadListEnd(ctx); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *AnsCategoryList) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "AnsCategoryList"); err != nil {
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

func (p *AnsCategoryList) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "request_key", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request_key: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.RequestKey)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.request_key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request_key: ", p), err) }
  return err
}

func (p *AnsCategoryList) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "result", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:result: ", p), err) }
  if err := p.Result_.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Result_), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:result: ", p), err) }
  return err
}

func (p *AnsCategoryList) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "categories", thrift.LIST, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:categories: ", p), err) }
  if err := oprot.WriteListBegin(ctx, thrift.STRUCT, len(p.Categories)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.Categories {
    if err := v.Write(ctx, oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
    }
  }
  if err := oprot.WriteListEnd(ctx); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:categories: ", p), err) }
  return err
}

func (p *AnsCategoryList) Equals(other *AnsCategoryList) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.RequestKey != other.RequestKey { return false }
  if !p.Result_.Equals(other.Result_) { return false }
  if len(p.Categories) != len(other.Categories) { return false }
  for i, _tgt := range p.Categories {
    _src1 := other.Categories[i]
    if !_tgt.Equals(_src1) { return false }
  }
  return true
}

func (p *AnsCategoryList) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AnsCategoryList(%+v)", *p)
}

func (p *AnsCategoryList) Validate() error {
  return nil
}
// Attributes:
//  - RequestKey
//  - Result_
//  - RoomInfos
type AnsSearchRoom struct {
  RequestKey protocol.ClientRequestKey `thrift:"request_key,1,required" db:"request_key" json:"request_key"`
  Result_ *klperror.Error `thrift:"result,2,required" db:"result" json:"result"`
  RoomInfos []*protocol.RoomInfo `thrift:"room_infos,3,required" db:"room_infos" json:"room_infos"`
}

func NewAnsSearchRoom() *AnsSearchRoom {
  return &AnsSearchRoom{}
}


func (p *AnsSearchRoom) GetRequestKey() protocol.ClientRequestKey {
  return p.RequestKey
}
var AnsSearchRoom_Result__DEFAULT *klperror.Error
func (p *AnsSearchRoom) GetResult_() *klperror.Error {
  if !p.IsSetResult_() {
    return AnsSearchRoom_Result__DEFAULT
  }
return p.Result_
}

func (p *AnsSearchRoom) GetRoomInfos() []*protocol.RoomInfo {
  return p.RoomInfos
}
func (p *AnsSearchRoom) IsSetResult_() bool {
  return p.Result_ != nil
}

func (p *AnsSearchRoom) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetRequestKey bool = false;
  var issetResult_ bool = false;
  var issetRoomInfos bool = false;

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
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetRoomInfos = true
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
  if !issetRoomInfos{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field RoomInfos is not set"));
  }
  return nil
}

func (p *AnsSearchRoom)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := protocol.ClientRequestKey(v)
  p.RequestKey = temp
}
  return nil
}

func (p *AnsSearchRoom)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  p.Result_ = &klperror.Error{}
  if err := p.Result_.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Result_), err)
  }
  return nil
}

func (p *AnsSearchRoom)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin(ctx)
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*protocol.RoomInfo, 0, size)
  p.RoomInfos =  tSlice
  for i := 0; i < size; i ++ {
    _elem2 := &protocol.RoomInfo{}
    if err := _elem2.Read(ctx, iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem2), err)
    }
    p.RoomInfos = append(p.RoomInfos, _elem2)
  }
  if err := iprot.ReadListEnd(ctx); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *AnsSearchRoom) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "AnsSearchRoom"); err != nil {
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

func (p *AnsSearchRoom) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "request_key", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request_key: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.RequestKey)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.request_key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request_key: ", p), err) }
  return err
}

func (p *AnsSearchRoom) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "result", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:result: ", p), err) }
  if err := p.Result_.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Result_), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:result: ", p), err) }
  return err
}

func (p *AnsSearchRoom) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "room_infos", thrift.LIST, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:room_infos: ", p), err) }
  if err := oprot.WriteListBegin(ctx, thrift.STRUCT, len(p.RoomInfos)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.RoomInfos {
    if err := v.Write(ctx, oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
    }
  }
  if err := oprot.WriteListEnd(ctx); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:room_infos: ", p), err) }
  return err
}

func (p *AnsSearchRoom) Equals(other *AnsSearchRoom) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.RequestKey != other.RequestKey { return false }
  if !p.Result_.Equals(other.Result_) { return false }
  if len(p.RoomInfos) != len(other.RoomInfos) { return false }
  for i, _tgt := range p.RoomInfos {
    _src3 := other.RoomInfos[i]
    if !_tgt.Equals(_src3) { return false }
  }
  return true
}

func (p *AnsSearchRoom) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AnsSearchRoom(%+v)", *p)
}

func (p *AnsSearchRoom) Validate() error {
  return nil
}
// Attributes:
//  - RequestKey
//  - Result_
//  - RoomInfos
type AnsRoomInfoList struct {
  RequestKey protocol.ClientRequestKey `thrift:"request_key,1,required" db:"request_key" json:"request_key"`
  Result_ *klperror.Error `thrift:"result,2,required" db:"result" json:"result"`
  RoomInfos []*protocol.RoomInfo `thrift:"room_infos,3,required" db:"room_infos" json:"room_infos"`
}

func NewAnsRoomInfoList() *AnsRoomInfoList {
  return &AnsRoomInfoList{}
}


func (p *AnsRoomInfoList) GetRequestKey() protocol.ClientRequestKey {
  return p.RequestKey
}
var AnsRoomInfoList_Result__DEFAULT *klperror.Error
func (p *AnsRoomInfoList) GetResult_() *klperror.Error {
  if !p.IsSetResult_() {
    return AnsRoomInfoList_Result__DEFAULT
  }
return p.Result_
}

func (p *AnsRoomInfoList) GetRoomInfos() []*protocol.RoomInfo {
  return p.RoomInfos
}
func (p *AnsRoomInfoList) IsSetResult_() bool {
  return p.Result_ != nil
}

func (p *AnsRoomInfoList) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetRequestKey bool = false;
  var issetResult_ bool = false;
  var issetRoomInfos bool = false;

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
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetRoomInfos = true
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
  if !issetRoomInfos{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field RoomInfos is not set"));
  }
  return nil
}

func (p *AnsRoomInfoList)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := protocol.ClientRequestKey(v)
  p.RequestKey = temp
}
  return nil
}

func (p *AnsRoomInfoList)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  p.Result_ = &klperror.Error{}
  if err := p.Result_.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Result_), err)
  }
  return nil
}

func (p *AnsRoomInfoList)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin(ctx)
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*protocol.RoomInfo, 0, size)
  p.RoomInfos =  tSlice
  for i := 0; i < size; i ++ {
    _elem4 := &protocol.RoomInfo{}
    if err := _elem4.Read(ctx, iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem4), err)
    }
    p.RoomInfos = append(p.RoomInfos, _elem4)
  }
  if err := iprot.ReadListEnd(ctx); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *AnsRoomInfoList) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "AnsRoomInfoList"); err != nil {
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

func (p *AnsRoomInfoList) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "request_key", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request_key: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.RequestKey)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.request_key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request_key: ", p), err) }
  return err
}

func (p *AnsRoomInfoList) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "result", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:result: ", p), err) }
  if err := p.Result_.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Result_), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:result: ", p), err) }
  return err
}

func (p *AnsRoomInfoList) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "room_infos", thrift.LIST, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:room_infos: ", p), err) }
  if err := oprot.WriteListBegin(ctx, thrift.STRUCT, len(p.RoomInfos)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.RoomInfos {
    if err := v.Write(ctx, oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
    }
  }
  if err := oprot.WriteListEnd(ctx); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:room_infos: ", p), err) }
  return err
}

func (p *AnsRoomInfoList) Equals(other *AnsRoomInfoList) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.RequestKey != other.RequestKey { return false }
  if !p.Result_.Equals(other.Result_) { return false }
  if len(p.RoomInfos) != len(other.RoomInfos) { return false }
  for i, _tgt := range p.RoomInfos {
    _src5 := other.RoomInfos[i]
    if !_tgt.Equals(_src5) { return false }
  }
  return true
}

func (p *AnsRoomInfoList) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AnsRoomInfoList(%+v)", *p)
}

func (p *AnsRoomInfoList) Validate() error {
  return nil
}
// Attributes:
//  - ContentsMessage
type ForwardFromLobbyContents struct {
  ContentsMessage protocol.Buffer `thrift:"contents_message,1,required" db:"contents_message" json:"contents_message"`
}

func NewForwardFromLobbyContents() *ForwardFromLobbyContents {
  return &ForwardFromLobbyContents{}
}


func (p *ForwardFromLobbyContents) GetContentsMessage() protocol.Buffer {
  return p.ContentsMessage
}
func (p *ForwardFromLobbyContents) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetContentsMessage bool = false;

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
        issetContentsMessage = true
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
  if !issetContentsMessage{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ContentsMessage is not set"));
  }
  return nil
}

func (p *ForwardFromLobbyContents)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := protocol.Buffer(v)
  p.ContentsMessage = temp
}
  return nil
}

func (p *ForwardFromLobbyContents) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "ForwardFromLobbyContents"); err != nil {
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

func (p *ForwardFromLobbyContents) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "contents_message", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:contents_message: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.ContentsMessage); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.contents_message (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:contents_message: ", p), err) }
  return err
}

func (p *ForwardFromLobbyContents) Equals(other *ForwardFromLobbyContents) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if bytes.Compare(p.ContentsMessage, other.ContentsMessage) != 0 { return false }
  return true
}

func (p *ForwardFromLobbyContents) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ForwardFromLobbyContents(%+v)", *p)
}

func (p *ForwardFromLobbyContents) Validate() error {
  return nil
}
