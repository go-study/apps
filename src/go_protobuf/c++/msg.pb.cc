// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: msg.proto

#define INTERNAL_SUPPRESS_PROTOBUF_FIELD_DEPRECATION
#include "msg.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/stubs/once.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)

namespace hello {

namespace {

const ::google::protobuf::Descriptor* ImgUrl_descriptor_ = NULL;
const ::google::protobuf::internal::GeneratedMessageReflection*
  ImgUrl_reflection_ = NULL;
const ::google::protobuf::Descriptor* User_descriptor_ = NULL;
const ::google::protobuf::internal::GeneratedMessageReflection*
  User_reflection_ = NULL;

}  // namespace


void protobuf_AssignDesc_msg_2eproto() {
  protobuf_AddDesc_msg_2eproto();
  const ::google::protobuf::FileDescriptor* file =
    ::google::protobuf::DescriptorPool::generated_pool()->FindFileByName(
      "msg.proto");
  GOOGLE_CHECK(file != NULL);
  ImgUrl_descriptor_ = file->message_type(0);
  static const int ImgUrl_offsets_[1] = {
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(ImgUrl, url_),
  };
  ImgUrl_reflection_ =
    new ::google::protobuf::internal::GeneratedMessageReflection(
      ImgUrl_descriptor_,
      ImgUrl::default_instance_,
      ImgUrl_offsets_,
      GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(ImgUrl, _has_bits_[0]),
      GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(ImgUrl, _unknown_fields_),
      -1,
      ::google::protobuf::DescriptorPool::generated_pool(),
      ::google::protobuf::MessageFactory::generated_factory(),
      sizeof(ImgUrl));
  User_descriptor_ = file->message_type(1);
  static const int User_offsets_[3] = {
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(User, uid_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(User, uname_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(User, imgurl_),
  };
  User_reflection_ =
    new ::google::protobuf::internal::GeneratedMessageReflection(
      User_descriptor_,
      User::default_instance_,
      User_offsets_,
      GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(User, _has_bits_[0]),
      GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(User, _unknown_fields_),
      -1,
      ::google::protobuf::DescriptorPool::generated_pool(),
      ::google::protobuf::MessageFactory::generated_factory(),
      sizeof(User));
}

namespace {

GOOGLE_PROTOBUF_DECLARE_ONCE(protobuf_AssignDescriptors_once_);
inline void protobuf_AssignDescriptorsOnce() {
  ::google::protobuf::GoogleOnceInit(&protobuf_AssignDescriptors_once_,
                 &protobuf_AssignDesc_msg_2eproto);
}

void protobuf_RegisterTypes(const ::std::string&) {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedMessage(
    ImgUrl_descriptor_, &ImgUrl::default_instance());
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedMessage(
    User_descriptor_, &User::default_instance());
}

}  // namespace

void protobuf_ShutdownFile_msg_2eproto() {
  delete ImgUrl::default_instance_;
  delete ImgUrl_reflection_;
  delete User::default_instance_;
  delete User_reflection_;
}

void protobuf_AddDesc_msg_2eproto() {
  static bool already_here = false;
  if (already_here) return;
  already_here = true;
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  ::google::protobuf::DescriptorPool::InternalAddGeneratedFile(
    "\n\tmsg.proto\022\005hello\"\025\n\006ImgUrl\022\013\n\003url\030\001 \002("
    "\t\"A\n\004User\022\013\n\003uid\030\001 \002(\005\022\r\n\005uname\030\002 \001(\t\022\035\n"
    "\006imgurl\030\003 \003(\0132\r.hello.ImgUrl", 108);
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedFile(
    "msg.proto", &protobuf_RegisterTypes);
  ImgUrl::default_instance_ = new ImgUrl();
  User::default_instance_ = new User();
  ImgUrl::default_instance_->InitAsDefaultInstance();
  User::default_instance_->InitAsDefaultInstance();
  ::google::protobuf::internal::OnShutdown(&protobuf_ShutdownFile_msg_2eproto);
}

// Force AddDescriptors() to be called at static initialization time.
struct StaticDescriptorInitializer_msg_2eproto {
  StaticDescriptorInitializer_msg_2eproto() {
    protobuf_AddDesc_msg_2eproto();
  }
} static_descriptor_initializer_msg_2eproto_;

// ===================================================================

#ifndef _MSC_VER
const int ImgUrl::kUrlFieldNumber;
#endif  // !_MSC_VER

ImgUrl::ImgUrl()
  : ::google::protobuf::Message() {
  SharedCtor();
}

void ImgUrl::InitAsDefaultInstance() {
}

ImgUrl::ImgUrl(const ImgUrl& from)
  : ::google::protobuf::Message() {
  SharedCtor();
  MergeFrom(from);
}

void ImgUrl::SharedCtor() {
  _cached_size_ = 0;
  url_ = const_cast< ::std::string*>(&::google::protobuf::internal::kEmptyString);
  ::memset(_has_bits_, 0, sizeof(_has_bits_));
}

ImgUrl::~ImgUrl() {
  SharedDtor();
}

void ImgUrl::SharedDtor() {
  if (url_ != &::google::protobuf::internal::kEmptyString) {
    delete url_;
  }
  if (this != default_instance_) {
  }
}

void ImgUrl::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* ImgUrl::descriptor() {
  protobuf_AssignDescriptorsOnce();
  return ImgUrl_descriptor_;
}

const ImgUrl& ImgUrl::default_instance() {
  if (default_instance_ == NULL) protobuf_AddDesc_msg_2eproto();
  return *default_instance_;
}

ImgUrl* ImgUrl::default_instance_ = NULL;

ImgUrl* ImgUrl::New() const {
  return new ImgUrl;
}

void ImgUrl::Clear() {
  if (_has_bits_[0 / 32] & (0xffu << (0 % 32))) {
    if (has_url()) {
      if (url_ != &::google::protobuf::internal::kEmptyString) {
        url_->clear();
      }
    }
  }
  ::memset(_has_bits_, 0, sizeof(_has_bits_));
  mutable_unknown_fields()->Clear();
}

bool ImgUrl::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!(EXPRESSION)) return false
  ::google::protobuf::uint32 tag;
  while ((tag = input->ReadTag()) != 0) {
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // required string url = 1;
      case 1: {
        if (::google::protobuf::internal::WireFormatLite::GetTagWireType(tag) ==
            ::google::protobuf::internal::WireFormatLite::WIRETYPE_LENGTH_DELIMITED) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadString(
                input, this->mutable_url()));
          ::google::protobuf::internal::WireFormat::VerifyUTF8String(
            this->url().data(), this->url().length(),
            ::google::protobuf::internal::WireFormat::PARSE);
        } else {
          goto handle_uninterpreted;
        }
        if (input->ExpectAtEnd()) return true;
        break;
      }

