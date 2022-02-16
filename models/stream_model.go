package models

type StreamModel struct {
	Id          string `json:"id" redis:"id"`
	Brand       string `json:"brand" redis:"brand"`
	Name        string `json:"name" redis:"name"`
	RtspAddress string `json:"rtsp_address" redis:"rtsp_address"`

	Pid         int    `json:"pid" redis:"pid"`
	CreatedAt   string `json:"created_at" redis:"created_at"`
	Args        string `json:"args" redis:"args"`
	FailedCount int    `json:"failedCount" redis:"failedCount"`

	StreamType              int    `json:"stream_type" redis:"stream_type"`
	RtmpServerInitialized   bool   `json:"rtmp_server_initialized" redis:"rtmp_server_initialized"`
	RtmpServerType          int    `json:"rtmp_server_type" redis:"rtmp_server_type"`
	FlvPlayerConnectionType int    `json:"flv_player_connection_type" redis:"flv_player_connection_type"`
	NeedReloadInterval      int    `json:"need_reload_interval" redis:"need_reload_interval"`
	RtmpImageName           string `json:"rtmp_image_name" redis:"rtmp_image_name"`
	RtmpContainerName       string `json:"rtmp_container_name" redis:"rtmp_container_name"`
	RtmpAddress             string `json:"rtmp_address" redis:"rtmp_address"`
	RtmpFlvAddress          string `json:"rtmp_flv_address" redis:"rtmp_flv_address"`
	RtmpContainerPorts      string `json:"rtmp_container_ports" redis:"rtmp_container_ports"`
	RtmpContainerCommands   string `json:"rtmp_container_commands" redis:"rtmp_container_commands"`

	Record         bool `json:"record" redis:"record"`
	RecordDuration int  `json:"record_duration" redis:"record_duration"`

	HlsOutputPath         string `json:"hls_output_path" redis:"hls_output_path"`
	ReadJpegOutputPath    string `json:"read_jpeg_output_path" redis:"read_jpeg_output_path"`
	RecorOutputFolderPath string `json:"record_output_folder_path" redis:"record_output_folder_path"`
}
