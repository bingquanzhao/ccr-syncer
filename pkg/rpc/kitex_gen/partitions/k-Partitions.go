// Code generated by Kitex v0.8.0. DO NOT EDIT.

package partitions

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"

	"github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/exprs"
	"github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/types"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
	_ = exprs.KitexUnusedProtection
	_ = types.KitexUnusedProtection
)

func (p *TPartitionKey) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetSign bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I16 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetSign = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetSign {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_TPartitionKey[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_TPartitionKey[fieldId]))
}

func (p *TPartitionKey) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI16(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Sign = v

	}
	return offset, nil
}

func (p *TPartitionKey) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		tmp := types.TPrimitiveType(v)
		p.Type = &tmp

	}
	return offset, nil
}

func (p *TPartitionKey) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.Key = &v

	}
	return offset, nil
}

// for compatibility
func (p *TPartitionKey) FastWrite(buf []byte) int {
	return 0
}

func (p *TPartitionKey) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "TPartitionKey")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *TPartitionKey) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("TPartitionKey")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *TPartitionKey) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "sign", thrift.I16, 1)
	offset += bthrift.Binary.WriteI16(buf[offset:], p.Sign)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TPartitionKey) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetType() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "type", thrift.I32, 2)
		offset += bthrift.Binary.WriteI32(buf[offset:], int32(*p.Type))

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TPartitionKey) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetKey() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "key", thrift.STRING, 3)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.Key)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TPartitionKey) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("sign", thrift.I16, 1)
	l += bthrift.Binary.I16Length(p.Sign)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TPartitionKey) field2Length() int {
	l := 0
	if p.IsSetType() {
		l += bthrift.Binary.FieldBeginLength("type", thrift.I32, 2)
		l += bthrift.Binary.I32Length(int32(*p.Type))

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TPartitionKey) field3Length() int {
	l := 0
	if p.IsSetKey() {
		l += bthrift.Binary.FieldBeginLength("key", thrift.STRING, 3)
		l += bthrift.Binary.StringLengthNocopy(*p.Key)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TPartitionRange) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetStartKey bool = false
	var issetEndKey bool = false
	var issetIncludeStartKey bool = false
	var issetIncludeEndKey bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetStartKey = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetEndKey = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.BOOL {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetIncludeStartKey = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.BOOL {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetIncludeEndKey = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetStartKey {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetEndKey {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetIncludeStartKey {
		fieldId = 3
		goto RequiredFieldNotSetError
	}

	if !issetIncludeEndKey {
		fieldId = 4
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_TPartitionRange[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_TPartitionRange[fieldId]))
}

func (p *TPartitionRange) FastReadField1(buf []byte) (int, error) {
	offset := 0

	tmp := NewTPartitionKey()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.StartKey = tmp
	return offset, nil
}

func (p *TPartitionRange) FastReadField2(buf []byte) (int, error) {
	offset := 0

	tmp := NewTPartitionKey()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.EndKey = tmp
	return offset, nil
}

func (p *TPartitionRange) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadBool(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.IncludeStartKey = v

	}
	return offset, nil
}

func (p *TPartitionRange) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadBool(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.IncludeEndKey = v

	}
	return offset, nil
}

// for compatibility
func (p *TPartitionRange) FastWrite(buf []byte) int {
	return 0
}

func (p *TPartitionRange) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "TPartitionRange")
	if p != nil {
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *TPartitionRange) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("TPartitionRange")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *TPartitionRange) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "start_key", thrift.STRUCT, 1)
	offset += p.StartKey.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TPartitionRange) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "end_key", thrift.STRUCT, 2)
	offset += p.EndKey.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TPartitionRange) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "include_start_key", thrift.BOOL, 3)
	offset += bthrift.Binary.WriteBool(buf[offset:], p.IncludeStartKey)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TPartitionRange) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "include_end_key", thrift.BOOL, 4)
	offset += bthrift.Binary.WriteBool(buf[offset:], p.IncludeEndKey)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TPartitionRange) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("start_key", thrift.STRUCT, 1)
	l += p.StartKey.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TPartitionRange) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("end_key", thrift.STRUCT, 2)
	l += p.EndKey.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TPartitionRange) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("include_start_key", thrift.BOOL, 3)
	l += bthrift.Binary.BoolLength(p.IncludeStartKey)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TPartitionRange) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("include_end_key", thrift.BOOL, 4)
	l += bthrift.Binary.BoolLength(p.IncludeEndKey)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TRangePartition) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetPartitionId bool = false
	var issetRange bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetPartitionId = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetRange = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetPartitionId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetRange {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_TRangePartition[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_TRangePartition[fieldId]))
}

func (p *TRangePartition) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PartitionId = v

	}
	return offset, nil
}

func (p *TRangePartition) FastReadField2(buf []byte) (int, error) {
	offset := 0

	tmp := NewTPartitionRange()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Range = tmp
	return offset, nil
}

