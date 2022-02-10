package quotes

/*
Delete a tag
 git tag -d release/aug2002

  git push origin :refs/tags/release/aug2002
*/
func Favs() []string {
	return []string{"Love For All, Hatred For None.", "Change the world by being yourself.", "Every moment is a fresh beginning."}
}
