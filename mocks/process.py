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

import sys
import time
from avalon import Avalon, SchedulerForwardReq
import warnings

if __name__ == '__main__':
    avalon = Avalon(working_root='.')
    print("waiting 3s")
    time.sleep(3)
    avalon.step_forward(SchedulerForwardReq.SchedulerSubmission(job_dir='job1', hostname='ism-gpu-a100'),
                        SchedulerForwardReq.SchedulerSubmission(job_dir='job2', hostname='ism-gpu-a100'))
    # warnings.warn("deprecated", DeprecationWarning)
