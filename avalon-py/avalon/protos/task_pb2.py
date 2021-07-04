# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: task.protobuf
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import enum_type_wrapper

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()

from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2

DESCRIPTOR = _descriptor.FileDescriptor(
    name='task.protobuf',
    package='task',
    syntax='proto3',
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
    serialized_pb=b'\n\ntask.protobuf\x12\x04task\x1a\x1fgoogle/protobuf/timestamp.protobuf\"\x18\n\x08TaskRoot\x12\x0c\n\x04path\x18\x01 \x01(\t\"\x14\n\x06TaskId\x12\n\n\x02id\x18\x01 \x01(\r\"\xde\x01\n\x04Task\x12\n\n\x02id\x18\x01 \x01(\r\x12\x11\n\ttask_root\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x18\n\x0b\x64\x65scription\x18\x04 \x01(\tH\x00\x88\x01\x01\x12.\n\ncreated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0b\x66inished_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x1e\n\x05state\x18\x07 \x01(\x0e\x32\x0f.task.TaskStateB\x0e\n\x0c_description*N\n\tTaskState\x12\x0f\n\x0bINITIALIZED\x10\x00\x12\x0b\n\x07RUNNING\x10\x01\x12\x0e\n\nTERMINATED\x10\x02\x12\t\n\x05\x45RROR\x10\x03\x12\x08\n\x04\x44ONE\x10\x04\x32[\n\tQueryTask\x12)\n\tGetTaskId\x12\x0e.task.TaskRoot\x1a\x0c.task.TaskId\x12#\n\x07GetTask\x12\x0c.task.TaskId\x1a\n.task.Taskb\x06proto3'
    ,
    dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR, ])

_TASKSTATE = _descriptor.EnumDescriptor(
    name='TaskState',
    full_name='task.TaskState',
    filename=None,
    file=DESCRIPTOR,
    create_key=_descriptor._internal_create_key,
    values=[
        _descriptor.EnumValueDescriptor(
            name='INITIALIZED', index=0, number=0,
            serialized_options=None,
            type=None,
            create_key=_descriptor._internal_create_key),
        _descriptor.EnumValueDescriptor(
            name='RUNNING', index=1, number=1,
            serialized_options=None,
            type=None,
            create_key=_descriptor._internal_create_key),
        _descriptor.EnumValueDescriptor(
            name='TERMINATED', index=2, number=2,
            serialized_options=None,
            type=None,
            create_key=_descriptor._internal_create_key),
        _descriptor.EnumValueDescriptor(
            name='ERROR', index=3, number=3,
            serialized_options=None,
            type=None,
            create_key=_descriptor._internal_create_key),
        _descriptor.EnumValueDescriptor(
            name='DONE', index=4, number=4,
            serialized_options=None,
            type=None,
            create_key=_descriptor._internal_create_key),
    ],
    containing_type=None,
    serialized_options=None,
    serialized_start=326,
    serialized_end=404,
)
_sym_db.RegisterEnumDescriptor(_TASKSTATE)

TaskState = enum_type_wrapper.EnumTypeWrapper(_TASKSTATE)
INITIALIZED = 0
RUNNING = 1
TERMINATED = 2
ERROR = 3
DONE = 4

_TASKROOT = _descriptor.Descriptor(
    name='TaskRoot',
    full_name='task.TaskRoot',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name='path', full_name='task.TaskRoot.path', index=0,
            number=1, type=9, cpp_type=9, label=1,
            has_default_value=False, default_value=b"".decode('utf-8'),
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR, create_key=_descriptor._internal_create_key),
    ],
    extensions=[
    ],
    nested_types=[],
    enum_types=[
    ],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[
    ],
    serialized_start=53,
    serialized_end=77,
)

_TASKID = _descriptor.Descriptor(
    name='TaskId',
    full_name='task.TaskId',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name='id', full_name='task.TaskId.id', index=0,
            number=1, type=13, cpp_type=3, label=1,
            has_default_value=False, default_value=0,
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR, create_key=_descriptor._internal_create_key),
    ],
    extensions=[
    ],
    nested_types=[],
    enum_types=[
    ],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[
    ],
    serialized_start=79,
    serialized_end=99,
)

