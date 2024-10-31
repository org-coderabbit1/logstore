package httpreq

import (
	"context"
	"net/http"
	"strings"

	"example.com/acme/kit/httpgrpc"
)

type EncodingFlag string

type EncodingFlags map[EncodingFlag]struct{}

func NewEncodingFlags(flags ...EncodingFlag) EncodingFlags {
	var ef EncodingFlags
	ef.Set(flags...)
	return ef
}

func (ef *EncodingFlags) Set(flags ...EncodingFlag) {
	if *ef == nil {
		*ef = make(EncodingFlags, len(flags))
	}

	for _, flag := range flags {
		(*ef)[flag] = struct{}{}
	}
}

func (ef *EncodingFlags) Has(flag EncodingFlag) bool {
	_, ok := (*ef)[flag]
	return ok
}

func (ef *EncodingFlags) String() string {
	var sb strings.Builder
	var i int
	for flag := range *ef {
		if i > 0 {
			sb.WriteString(EncodeFlagsDelimiter)
		}
		sb.WriteString(string(flag))
		i++
	}
	return sb.String()
}

const (
	LogstoreEncodingFlagsHeader              = "X-Logstore-Response-Encoding-Flags"
	FlagCategorizeLabels    EncodingFlag = "categorize-labels"

	EncodeFlagsDelimiter = ","
)

func AddEncodingFlags(req *http.Request, flags EncodingFlags) {
	if len(flags) == 0 {
		return
	}

	req.Header.Set(LogstoreEncodingFlagsHeader, flags.String())
}

func AddEncodingFlagsToContext(ctx context.Context, flags EncodingFlags) context.Context {
	if len(flags) == 0 {
		return ctx
	}

	return context.WithValue(ctx, headerContextKey(LogstoreEncodingFlagsHeader), flags.String())
}

func ExtractEncodingFlags(req *http.Request) EncodingFlags {
	rawValue := req.Header.Get(LogstoreEncodingFlagsHeader)
	return ParseEncodingFlags(rawValue)
}

func ExtractEncodingFlagsFromProto(req *httpgrpc.HTTPRequest) EncodingFlags {
	var rawValue string
	for _, header := range req.GetHeaders() {
		if header.GetKey() == LogstoreEncodingFlagsHeader {
			rawValue = header.GetValues()[0]
			return ParseEncodingFlags(rawValue)
		}
	}

	return nil
}

func ExtractEncodingFlagsFromCtx(ctx context.Context) EncodingFlags {
	rawValue := ExtractHeader(ctx, LogstoreEncodingFlagsHeader)
	if rawValue == "" {
		return nil
	}

	return ParseEncodingFlags(rawValue)
}

func ParseEncodingFlags(rawFlags string) EncodingFlags {
	if rawFlags == "" {
		return nil
	}

	split := strings.Split(rawFlags, EncodeFlagsDelimiter)
	flags := make(EncodingFlags, len(split))
	for _, rawFlag := range split {
		flags.Set(EncodingFlag(rawFlag))
	}
	return flags
}
