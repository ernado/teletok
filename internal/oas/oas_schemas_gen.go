// Code generated by ogen, DO NOT EDIT.

package oas

// Ref: #/components/schemas/Data
type Data struct {
	AnalyzeTime               OptString      "json:\"analyze_time\""
	APIURL                    OptString      "json:\"api_url\""
	NwmVideoURL               OptString      "json:\"nwm_video_url\""
	OriginalURL               OptString      "json:\"original_url\""
	Platform                  OptString      "json:\"platform\""
	Status                    string         "json:\"status\""
	URLType                   OptString      "json:\"url_type\""
	VideoAuthorDiggCount      OptMaybeNumber "json:\"video_author_diggCount\""
	VideoAuthorFollowerCount  OptMaybeNumber "json:\"video_author_followerCount\""
	VideoAuthorFollowingCount OptMaybeNumber "json:\"video_author_followingCount\""
	VideoAuthorHeartCount     OptMaybeNumber "json:\"video_author_heartCount\""
	VideoAuthorID             OptString      "json:\"video_author_id\""
	VideoAuthorNickname       OptString      "json:\"video_author_nickname\""
	VideoAuthorVideoCount     OptMaybeNumber "json:\"video_author_videoCount\""
	VideoAwemeID              OptString      "json:\"video_aweme_id\""
	VideoCommentCount         OptMaybeNumber "json:\"video_comment_count\""
	VideoCover                OptString      "json:\"video_cover\""
	VideoCreateTime           OptInt         "json:\"video_create_time\""
	VideoDiggCount            OptMaybeNumber "json:\"video_digg_count\""
	VideoDownloadCount        OptMaybeNumber "json:\"video_download_count\""
	VideoDynamicCover         OptString      "json:\"video_dynamic_cover\""
	VideoHashtags             []string       "json:\"video_hashtags\""
	VideoMusicAuthor          OptString      "json:\"video_music_author\""
	VideoMusicID              OptInt64       "json:\"video_music_id\""
	VideoMusicTitle           OptString      "json:\"video_music_title\""
	VideoMusicURL             OptString      "json:\"video_music_url\""
	VideoOriginCover          OptString      "json:\"video_origin_cover\""
	VideoPlayCount            OptMaybeNumber "json:\"video_play_count\""
	VideoShareCount           OptMaybeNumber "json:\"video_share_count\""
	VideoTitle                OptString      "json:\"video_title\""
	WmVideoURL                OptString      "json:\"wm_video_url\""
}

// Ref: #/components/schemas/MaybeNumber
// MaybeNumber represents sum type.
type MaybeNumber struct {
	Type  MaybeNumberType // switch on this field
	Int64 int64
	None  None
}

// MaybeNumberType is oneOf type of MaybeNumber.
type MaybeNumberType string

// Possible values for MaybeNumberType.
const (
	Int64MaybeNumber MaybeNumberType = "int64"
	NoneMaybeNumber  MaybeNumberType = "None"
)

// IsInt64 reports whether MaybeNumber is int64.
func (s MaybeNumber) IsInt64() bool { return s.Type == Int64MaybeNumber }

// IsNone reports whether MaybeNumber is None.
func (s MaybeNumber) IsNone() bool { return s.Type == NoneMaybeNumber }

// SetInt64 sets MaybeNumber to int64.
func (s *MaybeNumber) SetInt64(v int64) {
	s.Type = Int64MaybeNumber
	s.Int64 = v
}

// GetInt64 returns int64 and true boolean if MaybeNumber is int64.
func (s MaybeNumber) GetInt64() (v int64, ok bool) {
	if !s.IsInt64() {
		return v, false
	}
	return s.Int64, true
}

// NewInt64MaybeNumber returns new MaybeNumber from int64.
func NewInt64MaybeNumber(v int64) MaybeNumber {
	var s MaybeNumber
	s.SetInt64(v)
	return s
}

// SetNone sets MaybeNumber to None.
func (s *MaybeNumber) SetNone(v None) {
	s.Type = NoneMaybeNumber
	s.None = v
}

// GetNone returns None and true boolean if MaybeNumber is None.
func (s MaybeNumber) GetNone() (v None, ok bool) {
	if !s.IsNone() {
		return v, false
	}
	return s.None, true
}

// NewNoneMaybeNumber returns new MaybeNumber from None.
func NewNoneMaybeNumber(v None) MaybeNumber {
	var s MaybeNumber
	s.SetNone(v)
	return s
}

// Ref: #/components/schemas/None
type None string

const (
	NoneNone None = "None"
)

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt64 returns new OptInt64 with value set to v.
func NewOptInt64(v int64) OptInt64 {
	return OptInt64{
		Value: v,
		Set:   true,
	}
}

// OptInt64 is optional int64.
type OptInt64 struct {
	Value int64
	Set   bool
}

// IsSet returns true if OptInt64 was set.
func (o OptInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt64) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt64) Get() (v int64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt64) Or(d int64) int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMaybeNumber returns new OptMaybeNumber with value set to v.
func NewOptMaybeNumber(v MaybeNumber) OptMaybeNumber {
	return OptMaybeNumber{
		Value: v,
		Set:   true,
	}
}

// OptMaybeNumber is optional MaybeNumber.
type OptMaybeNumber struct {
	Value MaybeNumber
	Set   bool
}

// IsSet returns true if OptMaybeNumber was set.
func (o OptMaybeNumber) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMaybeNumber) Reset() {
	var v MaybeNumber
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMaybeNumber) SetTo(v MaybeNumber) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMaybeNumber) Get() (v MaybeNumber, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMaybeNumber) Or(d MaybeNumber) MaybeNumber {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}
