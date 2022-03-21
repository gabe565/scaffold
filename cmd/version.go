package cmd

var (
	version = "next"
	commit  = ""
)

func buildVersion() string {
	result := version
	if commit != "" {
		result += " (" + commit + ")"
	}
	return result
}