      default: {
      handle_uninterpreted:
        if (::google::protobuf::internal::WireFormatLite::GetTagWireType(tag) ==
            ::google::protobuf::internal::WireFormatLite::WIRETYPE_END_GROUP) {
          return true;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, mutable_unknown_fields()));
        break;
      }
    }
  }
  return true;
#undef DO_
}

void ImgUrl::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // required string url = 1;
  if (has_url()) {
    ::google::protobuf::internal::WireFormat::VerifyUTF8String(
      this->url().data(), this->url().length(),
      ::google::protobuf::internal::WireFormat::SERIALIZE);
    ::google::protobuf::internal::WireFormatLite::WriteString(
      1, this->url(), output);
  }

  if (!unknown_fields().empty()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        unknown_fields(), output);
  }
}

::google::protobuf::uint8* ImgUrl::SerializeWithCachedSizesToArray(
    ::google::protobuf::uint8* target) const {
  // required string url = 1;
  if (has_url()) {
    ::google::protobuf::internal::WireFormat::VerifyUTF8String(
      this->url().data(), this->url().length(),
      ::google::protobuf::internal::WireFormat::SERIALIZE);
    target =
      ::google::protobuf::internal::WireFormatLite::WriteStringToArray(
        1, this->url(), target);
  }

  if (!unknown_fields().empty()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        unknown_fields(), target);
  }
  return target;
}

int ImgUrl::ByteSize() const {
  int total_size = 0;

  if (_has_bits_[0 / 32] & (0xffu << (0 % 32))) {
    // required string url = 1;
    if (has_url()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::StringSize(
          this->url());
    }

  }
  if (!unknown_fields().empty()) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        unknown_fields());
  }
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = total_size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
  return total_size;
}

void ImgUrl::MergeFrom(const ::google::protobuf::Message& from) {
  GOOGLE_CHECK_NE(&from, this);
  const ImgUrl* source =
    ::google::protobuf::internal::dynamic_cast_if_available<const ImgUrl*>(
      &from);
  if (source == NULL) {
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
    MergeFrom(*source);
  }
}

void ImgUrl::MergeFrom(const ImgUrl& from) {
  GOOGLE_CHECK_NE(&from, this);
  if (from._has_bits_[0 / 32] & (0xffu << (0 % 32))) {
    if (from.has_url()) {
      set_url(from.url());
    }
  }
  mutable_unknown_fields()->MergeFrom(from.unknown_fields());
}

