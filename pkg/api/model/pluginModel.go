// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package model

//CreatePluginStruct CreatePluginStruct
//swagger:parameters createPlugin
type CreatePluginStruct struct {
	// in: path
	// required: true
	TenantName string `json:"tenant_name"`
	// in: body
	Body struct {
		//in: body
		//required: true
		PluginID string `json:"plugin_id" validate:"plugin_id|required"`
		//in: body
		//required: true
		PluginName string `json:"plugin_name" validate:"plugin_name|required"`
		//插件用途描述
		//in: body
		//required: false
		PluginInfo string `json:"plugin_info" validate:"plugin_info"`
		//插件docker地址
		//in: body
		//required: false
		ImageURL string `json:"image_url" validate:"image_url"`
		//插件goodrain地址
		//in: body
		//required: false
		ImageLocal string `json:"image_local" validate:"image_local"`
		//带分支信息的git地址
		//in: body
		//required: false
		Repo string `json:"repo" validate:"repo"`
		//git地址
		//in: body
		//required: false
		GitURL string `json:"git_url" validate:"git_url"`
		//构建模式
		//in: body
		//required: false
		BuildModel string `json:"build_model" validate:"build_model"`
		//插件模式
		//in: body
		//required: false
		PluginModel string `json:"plugin_model" validate:"plugin_model"`
		//插件启动命令
		//in: body
		//required: false
		PluginCMD string `json:"plugin_cmd" validate:"plugin_cmd"`
		//in: body
		//required: false
		TenantID string `json:"tenant_id" validate:"tenant_id"`
		//in: body
		//required: false
		EVNInfo []*PluginDefaultENV `json:"env_info" validate:"env_info"`
	}
}

//UpdatePluginStruct UpdatePluginStruct
//swagger:parameters updatePlugin
type UpdatePluginStruct struct {
	// in: path
	// required: true
	TenantName string `json:"tenant_name"`
	// in: path
	// required: true
	PluginID string `json:"plugin_id"`
	// in: body
	Body struct {
		//插件名称
		//in: body
		//required: false
		PluginName string `json:"plugin_name" validate:"plugin_name"`
		//插件用途描述
		//in: body
		//required: false
		PluginInfo string `json:"plugin_info" validate:"plugin_info"`
		//插件docker地址
		//in: body
		//required: false
		ImageURL string `json:"image_url" validate:"image_url"`
		//插件goodrain地址
		//in: body
		//required: false
		ImageLocal string `json:"image_local" validate:"image_local"`
		//带分支信息的git地址
		//in: body
		//required: false
		Repo string `json:"repo" validate:"repo"`
		//git地址
		//in: body
		//required: false
		GitURL string `json:"git_url" validate:"git_url"`
		//构建模式
		//in: body
		//required: false
		BuildModel string `json:"build_model" validate:"build_model"`
		//插件模式
		//in: body
		//required: false
		PluginModel string `json:"plugin_model" validate:"plugin_model"`
		//插件启动命令
		//in: body
		//required: false
		PluginCMD string `json:"plugin_cmd" validate:"plugin_cmd"`
	}
}

//deletePluginStruct deletePluginStruct
//swagger:parameters deletePlugin
type deletePluginStruct struct {
	// in: path
	// required: true
	TenantName string `json:"tenant_name"`
	// in: path
	// required: true
	PluginID string `json:"plugin_id"`
}

//ENVStruct ENVStruct
//swagger:parameters adddefaultenv updatedefaultenv
type ENVStruct struct {
	// in: path
	// required: true
	TenantName string `json:"tenant_name"`
	// in: path
	// required: true
	PluginID string `json:"plugin_id"`
	//in : body
	Body struct {
		//in: body
		//required: true
		EVNInfo []*PluginDefaultENV
	}
}

//DeleteENVstruct DeleteENVstruct
//swagger:parameters deletedefaultenv
type DeleteENVstruct struct {
	// in: path
	// required: true
	TenantName string `json:"tenant_name"`
	// in: path
	// required: true
	PluginID string `json:"plugin_id"`
	//配置项名称
	//in: path
	//required: true
	ENVName string `json:"env_name" validate:"env_name"`
}

//PluginDefaultENV 插件默认环境变量
type PluginDefaultENV struct {
	//对应插件id
	//in: path
	//required: false
	PluginID string `json:"plugin_id" validate:"plugin_id"`
	//配置项名称
	//in: path
	//required: true
	ENVName string `json:"env_name" validate:"env_name"`
	//配置项值
	//in: path
	//required: true
	ENVValue string `json:"env_value" validate:"env_value"`
	//是否可以被使用者修改
	//in :path
	//required: false
	IsChange bool `json:"is_change" validate:"is_change|bool"`
}

