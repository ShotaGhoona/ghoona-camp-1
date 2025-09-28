package value

// Platform represents supported social media platforms
type Platform string

const (
	PlatformTwitter   Platform = "twitter"
	PlatformGitHub    Platform = "github"
	PlatformLinkedIn  Platform = "linkedin"
	PlatformWebsite   Platform = "website"
	PlatformBlog      Platform = "blog"
	PlatformYouTube   Platform = "youtube"
	PlatformInstagram Platform = "instagram"
)

// IsValid checks if the platform is supported
func (p Platform) IsValid() bool {
	switch p {
	case PlatformTwitter, PlatformGitHub, PlatformLinkedIn, 
		 PlatformWebsite, PlatformBlog, PlatformYouTube, PlatformInstagram:
		return true
	default:
		return false
	}
}

// String returns the string representation of the platform
func (p Platform) String() string {
	return string(p)
}