void ImgUrl::CopyFrom(const ::google::protobuf::Message& from) {
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void ImgUrl::CopyFrom(const ImgUrl& from) {
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool ImgUrl::IsInitialized() const {
  if ((_has_bits_[0] & 0x00000001) != 0x00000001) return false;

  return true;
}

void ImgUrl::Swap(ImgUrl* other) {
  if (other != this) {
    std::swap(url_, other->url_);
    std::swap(_has_bits_[0], other->_has_bits_[0]);
    _unknown_fields_.Swap(&other->_unknown_fields_);
    std::swap(_cached_size_, other->_cached_size_);
  }
}

::google::protobuf::Metadata ImgUrl::GetMetadata() const {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::Metadata metadata;
  metadata.descriptor = ImgUrl_descriptor_;
  metadata.reflection = ImgUrl_reflection_;
  return metadata;
}


// ===================================================================

#ifndef _MSC_VER
const int User::kUidFieldNumber;
const int User::kUnameFieldNumber;
const int User::kImgurlFieldNumber;
#endif  // !_MSC_VER

User::User()
  : ::google::protobuf::Message() {
  SharedCtor();
}

void User::InitAsDefaultInstance() {
}

User::User(const User& from)
  : ::google::protobuf::Message() {
  SharedCtor();
  MergeFrom(from);
}

void User::SharedCtor() {
  _cached_size_ = 0;
  uid_ = 0;
  uname_ = const_cast< ::std::string*>(&::google::protobuf::internal::kEmptyString);
  ::memset(_has_bits_, 0, sizeof(_has_bits_));
}

User::~User() {
  SharedDtor();
}

void User::SharedDtor() {
  if (uname_ != &::google::protobuf::internal::kEmptyString) {
    delete uname_;
  }
  if (this != default_instance_) {
  }
}

void User::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* User::descriptor() {
  protobuf_AssignDescriptorsOnce();
  return User_descriptor_;
}

const User& User::default_instance() {
  if (default_instance_ == NULL) protobuf_AddDesc_msg_2eproto();
  return *default_instance_;
}

User* User::default_instance_ = NULL;

User* User::New() const {
  return new User;
}

void User::Clear() {
  if (_has_bits_[0 / 32] & (0xffu << (0 % 32))) {
    uid_ = 0;
    if (has_uname()) {
      if (uname_ != &::google::protobuf::internal::kEmptyString) {
        uname_->clear();
      }
    }
  }
  imgurl_.Clear();
  ::memset(_has_bits_, 0, sizeof(_has_bits_));
  mutable_unknown_fields()->Clear();
}

bool User::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!(EXPRESSION)) return false
  ::google::protobuf::uint32 tag;
  while ((tag = input->ReadTag()) != 0) {
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // required int32 uid = 1;
      case 1: {
        if (::google::protobuf::internal::WireFormatLite::GetTagWireType(tag) ==
            ::google::protobuf::internal::WireFormatLite::WIRETYPE_VARINT) {
          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int32, ::google::protobuf::internal::WireFormatLite::TYPE_INT32>(
                 input, &uid_)));
          set_has_uid();
        } else {
          goto handle_uninterpreted;
        }
        if (input->ExpectTag(18)) goto parse_uname;
        break;
      }

      // optional string uname = 2;
      case 2: {
        if (::google::protobuf::internal::WireFormatLite::GetTagWireType(tag) ==
            ::google::protobuf::internal::WireFormatLite::WIRETYPE_LENGTH_DELIMITED) {
         parse_uname:
          DO_(::google::protobuf::internal::WireFormatLite::ReadString(
                input, this->mutable_uname()));
          ::google::protobuf::internal::WireFormat::VerifyUTF8String(
            this->uname().data(), this->uname().length(),
            ::google::protobuf::internal::WireFormat::PARSE);
        } else {
          goto handle_uninterpreted;
        }
        if (input->ExpectTag(26)) goto parse_imgurl;
        break;
      }

      // repeated .hello.ImgUrl imgurl = 3;
      case 3: {
        if (::google::protobuf::internal::WireFormatLite::GetTagWireType(tag) ==
            ::google::protobuf::internal::WireFormatLite::WIRETYPE_LENGTH_DELIMITED) {
         parse_imgurl:
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessageNoVirtual(
                input, add_imgurl()));
        } else {
          goto handle_uninterpreted;
        }
        if (input->ExpectTag(26)) goto parse_imgurl;
        if (input->ExpectAtEnd()) return true;
        break;
      }

      default: {
      handle_uninterpreted:
        if (::google::protobuf::internal::WireFormatLite::GetTagWireType(tag) ==
            ::google::protobuf::internal::WireFormatLite::WIRETYPE_END_GROUP) {
          return true;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, mutable_unknown_fields()));
        break;
      }
    }
  }
  return true;
