# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: scheduler.protobuf
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()

import task_pb2 as task__pb2

DESCRIPTOR = _descriptor.FileDescriptor(
    name='scheduler.protobuf',
    package='scheduler',
    syntax='proto3',
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
    serialized_pb=
    b'\n\x0fscheduler.protobuf\x12\tscheduler\x1a\ntask.protobuf\"/\n\nSchedulerSubmission\x12\x0f\n\x07job_dir\x18\x01 \x01(\t\x12\x10\n\x08hostname\x18\x02 \x01(\t\"p\n\x08SchedulerForwardReq\x12\x0f\n\x07task_id\x18\x01 \x01(\r\x12\'\n\x04\x66lag\x18\x02 \x01(\x0e\x32\x19.scheduler.SchedulerForwardFlag\x12*\n\x0bsubmissions\x18\x04 \x03(\x0b\x32\x15.scheduler.SchedulerSubmission\"a\n\x10\x46orwardingResult\x12\x0f\n\x07task_id\x18\x01 \x01(\r\x12\'\n\x04\x66lag\x18\x02 \x01(\x0e\x32\x19.scheduler.SchedulerForwardFlag\x12\x13\n\x07job_ids\x18\x03 \x03(\rB\x02\x10\x01*8\n\x0e\x46orwardingFlag\x12\x0b\n\x07\x46ORWARD\x10\x00\x12\n\n\x06\x46INISH\x10\x01\x12\r\n\tTERMINATE\x10\x02\x32t\n\tScheduler\x12&\n\x08InitFrom\x12\x0e.task.TaskRoot\x1a\n.task.Task\x12?\n\x0bStepForward\x12\x13.scheduler.SchedulerForwardReq\x1a\x1b.scheduler.SchedulerForwardResb\x06proto3',
    dependencies=[
        task__pb2.DESCRIPTOR,
    ])

_SchedulerForwardFlag = _descriptor.EnumDescriptor(
    name='SchedulerForwardFlag',
    full_name='scheduler.SchedulerForwardFlag',
    filename=None,
    file=DESCRIPTOR,
    create_key=_descriptor._internal_create_key,
    values=[
        _descriptor.EnumValueDescriptor(name='FORWARD',
                                        index=0,
                                        number=0,
                                        serialized_options=None,
                                        type=None,
                                        create_key=_descriptor._internal_create_key),
        _descriptor.EnumValueDescriptor(name='FINISH',
                                        index=1,
                                        number=1,
                                        serialized_options=None,
                                        type=None,
                                        create_key=_descriptor._internal_create_key),
        _descriptor.EnumValueDescriptor(name='TERMINATE',
                                        index=2,
                                        number=2,
                                        serialized_options=None,
                                        type=None,
                                        create_key=_descriptor._internal_create_key),
    ],
    containing_type=None,
    serialized_options=None,
    serialized_start=304,
    serialized_end=360,
)
_sym_db.RegisterEnumDescriptor(_SchedulerForwardFlag)

SchedulerForwardFlag = enum_type_wrapper.EnumTypeWrapper(_SchedulerForwardFlag)
FORWARD = 0
FINISH = 1
TERMINATE = 2

_SchedulerSubmission = _descriptor.Descriptor(
    name='SchedulerSubmission',
    full_name='scheduler.SchedulerSubmission',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(name='job_dir',
                                    full_name='scheduler.SchedulerSubmission.job_dir',
                                    index=0,
                                    number=1,
                                    type=9,
                                    cpp_type=9,
                                    label=1,
                                    has_default_value=False,
                                    default_value=b"".decode('utf-8'),
                                    message_type=None,
                                    enum_type=None,
                                    containing_type=None,
                                    is_extension=False,
                                    extension_scope=None,
                                    serialized_options=None,
                                    file=DESCRIPTOR,
                                    create_key=_descriptor._internal_create_key),
        _descriptor.FieldDescriptor(name='hostname',
                                    full_name='scheduler.SchedulerSubmission.hostname',
                                    index=1,
                                    number=2,
                                    type=9,
                                    cpp_type=9,
                                    label=1,
                                    has_default_value=False,
                                    default_value=b"".decode('utf-8'),
                                    message_type=None,
                                    enum_type=None,
                                    containing_type=None,
                                    is_extension=False,
                                    extension_scope=None,
                                    serialized_options=None,
                                    file=DESCRIPTOR,
                                    create_key=_descriptor._internal_create_key),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[],
    serialized_start=42,
    serialized_end=89,
)

