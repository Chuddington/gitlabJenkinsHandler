package data

const (
	// IssueHook is a numerical ID for the Gitlab Hook for 'Issues'.
	IssueHook = 1000 << iota
	// MergeRequestHook is a numerical ID for the Gitlab Hook for
	// 'Merge Requests'.
	MergeRequestHook
	// NoteHook is a numerical ID for the Gitlab Hook for 'Notes'.
	NoteHook
	// PushHook is a numerical ID for the Gitlab Hook for 'Push' Events.
	PushHook
	// TagPushHook is a numerical ID for the Gitlab Hook for 'Tag Push' Events.
	TagPushHook
	// WikiPageHook is a numerical ID for the Gitlab Hook for 'Wiki Page' Updates.
	WikiPageHook
	// UnknownHook is a numerical ID for an Event not known to this program.
	UnknownHook
)

const (
	// NoMergeRequestsFound is a template string for use when a Merge Request
	// could not be found by any applicable Identifier.
	NoMergeRequestsFound = "No Merge Requests have been found"
	// GitlabDefaultURL is a default URL which is used when a Gitlab API URL is
	// not defined.  This variable is a placeholder for when the information has
	// not been defined in a configuration file.
	GitlabDefaultURL = "https://gitlab.com/api/v3/"
)