#undef DO_
}

void User::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // required int32 uid = 1;
  if (has_uid()) {
    ::google::protobuf::internal::WireFormatLite::WriteInt32(1, this->uid(), output);
  }

  // optional string uname = 2;
  if (has_uname()) {
    ::google::protobuf::internal::WireFormat::VerifyUTF8String(
      this->uname().data(), this->uname().length(),
      ::google::protobuf::internal::WireFormat::SERIALIZE);
    ::google::protobuf::internal::WireFormatLite::WriteString(
      2, this->uname(), output);
  }

  // repeated .hello.ImgUrl imgurl = 3;
  for (int i = 0; i < this->imgurl_size(); i++) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      3, this->imgurl(i), output);
  }

  if (!unknown_fields().empty()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        unknown_fields(), output);
  }
}

::google::protobuf::uint8* User::SerializeWithCachedSizesToArray(
    ::google::protobuf::uint8* target) const {
  // required int32 uid = 1;
  if (has_uid()) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt32ToArray(1, this->uid(), target);
  }

  // optional string uname = 2;
  if (has_uname()) {
    ::google::protobuf::internal::WireFormat::VerifyUTF8String(
      this->uname().data(), this->uname().length(),
      ::google::protobuf::internal::WireFormat::SERIALIZE);
    target =
      ::google::protobuf::internal::WireFormatLite::WriteStringToArray(
        2, this->uname(), target);
  }

  // repeated .hello.ImgUrl imgurl = 3;
  for (int i = 0; i < this->imgurl_size(); i++) {
    target = ::google::protobuf::internal::WireFormatLite::
      WriteMessageNoVirtualToArray(
        3, this->imgurl(i), target);
  }

  if (!unknown_fields().empty()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        unknown_fields(), target);
  }
  return target;
}

int User::ByteSize() const {
  int total_size = 0;

  if (_has_bits_[0 / 32] & (0xffu << (0 % 32))) {
    // required int32 uid = 1;
    if (has_uid()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::Int32Size(
          this->uid());
    }

    // optional string uname = 2;
    if (has_uname()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::StringSize(
          this->uname());
    }

  }
  // repeated .hello.ImgUrl imgurl = 3;
  total_size += 1 * this->imgurl_size();
  for (int i = 0; i < this->imgurl_size(); i++) {
    total_size +=
      ::google::protobuf::internal::WireFormatLite::MessageSizeNoVirtual(
        this->imgurl(i));
  }

  if (!unknown_fields().empty()) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        unknown_fields());
  }
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = total_size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
  return total_size;
}

void User::MergeFrom(const ::google::protobuf::Message& from) {
  GOOGLE_CHECK_NE(&from, this);
  const User* source =
    ::google::protobuf::internal::dynamic_cast_if_available<const User*>(
      &from);
  if (source == NULL) {
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
    MergeFrom(*source);
  }
}

void User::MergeFrom(const User& from) {
  GOOGLE_CHECK_NE(&from, this);
  imgurl_.MergeFrom(from.imgurl_);
  if (from._has_bits_[0 / 32] & (0xffu << (0 % 32))) {
    if (from.has_uid()) {
      set_uid(from.uid());
    }
    if (from.has_uname()) {
      set_uname(from.uname());
    }
  }
  mutable_unknown_fields()->MergeFrom(from.unknown_fields());
}

void User::CopyFrom(const ::google::protobuf::Message& from) {
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void User::CopyFrom(const User& from) {
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool User::IsInitialized() const {
  if ((_has_bits_[0] & 0x00000001) != 0x00000001) return false;

  for (int i = 0; i < imgurl_size(); i++) {
    if (!this->imgurl(i).IsInitialized()) return false;
  }
  return true;
}

void User::Swap(User* other) {
  if (other != this) {
    std::swap(uid_, other->uid_);
    std::swap(uname_, other->uname_);
    imgurl_.Swap(&other->imgurl_);
    std::swap(_has_bits_[0], other->_has_bits_[0]);
    _unknown_fields_.Swap(&other->_unknown_fields_);
    std::swap(_cached_size_, other->_cached_size_);
  }
}

::google::protobuf::Metadata User::GetMetadata() const {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::Metadata metadata;
  metadata.descriptor = User_descriptor_;
  metadata.reflection = User_reflection_;
  return metadata;
}


// @@protoc_insertion_point(namespace_scope)

}  // namespace hello

// @@protoc_insertion_point(global_scope)