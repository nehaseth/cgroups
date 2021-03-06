/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package v2

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestCgroupv2CpuStats(t *testing.T) {
	checkCgroupMode(t)
	group := "/cpu-test-cg"
	groupPath := fmt.Sprintf("%s-%d", group, os.Getpid())
	var (
		quota  int64  = 10000
		period uint64 = 8000
		weight uint64 = 100
	)
	max := "10000 8000"
	res := Resources{
		CPU: &CPU{
			Weight: &weight,
			Max:    NewCPUMax(&quota, &period),
			Cpus:   "0",
			Mems:   "0",
		},
	}
	c, err := NewManager(defaultCgroup2Path, groupPath, &res)
	if err != nil {
		t.Fatal("failed to init new cgroup manager: ", err)
	}
	defer os.Remove(c.path)

	checkFileContent(t, c.path, "cpu.weight", strconv.FormatUint(weight, 10))
	checkFileContent(t, c.path, "cpu.max", max)
	checkFileContent(t, c.path, "cpuset.cpus", "0")
	checkFileContent(t, c.path, "cpuset.mems", "0")
}
