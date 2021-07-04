# Copyright 2021 TsumiNa
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from typing import Dict, Union

from pathlib import Path
import grpc
import yaml

from .protos.scheduler_pb2_grpc import SchedulerStub
from .protos.scheduler_pb2 import SchedulerForwardReq, SchedulerForwardFlag


class Avalon(object):

    def __init__(self, working_root: Union[str, Path], *, address: str = 'localhost:50051'):
        self._working_root = Path(working_root)
        self._address = address
        with open(self._working_root / 'logs' / 'monitor.yml') as f:
            task_state = yaml.load(f, Loader=yaml.FullLoader)
        self._task_id = task_state["task_id"]
        self._job_ids = task_state["job_ids"]

    def step_forward(
        self,
        *submissions: SchedulerForwardReq.SchedulerSubmission,
        flag: SchedulerForwardFlag = SchedulerForwardFlag.FORWARD,
        variables: Dict[str, str] = None,
    ):
        # NOTE(gRPC Python Team): .close() is possible on a channel and should be
        # used in circumstances in which the with statement does not fit the needs
        # of the code.
        with grpc.insecure_channel(self._address) as channel:
            stub = SchedulerStub(channel)
            response = stub.StepForward(
                SchedulerForwardReq(flag=flag, variables=variables, task_id=self._task_id, submissions=submissions))
        print("Forwarding with: ", response)
        return response
