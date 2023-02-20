/*
*  Copyright (c) 2023 NetEase Inc.
*
*  Licensed under the Apache License, Version 2.0 (the "License");
*  you may not use this file except in compliance with the License.
*  You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
 */

/*
* Project: Curve-Manager
* Created Date: 2023-02-11
* Author: wanghai (SeanHai)
 */

package manager

import (
	"github.com/mcuadros/go-defaults"
	"github.com/opencurve/curve-manager/api/curvebs/agent"
	"github.com/opencurve/curve-manager/api/curvebs/core"
	"github.com/opencurve/curve-manager/internal/errno"
	"github.com/opencurve/pigeon"
)

func ListVolume(r *pigeon.Request, ctx *Context) bool {
	data := ctx.Data.(*ListVolumeRequest)
	defaults.SetDefaults(data)
	volumes, err := agent.ListVolume(r, data.Size, data.Page, data.Path, data.SortKey, data.SortDirection)
	if err != errno.OK {
		return core.Exit(r, err)
	}
	return core.ExitSuccessWithData(r, volumes)
}

func GetVolume(r *pigeon.Request, ctx *Context) bool {
	data := ctx.Data.(*GetVolumeRequest)
	volume, err := agent.GetVolume(r, data.VolumeName)
	if err != errno.OK {
		return core.Exit(r, err)
	}
	return core.ExitSuccessWithData(r, volume)
}
