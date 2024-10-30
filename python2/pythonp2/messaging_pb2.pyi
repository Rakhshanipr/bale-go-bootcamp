from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class RequestAddUser(_message.Message):
    __slots__ = ("user_name", "file_id")
    USER_NAME_FIELD_NUMBER: _ClassVar[int]
    FILE_ID_FIELD_NUMBER: _ClassVar[int]
    user_name: str
    file_id: str
    def __init__(self, user_name: _Optional[str] = ..., file_id: _Optional[str] = ...) -> None: ...

class ResponseAddUser(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class TextMessage(_message.Message):
    __slots__ = ("text",)
    TEXT_FIELD_NUMBER: _ClassVar[int]
    text: str
    def __init__(self, text: _Optional[str] = ...) -> None: ...

class FileMessage(_message.Message):
    __slots__ = ("file_id",)
    FILE_ID_FIELD_NUMBER: _ClassVar[int]
    file_id: str
    def __init__(self, file_id: _Optional[str] = ...) -> None: ...

class ImageMessage(_message.Message):
    __slots__ = ("file_id",)
    FILE_ID_FIELD_NUMBER: _ClassVar[int]
    file_id: str
    def __init__(self, file_id: _Optional[str] = ...) -> None: ...

class MessageContent(_message.Message):
    __slots__ = ("text", "file", "image")
    TEXT_FIELD_NUMBER: _ClassVar[int]
    FILE_FIELD_NUMBER: _ClassVar[int]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    text: TextMessage
    file: FileMessage
    image: ImageMessage
    def __init__(self, text: _Optional[_Union[TextMessage, _Mapping]] = ..., file: _Optional[_Union[FileMessage, _Mapping]] = ..., image: _Optional[_Union[ImageMessage, _Mapping]] = ...) -> None: ...

class RequestSendMessage(_message.Message):
    __slots__ = ("sender_id", "receiver_id", "content")
    SENDER_ID_FIELD_NUMBER: _ClassVar[int]
    RECEIVER_ID_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    sender_id: int
    receiver_id: int
    content: MessageContent
    def __init__(self, sender_id: _Optional[int] = ..., receiver_id: _Optional[int] = ..., content: _Optional[_Union[MessageContent, _Mapping]] = ...) -> None: ...

class ResponseSendMessage(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class RequestFetchMessage(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class ResponseFetchMessage(_message.Message):
    __slots__ = ("content",)
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    content: MessageContent
    def __init__(self, content: _Optional[_Union[MessageContent, _Mapping]] = ...) -> None: ...

class RequestGetUserMessages(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class ResponseGetUserMessages(_message.Message):
    __slots__ = ("contents",)
    CONTENTS_FIELD_NUMBER: _ClassVar[int]
    contents: _containers.RepeatedCompositeFieldContainer[MessageContent]
    def __init__(self, contents: _Optional[_Iterable[_Union[MessageContent, _Mapping]]] = ...) -> None: ...
