# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import job_pb2 as job__pb2
import task_pb2 as task__pb2


class QueryJobStub(object):
    """Missing associated documentation comment in .protobuf file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetJobs = channel.unary_unary(
                '/job.QueryJob/GetJobs',
                request_serializer=task__pb2.TaskId.SerializeToString,
                response_deserializer=job__pb2.JobList.FromString,
                )
        self.GetJobIds = channel.unary_unary(
                '/job.QueryJob/GetJobIds',
                request_serializer=task__pb2.TaskId.SerializeToString,
                response_deserializer=job__pb2.JobIds.FromString,
                )
        self.GetJob = channel.unary_unary(
                '/job.QueryJob/GetJob',
                request_serializer=job__pb2.JobId.SerializeToString,
                response_deserializer=job__pb2.Job.FromString,
                )


class QueryJobServicer(object):
    """Missing associated documentation comment in .protobuf file."""

    def GetJobs(self, request, context):
        """Missing associated documentation comment in .protobuf file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetJobIds(self, request, context):
        """Missing associated documentation comment in .protobuf file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetJob(self, request, context):
        """Missing associated documentation comment in .protobuf file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_QueryJobServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetJobs': grpc.unary_unary_rpc_method_handler(
                    servicer.GetJobs,
                    request_deserializer=task__pb2.TaskId.FromString,
                    response_serializer=job__pb2.JobList.SerializeToString,
            ),
            'GetJobIds': grpc.unary_unary_rpc_method_handler(
                    servicer.GetJobIds,
                    request_deserializer=task__pb2.TaskId.FromString,
                    response_serializer=job__pb2.JobIds.SerializeToString,
            ),
            'GetJob': grpc.unary_unary_rpc_method_handler(
                    servicer.GetJob,
                    request_deserializer=job__pb2.JobId.FromString,
                    response_serializer=job__pb2.Job.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'job.QueryJob', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class QueryJob(object):
    """Missing associated documentation comment in .protobuf file."""

    @staticmethod
    def GetJobs(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/job.QueryJob/GetJobs',
            task__pb2.TaskId.SerializeToString,
            job__pb2.JobList.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetJobIds(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/job.QueryJob/GetJobIds',
            task__pb2.TaskId.SerializeToString,
            job__pb2.JobIds.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetJob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/job.QueryJob/GetJob',
            job__pb2.JobId.SerializeToString,
            job__pb2.Job.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)