_SchedulerForwardReq = _descriptor.Descriptor(
    name='SchedulerForwardReq',
    full_name='scheduler.SchedulerForwardReq',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(name='task_id',
                                    full_name='scheduler.SchedulerForwardReq.task_id',
                                    index=0,
                                    number=1,
                                    type=13,
                                    cpp_type=3,
                                    label=1,
                                    has_default_value=False,
                                    default_value=0,
                                    message_type=None,
                                    enum_type=None,
                                    containing_type=None,
                                    is_extension=False,
                                    extension_scope=None,
                                    serialized_options=None,
                                    file=DESCRIPTOR,
                                    create_key=_descriptor._internal_create_key),
        _descriptor.FieldDescriptor(name='flag',
                                    full_name='scheduler.SchedulerForwardReq.flag',
                                    index=1,
                                    number=2,
                                    type=14,
                                    cpp_type=8,
                                    label=1,
                                    has_default_value=False,
                                    default_value=0,
                                    message_type=None,
                                    enum_type=None,
                                    containing_type=None,
                                    is_extension=False,
                                    extension_scope=None,
                                    serialized_options=None,
                                    file=DESCRIPTOR,
                                    create_key=_descriptor._internal_create_key),
        _descriptor.FieldDescriptor(name='submissions',
                                    full_name='scheduler.SchedulerForwardReq.submissions',
                                    index=2,
                                    number=4,
                                    type=11,
                                    cpp_type=10,
                                    label=3,
                                    has_default_value=False,
                                    default_value=[],
                                    message_type=None,
                                    enum_type=None,
                                    containing_type=None,
                                    is_extension=False,
                                    extension_scope=None,
                                    serialized_options=None,
                                    file=DESCRIPTOR,
                                    create_key=_descriptor._internal_create_key),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[],
    serialized_start=91,
    serialized_end=203,
)

_SchedulerForwardRes = _descriptor.Descriptor(
    name='SchedulerForwardRes',
    full_name='scheduler.SchedulerForwardRes',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(name='task_id',
                                    full_name='scheduler.SchedulerForwardRes.task_id',
                                    index=0,
                                    number=1,
                                    type=13,
                                    cpp_type=3,
                                    label=1,
                                    has_default_value=False,
                                    default_value=0,
                                    message_type=None,
                                    enum_type=None,
                                    containing_type=None,
                                    is_extension=False,
                                    extension_scope=None,
                                    serialized_options=None,
                                    file=DESCRIPTOR,
                                    create_key=_descriptor._internal_create_key),
        _descriptor.FieldDescriptor(name='flag',
                                    full_name='scheduler.SchedulerForwardRes.flag',
                                    index=1,
                                    number=2,
                                    type=14,
                                    cpp_type=8,
                                    label=1,
                                    has_default_value=False,
                                    default_value=0,
                                    message_type=None,
                                    enum_type=None,
                                    containing_type=None,
                                    is_extension=False,
                                    extension_scope=None,
                                    serialized_options=None,
                                    file=DESCRIPTOR,
                                    create_key=_descriptor._internal_create_key),
        _descriptor.FieldDescriptor(name='job_ids',
                                    full_name='scheduler.SchedulerForwardRes.job_ids',
                                    index=2,
                                    number=3,
                                    type=13,
                                    cpp_type=3,
                                    label=3,
                                    has_default_value=False,
                                    default_value=[],
                                    message_type=None,
                                    enum_type=None,
                                    containing_type=None,
                                    is_extension=False,
                                    extension_scope=None,
                                    serialized_options=b'\020\001',
                                    file=DESCRIPTOR,
                                    create_key=_descriptor._internal_create_key),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[],
    serialized_start=205,
    serialized_end=302,
)

