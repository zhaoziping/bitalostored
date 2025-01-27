// Copyright 2019-2024 Xu Ruibo (hustxurb@163.com) and Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build linux

package utils

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"time"
)

/*
#include <unistd.h>
*/
import "C"

type Usage struct {
	Utime      time.Duration `json:"utime"`
	Stime      time.Duration `json:"stime"`
	Cutime     time.Duration `json:"cutime"`
	Cstime     time.Duration `json:"cstime"`
	NumThreads int           `json:"num_threads"`
	VmSize     int64         `json:"vm_size"`
	VmRss      int64         `json:"vm_rss"`
}

func (u *Usage) MemTotal() int64 {
	return u.VmRss
}

func (u *Usage) CPUTotal() time.Duration {
	return time.Duration(u.Utime + u.Stime + u.Cutime + u.Cstime)
}

func GetUsage() (*Usage, error) {
	f, err := os.Open("/proc/self/stat")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ignore struct {
		s string
		d int64
	}

	r := bufio.NewReader(f)
	u := &Usage{}
	if _, err := fmt.Fscanf(r, "%d %s %s %d %d %d",
		&ignore.d, &ignore.s, &ignore.s, &ignore.d, &ignore.d, &ignore.d); err != nil {
		return nil, err
	}
	if _, err := fmt.Fscanf(r, "%d %d %d",
		&ignore.d, &ignore.d, &ignore.d); err != nil {
		return nil, err
	}
	if _, err := fmt.Fscanf(r, "%d %d %d %d",
		&ignore.d, &ignore.d, &ignore.d, &ignore.d); err != nil {
		return nil, err
	}

	var ticks struct {
		u int64
		s int64
	}
	unit := time.Second / time.Duration(C.sysconf(C._SC_CLK_TCK))

	if _, err := fmt.Fscanf(r, "%d %d",
		&ticks.u, &ticks.s); err != nil {
		return nil, err
	}
	u.Utime = time.Duration(ticks.u) * unit
	u.Stime = time.Duration(ticks.s) * unit

	if _, err := fmt.Fscanf(r, "%d %d",
		&ticks.u, &ticks.s); err != nil {
		return nil, err
	}
	u.Cutime = time.Duration(ticks.u) * unit
	u.Cstime = time.Duration(ticks.s) * unit

	if _, err := fmt.Fscanf(r, "%d %d %d %d %d",
		&ignore.d, &ignore.d, &u.NumThreads, &ignore.d, &ignore.d); err != nil {
		return nil, err
	}
	if _, err := fmt.Fscanf(r, "%d %d",
		&u.VmSize, &u.VmRss); err != nil {
		return nil, err
	}
	u.VmRss *= int64(syscall.Getpagesize())
	return u, nil
}
