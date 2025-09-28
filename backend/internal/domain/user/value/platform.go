package value

// Platform はサポートされているソーシャルメディアプラットフォームを表す
type Platform string

const (
	PlatformTwitter   Platform = "twitter"   // Twitter
	PlatformGitHub    Platform = "github"    // GitHub
	PlatformLinkedIn  Platform = "linkedin"  // LinkedIn
	PlatformWebsite   Platform = "website"   // ウェブサイト
	PlatformBlog      Platform = "blog"      // ブログ
)

// IsValid はプラットフォームがサポートされているかどうかを確認する
func (p Platform) IsValid() bool {
	switch p {
	case PlatformTwitter, PlatformGitHub, PlatformLinkedIn, 
		 PlatformWebsite, PlatformBlog:
		return true
	default:
		return false
	}
}

// String はプラットフォームの文字列表現を返す
func (p Platform) String() string {
	return string(p)
}