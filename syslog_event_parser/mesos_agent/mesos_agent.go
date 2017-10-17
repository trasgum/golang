package mesos_agent

import (
	"net/http"
	"github.com/projectcalico/go-json/json"
)

type MesosAgent struct {
	Version   string  `json:"version"`
	GitSha    string  `json:"git_sha"`
	BuildDate string  `json:"build_date"`
	BuildTime float64 `json:"build_time"`
	BuildUser string  `json:"build_user"`
	StartTime float64 `json:"start_time"`
	ID        string  `json:"id"`
	Pid       string  `json:"pid"`
	Hostname  string  `json:"hostname"`
	Resources struct {
		Disk  float64 `json:"disk"`
		Mem   float64 `json:"mem"`
		Gpus  float64 `json:"gpus"`
		Cpus  float64 `json:"cpus"`
		Ports string  `json:"ports"`
	} `json:"resources"`
	Attributes []struct {
		Label string `json:"label"`
	} `json:"attributes"`
	MasterHostname  string `json:"master_hostname"`
	ExternalLogFile string `json:"external_log_file"`
	Flags           struct {
		SlaveSubsystems                  string `json:"slave_subsystems"`
		AppcSimpleDiscoveryURIPrefix     string `json:"appc_simple_discovery_uri_prefix"`
		AppcStoreDir                     string `json:"appc_store_dir"`
		Attributes                       string `json:"attributes"`
		AuthenticateHTTPReadonly         string `json:"authenticate_http_readonly"`
		AuthenticateHTTPReadwrite        string `json:"authenticate_http_readwrite"`
		Authenticatee                    string `json:"authenticatee"`
		AuthenticationBackoffFactor      string `json:"authentication_backoff_factor"`
		Authorizer                       string `json:"authorizer"`
		CgroupsCPUEnablePidsAndTidsCount string `json:"cgroups_cpu_enable_pids_and_tids_count"`
		CgroupsEnableCfs                 string `json:"cgroups_enable_cfs"`
		CgroupsHierarchy                 string `json:"cgroups_hierarchy"`
		CgroupsLimitSwap                 string `json:"cgroups_limit_swap"`
		CgroupsRoot                      string `json:"cgroups_root"`
		ContainerDiskWatchInterval       string `json:"container_disk_watch_interval"`
		ContainerLogger                  string `json:"container_logger"`
		Containerizers                   string `json:"containerizers"`
		Credential                       string `json:"credential"`
		DefaultRole                      string `json:"default_role"`
		DiskWatchInterval                string `json:"disk_watch_interval"`
		Docker                           string `json:"docker"`
		DockerKillOrphans                string `json:"docker_kill_orphans"`
		DockerRegistry                   string `json:"docker_registry"`
		DockerRemoveDelay                string `json:"docker_remove_delay"`
		DockerSocket                     string `json:"docker_socket"`
		DockerStopTimeout                string `json:"docker_stop_timeout"`
		DockerStoreDir                   string `json:"docker_store_dir"`
		DockerVolumeCheckpointDir        string `json:"docker_volume_checkpoint_dir"`
		EnforceContainerDiskQuota        string `json:"enforce_container_disk_quota"`
		ExecutorEnvironmentVariables     string `json:"executor_environment_variables"`
		ExecutorRegistrationTimeout      string `json:"executor_registration_timeout"`
		ExecutorShutdownGracePeriod      string `json:"executor_shutdown_grace_period"`
		ExternalLogFile                  string `json:"external_log_file"`
		FetcherCacheDir                  string `json:"fetcher_cache_dir"`
		FetcherCacheSize                 string `json:"fetcher_cache_size"`
		FrameworksHome                   string `json:"frameworks_home"`
		GcDelay                          string `json:"gc_delay"`
		GcDiskHeadroom                   string `json:"gc_disk_headroom"`
		HadoopHome                       string `json:"hadoop_home"`
		Help                             string `json:"help"`
		HostnameLookup                   string `json:"hostname_lookup"`
		HTTPAuthenticators               string `json:"http_authenticators"`
		HTTPCommandExecutor              string `json:"http_command_executor"`
		HTTPCredentials                  string `json:"http_credentials"`
		ImageProviders                   string `json:"image_providers"`
		ImageProvisionerBackend          string `json:"image_provisioner_backend"`
		InitializeDriverLogging          string `json:"initialize_driver_logging"`
		IPDiscoveryCommand               string `json:"ip_discovery_command"`
		Isolation                        string `json:"isolation"`
		LauncherDir                      string `json:"launcher_dir"`
		Logbufsecs                       string `json:"logbufsecs"`
		LoggingLevel                     string `json:"logging_level"`
		Master                           string `json:"master"`
		ModulesDir                       string `json:"modules_dir"`
		NetworkCniConfigDir              string `json:"network_cni_config_dir"`
		NetworkCniPluginsDir             string `json:"network_cni_plugins_dir"`
		OversubscribedResourcesInterval  string `json:"oversubscribed_resources_interval"`
		PerfDuration                     string `json:"perf_duration"`
		PerfInterval                     string `json:"perf_interval"`
		Port                             string `json:"port"`
		QosCorrectionIntervalMin         string `json:"qos_correction_interval_min"`
		Quiet                            string `json:"quiet"`
		Recover                          string `json:"recover"`
		RecoveryTimeout                  string `json:"recovery_timeout"`
		RegistrationBackoffFactor        string `json:"registration_backoff_factor"`
		Resources                        string `json:"resources"`
		RevocableCPULowPriority          string `json:"revocable_cpu_low_priority"`
		SandboxDirectory                 string `json:"sandbox_directory"`
		Strict                           string `json:"strict"`
		SwitchUser                       string `json:"switch_user"`
		SystemdEnableSupport             string `json:"systemd_enable_support"`
		SystemdRuntimeDirectory          string `json:"systemd_runtime_directory"`
		Version                          string `json:"version"`
		WorkDir                          string `json:"work_dir"`
	} `json:"flags"`
	Frameworks []struct {
		ID              string  `json:"id"`
		Name            string  `json:"name"`
		User            string  `json:"user"`
		FailoverTimeout float64 `json:"failover_timeout"`
		Checkpoint      bool    `json:"checkpoint"`
		Role            string  `json:"role"`
		Hostname        string  `json:"hostname"`
		Executors       []struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			Source    string `json:"source"`
			Container string `json:"container"`
			Directory string `json:"directory"`
			Resources struct {
				Disk float64 `json:"disk"`
				Mem  float64 `json:"mem"`
				Gpus float64 `json:"gpus"`
				Cpus float64 `json:"cpus"`
			} `json:"resources"`
			Labels []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"labels"`
			Tasks []struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				FrameworkID string `json:"framework_id"`
				ExecutorID  string `json:"executor_id"`
				SlaveID     string `json:"slave_id"`
				State       string `json:"state"`
				Resources   struct {
					Disk float64 `json:"disk"`
					Mem  float64 `json:"mem"`
					Gpus float64 `json:"gpus"`
					Cpus float64 `json:"cpus"`
				} `json:"resources"`
				Statuses []struct {
					State     string  `json:"state"`
					Timestamp float64 `json:"timestamp"`
					Labels    []struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"labels"`
					ContainerStatus struct {
						NetworkInfos []struct {
							Labels      []interface{} `json:"labels"`
							IPAddresses []struct {
								IPAddress string `json:"ip_address"`
							} `json:"ip_addresses"`
							Name string `json:"name"`
						} `json:"network_infos"`
					} `json:"container_status"`
				} `json:"statuses"`
				Labels []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"labels"`
				Discovery struct {
					Visibility string `json:"visibility"`
					Name       string `json:"name"`
					Ports      struct {
						Ports []struct {
							Number   int    `json:"number"`
							Name     string `json:"name"`
							Protocol string `json:"protocol"`
							Labels   struct {
								Labels []struct {
									Key   string `json:"key"`
									Value string `json:"value"`
								} `json:"labels"`
							} `json:"labels"`
						} `json:"ports"`
					} `json:"ports"`
				} `json:"discovery"`
				Container struct {
					Type   string `json:"type"`
					Docker struct {
						Image      string `json:"image"`
						Network    string `json:"network"`
						Privileged bool   `json:"privileged"`
						Parameters []struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"parameters"`
						ForcePullImage bool `json:"force_pull_image"`
					} `json:"docker"`
					NetworkInfos []struct {
						IPAddresses []struct {
						} `json:"ip_addresses"`
						Name   string `json:"name"`
						Labels struct {
						} `json:"labels"`
					} `json:"network_infos"`
				} `json:"container"`
			} `json:"tasks"`
			QueuedTasks    []interface{} `json:"queued_tasks"`
			CompletedTasks []interface{} `json:"completed_tasks"`
		} `json:"executors"`
		CompletedExecutors []interface{} `json:"completed_executors"`
	} `json:"frameworks"`
	CompletedFrameworks []interface{} `json:"completed_frameworks"`
}

func (r *MesosAgent) Poll(IPAddr string) error {
	resp, err := http.Get("http://" + IPAddr + ":5051/state")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

    	return json.NewDecoder(resp.Body).Decode(r)
}