func (p *TRangePartition) FastReadField3(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.DistributedExprs = make([]*exprs.TExpr, 0, size)
	for i := 0; i < size; i++ {
		_elem := exprs.NewTExpr()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		p.DistributedExprs = append(p.DistributedExprs, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

func (p *TRangePartition) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.DistributeBucket = &v

	}
	return offset, nil
}

// for compatibility
func (p *TRangePartition) FastWrite(buf []byte) int {
	return 0
}

func (p *TRangePartition) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "TRangePartition")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *TRangePartition) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("TRangePartition")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *TRangePartition) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "partition_id", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.PartitionId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TRangePartition) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "range", thrift.STRUCT, 2)
	offset += p.Range.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TRangePartition) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetDistributedExprs() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "distributed_exprs", thrift.LIST, 3)
		listBeginOffset := offset
		offset += bthrift.Binary.ListBeginLength(thrift.STRUCT, 0)
		var length int
		for _, v := range p.DistributedExprs {
			length++
			offset += v.FastWriteNocopy(buf[offset:], binaryWriter)
		}
		bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
		offset += bthrift.Binary.WriteListEnd(buf[offset:])
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TRangePartition) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetDistributeBucket() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "distribute_bucket", thrift.I32, 4)
		offset += bthrift.Binary.WriteI32(buf[offset:], *p.DistributeBucket)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TRangePartition) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("partition_id", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.PartitionId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TRangePartition) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("range", thrift.STRUCT, 2)
	l += p.Range.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TRangePartition) field3Length() int {
	l := 0
	if p.IsSetDistributedExprs() {
		l += bthrift.Binary.FieldBeginLength("distributed_exprs", thrift.LIST, 3)
		l += bthrift.Binary.ListBeginLength(thrift.STRUCT, len(p.DistributedExprs))
		for _, v := range p.DistributedExprs {
			l += v.BLength()
		}
		l += bthrift.Binary.ListEndLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TRangePartition) field4Length() int {
	l := 0
	if p.IsSetDistributeBucket() {
		l += bthrift.Binary.FieldBeginLength("distribute_bucket", thrift.I32, 4)
		l += bthrift.Binary.I32Length(*p.DistributeBucket)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TDataPartition) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetType bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetType = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetType {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_TDataPartition[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_TDataPartition[fieldId]))
}

func (p *TDataPartition) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Type = TPartitionType(v)

	}
	return offset, nil
}

func (p *TDataPartition) FastReadField2(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.PartitionExprs = make([]*exprs.TExpr, 0, size)
	for i := 0; i < size; i++ {
		_elem := exprs.NewTExpr()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		p.PartitionExprs = append(p.PartitionExprs, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

func (p *TDataPartition) FastReadField3(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.PartitionInfos = make([]*TRangePartition, 0, size)
	for i := 0; i < size; i++ {
		_elem := NewTRangePartition()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		p.PartitionInfos = append(p.PartitionInfos, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

// for compatibility
func (p *TDataPartition) FastWrite(buf []byte) int {
	return 0
}

func (p *TDataPartition) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "TDataPartition")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *TDataPartition) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("TDataPartition")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *TDataPartition) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "type", thrift.I32, 1)
	offset += bthrift.Binary.WriteI32(buf[offset:], int32(p.Type))

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TDataPartition) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetPartitionExprs() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "partition_exprs", thrift.LIST, 2)
		listBeginOffset := offset
		offset += bthrift.Binary.ListBeginLength(thrift.STRUCT, 0)
		var length int
		for _, v := range p.PartitionExprs {
			length++
			offset += v.FastWriteNocopy(buf[offset:], binaryWriter)
		}
		bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
		offset += bthrift.Binary.WriteListEnd(buf[offset:])
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TDataPartition) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetPartitionInfos() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "partition_infos", thrift.LIST, 3)
		listBeginOffset := offset
		offset += bthrift.Binary.ListBeginLength(thrift.STRUCT, 0)
		var length int
		for _, v := range p.PartitionInfos {
			length++
			offset += v.FastWriteNocopy(buf[offset:], binaryWriter)
		}
		bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
		offset += bthrift.Binary.WriteListEnd(buf[offset:])
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TDataPartition) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("type", thrift.I32, 1)
	l += bthrift.Binary.I32Length(int32(p.Type))

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TDataPartition) field2Length() int {
	l := 0
	if p.IsSetPartitionExprs() {
		l += bthrift.Binary.FieldBeginLength("partition_exprs", thrift.LIST, 2)
		l += bthrift.Binary.ListBeginLength(thrift.STRUCT, len(p.PartitionExprs))
		for _, v := range p.PartitionExprs {
			l += v.BLength()
		}
		l += bthrift.Binary.ListEndLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TDataPartition) field3Length() int {
	l := 0
	if p.IsSetPartitionInfos() {
		l += bthrift.Binary.FieldBeginLength("partition_infos", thrift.LIST, 3)
		l += bthrift.Binary.ListBeginLength(thrift.STRUCT, len(p.PartitionInfos))
		for _, v := range p.PartitionInfos {
			l += v.BLength()
		}
		l += bthrift.Binary.ListEndLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}
