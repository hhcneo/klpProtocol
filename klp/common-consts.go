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

const KlpProtocolVersion = 14
const KInvalidServerType = -1
const KLobbyServer = 1
const KRoomServer = 2
const KWebGatewayServer = 3
const KFriendsServer = 4
const KTnmtServer = 5
const KUserDefinedBegin = 100
const KMaxIntegralSearchConditionCount = 10
const KMaxStringSearchConditionCount = 0

func init() {
}
