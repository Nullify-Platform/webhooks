package azuredevops

import (
	"fmt"
	"strings"
	"time"
)

// https://docs.microsoft.com/en-us/azure/devops/service-hooks/events

// azure devops does not send an event header, this BasicEvent is provided to get the EventType

type BasicEvent struct {
	ID          string `json:"id"`
	EventType   Event  `json:"eventType"`
	PublisherID string `json:"publisherId"`
	Scope       string `json:"scope"`
	CreatedDate Date   `json:"createdDate"`
}

// git.pullrequest.*
// git.pullrequest.created
// git.pullrequest.merged
// git.pullrequest.updated

type GitPullRequestEvent struct {
	ID                 string      `json:"id"`
	EventType          Event       `json:"eventType"`
	PublisherID        string      `json:"publisherId"`
	Scope              string      `json:"scope"`
	Message            Message     `json:"message"`
	DetailedMessage    Message     `json:"detailedMessage"`
	Resource           PullRequest `json:"resource"`
	ResourceVersion    string      `json:"resourceVersion"`
	ResourceContainers interface{} `json:"resourceContainers"`
	CreatedDate        Date        `json:"createdDate"`
}

// git.push

type GitPushEvent struct {
	CreatedDate        string             `json:"createdDate"`
	DetailedMessage    Message            `json:"detailedMessage"`
	EventType          string             `json:"eventType"`
	ID                 string             `json:"id"`
	Message            Message            `json:"message"`
	PublisherID        string             `json:"publisherId"`
	Resource           Resource           `json:"resource"`
	ResourceContainers ResourceContainers `json:"resourceContainers"`
	ResourceVersion    string             `json:"resourceVersion"`
	Scope              string             `json:"scope"`
}

// "ms.vss-code.git-pullrequest-comment-event"

type GitPullRequestCommentEvent struct {
	ID          string             `json:"id"`
	EventType   Event              `json:"eventType"`
	PublisherID string             `json:"publisherId"`
	Scope       string             `json:"scope"`
	Message     Message            `json:"message"`
	Resource    PullRequestComment `json:"resource"`
}

// build.complete

type BuildCompleteEvent struct {
	ID                 string      `json:"id"`
	EventType          Event       `json:"eventType"`
	PublisherID        string      `json:"publisherId"`
	Scope              string      `json:"scope"`
	Message            Message     `json:"message"`
	DetailedMessage    Message     `json:"detailedMessage"`
	Resource           Build       `json:"resource"`
	ResourceVersion    string      `json:"resourceVersion"`
	ResourceContainers interface{} `json:"resourceContainers"`
	CreatedDate        Date        `json:"createdDate"`
}

// -----------------------

type Message struct {
	Text     string `json:"text"`
	HTML     string `json:"html"`
	Markdown string `json:"markdown"`
}

type Commit struct {
	CommitID string `json:"commitId"`
	URL      string `json:"url"`
}

type PullRequest struct {
	Repository            Repository `json:"repository"`
	PullRequestID         int        `json:"pullRequestId"`
	Status                string     `json:"status"`
	CreatedBy             User       `json:"createdBy"`
	CreationDate          Date       `json:"creationDate"`
	ClosedDate            Date       `json:"closedDate"`
	Title                 string     `json:"title"`
	Description           string     `json:"description"`
	SourceRefName         string     `json:"sourceRefName"`
	TargetRefName         string     `json:"targetRefName"`
	MergeStatus           string     `json:"mergeStatus"`
	MergeID               string     `json:"mergeId"`
	LastMergeSourceCommit Commit     `json:"lastMergeSourceCommit"`
	LastMergeTargetCommit Commit     `json:"lastMergeTargetCommit"`
	LastMergeCommit       Commit     `json:"lastMergeCommit"`
	Reviewers             []Reviewer `json:"reviewers"`
	Commits               []Commit   `json:"commits"`
	URL                   string     `json:"url"`
}

type PullRequestComment struct {
	PullRequest PullRequest `json:"pullRequest"`
	Comment     Comment     `json:"comment"`
}

