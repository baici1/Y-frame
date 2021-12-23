package configs

type FileUploadSetting struct {
	Size               int64    `yaml:"Size" json:"size,omitempty" mapstructure:"Size"`
	UploadFileField    string   `yaml:"UploadFileField" json:"upload_file_field,omitempty" mapstructure:"UploadFileField"`
	UploadFileSavePath string   `yaml:"UploadFileSavePath" json:"upload_file_save_path,omitempty" mapstructure:"UploadFileSavePath"`
	AllowMimeType      []string `yaml:"AllowMimeType" json:"allow_mime_type,omitempty" mapstructure:"AllowMimeType"`
}