//BuildPluginStruct BuildPluginStruct
//swagger:parameters buildPlugin
type BuildPluginStruct struct {
	// in: path
	// required: true
	TenantName string `json:"tenant_name" validate:"tenant_name"`
	// in: path
	// required: true
	PluginID string `json:"plugin_id" validate:"plugin_id"`
	//in: body
	Body struct {
		// the event id
		// in: body
		// required: true
		EventID string `json:"event_id" validate:"event_id|required"`
		// 插件构建类型, image/dockerfile
		// in: body
		// required: true
		Kind string `json:"kind" validate:"kind|required"`
		// 插件CPU权重, 默认500
		// in: body
		// required: false
		PluginCPU int `json:"plugin_cpu" validate:"plugin_cpu"`
		// 插件最大内存, 默认128
		// in: body
		// required: false
		PluginMemory int `json:"plugin_memory" validate:"plugin_memory"`
		// 镜像地址
		// in: body
		// required: false
		ImageURL string `json:"image_url" validate:"image_url"`
		// 部署的版本号
		// in: body
		// required: true
		DeployVersion string `json:"deploy_version" validate:"deploy_version|required"`
		// git地址 带版本
		// in: body
		// required: false
		RepoURL string `json:"repo_url" validate:"repo_url"`
		// git地址
		// in: body
		// required: false
		GitURL string `json:"git_url" validate:"git_url"`
		// 版本信息, 协助选择插件版本
		// in:body
		// required: true
		Info string `json:"info" validate:"info"`
		// 操作人
		// in: body
		// required: false
		Operator string `json:"operator" validate:"operator"`
		//租户id
		// in: body
		// required: false
		TenantID string `json:"tenant_id" validate:"tenant_id"`
	}
}

//PluginBuildVersionStruct PluginBuildVersionStruct
//swagger:parameters deletePluginVersion pluginVersion
type PluginBuildVersionStruct struct {
	// in: path
	// required: true
	TenantName string `json:"tenant_name" validate:"tenant_name"`
	// in: path
	// required: true
	PluginID string `json:"plugin_id" validate:"plugin_id"`
	//in : path
	//required: true
	VersionID string `json:"version_id" validate:"version_id"`
}

//AllPluginBuildVersionStruct AllPluginBuildVersionStruct
//swagger:parameters allPluginVersions
type AllPluginBuildVersionStruct struct {
	// in: path
	// required: true
	TenantName string `json:"tenant_name" validate:"tenant_name"`
	// in: path
	// required: true
	PluginID string `json:"plugin_id" validate:"plugin_id"`
}

//PluginSetStruct PluginSetStruct
//swagger:parameters updatePluginSet addPluginSet
type PluginSetStruct struct {
	// in: path
	// required: true
	TenantName string `json:"tenant_name"`
	// in: path
	// required: true
	ServiceAlias string `json:"service_alias"`
	// in: body
	Body struct {
		// 插件id
		// in: body
		// required: true
		PluginID string `json:"plugin_id" validate:"plugin_id"`
		// 插件版本
		// in: body
		// required: true
		VersionID string `json:"version_id" validate:"version_id"`
		// 开关
		// in: body
		//required: false
		Switch bool `json:"switch" validate:"switch|bool"`
	}
}

//GetPluginSetStruct GetPluginSetStruct
//swagger:parameters getPluginSet
type GetPluginSetStruct struct {
	// in: path
	// required: true
	TenantName string `json:"tenant_name"`
	// in: path
	// required: true
	ServiceAlias string `json:"service_alias"`
}

//DeletePluginSetStruct DeletePluginSetStruct
//swagger:parameters deletePluginRelation
type DeletePluginSetStruct struct {
	// in: path
	// required: true
	TenantName string `json:"tenant_name"`
	// in: path
	// required: true
	ServiceAlias string `json:"service_alias"`
	// 插件id
	// in: path
	// required: true
	PluginID string `json:"plugin_id"`
}

//GetPluginEnvStruct GetPluginEnvStruct
//swagger:parameters getPluginEnv
type GetPluginEnvStruct struct {
	// in: path
	// required: true
	TenantName string `json:"tenant_name"`
	// in: path
	// required: true
	PluginID string `json:"plugin_id"`
}