package model

import (
	"encoding/json"
	"strings"
	"time"
)

func NewRecord(title, description, priority string) Record {
	return Record{ID: title + "-" + time.Now().Format("150405.000"), Title: strings.TrimSpace(title), Description: strings.TrimSpace(description), Priority: priority, Status: StatusOpen, CreatedAt: time.Now(), UpdatedAt: time.Now()}
}
func EncodeRecord(r Record) ([]byte, error) { return json.Marshal(r) }
func DecodeRecord(b []byte) (Record, error) { var r Record; e := json.Unmarshal(b, &r); return r, e }
func NormalizePriority(p string) string {
	switch strings.ToLower(strings.TrimSpace(p)) {
	case "urgent", "high":
		return "high"
	case "low":
		return "low"
	default:
		return "normal"
	}
}
func CloneRecord(r Record) Record { r.Tags = append([]string(nil), r.Tags...); return r }
func AddTag(r *Record, tag string) bool {
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return false
	}
	for _, x := range r.Tags {
		if x == tag {
			return false
		}
	}
	r.Tags = append(r.Tags, tag)
	return true
}