type Comment struct {
	ID                     int    `json:"id"`
	ParentCommentID        int    `json:"parentCommentId"`
	Content                string `json:"content"`
	Author                 User   `json:"author"`
	PublishedDate          Date   `json:"publishedDate"`
	LastUpdatedDate        Date   `json:"lastUpdatedDate"`
	LastContentUpdatedDate Date   `json:"lastContentUpdatedDate"`
	CommentType            string `json:"commentType"`
}

type Repository struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	URL           string  `json:"url"`
	Project       Project `json:"project"`
	DefaultBranch string  `json:"defaultBranch"`
	RemoteURL     string  `json:"remoteUrl"`
}

type Project struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	URL   string `json:"url"`
	State string `json:"state"`
}

type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	UniqueName  string `json:"uniqueName"`
	URL         string `json:"url"`
	ImageURL    string `json:"imageUrl"`
}

type Reviewer struct {
	ReviewerURL string `json:"reviewerUrl"`
	Vote        int    `json:"vote"`
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	UniqueName  string `json:"uniqueName"`
	URL         string `json:"url"`
	ImageURL    string `json:"imageUrl"`
	IsContainer bool   `json:"isContainer"`
}

type Build struct {
	URI                string          `json:"uri"`
	ID                 int             `json:"id"`
	BuildNumber        string          `json:"buildNumber"`
	URL                string          `json:"url"`
	StartTime          Date            `json:"startTime"`
	FinishTime         Date            `json:"finishTime"`
	Reason             string          `json:"reason"`
	Status             string          `json:"status"`
	DropLocation       string          `json:"dropLocation"`
	Drop               Drop            `json:"drop"`
	Log                Log             `json:"log"`
	SourceGetVersion   string          `json:"sourceGetVersion"`
	LastChangedBy      User            `json:"lastChangedBy"`
	RetainIndefinitely bool            `json:"retainIndefinitely"`
	HasDiagnostics     bool            `json:"hasDiagnostics"`
	Definition         BuildDefinition `json:"definition"`
	Queue              Queue           `json:"queue"`
	Requests           []Request       `json:"requests"`
}

type Drop struct {
	Location    string `json:"location"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	DownloadURL string `json:"downloadUrl"`
}

type Log struct {
	Type        string `json:"type"`
	URL         string `json:"url"`
	DownloadURL string `json:"downloadUrl"`
}

type BuildDefinition struct {
	BatchSize      int    `json:"batchSize"`
	TriggerType    string `json:"triggerType"`
	DefinitionType string `json:"definitionType"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	URL            string `json:"url"`
}

type Queue struct {
	QueueType string `json:"queueType"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
}

type Request struct {
	ID           int    `json:"id"`
	URL          string `json:"url"`
	RequestedFor User   `json:"requestedFor"`
}

type Resource struct {
	Commits    []Commit    `json:"commits"`
	Date       string      `json:"date"`
	PushedBy   PushedBy    `json:"pushedBy"`
	PushID     int         `json:"pushId"`
	RefUpdates []RefUpdate `json:"refUpdates"`
	Repository Repository  `json:"repository"`
	URL        string      `json:"url"`
}

type RefUpdate struct {
	Name        string `json:"name"`
	NewObjectID string `json:"newObjectId"`
	OldObjectID string `json:"oldObjectId"`
}

type PushedBy struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
	UniqueName  string `json:"uniqueName"`
}

type ResourceContainers struct {
	Account    Account `json:"account"`
	Collection Account `json:"collection"`
	Project    Account `json:"project"`
}

type Account struct {
	ID string `json:"id"`
}

type Date time.Time

func (b *Date) UnmarshalJSON(p []byte) error {
	t, err := time.Parse(time.RFC3339Nano, strings.Replace(string(p), "\"", "", -1))
	if err != nil {
		return err
	}
	*b = Date(t)
	return nil
}

func (b Date) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(b).Format(time.RFC3339Nano))
	return []byte(stamp), nil
}
