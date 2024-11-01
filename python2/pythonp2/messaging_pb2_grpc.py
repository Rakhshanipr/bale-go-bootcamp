# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import python2.pythonp2.messaging_pb2 as messaging__pb2

GRPC_GENERATED_VERSION = '1.67.1'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The gogrpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in messaging_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your gogrpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class MessagingStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A gogrpc.Channel.
        """
        self.AddUser = channel.unary_unary(
                '/Messaging/AddUser',
                request_serializer=messaging__pb2.RequestAddUser.SerializeToString,
                response_deserializer=messaging__pb2.ResponseAddUser.FromString,
                _registered_method=True)
        self.SendMessage = channel.unary_unary(
                '/Messaging/SendMessage',
                request_serializer=messaging__pb2.RequestSendMessage.SerializeToString,
                response_deserializer=messaging__pb2.ResponseSendMessage.FromString,
                _registered_method=True)
        self.FetchMessage = channel.unary_unary(
                '/Messaging/FetchMessage',
                request_serializer=messaging__pb2.RequestFetchMessage.SerializeToString,
                response_deserializer=messaging__pb2.ResponseFetchMessage.FromString,
                _registered_method=True)
        self.GetUserMessages = channel.unary_unary(
                '/Messaging/GetUserMessages',
                request_serializer=messaging__pb2.RequestGetUserMessages.SerializeToString,
                response_deserializer=messaging__pb2.ResponseGetUserMessages.FromString,
                _registered_method=True)


class MessagingServicer(object):
    """Missing associated documentation comment in .proto file."""

    def AddUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendMessage(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def FetchMessage(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUserMessages(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MessagingServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'AddUser': grpc.unary_unary_rpc_method_handler(
                    servicer.AddUser,
                    request_deserializer=messaging__pb2.RequestAddUser.FromString,
                    response_serializer=messaging__pb2.ResponseAddUser.SerializeToString,
            ),
            'SendMessage': grpc.unary_unary_rpc_method_handler(
                    servicer.SendMessage,
                    request_deserializer=messaging__pb2.RequestSendMessage.FromString,
                    response_serializer=messaging__pb2.ResponseSendMessage.SerializeToString,
            ),
            'FetchMessage': grpc.unary_unary_rpc_method_handler(
                    servicer.FetchMessage,
                    request_deserializer=messaging__pb2.RequestFetchMessage.FromString,
                    response_serializer=messaging__pb2.ResponseFetchMessage.SerializeToString,
            ),
            'GetUserMessages': grpc.unary_unary_rpc_method_handler(
                    servicer.GetUserMessages,
                    request_deserializer=messaging__pb2.RequestGetUserMessages.FromString,
                    response_serializer=messaging__pb2.ResponseGetUserMessages.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Messaging', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('Messaging', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Messaging(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def AddUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/Messaging/AddUser',
            messaging__pb2.RequestAddUser.SerializeToString,
            messaging__pb2.ResponseAddUser.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SendMessage(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/Messaging/SendMessage',
            messaging__pb2.RequestSendMessage.SerializeToString,
            messaging__pb2.ResponseSendMessage.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def FetchMessage(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/Messaging/FetchMessage',
            messaging__pb2.RequestFetchMessage.SerializeToString,
            messaging__pb2.ResponseFetchMessage.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetUserMessages(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/Messaging/GetUserMessages',
            messaging__pb2.RequestGetUserMessages.SerializeToString,
            messaging__pb2.ResponseGetUserMessages.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
