# Introducing Rec Files, YAML Data About Recordings

As the amount of recorded video and audio data increases the need for a
succinct way to document these recordings continues increase. Yet little
is being done to catalog and document these recordings in an organized
way that can be easily maintained *during* the recording itself without
too much interference (for example, by a live streamer during an
educational stream).

YouTube has partially responded to this need by allowing chapter markers
to be provided as offset time stamps, but the limited and unstable
format is insufficient to provide a sustainable solution that can be
archived, transferred, and queried. Hence this specification suggestion.

## "Rec Files"

A colloquial way to refer to these YAML data files about recordings is
"rec files," as in, "did you make a rec file for that live stream you
captured?" "Meta" and "YAML" are terms that are either too vague or
played out. "Data" file would be specific as well.

## Specification

### Integration with Chronological Logs

Recordings are usually associated with a specific span of time. The data
about them is, therefore, frequently organized with other log data. This
makes one view of day, for example, a convenient collection of log
entries about that day and any live recordings that were made on that
day. Combining this information often provides the best context.

### Suggest Directory Name: `rec`

The name `rec` is suggested when combined with other log data. This
allows the recordings to be further documented in a `README.md` file
that could easily be generated from the YAML structured data.
## YAML Example

See the [example](2021-01-21_05-18-51.yml) in this directory.

### File Name

Files are named using the standard name from OBS (the File field) with
"no spaces in filename" option but with the `.yml` suffix:

```
2021-01-21_12-08-38.yml
```

This format can easily be converted into an ISO8601 time stamp but
should *not* be depended on since it lacks a timezone information that
is only available in the YAML data as *Recorded* (which is also used for
*Marks*.)

### Schema

```yaml
%YAML 1.2
---
!!map {
  !!str "Title"    : !!str
  ? !!str "Recorded" : !!timestamp
File:        !!str
Description: !!str
Services:    !!seq
- Service:   !!str
  ID:        !!str
Marks:       !!seq
- Time:      !!timestamp
  Note:      !!str
```

Fields are denoted in their `yq` query format for specificity and convenience.

#### `.Title`

#### `.Recorded`

#### `.File`

#### `.Description`

#### `.Services[]`

#### `.Services[].Service`

#### `.Services[].ID`

#### `.Marks[]`

#### `.Marks[].Time`

#### `.Marks[].Note`

## FAQ

### What's the Difference Between "Video" and "Recording"?

A video generally means something that was produced for publishing
(despite the generic term). Therefore, it seems like "recording" is the
most accurate term for, well, recording something whether it be audio or
video.

#### What About Stream?

A recording can't really be called a "stream" because it's a recording
of a stream.

#### What About Vlog?

Even though a "vlog" is a recording, the term "vlog" has been taken by
people talking to the camera mostly. The term "recording" is more broad
than that.