_TASK = _descriptor.Descriptor(
    name='Task',
    full_name='task.Task',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name='id', full_name='task.Task.id', index=0,
            number=1, type=13, cpp_type=3, label=1,
            has_default_value=False, default_value=0,
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR, create_key=_descriptor._internal_create_key),
        _descriptor.FieldDescriptor(
            name='task_root', full_name='task.Task.task_root', index=1,
            number=2, type=9, cpp_type=9, label=1,
            has_default_value=False, default_value=b"".decode('utf-8'),
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR, create_key=_descriptor._internal_create_key),
        _descriptor.FieldDescriptor(
            name='hello', full_name='task.Task.hello', index=2,
            number=3, type=9, cpp_type=9, label=1,
            has_default_value=False, default_value=b"".decode('utf-8'),
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR, create_key=_descriptor._internal_create_key),
        _descriptor.FieldDescriptor(
            name='description', full_name='task.Task.description', index=3,
            number=4, type=9, cpp_type=9, label=1,
            has_default_value=False, default_value=b"".decode('utf-8'),
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR, create_key=_descriptor._internal_create_key),
        _descriptor.FieldDescriptor(
            name='created_at', full_name='task.Task.created_at', index=4,
            number=5, type=11, cpp_type=10, label=1,
            has_default_value=False, default_value=None,
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR, create_key=_descriptor._internal_create_key),
        _descriptor.FieldDescriptor(
            name='finished_at', full_name='task.Task.finished_at', index=5,
            number=6, type=11, cpp_type=10, label=1,
            has_default_value=False, default_value=None,
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR, create_key=_descriptor._internal_create_key),
        _descriptor.FieldDescriptor(
            name='state', full_name='task.Task.state', index=6,
            number=7, type=14, cpp_type=8, label=1,
            has_default_value=False, default_value=0,
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR, create_key=_descriptor._internal_create_key),
    ],
    extensions=[
    ],
    nested_types=[],
    enum_types=[
    ],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[
        _descriptor.OneofDescriptor(
            name='_description', full_name='task.Task._description',
            index=0, containing_type=None,
            create_key=_descriptor._internal_create_key,
            fields=[]),
    ],
    serialized_start=102,
    serialized_end=324,
)

_TASK.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TASK.fields_by_name['finished_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TASK.fields_by_name['state'].enum_type = _TASKSTATE
_TASK.oneofs_by_name['_description'].fields.append(
    _TASK.fields_by_name['description'])
_TASK.fields_by_name['description'].containing_oneof = _TASK.oneofs_by_name['_description']
DESCRIPTOR.message_types_by_name['TaskRoot'] = _TASKROOT
DESCRIPTOR.message_types_by_name['TaskId'] = _TASKID
DESCRIPTOR.message_types_by_name['Task'] = _TASK
DESCRIPTOR.enum_types_by_name['TaskState'] = _TASKSTATE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TaskRoot = _reflection.GeneratedProtocolMessageType('TaskRoot', (_message.Message,), {
    'DESCRIPTOR': _TASKROOT,
    '__module__': 'task_pb2'
    # @@protoc_insertion_point(class_scope:task.TaskRoot)
})
_sym_db.RegisterMessage(TaskRoot)

TaskId = _reflection.GeneratedProtocolMessageType('TaskId', (_message.Message,), {
    'DESCRIPTOR': _TASKID,
    '__module__': 'task_pb2'
    # @@protoc_insertion_point(class_scope:task.TaskId)
})
_sym_db.RegisterMessage(TaskId)

Task = _reflection.GeneratedProtocolMessageType('Task', (_message.Message,), {
    'DESCRIPTOR': _TASK,
    '__module__': 'task_pb2'
    # @@protoc_insertion_point(class_scope:task.Task)
})
_sym_db.RegisterMessage(Task)

_QUERYTASK = _descriptor.ServiceDescriptor(
    name='QueryTask',
    full_name='task.QueryTask',
    file=DESCRIPTOR,
    index=0,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
    serialized_start=406,
    serialized_end=497,
    methods=[
        _descriptor.MethodDescriptor(
            name='GetTaskId',
            full_name='task.QueryTask.GetTaskId',
            index=0,
            containing_service=None,
            input_type=_TASKROOT,
            output_type=_TASKID,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.MethodDescriptor(
            name='GetTask',
            full_name='task.QueryTask.GetTask',
            index=1,
            containing_service=None,
            input_type=_TASKID,
            output_type=_TASK,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
    ])
_sym_db.RegisterServiceDescriptor(_QUERYTASK)

DESCRIPTOR.services_by_name['QueryTask'] = _QUERYTASK

# @@protoc_insertion_point(module_scope)
