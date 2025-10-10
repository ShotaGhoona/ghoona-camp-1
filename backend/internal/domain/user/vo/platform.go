package vo

import "errors"

// Platform ソーシャルリンクのプラットフォーム種別を表すValue Object
type Platform string

const (
	PlatformTwitter   Platform = "twitter"   // Twitter
	PlatformInstagram Platform = "instagram" // Instagram
	PlatformGithub    Platform = "github"    // GitHub
	PlatformLinkedin  Platform = "linkedin"  // LinkedIn
	PlatformFacebook  Platform = "facebook"  // Facebook
	PlatformYoutube   Platform = "youtube"   // YouTube
	PlatformWebsite   Platform = "website"   // Webサイト
	PlatformBlog      Platform = "blog"      // ブログ
	PlatformPortfolio Platform = "portfolio" // ポートフォリオ
	PlatformOther     Platform = "other"     // その他
)

// NewPlatform Platformを作成する
func NewPlatform(value string) (Platform, error) {
	platform := Platform(value)
	if !platform.isValid() {
		return "", errors.New("不正なプラットフォームです")
	}
	return platform, nil
}

// isValid 有効なプラットフォーム値かどうかを検証する
func (p Platform) isValid() bool {
	switch p {
	case PlatformTwitter, PlatformInstagram, PlatformGithub, PlatformLinkedin,
		PlatformFacebook, PlatformYoutube, PlatformWebsite, PlatformBlog,
		PlatformPortfolio, PlatformOther:
		return true
	default:
		return false
	}
}

// String 文字列表現を返す
func (p Platform) String() string {
	return string(p)
}

// Value 内部の文字列値を返す
func (p Platform) Value() string {
	return string(p)
}