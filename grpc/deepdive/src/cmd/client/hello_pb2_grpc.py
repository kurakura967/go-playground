# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import hello_pb2 as hello__pb2


class GreetingServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Hello = channel.unary_unary(
                '/myapp.GreetingService/Hello',
                request_serializer=hello__pb2.HelloRequest.SerializeToString,
                response_deserializer=hello__pb2.HelloResponse.FromString,
                )
        self.HelloServerStream = channel.unary_stream(
                '/myapp.GreetingService/HelloServerStream',
                request_serializer=hello__pb2.HelloRequest.SerializeToString,
                response_deserializer=hello__pb2.HelloResponse.FromString,
                )
        self.HelloClientStream = channel.stream_unary(
                '/myapp.GreetingService/HelloClientStream',
                request_serializer=hello__pb2.HelloRequest.SerializeToString,
                response_deserializer=hello__pb2.HelloResponse.FromString,
                )
        self.HelloBiStreams = channel.stream_stream(
                '/myapp.GreetingService/HelloBiStreams',
                request_serializer=hello__pb2.HelloRequest.SerializeToString,
                response_deserializer=hello__pb2.HelloResponse.FromString,
                )


class GreetingServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Hello(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def HelloServerStream(self, request, context):
        """サーバーストリーミングRPC
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def HelloClientStream(self, request_iterator, context):
        """クライアントストリーミングRPC
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def HelloBiStreams(self, request_iterator, context):
        """双方向ストリーミングRPC
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_GreetingServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Hello': grpc.unary_unary_rpc_method_handler(
                    servicer.Hello,
                    request_deserializer=hello__pb2.HelloRequest.FromString,
                    response_serializer=hello__pb2.HelloResponse.SerializeToString,
            ),
            'HelloServerStream': grpc.unary_stream_rpc_method_handler(
                    servicer.HelloServerStream,
                    request_deserializer=hello__pb2.HelloRequest.FromString,
                    response_serializer=hello__pb2.HelloResponse.SerializeToString,
            ),
            'HelloClientStream': grpc.stream_unary_rpc_method_handler(
                    servicer.HelloClientStream,
                    request_deserializer=hello__pb2.HelloRequest.FromString,
                    response_serializer=hello__pb2.HelloResponse.SerializeToString,
            ),
            'HelloBiStreams': grpc.stream_stream_rpc_method_handler(
                    servicer.HelloBiStreams,
                    request_deserializer=hello__pb2.HelloRequest.FromString,
                    response_serializer=hello__pb2.HelloResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'myapp.GreetingService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class GreetingService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Hello(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/myapp.GreetingService/Hello',
            hello__pb2.HelloRequest.SerializeToString,
            hello__pb2.HelloResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def HelloServerStream(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/myapp.GreetingService/HelloServerStream',
            hello__pb2.HelloRequest.SerializeToString,
            hello__pb2.HelloResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def HelloClientStream(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/myapp.GreetingService/HelloClientStream',
            hello__pb2.HelloRequest.SerializeToString,
            hello__pb2.HelloResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def HelloBiStreams(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(request_iterator, target, '/myapp.GreetingService/HelloBiStreams',
            hello__pb2.HelloRequest.SerializeToString,
            hello__pb2.HelloResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
