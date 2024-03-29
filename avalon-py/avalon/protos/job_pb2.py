# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: job.protobuf
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
import task_pb2 as task__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='job.protobuf',
  package='job',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\tjob.protobuf\x12\x03job\x1a\x1fgoogle/protobuf/timestamp.protobuf\x1a\ntask.protobuf\"!\n\x07JobList\x12\x16\n\x04jobs\x18\x01 \x03(\x0b\x32\x08.job.Job\"\x13\n\x05JobId\x12\n\n\x02id\x18\x01 \x01(\r\"\x18\n\x06JobIds\x12\x0e\n\x02id\x18\x01 \x03(\x05\x42\x02\x10\x01\"\xd5\x01\n\x03Job\x12\n\n\x02id\x18\x01 \x01(\r\x12\x0f\n\x07job_dir\x18\x02 \x01(\t\x12\x10\n\x08hostname\x18\x03 \x01(\t\x12.\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0b\x66inished_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x1c\n\x05state\x18\x06 \x01(\x0e\x32\r.job.JobState\x12\x14\n\x07message\x18\x07 \x01(\tH\x00\x88\x01\x01\x42\n\n\x08_message*b\n\x08JobState\x12\x0b\n\x07STANDBY\x10\x00\x12\n\n\x06MOVING\x10\x01\x12\x0b\n\x07PENDING\x10\x02\x12\x0b\n\x07RUNNING\x10\x03\x12\x08\n\x04\x44ONE\x10\x04\x12\t\n\x05\x45RROR\x10\x05\x12\x0e\n\nTERMINATED\x10\x06\x32y\n\x08QueryJob\x12%\n\x07GetJobs\x12\x0c.task.TaskId\x1a\x0c.job.JobList\x12&\n\tGetJobIds\x12\x0c.task.TaskId\x1a\x0b.job.JobIds\x12\x1e\n\x06GetJob\x12\n.job.JobId\x1a\x08.job.Jobb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,task__pb2.DESCRIPTOR,])

_JOBSTATE = _descriptor.EnumDescriptor(
  name='JobState',
  full_name='job.JobState',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='STANDBY', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='MOVING', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='PENDING', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='RUNNING', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DONE', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ERROR', index=5, number=5,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TERMINATED', index=6, number=6,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=361,
  serialized_end=459,
)
_sym_db.RegisterEnumDescriptor(_JOBSTATE)

JobState = enum_type_wrapper.EnumTypeWrapper(_JOBSTATE)
STANDBY = 0
MOVING = 1
PENDING = 2
RUNNING = 3
DONE = 4
ERROR = 5
TERMINATED = 6



_JOBLIST = _descriptor.Descriptor(
  name='JobList',
  full_name='job.JobList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='jobs', full_name='job.JobList.jobs', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=63,
  serialized_end=96,
)


_JOBID = _descriptor.Descriptor(
  name='JobId',
  full_name='job.JobId',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='job.JobId.id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=98,
  serialized_end=117,
)


_JOBIDS = _descriptor.Descriptor(
  name='JobIds',
  full_name='job.JobIds',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='job.JobIds.id', index=0,
      number=1, type=5, cpp_type=1, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\020\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=119,
  serialized_end=143,
)


_JOB = _descriptor.Descriptor(
  name='Job',
  full_name='job.Job',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='job.Job.id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='job_dir', full_name='job.Job.job_dir', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='hostname', full_name='job.Job.hostname', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='job.Job.created_at', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='finished_at', full_name='job.Job.finished_at', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='state', full_name='job.Job.state', index=5,
      number=6, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='job.Job.message', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
      name='_message', full_name='job.Job._message',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=146,
  serialized_end=359,
)

_JOBLIST.fields_by_name['jobs'].message_type = _JOB
_JOB.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_JOB.fields_by_name['finished_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_JOB.fields_by_name['state'].enum_type = _JOBSTATE
_JOB.oneofs_by_name['_message'].fields.append(
  _JOB.fields_by_name['message'])
_JOB.fields_by_name['message'].containing_oneof = _JOB.oneofs_by_name['_message']
DESCRIPTOR.message_types_by_name['JobList'] = _JOBLIST
DESCRIPTOR.message_types_by_name['JobId'] = _JOBID
DESCRIPTOR.message_types_by_name['JobIds'] = _JOBIDS
DESCRIPTOR.message_types_by_name['Job'] = _JOB
DESCRIPTOR.enum_types_by_name['JobState'] = _JOBSTATE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

JobList = _reflection.GeneratedProtocolMessageType('JobList', (_message.Message,), {
  'DESCRIPTOR' : _JOBLIST,
  '__module__' : 'job_pb2'
  # @@protoc_insertion_point(class_scope:job.JobList)
  })
_sym_db.RegisterMessage(JobList)

JobId = _reflection.GeneratedProtocolMessageType('JobId', (_message.Message,), {
  'DESCRIPTOR' : _JOBID,
  '__module__' : 'job_pb2'
  # @@protoc_insertion_point(class_scope:job.JobId)
  })
_sym_db.RegisterMessage(JobId)

JobIds = _reflection.GeneratedProtocolMessageType('JobIds', (_message.Message,), {
  'DESCRIPTOR' : _JOBIDS,
  '__module__' : 'job_pb2'
  # @@protoc_insertion_point(class_scope:job.JobIds)
  })
_sym_db.RegisterMessage(JobIds)

Job = _reflection.GeneratedProtocolMessageType('Job', (_message.Message,), {
  'DESCRIPTOR' : _JOB,
  '__module__' : 'job_pb2'
  # @@protoc_insertion_point(class_scope:job.Job)
  })
_sym_db.RegisterMessage(Job)


_JOBIDS.fields_by_name['id']._options = None

_QUERYJOB = _descriptor.ServiceDescriptor(
  name='QueryJob',
  full_name='job.QueryJob',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=461,
  serialized_end=582,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetJobs',
    full_name='job.QueryJob.GetJobs',
    index=0,
    containing_service=None,
    input_type=task__pb2._TASKID,
    output_type=_JOBLIST,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetJobIds',
    full_name='job.QueryJob.GetJobIds',
    index=1,
    containing_service=None,
    input_type=task__pb2._TASKID,
    output_type=_JOBIDS,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetJob',
    full_name='job.QueryJob.GetJob',
    index=2,
    containing_service=None,
    input_type=_JOBID,
    output_type=_JOB,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_QUERYJOB)

DESCRIPTOR.services_by_name['QueryJob'] = _QUERYJOB

# @@protoc_insertion_point(module_scope)
