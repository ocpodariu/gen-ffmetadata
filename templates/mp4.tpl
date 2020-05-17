;FFMETADATA1
title={{ .Title }}
artist={{ .Artist }}
episode_id={{ .Episode }}
year={{ .Date.Year }}
description={{ .Date.Format "Mon, 02-Jan-2006"}}

{{ range .Chapters }}
[CHAPTER]
{{- /* start and end are in nanoseconds by default */ -}}
{{/* this timebase changes them to be expressed in milliseconds */}}
TIMEBASE=1/1000
START={{ .StartTime.Milliseconds }}
END={{ .EndTime.Milliseconds }}
title={{ .Title }}
{{ end }}

[STREAM]
title={{ .Title }}
