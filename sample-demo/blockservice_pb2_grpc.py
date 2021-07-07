# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import blockservice_pb2 as blockservice__pb2


class BlockServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetMCBlockByNumber = channel.unary_unary(
                '/bsp_transaction.BlockService/GetMCBlockByNumber',
                request_serializer=blockservice__pb2.BspBlockRequest.SerializeToString,
                response_deserializer=blockservice__pb2.BlockResponse.FromString,
                )
        self.GetECBlockByNumber = channel.unary_unary(
                '/bsp_transaction.BlockService/GetECBlockByNumber',
                request_serializer=blockservice__pb2.BspBlockRequest.SerializeToString,
                response_deserializer=blockservice__pb2.BlockResponse.FromString,
                )


class BlockServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetMCBlockByNumber(self, request, context):
        """GetMCBlockByNumber actually returns MCHeader, rather than real body of MC-Block
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetECBlockByNumber(self, request, context):
        """GetECBlockByNumber returns a requested EC-Block
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_BlockServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetMCBlockByNumber': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMCBlockByNumber,
                    request_deserializer=blockservice__pb2.BspBlockRequest.FromString,
                    response_serializer=blockservice__pb2.BlockResponse.SerializeToString,
            ),
            'GetECBlockByNumber': grpc.unary_unary_rpc_method_handler(
                    servicer.GetECBlockByNumber,
                    request_deserializer=blockservice__pb2.BspBlockRequest.FromString,
                    response_serializer=blockservice__pb2.BlockResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'bsp_transaction.BlockService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class BlockService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetMCBlockByNumber(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/bsp_transaction.BlockService/GetMCBlockByNumber',
            blockservice__pb2.BspBlockRequest.SerializeToString,
            blockservice__pb2.BlockResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetECBlockByNumber(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/bsp_transaction.BlockService/GetECBlockByNumber',
            blockservice__pb2.BspBlockRequest.SerializeToString,
            blockservice__pb2.BlockResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
