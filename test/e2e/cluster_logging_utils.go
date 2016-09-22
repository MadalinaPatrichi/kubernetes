/*
Copyright 2015 The Kubernetes Authors.

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

package e2e

import (
	"fmt"
	"time"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/test/e2e/framework"
)

const (
	// ingestionTimeout is how long to keep retrying to wait for all the
	// logs to be ingested.
	ingestionTimeout = 10 * time.Minute
	// ingestionRetryDelay is how long test should wait between
	// two attempts to check for ingestion
	ingestionRetryDelay = 25 * time.Second

	synthLoggerPodName = "synthlogger"

	// expectedLinesCount is the number of log lines emitted (and checked) for each synthetic logging pod.
	expectedLinesCount = 100
)

func createSynthLogger(f *framework.Framework, linesCount int) {
	f.PodClient().Create(&api.Pod{
		ObjectMeta: api.ObjectMeta{
			Name:      synthLoggerPodName,
			Namespace: f.Namespace.Name,
		},
		Spec: api.PodSpec{
			Containers: []api.Container{
				{
					Name:  synthLoggerPodName,
					Image: "gcr.io/google_containers/busybox:1.24",
					// notice: the subshell syntax is escaped with `$$`
					Command: []string{"/bin/sh", "-c", fmt.Sprintf("i=0; while [ $i -lt %d ]; do echo $i; i=`expr $i + 1`; done", linesCount)},
				},
			},
		},
	})
}
