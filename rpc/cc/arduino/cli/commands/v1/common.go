// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package commands

import (
	"sort"

	semver "go.bug.st/relaxed-semver"
)

// DownloadProgressCB is a callback to get updates on download progress
type DownloadProgressCB func(curr *DownloadProgress)

// Start sends a "start" DownloadProgress message to the callback function
func (d DownloadProgressCB) Start(url, label string) {
	d(&DownloadProgress{
		Message: &DownloadProgress_Start{
			Start: &DownloadProgressStart{
				Url:   url,
				Label: label,
			},
		},
	})
}

// Update sends an "update" DownloadProgress message to the callback function
func (d DownloadProgressCB) Update(downloaded int64, totalSize int64) {
	d(&DownloadProgress{
		Message: &DownloadProgress_Update{
			Update: &DownloadProgressUpdate{
				Downloaded: downloaded,
				TotalSize:  totalSize,
			},
		},
	})
}

// End sends an "end" DownloadProgress message to the callback function
func (d DownloadProgressCB) End(success bool, message string) {
	d(&DownloadProgress{
		Message: &DownloadProgress_End{
			End: &DownloadProgressEnd{
				Success: success,
				Message: message,
			},
		},
	})
}

// TaskProgressCB is a callback to receive progress messages
type TaskProgressCB func(msg *TaskProgress)

// InstanceCommand is an interface that represents a gRPC command with
// a gRPC Instance.
type InstanceCommand interface {
	GetInstance() *Instance
}

// GetLatestRelease returns the latest release in this PlatformSummary,
// or nil if not available.
func (s *PlatformSummary) GetLatestRelease() *PlatformRelease {
	if s.GetLatestVersion() == "" {
		return nil
	}
	return s.GetReleases()[s.GetLatestVersion()]
}

// GetInstalledRelease returns the latest release in this PlatformSummary,
// or nil if not available.
func (s *PlatformSummary) GetInstalledRelease() *PlatformRelease {
	if s.GetInstalledVersion() == "" {
		return nil
	}
	return s.GetReleases()[s.GetInstalledVersion()]
}

// GetSortedReleases returns the releases in order of version.
func (s *PlatformSummary) GetSortedReleases() []*PlatformRelease {
	res := []*PlatformRelease{}
	for _, release := range s.GetReleases() {
		res = append(res, release)
	}
	sort.SliceStable(res, func(i, j int) bool {
		return semver.ParseRelaxed(res[i].GetVersion()).LessThan(semver.ParseRelaxed(res[j].GetVersion()))
	})
	return res
}
