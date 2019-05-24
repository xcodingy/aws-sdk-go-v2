// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/GetHLSStreamingSessionURLInput
type GetHLSStreamingSessionURLInput struct {
	_ struct{} `type:"structure"`

	// Specifies which format should be used for packaging the media. Specifying
	// the FRAGMENTED_MP4 container format packages the media into MP4 fragments
	// (fMP4 or CMAF). This is the recommended packaging because there is minimal
	// packaging overhead. The other container format option is MPEG_TS. HLS has
	// supported MPEG TS chunks since it was released and is sometimes the only
	// supported packaging on older HLS players. MPEG TS typically has a 5-25 percent
	// packaging overhead. This means MPEG TS typically requires 5-25 percent more
	// bandwidth and cost than fMP4.
	//
	// The default is FRAGMENTED_MP4.
	ContainerFormat ContainerFormat `type:"string" enum:"true"`

	// Specifies when flags marking discontinuities between fragments will be added
	// to the media playlists. The default is ALWAYS when HLSFragmentSelector is
	// SERVER_TIMESTAMP, and NEVER when it is PRODUCER_TIMESTAMP.
	//
	// Media players typically build a timeline of media content to play, based
	// on the timestamps of each fragment. This means that if there is any overlap
	// between fragments (as is typical if HLSFragmentSelector is SERVER_TIMESTAMP),
	// the media player timeline has small gaps between fragments in some places,
	// and overwrites frames in other places. When there are discontinuity flags
	// between fragments, the media player is expected to reset the timeline, resulting
	// in the fragment being played immediately after the previous fragment. We
	// recommend that you always have discontinuity flags between fragments if the
	// fragment timestamps are not accurate or if fragments might be missing. You
	// should not place discontinuity flags between fragments for the player timeline
	// to accurately map to the producer timestamps.
	DiscontinuityMode DiscontinuityMode `type:"string" enum:"true"`

	// Specifies when the fragment start timestamps should be included in the HLS
	// media playlist. Typically, media players report the playhead position as
	// a time relative to the start of the first fragment in the playback session.
	// However, when the start timestamps are included in the HLS media playlist,
	// some media players might report the current playhead as an absolute time
	// based on the fragment timestamps. This can be useful for creating a playback
	// experience that shows viewers the wall-clock time of the media.
	//
	// The default is NEVER. When HLSFragmentSelector is SERVER_TIMESTAMP, the timestamps
	// will be the server start timestamps. Similarly, when HLSFragmentSelector
	// is PRODUCER_TIMESTAMP, the timestamps will be the producer start timestamps.
	DisplayFragmentTimestamp DisplayFragmentTimestamp `type:"string" enum:"true"`

	// The time in seconds until the requested session expires. This value can be
	// between 300 (5 minutes) and 43200 (12 hours).
	//
	// When a session expires, no new calls to GetHLSMasterPlaylist, GetHLSMediaPlaylist,
	// GetMP4InitFragment, or GetMP4MediaFragment can be made for that session.
	//
	// The default is 300 (5 minutes).
	Expires *int64 `min:"300" type:"integer"`

	// The time range of the requested fragment, and the source of the timestamps.
	//
	// This parameter is required if PlaybackMode is ON_DEMAND. This parameter is
	// optional if PlaybackMode is LIVE. If PlaybackMode is LIVE, the FragmentSelectorType
	// can be set, but the TimestampRange should not be set. If PlaybackMode is
	// ON_DEMAND, both FragmentSelectorType and TimestampRange must be set.
	HLSFragmentSelector *HLSFragmentSelector `type:"structure"`

	// The maximum number of fragments that are returned in the HLS media playlists.
	//
	// When the PlaybackMode is LIVE, the most recent fragments are returned up
	// to this value. When the PlaybackMode is ON_DEMAND, the oldest fragments are
	// returned, up to this maximum number.
	//
	// When there are a higher number of fragments available in a live HLS media
	// playlist, video players often buffer content before starting playback. Increasing
	// the buffer size increases the playback latency, but it decreases the likelihood
	// that rebuffering will occur during playback. We recommend that a live HLS
	// media playlist have a minimum of 3 fragments and a maximum of 10 fragments.
	//
	// The default is 5 fragments if PlaybackMode is LIVE, and 1,000 if PlaybackMode
	// is ON_DEMAND.
	//
	// The maximum value of 1,000 fragments corresponds to more than 16 minutes
	// of video on streams with 1-second fragments, and more than 2 1/2 hours of
	// video on streams with 10-second fragments.
	MaxMediaPlaylistFragmentResults *int64 `min:"1" type:"long"`

	// Whether to retrieve live or archived, on-demand data.
	//
	// Features of the two types of session include the following:
	//
	//    * LIVE: For sessions of this type, the HLS media playlist is continually
	//    updated with the latest fragments as they become available. We recommend
	//    that the media player retrieve a new playlist on a one-second interval.
	//    When this type of session is played in a media player, the user interface
	//    typically displays a "live" notification, with no scrubber control for
	//    choosing the position in the playback window to display.
	//
	// In LIVE mode, the newest available fragments are included in an HLS media
	//    playlist, even if there is a gap between fragments (that is, if a fragment
	//    is missing). A gap like this might cause a media player to halt or cause
	//    a jump in playback. In this mode, fragments are not added to the HLS media
	//    playlist if they are older than the newest fragment in the playlist. If
	//    the missing fragment becomes available after a subsequent fragment is
	//    added to the playlist, the older fragment is not added, and the gap is
	//    not filled.
	//
	//    * ON_DEMAND: For sessions of this type, the HLS media playlist contains
	//    all the fragments for the session, up to the number that is specified
	//    in MaxMediaPlaylistFragmentResults. The playlist must be retrieved only
	//    once for each session. When this type of session is played in a media
	//    player, the user interface typically displays a scrubber control for choosing
	//    the position in the playback window to display.
	//
	// In both playback modes, if FragmentSelectorType is PRODUCER_TIMESTAMP, and
	// if there are multiple fragments with the same start timestamp, the fragment
	// that has the larger fragment number (that is, the newer fragment) is included
	// in the HLS media playlist. The other fragments are not included. Fragments
	// that have different timestamps but have overlapping durations are still included
	// in the HLS media playlist. This can lead to unexpected behavior in the media
	// player.
	//
	// The default is LIVE.
	PlaybackMode PlaybackMode `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the stream for which to retrieve the HLS
	// master playlist URL.
	//
	// You must specify either the StreamName or the StreamARN.
	StreamARN *string `min:"1" type:"string"`

	// The name of the stream for which to retrieve the HLS master playlist URL.
	//
	// You must specify either the StreamName or the StreamARN.
	StreamName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetHLSStreamingSessionURLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetHLSStreamingSessionURLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetHLSStreamingSessionURLInput"}
	if s.Expires != nil && *s.Expires < 300 {
		invalidParams.Add(aws.NewErrParamMinValue("Expires", 300))
	}
	if s.MaxMediaPlaylistFragmentResults != nil && *s.MaxMediaPlaylistFragmentResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxMediaPlaylistFragmentResults", 1))
	}
	if s.StreamARN != nil && len(*s.StreamARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamARN", 1))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetHLSStreamingSessionURLInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.ContainerFormat) > 0 {
		v := s.ContainerFormat

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ContainerFormat", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.DiscontinuityMode) > 0 {
		v := s.DiscontinuityMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DiscontinuityMode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.DisplayFragmentTimestamp) > 0 {
		v := s.DisplayFragmentTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DisplayFragmentTimestamp", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Expires != nil {
		v := *s.Expires

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Expires", protocol.Int64Value(v), metadata)
	}
	if s.HLSFragmentSelector != nil {
		v := s.HLSFragmentSelector

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "HLSFragmentSelector", v, metadata)
	}
	if s.MaxMediaPlaylistFragmentResults != nil {
		v := *s.MaxMediaPlaylistFragmentResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxMediaPlaylistFragmentResults", protocol.Int64Value(v), metadata)
	}
	if len(s.PlaybackMode) > 0 {
		v := s.PlaybackMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PlaybackMode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StreamARN != nil {
		v := *s.StreamARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamName != nil {
		v := *s.StreamName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/GetHLSStreamingSessionURLOutput
type GetHLSStreamingSessionURLOutput struct {
	_ struct{} `type:"structure"`

	// The URL (containing the session token) that a media player can use to retrieve
	// the HLS master playlist.
	HLSStreamingSessionURL *string `type:"string"`
}

// String returns the string representation
func (s GetHLSStreamingSessionURLOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetHLSStreamingSessionURLOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.HLSStreamingSessionURL != nil {
		v := *s.HLSStreamingSessionURL

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "HLSStreamingSessionURL", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetHLSStreamingSessionURL = "GetHLSStreamingSessionURL"

// GetHLSStreamingSessionURLRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams Archived Media.
//
// Retrieves an HTTP Live Streaming (HLS) URL for the stream. You can then open
// the URL in a browser or media player to view the stream contents.
//
// You must specify either the StreamName or the StreamARN.
//
// An Amazon Kinesis video stream has the following requirements for providing
// data through HLS:
//
//    * The media must contain h.264 encoded video and, optionally, AAC encoded
//    audio. Specifically, the codec id of track 1 should be V_MPEG/ISO/AVC.
//    Optionally, the codec id of track 2 should be A_AAC.
//
//    * Data retention must be greater than 0.
//
//    * The video track of each fragment must contain codec private data in
//    the Advanced Video Coding (AVC) for H.264 format (MPEG-4 specification
//    ISO/IEC 14496-15 (https://www.iso.org/standard/55980.html)). For information
//    about adapting stream data to a given format, see NAL Adaptation Flags
//    (http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/producer-reference-nal.html).
//
//    * The audio track (if present) of each fragment must contain codec private
//    data in the AAC format (AAC specification ISO/IEC 13818-7 (https://www.iso.org/standard/43345.html)).
//
// Kinesis Video Streams HLS sessions contain fragments in the fragmented MPEG-4
// form (also called fMP4 or CMAF), rather than the MPEG-2 form (also called
// TS chunks, which the HLS specification also supports). For more information
// about HLS fragment types, see the HLS specification (https://tools.ietf.org/html/draft-pantos-http-live-streaming-23).
//
// The following procedure shows how to use HLS with Kinesis Video Streams:
//
// Get an endpoint using GetDataEndpoint (http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/API_GetDataEndpoint.html),
// specifying GET_HLS_STREAMING_SESSION_URL for the APIName parameter.
//
// Retrieve the HLS URL using GetHLSStreamingSessionURL. Kinesis Video Streams
// creates an HLS streaming session to be used for accessing content in a stream
// using the HLS protocol. GetHLSStreamingSessionURL returns an authenticated
// URL (that includes an encrypted session token) for the session's HLS master
// playlist (the root resource needed for streaming with HLS).
//
// Don't share or store this token where an unauthorized entity could access
// it. The token provides access to the content of the stream. Safeguard the
// token with the same measures that you would use with your AWS credentials.
//
// The media that is made available through the playlist consists only of the
// requested stream, time range, and format. No other media data (such as frames
// outside the requested window or alternate bitrates) is made available.
//
// Provide the URL (containing the encrypted session token) for the HLS master
// playlist to a media player that supports the HLS protocol. Kinesis Video
// Streams makes the HLS media playlist, initialization fragment, and media
// fragments available through the master playlist URL. The initialization fragment
// contains the codec private data for the stream, and other data needed to
// set up the video or audio decoder and renderer. The media fragments contain
// H.264-encoded video frames or AAC-encoded audio samples.
//
// The media player receives the authenticated URL and requests stream metadata
// and media data normally. When the media player requests data, it calls the
// following actions:
//
// GetHLSMasterPlaylist: Retrieves an HLS master playlist, which contains a
// URL for the GetHLSMediaPlaylist action for each track, and additional metadata
// for the media player, including estimated bitrate and resolution.
//
// GetHLSMediaPlaylist: Retrieves an HLS media playlist, which contains a URL
// to access the MP4 initialization fragment with the GetMP4InitFragment action,
// and URLs to access the MP4 media fragments with the GetMP4MediaFragment actions.
// The HLS media playlist also contains metadata about the stream that the player
// needs to play it, such as whether the PlaybackMode is LIVE or ON_DEMAND.
// The HLS media playlist is typically static for sessions with a PlaybackType
// of ON_DEMAND. The HLS media playlist is continually updated with new fragments
// for sessions with a PlaybackType of LIVE. There is a distinct HLS media playlist
// for the video track and the audio track (if applicable) that contains MP4
// media URLs for the specific track.
//
// GetMP4InitFragment: Retrieves the MP4 initialization fragment. The media
// player typically loads the initialization fragment before loading any media
// fragments. This fragment contains the "fytp" and "moov" MP4 atoms, and the
// child atoms that are needed to initialize the media player decoder.
//
// The initialization fragment does not correspond to a fragment in a Kinesis
// video stream. It contains only the codec private data for the stream and
// respective track, which the media player needs to decode the media frames.
//
// GetMP4MediaFragment: Retrieves MP4 media fragments. These fragments contain
// the "moof" and "mdat" MP4 atoms and their child atoms, containing the encoded
// fragment's media frames and their timestamps.
//
// After the first media fragment is made available in a streaming session,
// any fragments that don't contain the same codec private data cause an error
// to be returned when those different media fragments are loaded. Therefore,
// the codec private data should not change between fragments in a session.
// This also means that the session fails if the fragments in a stream change
// from having only video to having both audio and video.
//
// Data retrieved with this action is billable. See Pricing (https://aws.amazon.com/kinesis/video-streams/pricing/)
// for details.
//
// GetTSFragment: Retrieves MPEG TS fragments containing both initialization
// and media data for all tracks in the stream.
//
// If the ContainerFormat is MPEG_TS, this API is used instead of GetMP4InitFragment
// and GetMP4MediaFragment to retrieve stream media.
//
// Data retrieved with this action is billable. For more information, see Kinesis
// Video Streams pricing (https://aws.amazon.com/kinesis/video-streams/pricing/).
//
// The following restrictions apply to HLS sessions:
//
// A streaming session URL should not be shared between players. The service
// might throttle a session if multiple media players are sharing it. For connection
// limits, see Kinesis Video Streams Limits (http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/limits.html).
//
// A Kinesis video stream can have a maximum of five active HLS streaming sessions.
// If a new session is created when the maximum number of sessions is already
// active, the oldest (earliest created) session is closed. The number of active
// GetMedia connections on a Kinesis video stream does not count against this
// limit, and the number of active HLS sessions does not count against the active
// GetMedia connection limit.
//
// You can monitor the amount of data that the media player consumes by monitoring
// the GetMP4MediaFragment.OutgoingBytes Amazon CloudWatch metric. For information
// about using CloudWatch to monitor Kinesis Video Streams, see Monitoring Kinesis
// Video Streams (http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/monitoring.html).
// For pricing information, see Amazon Kinesis Video Streams Pricing (https://aws.amazon.com/kinesis/video-streams/pricing/)
// and AWS Pricing (https://aws.amazon.com/pricing/). Charges for both HLS sessions
// and outgoing AWS data apply.
//
// For more information about HLS, see HTTP Live Streaming (https://developer.apple.com/streaming/)
// on the Apple Developer site (https://developer.apple.com).
//
//    // Example sending a request using GetHLSStreamingSessionURLRequest.
//    req := client.GetHLSStreamingSessionURLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/GetHLSStreamingSessionURL
func (c *Client) GetHLSStreamingSessionURLRequest(input *GetHLSStreamingSessionURLInput) GetHLSStreamingSessionURLRequest {
	op := &aws.Operation{
		Name:       opGetHLSStreamingSessionURL,
		HTTPMethod: "POST",
		HTTPPath:   "/getHLSStreamingSessionURL",
	}

	if input == nil {
		input = &GetHLSStreamingSessionURLInput{}
	}

	req := c.newRequest(op, input, &GetHLSStreamingSessionURLOutput{})
	return GetHLSStreamingSessionURLRequest{Request: req, Input: input, Copy: c.GetHLSStreamingSessionURLRequest}
}

// GetHLSStreamingSessionURLRequest is the request type for the
// GetHLSStreamingSessionURL API operation.
type GetHLSStreamingSessionURLRequest struct {
	*aws.Request
	Input *GetHLSStreamingSessionURLInput
	Copy  func(*GetHLSStreamingSessionURLInput) GetHLSStreamingSessionURLRequest
}

// Send marshals and sends the GetHLSStreamingSessionURL API request.
func (r GetHLSStreamingSessionURLRequest) Send(ctx context.Context) (*GetHLSStreamingSessionURLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetHLSStreamingSessionURLResponse{
		GetHLSStreamingSessionURLOutput: r.Request.Data.(*GetHLSStreamingSessionURLOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetHLSStreamingSessionURLResponse is the response type for the
// GetHLSStreamingSessionURL API operation.
type GetHLSStreamingSessionURLResponse struct {
	*GetHLSStreamingSessionURLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetHLSStreamingSessionURL request.
func (r *GetHLSStreamingSessionURLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
