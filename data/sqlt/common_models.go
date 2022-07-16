package sqlt

import (
	"time"
)

type BaseEntity struct {
	GroupId      string `json:"group_id" gorm:"index"`
	SourceId     string `json:"source_id"  gorm:"index:idx_query"` //Index
	CreatedAtStr string `json:"created_at_str"`

	ImageFileName        string     `json:"image_file_name"`
	VideoFileName        string     `json:"video_file_name" gorm:"index"` //Index
	VideoFileCreatedDate *time.Time `json:"video_file_created_date"`
	VideoFileDuration    int        `json:"video_file_duration"`

	AiClipEnabled           bool   `json:"ai_clip_enabled"`
	AiClipFileName          string `json:"ai_clip_file_name" gorm:"index"` //Index
	AiClipCreatedAtStr      string `json:"ai_clip_created_at_str"`
	AiClipLastModifiedAtStr string `json:"ai_clip_last_modified_at_str"`
	AiClipDuration          int    `json:"ai_clip_duration"`

	CreatedDate time.Time `json:"created_date" gorm:"index:idx_query"` //Index
}