_SchedulerForwardReq.fields_by_name['flag'].enum_type = _SchedulerForwardFlag
_SchedulerForwardReq.fields_by_name['submissions'].message_type = _SchedulerSubmission
_SchedulerForwardRes.fields_by_name['flag'].enum_type = _SchedulerForwardFlag
DESCRIPTOR.message_types_by_name['SchedulerSubmission'] = _SchedulerSubmission
DESCRIPTOR.message_types_by_name['SchedulerForwardReq'] = _SchedulerForwardReq
DESCRIPTOR.message_types_by_name['SchedulerForwardRes'] = _SchedulerForwardRes
DESCRIPTOR.enum_types_by_name['SchedulerForwardFlag'] = _SchedulerForwardFlag
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SchedulerSubmission = _reflection.GeneratedProtocolMessageType(
    'SchedulerSubmission',
    (_message.Message,),
    {
        'DESCRIPTOR': _SchedulerSubmission,
        '__module__': 'scheduler_pb2'
        # @@protoc_insertion_point(class_scope:scheduler.SchedulerSubmission)
    })
_sym_db.RegisterMessage(SchedulerSubmission)

SchedulerForwardReq = _reflection.GeneratedProtocolMessageType(
    'SchedulerForwardReq',
    (_message.Message,),
    {
        'DESCRIPTOR': _SchedulerForwardReq,
        '__module__': 'scheduler_pb2'
        # @@protoc_insertion_point(class_scope:scheduler.SchedulerForwardReq)
    })
_sym_db.RegisterMessage(SchedulerForwardReq)

SchedulerForwardRes = _reflection.GeneratedProtocolMessageType(
    'SchedulerForwardRes',
    (_message.Message,),
    {
        'DESCRIPTOR': _SchedulerForwardRes,
        '__module__': 'scheduler_pb2'
        # @@protoc_insertion_point(class_scope:scheduler.SchedulerForwardRes)
    })
_sym_db.RegisterMessage(SchedulerForwardRes)

_SchedulerForwardRes.fields_by_name['job_ids']._options = None

_SCHEDULER = _descriptor.ServiceDescriptor(name='Scheduler',
                                           full_name='scheduler.Scheduler',
                                           file=DESCRIPTOR,
                                           index=0,
                                           serialized_options=None,
                                           create_key=_descriptor._internal_create_key,
                                           serialized_start=362,
                                           serialized_end=478,
                                           methods=[
                                               _descriptor.MethodDescriptor(
                                                   name='InitFrom',
                                                   full_name='scheduler.Scheduler.InitFrom',
                                                   index=0,
                                                   containing_service=None,
                                                   input_type=task__pb2._TASKROOT,
                                                   output_type=task__pb2._TASK,
                                                   serialized_options=None,
                                                   create_key=_descriptor._internal_create_key,
                                               ),
                                               _descriptor.MethodDescriptor(
                                                   name='StepForward',
                                                   full_name='scheduler.Scheduler.StepForward',
                                                   index=1,
                                                   containing_service=None,
                                                   input_type=_SchedulerForwardReq,
                                                   output_type=_SchedulerForwardRes,
                                                   serialized_options=None,
                                                   create_key=_descriptor._internal_create_key,
                                               ),
                                           ])
_sym_db.RegisterServiceDescriptor(_SCHEDULER)

DESCRIPTOR.services_by_name['Scheduler'] = _SCHEDULER

# @@protoc_insertion_point(module_scope)