// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: statement.proto

package statement

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Transaction struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	RefNo             int32                  `protobuf:"varint,1,opt,name=ref_no,json=refNo,proto3" json:"ref_no,omitempty"`
	Sender            string                 `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver          string                 `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	TransactionType   string                 `protobuf:"bytes,4,opt,name=transaction_type,json=transactionType,proto3" json:"transaction_type,omitempty"` // Withdraw, Deposit
	Amount            float64                `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	DateOfTransaction string                 `protobuf:"bytes,6,opt,name=date_of_transaction,json=dateOfTransaction,proto3" json:"date_of_transaction,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	mi := &file_statement_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_statement_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_statement_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetRefNo() int32 {
	if x != nil {
		return x.RefNo
	}
	return 0
}

func (x *Transaction) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Transaction) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *Transaction) GetTransactionType() string {
	if x != nil {
		return x.TransactionType
	}
	return ""
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetDateOfTransaction() string {
	if x != nil {
		return x.DateOfTransaction
	}
	return ""
}

type Statement struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	AccountNumber  string                 `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StatementDate  string                 `protobuf:"bytes,3,opt,name=statement_date,json=statementDate,proto3" json:"statement_date,omitempty"`
	CurrentBalance float64                `protobuf:"fixed64,4,opt,name=current_balance,json=currentBalance,proto3" json:"current_balance,omitempty"`
	Balance        float64                `protobuf:"fixed64,5,opt,name=balance,proto3" json:"balance,omitempty"`
	Transactions   []*Transaction         `protobuf:"bytes,6,rep,name=transactions,proto3" json:"transactions,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Statement) Reset() {
	*x = Statement{}
	mi := &file_statement_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement) ProtoMessage() {}

func (x *Statement) ProtoReflect() protoreflect.Message {
	mi := &file_statement_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statement.ProtoReflect.Descriptor instead.
func (*Statement) Descriptor() ([]byte, []int) {
	return file_statement_proto_rawDescGZIP(), []int{1}
}

func (x *Statement) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *Statement) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Statement) GetStatementDate() string {
	if x != nil {
		return x.StatementDate
	}
	return ""
}

func (x *Statement) GetCurrentBalance() float64 {
	if x != nil {
		return x.CurrentBalance
	}
	return 0
}

func (x *Statement) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Statement) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type AddStatementRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountNumber string                 `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	StatementDate string                 `protobuf:"bytes,2,opt,name=statement_date,json=statementDate,proto3" json:"statement_date,omitempty"`
	StartDate     string                 `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate       string                 `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddStatementRequest) Reset() {
	*x = AddStatementRequest{}
	mi := &file_statement_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddStatementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStatementRequest) ProtoMessage() {}

func (x *AddStatementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_statement_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStatementRequest.ProtoReflect.Descriptor instead.
func (*AddStatementRequest) Descriptor() ([]byte, []int) {
	return file_statement_proto_rawDescGZIP(), []int{2}
}

func (x *AddStatementRequest) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *AddStatementRequest) GetStatementDate() string {
	if x != nil {
		return x.StatementDate
	}
	return ""
}

func (x *AddStatementRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *AddStatementRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type AddStatementResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Statement     *Statement             `protobuf:"bytes,2,opt,name=statement,proto3" json:"statement,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddStatementResponse) Reset() {
	*x = AddStatementResponse{}
	mi := &file_statement_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddStatementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStatementResponse) ProtoMessage() {}

func (x *AddStatementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_statement_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStatementResponse.ProtoReflect.Descriptor instead.
func (*AddStatementResponse) Descriptor() ([]byte, []int) {
	return file_statement_proto_rawDescGZIP(), []int{3}
}

func (x *AddStatementResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddStatementResponse) GetStatement() *Statement {
	if x != nil {
		return x.Statement
	}
	return nil
}

type GetStatementRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountNumber string                 `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	StatementDate string                 `protobuf:"bytes,2,opt,name=statement_date,json=statementDate,proto3" json:"statement_date,omitempty"`
	Format        string                 `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"` // for file types like --> pdf, csv
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStatementRequest) Reset() {
	*x = GetStatementRequest{}
	mi := &file_statement_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatementRequest) ProtoMessage() {}

func (x *GetStatementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_statement_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatementRequest.ProtoReflect.Descriptor instead.
func (*GetStatementRequest) Descriptor() ([]byte, []int) {
	return file_statement_proto_rawDescGZIP(), []int{4}
}

func (x *GetStatementRequest) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *GetStatementRequest) GetStatementDate() string {
	if x != nil {
		return x.StatementDate
	}
	return ""
}

func (x *GetStatementRequest) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

type GetStatementResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Statement     *Statement             `protobuf:"bytes,1,opt,name=statement,proto3" json:"statement,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStatementResponse) Reset() {
	*x = GetStatementResponse{}
	mi := &file_statement_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatementResponse) ProtoMessage() {}

func (x *GetStatementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_statement_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatementResponse.ProtoReflect.Descriptor instead.
func (*GetStatementResponse) Descriptor() ([]byte, []int) {
	return file_statement_proto_rawDescGZIP(), []int{5}
}

func (x *GetStatementResponse) GetStatement() *Statement {
	if x != nil {
		return x.Statement
	}
	return nil
}

var File_statement_proto protoreflect.FileDescriptor

var file_statement_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x72, 0x65, 0x66, 0x4e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2e, 0x0a, 0x13, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x66, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xe2, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x22, 0x5a, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x7b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x40, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x32,
	0x8c, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x41, 0x64, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x20,
	0x5a, 0x1e, 0x66, 0x69, 0x6e, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_statement_proto_rawDescOnce sync.Once
	file_statement_proto_rawDescData []byte
)

func file_statement_proto_rawDescGZIP() []byte {
	file_statement_proto_rawDescOnce.Do(func() {
		file_statement_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_statement_proto_rawDesc), len(file_statement_proto_rawDesc)))
	})
	return file_statement_proto_rawDescData
}

var file_statement_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_statement_proto_goTypes = []any{
	(*Transaction)(nil),          // 0: Transaction
	(*Statement)(nil),            // 1: Statement
	(*AddStatementRequest)(nil),  // 2: AddStatementRequest
	(*AddStatementResponse)(nil), // 3: AddStatementResponse
	(*GetStatementRequest)(nil),  // 4: GetStatementRequest
	(*GetStatementResponse)(nil), // 5: GetStatementResponse
}
var file_statement_proto_depIdxs = []int32{
	0, // 0: Statement.transactions:type_name -> Transaction
	1, // 1: AddStatementResponse.statement:type_name -> Statement
	1, // 2: GetStatementResponse.statement:type_name -> Statement
	2, // 3: StatementService.AddStatement:input_type -> AddStatementRequest
	4, // 4: StatementService.GetStatement:input_type -> GetStatementRequest
	3, // 5: StatementService.AddStatement:output_type -> AddStatementResponse
	5, // 6: StatementService.GetStatement:output_type -> GetStatementResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_statement_proto_init() }
func file_statement_proto_init() {
	if File_statement_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_statement_proto_rawDesc), len(file_statement_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_statement_proto_goTypes,
		DependencyIndexes: file_statement_proto_depIdxs,
		MessageInfos:      file_statement_proto_msgTypes,
	}.Build()
	File_statement_proto = out.File
	file_statement_proto_goTypes = nil
	file_statement_proto_depIdxs = nil
}
